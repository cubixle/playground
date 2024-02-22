package queue_test

import (
	"testing"

	queue "github.com/cubixle/playground/sorted_queue"
)

func TestQueue(t *testing.T) {
	q := queue.NewQueue()
	q.Put(&queue.Item{ID: "aaa", Priority: 6})
	q.Put(&queue.Item{ID: "bbb", Priority: 7})
	q.Put(&queue.Item{ID: "ccc", Priority: 2})
	q.Put(&queue.Item{ID: "ddd", Priority: 1})

	v := q.GetNext()
	if v == nil {
		t.Fatal("didn't get a item")
	}
	if v.ID != "bbb" {
		t.Fatal("id wasn't what we expected")
	}
	if v.ID != "aaa" {
		t.Fatal("id wasn't what we expected")
	}
	if v.ID != "ccc" {
		t.Fatal("id wasn't what we expected")
	}
	if v.ID != "ddd" {
		t.Fatal("id wasn't what we expected")
	}
}
