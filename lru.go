package main

type LRUCache[K comparable, V any] struct {
	capacity int
	store    map[K]*Node[K, V]
	list     *DoublyLinkedList[K, V]
}

func New[K comparable, V any](capacity int) *LRUCache[K, V] {
	list := newDoublyLinkedList[K, V]()
	return &LRUCache[K, V]{capacity: capacity, store: make(map[K]*Node[K, V]), list: list}
}

func (c *LRUCache[K, V]) Size() int {
	return len(c.store)
}

func (c *LRUCache[K, V]) Get(key K) V {
	node, ok := c.store[key]
	if !ok {
		var noop V
		return noop
	}
	c.list.remove(node)
	c.list.add(node)
	return node.value
}

func (c *LRUCache[K, V]) Put(key K, value V) {
	_, ok := c.store[key]
	if ok {
		delete(c.store, key)
	}

	if c.Size() == c.capacity {
		delete(c.store, c.list.tail.prev.key)
		c.list.remove(c.list.tail.prev)
	}

	n := new(Node[K, V])
	n.key = key
	n.value = value
	c.list.add(n)
	c.store[key] = n
}
