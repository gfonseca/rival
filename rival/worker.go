package rival

import (
	"errors"
	"time"
)

// Worker abstraction of go routine
type Worker struct {
	exec      func(raw interface{}, index int)
	lt        time.Duration
	startedAt time.Time
	index     int
	running   bool
}

// Run execute worker
func (w *Worker) run(raw interface{}, index int) {
	w.running = true
	w.startedAt = time.Now()
	w.exec(raw, index)
	w.lt = time.Now().Sub(w.startedAt)
	w.running = false
}

// MakeWorker build a new worker
func MakeWorker(f func(raw interface{}, index int)) *Worker {
	return &Worker{
		f,
		time.Duration(0),
		time.Now(),
		-1,
		false,
	}
}

// WorkerPool a container for controll the running workers
type WorkerPool struct {
	Capacity   int
	Terminated int
	Running    int
	Output     []interface{}
	workers    []*Worker
	index      int32
}

// MakeWorkerPool build a WorkerPool
func MakeWorkerPool(capacity int) WorkerPool {
	return WorkerPool{
		capacity,
		0,
		0,
		[]interface{}{},
		make([]*Worker, capacity),
		0,
	}
}

// Submit Add a worker to WorkerPool
func (wp *WorkerPool) Submit(w *Worker, raw interface{}) error {
	if wp.Running >= wp.Capacity {
		return errors.New("WorkerPool is full")
	}
	wp.workers = append(wp.workers, w)
	go w.run(raw, w.index)
	w.index++
	wp.Running++
	return nil
}
