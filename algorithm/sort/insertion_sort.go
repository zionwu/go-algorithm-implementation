package sort

// the implementation for insertion sort
func InsertionSort(items []Comparable) {
    for i := 1; i < len(items); i++ {
        it := items[i]
        j := i
        for ; j > 0; j-- {
            if it.CompareTo(items[j-1]) < 0 {
                items[j] = items[j-1]
            } else {
                items[j] = it 
                break
            }
        }

        if j == 0 {
            items[j] = it
        }
    }
}
