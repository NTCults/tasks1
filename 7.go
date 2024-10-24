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

func concurrentMap() {
	m := &ConcurrentMap{
		mu:       &sync.RWMutex{},
		innerMap: make(map[string]interface{}),
	}

	ctx, cancel := context.WithCancel(context.Background())

	for i := range 10 {
		go mapWorker(ctx, m, i)
	}

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	<-exit
	cancel()

	for k, v := range m.innerMap {
		fmt.Println(k, v)
	}

}

func mapWorker(ctx context.Context, m *ConcurrentMap, workerID int) {
	keyCounter := 0
	t := time.Tick(time.Second * 1)
	for {
		select {
		case <-t:
			key := fmt.Sprintf("%d:%d", workerID, keyCounter)
			keyCounter++
			m.Add(key, rand.IntN(100))
		case <-ctx.Done():
			return
		}
	}
}

type ConcurrentMap struct {
	mu       *sync.RWMutex
	innerMap map[string]interface{}
}

func (c *ConcurrentMap) Add(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.innerMap[key] = value
}

func (c *ConcurrentMap) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.innerMap, key)
}

func (c *ConcurrentMap) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.innerMap[key]
}
