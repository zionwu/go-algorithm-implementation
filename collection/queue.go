package collection


// The interface for Queue
type Queue interface {
    // add an item to the end of the queue
    Enqueue(item interface{})
    // remove an item from the beginning of the queue
    Dequeue() interface{}
    // if the queue is empty
    IsEmpty() bool
    // return the size of the queue
    Size() int
}
