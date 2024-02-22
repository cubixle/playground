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
		t.Fatalf("id wasn't what we expected bbb, got %s", v.ID)
	}
	v = q.GetNext()
	if v.ID != "aaa" {
		t.Fatalf("id wasn't what we expected aaa, got %s", v.ID)
	}
	v = q.GetNext()
	if v.ID != "ccc" {
		t.Fatalf("id wasn't what we expected ccc, got %s", v.ID)
	}
	v = q.GetNext()
	if v.ID != "ddd" {
		t.Fatalf("id wasn't what we expected ddd, got %s", v.ID)
	}
}
