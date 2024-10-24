package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func workers() {
	numWorkers := getNumWorkers()

	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}
	ch := make(chan interface{})
	for i := range numWorkers {
		wg.Add(1)
		go worker(ctx, ch, i, wg)
	}

	t := time.Tick(time.Millisecond * 500)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
		<-exit
		cancel()
	}()
	for {
		val := rand.IntN(100)

		select {
		case <-t:
			ch <- val
		case <-ctx.Done():
			wg.Wait()
			return
		}
	}

}

func getNumWorkers() int {
	fmt.Print("Enter num workers: ")

	var numWorkers int
	_, err := fmt.Scanf("%d", &numWorkers)
	if err != nil {
		fmt.Println(err.Error())
		numWorkers = getNumWorkers()
	}

	return numWorkers
}

func worker(ctx context.Context, ch chan interface{}, workerID int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case val := <-ch:
			fmt.Printf("Worker %d: %v\n", workerID, val)
		case <-ctx.Done():
			fmt.Printf("Worker %d stop\n", workerID)
			return
		}
	}
}
