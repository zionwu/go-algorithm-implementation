package collection

import (
    "./../util"
)


// The implementation for linkedlist Queue

// the definition of the linked list queue
type LinkedListQueue struct {
    head *node
    tail *node
    size int
}

// the definition of the iterator for linked list queeu
type LinkedListQueueIterator struct {
    current *node
}

// if the queue is empty
func (q *LinkedListQueue) IsEmpty() bool {
    return q.head == nil
}

// Return the size of the queue
func (q *LinkedListQueue) Size() int {
    return q.size
}

// Get a new instance  of the LinkedList Queue
func NewLinkedListQueue() Queue {
    q := new(LinkedListQueue)
    q.size = 0
    q.head = nil
    q.tail = nil
    return q
}

// Add an item to the queue
func (q *LinkedListQueue) Enqueue(item interface{}) {
    old_tail := q.tail
    new_node := new(node)
    new_node.item = item
    new_node.next = nil
    q.tail = new_node
    if q.IsEmpty() {
        q.head = new_node
    } else {
        old_tail.next = new_node
    }

    q.size++
}

// Get an item out of the queue
func (q *LinkedListQueue) Dequeue() interface{} {
    item := q.head.item
    q.head = q.head.next
    if q.IsEmpty() {
       q.tail = nil
    }
    q.size--;
    return item
}

// Get an iterator for the queue
func (q *LinkedListQueue) NewIterator() util.Iterator {
    iter := new(LinkedListQueueIterator)
    iter.current = q.head
    return iter
}

// if the next item exists
func (iter *LinkedListQueueIterator) HasNext() bool {
    if iter.current == nil {
       return false
    } else {
       return iter.current.next == nil 
    }
}

// Get the next item
func (iter *LinkedListQueueIterator) Next() interface{} {
    item := iter.current.item
    iter.current = iter.current.next
    return item
}



