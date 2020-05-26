package queue

import (
	"sync"
)

type (
	Queue interface {
		Len() int
		UnsafeLen() int
		Push(ele interface{})
		Pop() interface{}
		Del(i int)
		UnsafeDel(i int)
		Get(i int) interface{}
		UnsafeGet(i int) interface{}
		Set(index int, v interface{})
		UnsafeSet(index int, v interface{})
		Range(func(i int, v interface{}))
		UnsafeRange(func(i int, v interface{}))
		Index(data interface{}) (ok bool, i int)
		UnsafeIndex(data interface{}) (ok bool, i int)
	}

	queue struct {
		mu    *sync.RWMutex
		count int
		data  []interface{}
	}
)

var _ Queue = &queue{}

func NewQueue(defaultCap int) Queue {
	q := &queue{
		mu:    &sync.RWMutex{},
		count: 0,
		data:  make([]interface{}, 0, defaultCap),
	}
	return q
}

func (q *queue) init() {
	q.count = 0
	q.data = make([]interface{}, 0)
}

func (q *queue) checkLen(i int) {
	if q.count-1 < i {
		panic("out of range")
	}
}

func (q *queue) Len() int {
	q.mu.RLock()
	defer q.mu.RUnlock()
	return q.count
}

func (q *queue) UnsafeLen() int {
	return q.count
}

//add
func (q *queue) Push(ele interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.count++
	q.data = append(q.data, ele)
}

//get&remove
func (q *queue) Pop() interface{} {
	if q.Len() == 0 {
		return nil
	}
	q.mu.Lock()
	data := (q.data)[0]
	q.data = q.data[1:]
	q.count--
	q.mu.Unlock()
	return data
}

//update
func (q *queue) Set(i int, v interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.checkLen(i)
	q.data[i] = v
}

func (q *queue) UnsafeSet(i int, v interface{}) {
	q.checkLen(i)
	q.data[i] = v
}

//del
func (q *queue) Del(i int) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.checkLen(i)
	q.data = append(q.data[:i], q.data[i+1:]...)
}

func (q *queue) UnsafeDel(i int) {
	q.checkLen(i)
	q.data = append(q.data[:i], q.data[i+1:]...)
}

//Get
func (q *queue) Get(i int) interface{} {
	q.mu.RLock()
	defer q.mu.RUnlock()
	q.checkLen(i)
	d := q.data[i]
	return d
}

func (q *queue) UnsafeGet(i int) interface{} {
	q.checkLen(i)
	d := q.data[i]
	return d
}

//safe range
func (q *queue) Range(f func(i int, v interface{})) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for i, v := range q.data {
		f(i, v)
	}
}

//unsafe range
func (q *queue) UnsafeRange(f func(i int, v interface{})) {
	for i, v := range q.data {
		f(i, v)
	}
}

func (q *queue) Index(data interface{}) (bool, int) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	for i, v := range q.data {
		if v == data {
			return true, i
		}
	}
	return false, 0
}

func (q *queue) UnsafeIndex(data interface{}) (bool, int) {
	for i, v := range q.data {
		if v == data {
			return true, i
		}
	}
	return false, 0
}
