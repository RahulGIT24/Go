package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition - Rahul")

	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(4)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Go routine 1")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Go routine 2")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Go routine 3")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		defer wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		defer wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}
