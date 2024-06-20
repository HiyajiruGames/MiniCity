package resource

import (
	"fmt"
	"time"
)

// Queue implementation
type Queue struct {
	data []Resourcer
	size int
}

// NewQueue instantiates a new queue
func NewQueue(cap int) *Queue {
	return &Queue{data: make([]Resourcer, 0, cap), size: 0}
}

// Push adds a new element at the end of the queue
func (q *Queue) Push(r Resourcer) {
	q.data = append(q.data, r)
	q.size++
}

// Pop removes the first element from queue
func (q *Queue) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.size--
	q.data = q.data[1:]
	return true
}

// Front returns the first element of queue
func (q *Queue) Front() Resourcer {
	return q.data[0]
}

// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

// String implements Stringer interface
func (q *Queue) String() string {
	return fmt.Sprint(q.data)
}

type Loader struct {
	queue Queue
	ch    chan bool
}

func NewLoader() *Loader {
	l := new(Loader)
	l.queue = *NewQueue(10)
	return l
}

func (l *Loader) Load(r Resourcer) {
	l.queue.Push(r)
}

func (l *Loader) Start() {
	l.ch = make(chan bool)
	f := func() {
		for {
			if l.queue.IsEmpty() {
				time.Sleep(100 * time.Millisecond)
			} else {
				r := l.queue.Front()
				r.Load()
				l.queue.Pop()
			}
		}
	}
	go f()
}

func (l *Loader) Stop() {
	// Release has resources after load thread exits from loop.
}
