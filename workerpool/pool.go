package workerpool

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	inputChan chan string // Канал для передачи задач воркерам
	workers   map[int]*Worker // Словарь для хранения воркеров по их ID
	mu        sync.Mutex // Мьютекс для синхронизации доступа к пулу воркеров
	counter   int // Счетчик для генерации уникальных ID воркеров
}

func (p *WorkerPool) WorkerCount() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return len(p.workers)
}

func NewWorkerPool(inputChan chan string) *WorkerPool {
	return &WorkerPool{
		inputChan: inputChan,
		workers:   make(map[int]*Worker),
	}
}

// Добавляет новый воркер в пул
func (wp *WorkerPool) Add() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	wp.counter++
	id := wp.counter
	worker := NewWorker(id, wp.inputChan)
	wp.workers[id] = worker
	worker.Start()
}

// Убирает последний воркер из пула
func (wp *WorkerPool) Remove() {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	if len(wp.workers) == 0 {
		fmt.Println("worker pool is empty.")
		return
	}
	var lastID int
	for id := range wp.workers {
		lastID = id
		break
	}
	worker := wp.workers[lastID]
	worker.Stop()
	delete(wp.workers, lastID)
}
