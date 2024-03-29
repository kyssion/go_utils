package goroutine

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GetGIDWithStack 在上下文中获取GID
func GetGIDWithStack() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
