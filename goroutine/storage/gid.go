package storage

var stackTagPool = &idPool{}
var markLookup [16]func(uint, func())

func GetGoroutineId() (gid uint, ok bool) {
	return readStackTag()
}

func EnsureGoroutineId(cb func(gid uint)) {
	if gid, ok := readStackTag(); ok {
		cb(gid)
		return
	}
	gid := stackTagPool.Acquire()
	defer stackTagPool.Release(gid)
	addStackTag(gid, func() { cb(gid) })
}

func addStackTag(tag uint, contextCall func()) {
	if contextCall == nil {
		return
	}
	_m(tag, contextCall)
}

func _m(tagRemainder uint, cb func()) {
	if tagRemainder == 0 {
		cb()
	} else {
		markLookup[tagRemainder&0xf](tagRemainder>>bitWidth, cb)
	}
}
