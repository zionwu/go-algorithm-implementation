package algorithm

// This is the implementation for Weighted Quick Union Find algorithm
// described in <Algorithm> 1.5

type WeightedQuickUnionUF struct {
    id []int
    count int
}

func (w *WeightedQuickUnionUF) Count() int {
    return w.count
}

func (w *WeightedQuickUnionUF) Connected(p int, q int) bool {
    pid, _ := w.Find(p)
    qid, _ := w.Find(q)
    return pid == qid
}

func (w *WeightedQuickUnionUF) Find(p int) (int, int) {
    size := 0;
    for ;w.id[p] != p ;{
        p = w.id[p]
        size++;
    }
    return p, size
}

func (w *WeightedQuickUnionUF) Union(p int, q int) {
    pid, psize := w.Find(p)
    qid, qsize := w.Find(q)
    if pid != qid {
        if psize >= qsize {
            w.id[q] = p
            w.count--
        } else {
            w.id[p] = q
            w.count--
        }
    } 
}
