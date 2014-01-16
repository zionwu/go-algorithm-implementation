package algorithm

// This is the implementation for Weighted Quick Union Find algorithm
// described in <Algorithm> 1.5

// The struct used for the algorithm
type WeightedQuickUnionUF struct {
    //the index of the slice is the node, while the value is which node it is connected to
    id []int
    //the index is the node, the value is the depth of the graph the node belongs to 
    depth []int
    // The number of isolated graph
    count int
}

// Return a new instance of WeightedQuickUnionUF
func NewWeightedQuickUnionUF(size int) *WeightedQuickUnionUF {
    w := new (WeightedQuickUnionUF)
    w.count = size
    w.id = make([]int, size)
    w.depth = make([]int, size)
    for i := 0; i < size ; i++ {
        w.id[i] = i
        w.depth[i] = 0
    }
    return w
}

// Return the number of isolated graph
func (w *WeightedQuickUnionUF) Count() int {
    return w.count
}

// If the given two node p and q are connected
func (w *WeightedQuickUnionUF) Connected(p int, q int) bool {
    pid := w.Find(p)
    qid := w.Find(q)
    return pid == qid
}

// Return the root node of the graph that the given node p belongs to
func (w *WeightedQuickUnionUF) Find(p int) int {
    for ;w.id[p] != p ;{
        p = w.id[p]
    }
    return p
}

// Connect two graph that the node p and q belong to into one 
func (w *WeightedQuickUnionUF) Union(p int, q int) {
    pid := w.Find(p)
    qid := w.Find(q)
    if pid != qid {
        if w.depth[pid] >= w.depth[qid] {
            w.id[qid] = pid
            if w.depth[pid] == w.depth[qid] {
               w.depth[pid]++
            }
            w.count--
        } else {
            w.id[pid] = qid
            w.count--
        }
    }
}
