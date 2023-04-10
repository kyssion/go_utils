package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			fmt.Printf("x=%v\n", x)
			// x=0
			time.Sleep(time.Millisecond * 10)
		}
	}()
	go storeFunc()
	wg.Wait()
}

func test(num int) {
	var startTime = time.Now()
	for i := 0; i < 500; i++ {
		fibonacci(num)
	}
	fmt.Println(time.Since(startTime).Milliseconds())

}
func fibonacci(i int) int {
	if i < 2 {
		return i
	}
	return fibonacci(i-2) + fibonacci(i-1)
}

var x int64 = 0

func storeFunc() {
	for i := 0; ; i++ {
		// time.Sleep(time.Millisecond * 10)
		if i%2 == 0 {
			x = 2
		} else {
			x = 1
		}
	}
}
