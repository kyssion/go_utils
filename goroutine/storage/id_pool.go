package storage

import (
	"sync"
)

type idPool struct {
	mtx      sync.Mutex
	released []uint
	maxId    uint
}

func (p *idPool) Acquire() (id uint) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	if len(p.released) > 0 {
		id = p.released[len(p.released)-1]
		p.released = p.released[:len(p.released)-1]
		return id
	}
	id = p.maxId
	p.maxId++
	return id
}

func (p *idPool) Release(id uint) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.released = append(p.released, id)
}
