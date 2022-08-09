package storage

import (
	"sync"
)

var (
	// mgrRegistry 上下文记录注册表
	mgrRegistry = make(map[*ContextManager]bool)
	// mgrRegistryMtx 注册表操作读写锁
	mgrRegistryMtx sync.RWMutex
)

// Values : key -> value类型的映射
type Values map[interface{}]interface{}

type ContextManager struct {
	mtx    sync.Mutex
	values map[uint]Values
}

func NewContextManager() *ContextManager {
	mgr := &ContextManager{values: make(map[uint]Values)}
	mgrRegistryMtx.Lock()
	defer mgrRegistryMtx.Unlock()
	mgrRegistry[mgr] = true
	return mgr
}

func (m *ContextManager) Unregister() {
	mgrRegistryMtx.Lock()
	defer mgrRegistryMtx.Unlock()
	// 删除表中的注册记录
	delete(mgrRegistry, m)
}

// SetValues 设置上下文 ， contextCall 传入一个方法，设置完成之后进行回调三
func (m *ContextManager) SetValues(newValues Values, contextCall func()) {
	if len(newValues) == 0 {
		contextCall()
		return
	}

	mutatedKeys := make([]interface{}, 0, len(newValues))
	mutatedValues := make(Values, len(newValues))

	EnsureGoroutineId(func(gid uint) {
		m.mtx.Lock()
		state, found := m.values[gid]
		if !found {
			state = make(Values, len(newValues))
			m.values[gid] = state
		}
		m.mtx.Unlock()

		for key, newVal := range newValues {
			mutatedKeys = append(mutatedKeys, key)
			if oldVal, ok := state[key]; ok {
				mutatedValues[key] = oldVal
			}
			state[key] = newVal
		}

		defer func() {
			if !found {
				m.mtx.Lock()
				delete(m.values, gid)
				m.mtx.Unlock()
				return
			}

			for _, key := range mutatedKeys {
				if val, ok := mutatedValues[key]; ok {
					state[key] = val
				} else {
					delete(state, key)
				}
			}
		}()

		contextCall()
	})
}

func (m *ContextManager) GetValue(key interface{}) (
	value interface{}, ok bool) {
	gid, ok := GetGoroutineId()
	if !ok {
		return nil, false
	}

	m.mtx.Lock()
	state, found := m.values[gid]
	m.mtx.Unlock()

	if !found {
		return nil, false
	}
	value, ok = state[key]
	return value, ok
}

func (m *ContextManager) getValues() Values {
	gid, ok := GetGoroutineId()
	if !ok {
		return nil
	}
	m.mtx.Lock()
	state, _ := m.values[gid]
	m.mtx.Unlock()
	return state
}

func Go(cb func()) {
	mgrRegistryMtx.RLock()
	defer mgrRegistryMtx.RUnlock()

	for mgr := range mgrRegistry {
		values := mgr.getValues()
		if len(values) > 0 {
			cb = func(mgr *ContextManager, cb func()) func() {
				return func() { mgr.SetValues(values, cb) }
			}(mgr, cb)
		}
	}

	go cb()
}
