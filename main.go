package main

import (
	"fmt"
	"sync"
	"time"
)

func AsyncFunc(text string, delay time.Duration, wg *sync.WaitGroup) {
	async_txt := "Async-Fire"
	go func(prefix_text string) {
		defer wg.Done()
		time.Sleep(delay)
		fmt.Printf("Async Version: %s:%s", prefix_text, text)
	}(async_txt) // Note the parentheses. We must call the anonymous function.
}

func SyncFuncExecAsyncFunc(text string, delay time.Duration) {
	async_txt := "Async-Fire"

	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(prefix_text string, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(delay)
			fmt.Printf("\nSync Calling Async Version: [ %s:%s ]", prefix_text, text)
		}(async_txt, &wg) // Note the parentheses. We must call the anonymous function.
	}
	wg.Wait()
}

func main() {

	fmt.Println("In main()")
	//var wg sync.WaitGroup
	//wg.Add(2)
	var duration_secs time.Duration = (3) * time.Second
	//Publish("Fire", duration_secs, &wg)
	SyncFuncExecAsyncFunc("Sync-to-Async", duration_secs)
	//wg.Wait()
}
