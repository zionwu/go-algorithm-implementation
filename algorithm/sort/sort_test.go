package sort

import (
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

// Test selection sort
func TestSelectionSort(t *testing.T) {
    items := make([]Comparable, 10)
    var item *Item
    for i := 0; i < 10; i++ {
        item = new(Item)
        item.val = rand.Int()
        items[i] = item;
    }

    SelectionSort(items)

    for i := 0; i < 9; i++ {
        if items[i].(*Item).val > items[i+1].(*Item).val {
            t.Fatalf("The item %v should be less then item %v", items[i].(*Item).val, items[i+1].(*Item).val)
        }
    }
}

// Test insertion sort
func TestInsertionSort(t *testing.T) {
    items := make([]Comparable, 10)
    var item *Item
    for i := 0; i < 10; i++ {
        item = new(Item)
        item.val = rand.Int()
        items[i] = item;
    }

    InsertionSort(items)

    for i := 0; i < 9; i++ {
        if items[i].(*Item).val > items[i+1].(*Item).val {
            t.Fatalf("The item %v should be less then item %v", items[i].(*Item).val, items[i+1].(*Item).val)
        }
    }
}

// Test shell sort
func TestShellSort(t *testing.T) {
    items := make([]Comparable, 10)
    var item *Item
    for i := 0; i < 10; i++ {
        item = new(Item)
        item.val = rand.Int()
        items[i] = item;
    }

    ShellSort(items)

    for i := 0; i < 9; i++ {
        if items[i].(*Item).val > items[i+1].(*Item).val {
            t.Fatalf("The item %v should be less then item %v", items[i].(*Item).val, items[i+1].(*Item).val)
        }
    }
}

// Test top-down merge sort
func TestMergeSortTD(t *testing.T) {
    items := make([]Comparable, 10)
    var item *Item
    for i := 0; i < 10; i++ {
        item = new(Item)
        item.val = rand.Int()
        items[i] = item;
    }

    MergeSortTD(items)

    for i := 0; i < 9; i++ {
        if items[i].(*Item).val > items[i+1].(*Item).val {
            t.Fatalf("The item %v should be less then item %v", items[i].(*Item).val, items[i+1].(*Item).val)
        }
    }
}

// Test bottom-up merge sort
func TestMergeSortBU(t *testing.T) {
    items := make([]Comparable, 10)
    var item *Item
    for i := 0; i < 10; i++ {
        item = new(Item)
        item.val = rand.Int()
        items[i] = item;
    }

    MergeSortBU(items)

    for i := 0; i < 9; i++ {
        if items[i].(*Item).val > items[i+1].(*Item).val {
            t.Fatalf("The item %v should be less then item %v", items[i].(*Item).val, items[i+1].(*Item).val)
        }
    }
}

// Test quick sort
func TestQuickSort(t *testing.T) {
    items := make([]Comparable, 10)
    var item *Item
    for i := 0; i < 10; i++ {
        item = new(Item)
        item.val = rand.Int()
        items[i] = item;
    }

    QuickSort(items)

    for i := 0; i < 9; i++ {
        if items[i].(*Item).val > items[i+1].(*Item).val {
            t.Fatalf("The item %v should be less then item %v", items[i].(*Item).val, items[i+1].(*Item).val)
        }
    }
}

// Test heap sort
func TestHeapSort(t *testing.T) {
    items := make([]Comparable, 10)
    var item *Item
    for i := 0; i < 10; i++ {
        item = new(Item)
        item.val = rand.Int()
        items[i] = item;
    }

    HeapSort(items)

    for i := 0; i < 9; i++ {
        if items[i].(*Item).val < items[i+1].(*Item).val {
            t.Fatalf("The item %v should be less then item %v", items[i].(*Item).val, items[i+1].(*Item).val)
        }
    }
}
