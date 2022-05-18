package main

type Node[K comparable, V any] struct {
	key   K
	value V
	next  *Node[K, V]
	prev  *Node[K, V]
}

type DoublyLinkedList[K comparable, V any] struct {
	head *Node[K, V]
	tail *Node[K, V]
}

func newDoublyLinkedList[K comparable, V any]() *DoublyLinkedList[K, V] {
	list := &DoublyLinkedList[K, V]{head: new(Node[K, V]), tail: new(Node[K, V])}
	list.head.next = list.tail
	list.tail.prev = list.head
	return list
}

func (l *DoublyLinkedList[K, V]) add(n *Node[K, V]) {
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
	n.prev = l.head
}

func (l *DoublyLinkedList[K, V]) remove(n *Node[K, V]) {
	n.next.prev = n.prev
	n.prev.next = n.next
}
