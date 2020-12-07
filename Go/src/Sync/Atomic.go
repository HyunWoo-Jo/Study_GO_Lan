package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

/*
	atomic
	add: 더하고 결과를 리턴
	compareandSwap : 변수 A오 B를 비교해서 같으면 C를 대입 or bool형으로 반환
	Load: 변수에서 값을 가지고옴
	Store: 변수에 값을 저장함
	Swap : 변수에 새 값을 대입하고 이전 값을 리턴

*/
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var data int32 = 0
	wg := new(sync.WaitGroup)

	for i := 0; i < 2000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, 1)
			wg.Done()
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&data, -1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("data :", data)
}
