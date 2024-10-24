package main

import (
	"fmt"
	"sync"
)

func concurrentUnordered() {
	//в задании сказано, что это массив
	data := [...]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	for _, v := range data {
		wg.Add(1)
		go func() {
			fmt.Println(v * v)
			wg.Done()
		}()
	}
	wg.Wait()
}

func concurrentOrdered() {
	data := [...]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	result := [len(data)]int{}
	for i, v := range data {
		wg.Add(1)
		go func() {
			val := v * v
			result[i] = val
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(result)
	for _, i := range result {
		fmt.Println(i)
	}
}
