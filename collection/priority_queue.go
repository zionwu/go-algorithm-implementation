package collection

import "./../algorithm/sort"

// The implementation for Priority Queue

// The definition for Priority Queue
type PriorityQueue struct {
    // the underlining array to store elements
    heap []sort.Comparable
    // the size of the heap
    size int
}

// method to get a initailized PriorityQueue given the queue capacity
func NewPriorityQueue(capacity int) *PriorityQueue {
    q := new(PriorityQueue)
    q.heap = make([]sort.Comparable, capacity+1)
    q.size = 0
    return q
}

// If the queue is empty or not
func (q *PriorityQueue) IsEmpty() bool {
    return q.size == 0
}

// Get the size of the queue
func (q *PriorityQueue) Size() int {
    return q.size
}

// Insert a new element into the queue
func (q *PriorityQueue) Insert(key sort.Comparable) {
    q.size++
    q.heap[q.size] = key
    // call swim to ajust queue structure
    q.swim(q.size)
}

// Delete the maximum element from the queue
func (q *PriorityQueue) DelMax() sort.Comparable {
    max := q.heap[1] 
    q.heap[1] = q.heap[q.size]
    // set it to nil to avoid memory leak
    q.heap[q.size] = nil
    q.size--
    // call sink to adjust queue structure
    q.sink(1)
    return max
}

// method used to adjust queue structure. The k th element in the array will go up the heap if it is bigger then its parents
func (q *PriorityQueue) swim(k int) {
    for k/2 >= 1 {
        if q.heap[k].CompareTo(q.heap[k/2]) > 0 {
            tmp := q.heap[k/2]
            q.heap[k/2] = q.heap[k]
            q.heap[k] = tmp
            k = k/2
        } else {
            break
        }
    }
}

// method used to adjust queue structure. The first element in the array will go down to the heap if it is smaller then its child
func (q *PriorityQueue) sink(k int) {
    for 2*k <= q.size {
        t := 2*k
        if t < q.size && q.heap[t].CompareTo(q.heap[t+1]) < 0 {
           t += 1
        }

        if q.heap[k].CompareTo(q.heap[t]) < 0 {
            tmp := q.heap[k]
            q.heap[k] = q.heap[t]
            q.heap[t] = tmp
            k = t
        } else {
            break
        }
    }
}

