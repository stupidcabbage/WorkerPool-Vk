package main

import (
	"workerpool/workerpool"
	"fmt"
	"time"
)

func main() {
	inputChan := make(chan string, 10)
	pool := workerpool.NewWorkerPool(inputChan)

	pool.Add()
	pool.Add()

	go func() {
		for i := 0; i < 10; i++ {
			inputChan <- fmt.Sprintf("Task %d", i)
			time.Sleep(300 * time.Millisecond)
		}
	}()

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("Adding a new worker")
		pool.Add()
	})

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("Removing a worker")
		pool.Remove()
	})

	time.AfterFunc(5*time.Second, func() {
		fmt.Printf("Now working %d\n", pool.WorkerCount())
		pool.Remove()
	})

	time.Sleep(6 * time.Second)
	fmt.Println("Done")
}
