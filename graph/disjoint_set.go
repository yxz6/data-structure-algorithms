package graph

type DisjointSet struct {
	parents []int // given node x, parents[x] is the parent node
	ranks   []int // used for union merge
}

func NewDisjointSet(n int) *DisjointSet {
	ds := &DisjointSet{
		parents: make([]int, n),
		ranks:   make([]int, n),
	}
	for i := range n {
		ds.parents[i] = i // each node is a set initially
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

	if ds.ranks[x] < ds.ranks[y] {
		x, y = y, x // ensure x.rank >= y.rank
	}
	ds.parents[y] = x
	if ds.ranks[x] == ds.ranks[y] {
		ds.ranks[x] += 1
	}
}
