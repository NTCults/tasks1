package main

import (
	"fmt"
	"sync"
)

func squareSum() {
	sum := &summCounter{
		mu: &sync.Mutex{},
	}
	data := [...]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, v := range data {
		wg.Add(1)
		go func() {
			sum.add(v * v)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(sum.result())
}

type summCounter struct {
	mu   *sync.Mutex
	summ int
}

func (sc *summCounter) add(val int) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.summ += val
}

func (sc *summCounter) result() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.summ
}
