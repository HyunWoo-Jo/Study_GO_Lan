package main

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	fmt.Println("Push", x)
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	//list
	l := list.New()
	l.PushBack(10)
	l.PushBack(20)
	l.PushBack(30)

	fmt.Println("Front", l.Front().Value)
	fmt.Println("Back", l.Back().Value)

	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
	// heap

	data := new(MinHeap)

	heap.Init(data)
	heap.Push(data, 5)
	heap.Push(data, 2)
	heap.Push(data, 7)
	heap.Push(data, 3)

	fmt.Println(data, "최소값 : ", (*data)[0])
	// ring

	datas := []string{"Maria", "John", "Andrew", "James"}

	r := ring.New(len(datas))
	for i := 0; i < r.Len(); i++ {
		r.Value = datas[i]
		r = r.Next()
	}

	r.Do(func(x interface{}) {
		fmt.Println(x)
	})

	fmt.Println("Move forward :")
	r = r.Move(1)

	fmt.Println("Curr : ", r.Value)
	fmt.Println("Prev : ", r.Prev().Value)
	fmt.Println("Next : ", r.Next().Value)

}
