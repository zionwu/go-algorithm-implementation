package sort

func selection_sort(items []Comparable) {
    var min Comparable
    for i, _ := range items {
        min = items[i]
        for j := i+1; j < len(items); j++ {
            if items[j].CompareTo(min) < 0 {
                min = items[j]
            }
        }
        items[i] = min
    }
}
