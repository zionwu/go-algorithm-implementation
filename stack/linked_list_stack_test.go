package stack

import (
    "testing"
)

// test for the LinkedListStack 


// Test whether it implements the Stack inferface 
func TestNewLinkedListStack(t *testing.T) {
    s := NewLinkedListStack();
    if s.Size() != 0 {
        t.Fatalf("The initial size of the stack should be 0");
    }

    if !s.IsEmpty() {
        t.Fatalf("The new stack should be empty")
    }

    s.Push(1);
    s.Push(2);
    if s.Size() != 2 {
        t.Fatalf("The initial size of the stack is not 2")
    }

    if s.IsEmpty() {
        t.Fatalf("The new stack should not be empty")
    }

    if item := s.Pop(); item != 2 {
        t.Fatalf("the item is not 2")
    }

    s.Push(3);
    s.Push(4);
    s.Push(5);
    s.Push(6);
    s.Push(7);

    if s.Size() != 6 {
        t.Fatalf("The stack size should be 6")
    }
}


// Test for Iterator of the ResizingArrayStack
func TestLinkedListIterator(t *testing.T) {
    ns := NewLinkedListStack()
    s := ns.(*LinkedListStack)
    s.Push(1)
    s.Push(2)
    s.Push(3)
    iter := s.NewIterator()

    if iter.HasNext() {
        if iter.Next() != 1 {
            t.Fatalf("the first item is not 1")
        } 
    }

    if iter.HasNext() {
        if iter.Next() != 2 {
            t.Fatalf("the first item is not 2")
        } 
    }

    if iter.HasNext() {
        if iter.Next() != 3 {
            t.Fatalf("the first item is not 3")
        } 
    }

    if iter.HasNext() {
        t.Fatalf("HasNext should be false")
    }

}
