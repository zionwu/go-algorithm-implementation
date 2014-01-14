package collection 

import (
    "testing"
)

// test for the LinkedListStack 


// Test whether it implements the Queue inferface 
func TestNewLinkedListQueue(t *testing.T) {
    s := NewLinkedListQueue();
    if s.Size() != 0 {
        t.Fatalf("The initial size of the queu should be 0");
    }

    if !s.IsEmpty() {
        t.Fatalf("The new queue should be empty")
    }

    s.Enqueue(1);
    s.Enqueue(2);
    if s.Size() != 2 {
        t.Fatalf("The initial size of the stack is not 2")
    }

    if s.IsEmpty() {
        t.Fatalf("The new queue should not be empty")
    }

    if item := s.Dequeue(); item != 1 {
        t.Fatalf("the item is not 1")
    }

    s.Enqueue(3);
    s.Enqueue(4);
    s.Enqueue(5);
    s.Enqueue(6);

    if s.Size() != 5 {
        t.Fatalf("The stack size should be 5")
    }
}


// Test for Iterator of the LinkedListQueue 
func TestLinkedListQueueIterator(t *testing.T) {
    ns := NewLinkedListQueue()
    s := ns.(*LinkedListQueue)
    s.Enqueue(1);
    s.Enqueue(2);
    s.Enqueue(3);
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
