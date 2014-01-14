package collection 

import (
    "./../util"
)

// The implementation of LinkedList stack for <Algorithm> 1.2

// The definition for LinkedListStack
type LinkedListStack struct {
    head *node
    size int
}

// The definition for Iterator for LinkedListStack
type LinkedListStackIterator struct {
    current *node
}

// Return an instance of the LinkedListStack
func NewLinkedListStack() Stack {
    stack := new(LinkedListStack)
    stack.size = 0
    stack.head = nil
    return stack
}

// Push a new item into the stack
func (ls *LinkedListStack) Push(item interface{}) {
    new_node := new(node)
    new_node.item = item
    new_node.next = ls.head 
    ls.head = new_node
    ls.size++
}

// Pop a item our of the stack
func (ls *LinkedListStack) Pop() interface{} {
    item := ls.head.item
    ls.head = ls.head.next
    ls.size--
    return item;
}

// Get the value for the item at the top of the stack
func (ls *LinkedListStack) Peek() interface{} {
    return ls.head.item
}

// Return the size of the stack
func (ls *LinkedListStack) Size() int {
    return ls.size
}

// If the stack is empty or not
func (ls *LinkedListStack) IsEmpty() bool {
    return ls.head == nil 
}

// Get a iterator for the current stack
func (ls *LinkedListStack) NewIterator() util.Iterator {
    iter := new(LinkedListStackIterator)
    iter.current = ls.head
    return iter
}

// If there is a next item in the stack 
func (iter *LinkedListStackIterator) HasNext() bool {
    return iter.current.next == nil 
    
}

// Get the next item
func (iter *LinkedListStackIterator) Next() interface{} {
    item := iter.current.item
    iter.current = iter.current.next
    return item
}

