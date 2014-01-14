package collection 

// the defination for Stack

type Stack interface {
    // push an item into the stack
    Push(item interface{})
    // pop an item out of the stack
    Pop() interface{}
    // get the item at the top of the stack
    Peek() interface{}
    // the size of the stack
    Size() int
    // is the stack empty
    IsEmpty() bool
}
