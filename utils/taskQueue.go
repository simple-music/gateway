package utils

import (
	"fmt"
	"github.com/simple-music/gateway/logs"
	"sync"
	"time"
)

const (
	QueueTickInterval = time.Second
)

type TaskFunc func() bool

type Task struct {
	TaskFunc TaskFunc
}

type TaskQueue struct {
	tasks  []*Task
	ticker *time.Ticker
	mutex  *sync.Mutex
	logger *logs.Logger
}

func NewTaskQueue(logger *logs.Logger) *TaskQueue {
	return &TaskQueue{
		tasks:  make([]*Task, 0),
		ticker: time.NewTicker(QueueTickInterval),
		mutex:  &sync.Mutex{},
		logger: logger,
	}
}

func (q *TaskQueue) Run() {
	go func() {
		for range q.ticker.C {
			q.mutex.Lock()

			index := 0
			for index < len(q.tasks) {
				if q.tasks[index].TaskFunc() {
					q.tasks = append(q.tasks[:index], q.tasks[index+1:]...)
				}
				index++
			}

			n := len(q.tasks)
			if n > 0 {
				q.logger.Info(fmt.Sprintf("Task queue not empty. Length: %d", n))
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
