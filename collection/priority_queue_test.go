package collection

import (
    "./../algorithm/sort"
    "testing"
    "math/rand"
)

// Define a strcuture that a wrapper of int which implement Comparable interface so the it can be passed to sort algorithm
type Item struct {
    val int
}

// Implement the CompareTo method for Item
func (i *Item) CompareTo(e interface{}) int {
    elem := e.(*Item)
    if i.val > elem.val {
        return 1
    } else if i.val < elem.val {
        return -1
    } else {
        return 0
    }
}

func TestPriorityQueue(t *testing.T) {
    items := make([]sort.Comparable, 10)
    var item *Item
    for i := 0; i < 10; i++ {
        rand.Seed(42)
        item = new(Item)
        item.val = rand.Int()
        items[i] = item;
    }

    q := NewPriorityQueue(10)
    for _, v := range items {
        q.Insert(v)
    }

    for i, _ := range items {
        items[i] = q.DelMax()
    }

    for i := 0; i < 9; i++ {
        if items[i].(*Item).val > items[i+1].(*Item).val {
            t.Fatalf("The item %v should be less then item %v", items[i].(*Item).val, items[i+1].(*Item).val)
        }
    }


}
