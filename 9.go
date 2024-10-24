package main

import (
	"fmt"
)

func conveyor() {
	array := [...]int{1, 8, 35, 95, 82, 7, 64, 3}
	chan1 := make(chan int)
	go func() {
		for i := 0; i < len(array); i++ {
			chan1 <- array[i]
		}
		close(chan1)
		return
	}()

	chan2 := make(chan int)
	go func() {
		for {
			select {
			case i, ok := <-chan1:
				if ok {
					chan2 <- i * 2
				} else {
					close(chan2)
					return
				}
			}
		}
	}()
	for {
		select {
		case val, ok := <-chan2:
			if ok {
				fmt.Println(val)
			} else {
				return
			}
		}
	}
}
