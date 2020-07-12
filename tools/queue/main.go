package queue

// nolint:gochecknoglobals
var sq *Queue

func initQueue() {
	if sq != nil {
		return
	}

	sq = sq.Init()
}

// NewQueue init empty Queue.
func NewQueue() (q *Queue) {
	q = new(Queue)
	return q.Init()
}

// Put adds value to queue.
func Put(uid string, data interface{}) {
	initQueue()
	sq.Put(uid, data)
}

// Delete removes value from queue.
func Delete(uid string) {
	initQueue()
	sq.Delete(uid)
}

// GetFirst returns first element from Queue.
func GetFirst() interface{} {
	initQueue()
	return sq.GetFirst()
}

// GetLast returns last element from Queue.
func GetLast() interface{} {
	initQueue()
	return sq.GetLast()
}

// GetByID returns element from Queue by uid.
func GetByID(uid string) interface{} {
	initQueue()
	return sq.GetByID(uid)
}

// IsInQueue checks is element with given uid in Queue.
func IsInQueue(uid string) bool {
	initQueue()
	return sq.IsInQueue(uid)
}
