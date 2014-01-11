package stack

import (
    "./../util"
)

// The implementation of ResizingArrayStack for <Algorithm>

// The initial size for a new stack
const default_size = 5

// Definition for ResizingArrayStack
type ResizingArrayStack struct {
    stack []interface{}
    n int
}

// Definition for ResizingArrayStack's iterator
type ResizingArrayStackIterator struct {
    index int
    st *ResizingArrayStack
}

// return the Stack interface 
func New() Stack {
    stack := new(ResizingArrayStack)
    stack.stack = make([]interface{}, default_size)
    stack.n = 0
    return stack
}

// Push a new item into stack
func (s *ResizingArrayStack) Push(item interface{}) {
    if s.n == len(s.stack) {
       s.resize(2 * len(s.stack))
    }
    s.stack[s.n] = item
    s.n++
}

// Pop the top item from the stack
func (s *ResizingArrayStack) Pop() interface{} {
    s.n--
    item := s.stack[s.n]
    s.stack[s.n] = nil
    if s.n > 0 && s.n == len(s.stack)/4 {
        s.resize(len(s.stack)/2)
    }
    return item
}

// Get the top item from the stack without poping it
func (s *ResizingArrayStack) Peek() interface{} {
    index := s.n - 1
    return s.stack[index]
}

// Get the current size of the stack
func (s *ResizingArrayStack) Size() int {
    return s.n
}

// if the stack is empty
func (s *ResizingArrayStack) IsEmpty() bool {
    return s.n == 0
}

// private method used to re-allocate stack capacity
func (s *ResizingArrayStack) resize(size int) {
    if size >= default_size {
        new_stack := make([]interface{}, size)
        copy(new_stack, s.stack)
        s.stack = new_stack
    }
}

// Get a Iterator for the stack
func (s *ResizingArrayStack) NewIterator() util.Iterator {
    iter := new(ResizingArrayStackIterator)
    iter.st = s
    iter.index = 0
    return iter
}

// If next item exists
func (iter *ResizingArrayStackIterator) HasNext() bool {
    return iter.index < iter.st.Size() 
}

// Get the next item
func (iter *ResizingArrayStackIterator) Next() interface{} {
    item := iter.st.stack[iter.index]
    iter.index ++ 
    return item
}

