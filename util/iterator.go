package util

// This interface Iterator is used to iterate a collection

type Iterator interface {
     //if the next element exist
     HasNext() bool
     //get the next element
     Next() interface{}
} 
