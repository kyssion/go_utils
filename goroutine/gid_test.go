package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestGetGIDWithStack(t *testing.T) {
	fmt.Println("start", GetGIDWithStack())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("index : %d , gid : %d \n", i, GetGIDWithStack())
		}()
	}
	wg.Wait()
}
