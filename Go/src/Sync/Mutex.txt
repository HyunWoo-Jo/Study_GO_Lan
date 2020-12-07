package main

import (
	"fmt"
	_ "fmt"
	"runtime"
	_ "runtime"
	"sync"
	_ "sync"
	"time"
	_ "time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	data := 0
	rwMutex := new(sync.RWMutex)

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.Lock()
			data++
			fmt.Println("write : ", data)
			time.Sleep(10 * time.Millisecond)
			rwMutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 1 : ", data)
			time.Sleep(1 * time.Second)
			rwMutex.RUnlock()
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			rwMutex.RLock()
			fmt.Println("read 2 : ", data)
			time.Sleep(2 * time.Second)
			rwMutex.RUnlock()
		}
	}()
	time.Sleep(10 * time.Second)
}
