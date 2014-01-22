package sort

//Implementation for Merge Sort


// Merge Sort using Bottom-Up approach
func MergeSortBU(items []Comparable) {
    t := make([]Comparable, len(items))

    for size:=1; size < len(items); size *= 2 {
        for i := 0; i + size < len(items); i += 2 * size {
            var high int 
            if i+2*size > len(items) {
                high = len(items) -1
            } else {
                high = i+2*size-1
            }
            merge(items, t, i, i+size-1, high)
        }
    }
}

// Merge Sort using Top-Down approach
func MergeSortTD(items []Comparable) {
    tmp := make([]Comparable, len(items))
    merge_sort_topdown(items, tmp, 0, len(items)-1)
}

// private method, which is run recursively
func merge_sort_topdown(a []Comparable, t[]Comparable,
                        low int, high int) {
    if high >= low {
        return;
    }
    mid := low + (high - low)/2
    merge_sort_topdown(a, t, low, mid)
    merge_sort_topdown(a, t, mid+1, high)
    merge(a, t, low, mid, high)
}

// private method performing merge operation for two array
func merge(a []Comparable, t[]Comparable, low int, mid int, high int) {
    i := low
    j := mid + 1
    k := low
    for ; k <= high; k++ {
        t[k] = a[k]
    }

    for k <= high {
        if a[i].CompareTo(a[j]) <= 0 {
            a[k] = t[i]
            k++
            i++
        } else if a[i].CompareTo(a[j]) > 0 {
            a[k] = t[j]
            k++
            j++
        } else if i > mid {
            a[k] = t[j]
            k++
            j++
        } else if j > high {
            a[k] = t[i]
            k++
            i++
        }
    }
}
