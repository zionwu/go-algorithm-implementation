package sort

func HeapSort(items []Comparable) {
    size := len(items)
    q := make([]Comparable, size + 1)
    // the heap should starts at index 1 of array
    for i := 1; i < size + 1; i++ {
        q[i] = items[i-1]
    }

    for i := size/2; i >=1; i-- {
       sink(q, i, size)
    }

    for i, _ := range items {
        items[i] = q[1]
        q[1] = q[size]
        size-- 
        sink(q, 1, size)
    }

}

// method used to adjust queue structure. The first element in the array will go down to the heap if it is smaller then its child
func sink(heap []Comparable, k int, size int) {
    for 2*k <= size {
        t := 2*k
        if t < size && heap[t].CompareTo(heap[t+1]) < 0 {
           t += 1
        }

        if heap[k].CompareTo(heap[t]) < 0 {
            tmp := heap[k]
            heap[k] = heap[t]
            heap[t] = tmp
            k = t
        } else {
            break
        }
    }
}
