package utils

import (
	"sync"
	"time"
)

const (
	QueueTickInterval = time.Second
)

type TaskFunc func()
type CompleteFunc func() bool

type Task struct {
	TaskFunc     TaskFunc
	CompleteFunc CompleteFunc
}

type TaskQueue struct {
	tasks  []*Task
	ticker *time.Ticker
	mutex  *sync.Mutex
}

func NewTaskQueue() *TaskQueue {
	return &TaskQueue{
		tasks:  make([]*Task, 0),
		ticker: time.NewTicker(QueueTickInterval),
		mutex:  &sync.Mutex{},
	}
}

func (q *TaskQueue) Run() {
	go func() {
		for range q.ticker.C {
			q.mutex.Lock()

			index := 0
			for index < len(q.tasks) {
				q.tasks[index].TaskFunc()
				if q.tasks[index].CompleteFunc() {
					q.tasks = append(q.tasks[:index], q.tasks[index:]...)
				}
			}

			q.mutex.Unlock()
		}
	}()
}

func (q *TaskQueue) AddTask(task *Task) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	q.tasks = append(q.tasks, task)
}
