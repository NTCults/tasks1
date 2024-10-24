package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func chanReadWrite() {

	numSeconds := 2

	ch := make(chan int)

	go func() {
		for {
			select {
			case val, ok := <-ch:
				if ok {
					fmt.Println(val)
				} else {
					return
				}

			}
		}
	}()

	t := time.NewTimer(time.Second * time.Duration(numSeconds))
	for {
		select {
		case <-t.C:
			close(ch)
			return
		default:
			ch <- rand.Int()
		}
		time.Sleep(time.Millisecond * 500)
	}
}
