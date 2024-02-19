package lru

import "log"

type Cache struct {
	cap int
	m   map[string]*node
	l   *list
}

func NewCache(cap int) *Cache {
	return &Cache{
		cap: cap,
		m:   make(map[string]*node, 2),
		l:   newList(),
	}
}

func (c *Cache) Get(key string) string {
	v, ok := c.m[key]
	if !ok {
		return ""
	}

	if v == nil {
		return ""
	}

	return v.data
}

func (c *Cache) Set(key, val string) {
	node, ok := c.m[key]
	if !ok {
		node := c.l.pushFront(key, val)

		c.m[key] = node

		if len(c.m) > c.cap {
			n := c.l.back()
			c.l.remove(n)
			delete(c.m, n.key)
		}

		return
	}

	node.data = val
	c.l.moveToFront(node)
}

type node struct {
	key  string
	data string
	next *node
	prev *node
}

type list struct {
	head *node
}

func newList() *list {
	return &list{}
}

func (l *list) moveToFront(n *node) {
	if l.head == nil {
		n.prev = l.head
		l.head = n
		l.head.prev = n
	} else {
		n.next = l.head
		n.prev = l.head.prev
		l.head.prev = n
		l.head = n
	}
}

func (l *list) pushFront(key, data string) *node {
	n := &node{
		key:  key,
		data: data,
	}

	l.moveToFront(n)

	return n
}

func (l *list) remove(n *node) {
	if n == nil {
		return
	}

	n.prev.next = n.next
	if n.next != nil {
		n.next.prev = n.prev
	}

	n = nil
}

func (l *list) back() *node {
	return l.head.prev
}

func (l *list) debug() {
	log.Println("--------")
	start := l.head.data
	node := l.head
	for node != nil {
		log.Println(node)

		node = node.next
		if node != nil && node.data == start {
			break
		}
	}
	log.Println("--------")
}
