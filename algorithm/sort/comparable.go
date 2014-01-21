package sort


// The interface Comparable.
// Element that want to use sort related algirhtm must implement 
// the method CompareTo, which will be called in sort algorithm
type Comparable interface {
    // return positive number, 0 or negative number if the acceiving object is larger than, equals to or less than the specified object 
    CompareTo(e interface{}) int
}
