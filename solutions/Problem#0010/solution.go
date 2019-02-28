package main

import (
	"fmt"
	"sync"
	"time"
)

type function func()

var wg sync.WaitGroup

func main() {
	defer wg.Wait()
	fooFunc := func() { fmt.Println("100 ms") }
	runFuncAfterNMilli(fooFunc, 100)
	barFunc := func() { fmt.Println("80 ms") }
	runFuncAfterNMilli(barFunc, 80)
}

func runFuncAfterNMilli(callback function, n int64) {
	wg.Add(1)
	go func() {
		time.Sleep(time.Duration(n) * time.Millisecond)
		callback()
		defer wg.Done()
	}()
}
