package workerpool

import (
	"fmt"
)

type Worker struct {
	ID   int // Уникальный идентификатор воркера
	Quit chan struct{} // Канал для остановки воркера
	JobChan chan string // Канал для получения задач от пула воркеров
}

func NewWorker(id int, jobChan chan string) *Worker {
	return &Worker{
		ID:   id,
		Quit: make(chan struct{}),
		JobChan: jobChan,
	}
}

func (w *Worker) Start() {
	go func() {
		fmt.Printf("Worker %d started\n", w.ID)
		for {
			select {
			case job := <-w.JobChan:
				fmt.Printf("Worker %d processing job: %s\n", w.ID, job)
			case <-w.Quit:
				fmt.Printf("Worker %d stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w *Worker) Stop() {
	close(w.Quit)
}