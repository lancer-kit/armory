package queue

import (
	"container/list"
	"sync"
)

// Queue is a simple concurrency-safe FIFO implementation.
type Queue struct {
	list     *list.List
	index    map[string]struct{}
	elements map[string]*list.Element

	rwMutex sync.RWMutex
}

// Init inits empty Queue.
func (*Queue) Init() *Queue {
	q := &Queue{
		list:     list.New(),
		index:    make(map[string]struct{}),
		elements: make(map[string]*list.Element),
		rwMutex:  sync.RWMutex{},
	}
	return q
}

// Put adds value to queue.
func (q *Queue) Put(uid string, data interface{}) {
	if q.IsInQueue(uid) {
		return
	}

	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()

	q.elements[uid] = q.list.PushBack(data)
	q.index[uid] = struct{}{}
}

// Delete removes value from queue.
func (q *Queue) Delete(uid string) {
	if !q.IsInQueue(uid) {
		return
	}

	q.rwMutex.Lock()
	defer q.rwMutex.Unlock()

	el := q.elements[uid]
	q.list.Remove(el)

	delete(q.elements, uid)
	delete(q.index, uid)
}

// GetFirst returns first element from Queue.
func (q *Queue) GetFirst() interface{} {
	el := q.list.Front()
	if el == nil {
		return nil
	}
	return el.Value
}

// GetLast returns last element from Queue.
func (q *Queue) GetLast() interface{} {
	el := q.list.Back()
	if el == nil {
		return nil
	}
	return el.Value
}

// GetByID returns element from Queue by uid.
func (q *Queue) GetByID(uid string) interface{} {
	if !q.IsInQueue(uid) {
		return nil
	}

	q.rwMutex.RLock()
	defer q.rwMutex.RUnlock()

	el := q.elements[uid]
	if el == nil {
		return nil
	}
	return el.Value
}

// IsInQueue checks is element with given uid in Queue.
func (q *Queue) IsInQueue(uid string) bool {
	q.rwMutex.RLock()
	defer q.rwMutex.RUnlock()

	_, ok := q.index[uid]
	return ok
}
