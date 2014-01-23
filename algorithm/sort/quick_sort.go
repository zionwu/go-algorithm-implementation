package sort

import "math/rand"

func QuickSort(items []Comparable) {
    shuffle(items)
    quick_sort(items, 0, len(items)-1)
}

func quick_sort(items []Comparable, low int, high int) {
    if low >= high {
        return
    }

    p := partition(items, low, high)
    quick_sort(items, low, p-1)
    quick_sort(items, p+1, high)
}

func partition(items []Comparable, low int, high int) int {
    i := low
    j := high + 1
    pivot := items[low]
    var tmp Comparable

    for {
        for i++; pivot.CompareTo(items[i]) >= 0; i++  {
            if i == high {
                break
            }
        }
        for j--; pivot.CompareTo(items[j]) <= 0; j--  {
            if j == low {
                break
            }
        }

        if i >= j {
            break
        }

        tmp = items[i]
        items[i] = items[j]
        items[j] = tmp
    }

    items[low] = items[j]
    items[j] = pivot
    return j
}


// shuffle the array randomly. 
// TO-DO: move this as util method in the future 
//        so that the code can be re-used
func shuffle(items []Comparable) {
    dest := make([]Comparable, len(items))
    for i, v := range items {
        dest[i] = v
    }

    perm := rand.Perm(len(items))
    for i, v := range perm {
       items[v] = dest[i] 
    }
}
