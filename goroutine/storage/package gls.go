package storage

import "runtime"

const (
	bitWidth       = 4
	stackBatchSize = 16
)

var (
	pcLookup = make(map[uintptr]int8, 17)
)

func readStackTag() (tag uint, ok bool) {
	var currenTag uint
	offset := 0
	for {
		batch, nextOffset := getStack(offset, stackBatchSize)
		for _, pc := range batch {
			val, ok := pcLookup[pc]
			if !ok {
				continue
			}
			if val < 0 {
				return currenTag, true
			}
			currenTag <<= bitWidth
			currenTag += uint(val)
		}
		if nextOffset == 0 {
			break
		}
		offset = nextOffset
	}
	return 0, false
}

func getStack(offset, amount int) (stack []uintptr, nextOffset int) {
	stack = make([]uintptr, amount)
	stack = stack[:runtime.Callers(offset, stack)]
	if len(stack) < amount {
		return stack, 0
	}
	return stack, offset + len(stack)
}
