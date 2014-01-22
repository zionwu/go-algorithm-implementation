package sort

func ShellSort(items []Comparable) {
    h := 1
    for h < len(items){
        h =  3 * h + 1
    }

    for h >= 1 {
        for i := h; i < len(items); i++ {
            tmp := items[i]
            j := i
            for ; j > h-1; j = j - h {
                if tmp.CompareTo(items[j - h]) < 0 {
                    items[j] = items[j-h]
                } else {
                    items[j] = tmp
                    break
                }
            }

            if (j <= h-1) {
                items[j] = tmp
            }

        }
        h = h/3
    }
}
