package graph

// DisjointSet a disjoint-set data structure, also called a union–find data structure or merge–find set
type DisjointSet struct {
	parents []int // given node x, parents[x] is the parent node
	sizes   []int // used for union merge
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parents: make([]int, n),
		sizes:   make([]int, n),
	}
	for i := range n {
		ds.parents[i] = i // each node is a set initially
		ds.sizes[i] = 1
	}
	return ds
}

func (ds *DisjointSet) Find(x int) int {
	for ds.parents[x] != x { //root node's parent is itself
		// tree is compressed during find
		x, ds.parents[x] = ds.parents[x], ds.parents[ds.parents[x]]
	}
	return x
}

func (ds *DisjointSet) Union(x, y int) {
	x, y = ds.Find(x), ds.Find(y)
	if x == y {
		return
	}

	if ds.sizes[x] < ds.sizes[y] {
		x, y = y, x // ensure x.size >= y.size
	}
	ds.parents[y] = x
	ds.sizes[x] += ds.sizes[y]
}
