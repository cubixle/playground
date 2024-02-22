package queue

import (
	"log"
	"sync"
)

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

	qu := q.queue
	i := 1
	for i < len(qu) {
		j := i
		for j >= 1 && qu[j].Priority > qu[j-1].Priority {
			tmp := qu[j-1]
			qu[j-1] = qu[j]
			qu[j] = tmp

			j--
		}
		i++
	}

	q.queue = qu

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

func (q *queue) Display() {
	for _, i := range q.queue {
		log.Println(i)
	}
}
