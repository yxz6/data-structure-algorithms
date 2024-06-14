package leetcode

//https://leetcode.com/problems/longest-consecutive-sequence/description/

// nums allow duplicate numbers

func longestConsecutive(nums []int) int {
	ds := NewDisjointSet(len(nums))
	numToIndex := make(map[int]int, len(nums))
	for i, num := range nums {
		numToIndex[num] = i
	}

	for num, i := range numToIndex {
		if left, ok := numToIndex[num-1]; ok {
			ds.Union(i, left)
		}

		if right, ok := numToIndex[num+1]; ok {
			ds.Union(i, right)
		}
	}

	res := 0
	for _, v := range ds.sizes {
		if v > res {
			res = v
		}
	}
	return res
}

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
