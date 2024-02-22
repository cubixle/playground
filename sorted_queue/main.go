package queue

import "sync"

type Item struct {
	ID       string
	Data     string
	Priority int
}

type queue struct {
	m     map[string]struct{}
	queue []*Item
	mu    *sync.Mutex
}

func NewQueue() *queue {
	return &queue{
		m:     map[string]struct{}{},
		mu:    &sync.Mutex{},
		queue: []*Item{},
	}
}

// put on the queue with no duplicates
func (q *queue) Put(r *Item) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	_, ok := q.m[r.ID]
	if ok {
		return false
	}

	q.m[r.ID] = struct{}{}

	q.queue = append(q.queue, r)

	if len(q.queue) > 0 {
		// sort the queue
		for i := len(q.queue); i > 0; i-- {
			if q.queue[i].Priority > q.queue[i-1].Priority {
				q.queue[i] = q.queue[i-1]
				q.queue[i-1] = q.queue[i]
			}
		}
	}

	return true
}

func (q *queue) GetNext() *Item {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.m) == 0 {
		return nil
	}

	r := q.queue[0]

	// delete the id from the map
	delete(q.m, r.ID)

	// delete from the front of the list
	q.queue = q.queue[1:]

	return r
}
