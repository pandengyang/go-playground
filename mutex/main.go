package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func main() {
	wg.Add(10)

	go postUuids(1)
	go postUuids(2)
	go postUuids(3)
	go postUuids(4)
	go postUuids(5)
	go postUuids(6)
	go postUuids(7)
	go postUuids(8)
	go postUuids(9)
	go postUuids(10)

	wg.Wait()
}

func postUuids(id int) {
	defer wg.Done()

	mutex.Lock()
	defer mutex.Unlock()

	fmt.Printf("[%d] Lock\n", id)
	time.Sleep(time.Duration(id) * time.Second)
	fmt.Printf("[%d] POST /uuids\n", id)
	fmt.Printf("[%d] Unlock\n", id)
}
