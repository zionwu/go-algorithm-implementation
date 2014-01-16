package algorithm

import (
    "testing"
)

func TestWeightedQuickUnionUF(t *testing.T) {
    w := NewWeightedQuickUnionUF(15)
    if w.Count() != 15 {
        t.Fatalf("The inital number of graph should be 15")
    }

    if w.Connected(1, 2) {
        t.Fatalf("The node 1 and 2 should be disconnected")
    }

    w.Union(1, 2)
    if !w.Connected(1, 2) {
        t.Fatalf("The node 1 and 2 should be connected")
    }

    w.Union(3, 4)
    w.Union(4, 5)
    w.Union(5, 6)

    w.Union(1, 6)
    if !w.Connected(3, 2) {
        t.Fatalf("The node 3 and 2 should be disconnected")
    }

    p := w.Find(6) 
    if p != 1 {
        t.Fatalf("The root should be 1 instead of : %v", p)
    }

    if w.depth[1] != 2 {
        t.Fatalf("The depth of node 6 should be 2 instead of %v", w.depth[1])
    }

    if w.Count() != 10 {
        t.Fatalf("The number of graph should be 10 instead of %v", w.Count())
    }


}
