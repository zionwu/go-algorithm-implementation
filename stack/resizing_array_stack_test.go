package stack

import (
    "testing"
)

// test for the ResizingArrayStack


// Test whether it implements the Stack inferface 
func TestNewStack(t *testing.T) {
    s := New();
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

// Test for the interanl implementation logic
func TestStackImplementation(t *testing.T) {
    ns := New();
    s := ns.(*ResizingArrayStack)
    if s.Size() != 0 {
        t.Fatalf("The initial size of the stack should be 0");
    }

    if !s.IsEmpty() {
        t.Fatalf("The new stack should be empty")
    }

    if len(s.stack) != 5 {
        t.Fatalf("The underlining slice size should be 5")
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

    if len(s.stack) != 10 {
        t.Fatalf("The underlining slice size should be 10 : %d", len(s.stack) )
    }

    s.Pop()
    s.Pop()
    s.Pop()
    s.Pop()
    s.Pop()

    if len(s.stack) != 5 {
        t.Fatalf("The underlining slice size should be 5")
    }

    if s.Peek() != 1 {
        t.Fatalf("the top element is not 1")
    }
}

// Test for Iterator of the ResizingArrayStack
func TestIterator(t *testing.T) {
    ns := New()
    s := ns.(*ResizingArrayStack)
    iter := s.NewIterator()
    s.Push(1)
    s.Push(2)
    s.Push(3)

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
