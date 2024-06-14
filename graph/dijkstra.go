package graph

/*
Let us choose a starting node, and let the distance of node N be the distance from the starting node to N. Dijkstra's algorithm will initially start with infinite distances and will try to improve them step by step.

    Mark all nodes as unvisited. Create a set of all the unvisited nodes called the unvisited set.
    Assign to every node a distance from start value: for the starting node, it is zero, and for all other nodes, it is infinity, since initially no path is known to these nodes. During execution of the algorithm, the distance of a node N is the length of the shortest path discovered so far between the starting node and N.[17]
    From the unvisited set, select the current node to be the one with the smallest finite distance; initially, this will be the starting node, which has distance zero. If the unvisited set is empty, or contains only nodes with infinite distance (which are unreachable), then the algorithm terminates by going to step 6. If we are only concerned about the path to a target node, we may terminate here if the current node is the target node. Otherwise, we can continue to find the shortest paths to all reachable nodes.
    For the current node, consider all of its unvisited neighbors and update their distances through the current node; compare the newly calculated distance to the one currently assigned to the neighbor and assign it the smaller one. For example, if the current node A is marked with a distance of 6, and the edge connecting it with its neighbor B has length 2, then the distance to B through A is 6 + 2 = 8. If B was previously marked with a distance greater than 8, then update it to 8 (the path to B through A is shorter). Otherwise, keep its current distance (the path to B through A is not the shortest).
    When we are done considering all of the unvisited neighbors of the current node, mark the current node as visited and remove it from the unvisited set. This is so that a visited node is never checked again, which is correct because the distance recorded on the current node is minimal (as ensured in step 3), and thus final. Go back to step 3.
    Once the loop exits (steps 3–5), every visited node will contain its shortest distance from the starting node.
*/

//type DijkstraNode struct {
//	Index    int
//	Distance int
//}
//
//type DijkstraHeap struct {
//	nodes []*DijkstraNode
//}
//
//func (h *DijkstraHeap) Len() int {
//	return len(h.nodes)
//}
//
//func (h *DijkstraHeap) Less(i, j int) bool {
//	if h.nodes[i].Distance < 0 {
//		return false
//	}
//
//	if h.nodes[j].Distance < 0 {
//		return true
//	}
//
//	return h.nodes[i].Distance < h.nodes[j].Distance
//}
//
//func (h *DijkstraHeap) Swap(i, j int) {
//	h.nodes[i], h.nodes[j] = h.nodes[j], h.nodes[i]
//}
//
//func (h *DijkstraHeap) Push(x any) {
//	h.nodes = append(h.nodes, x.(*DijkstraNode))
//}
//
//func (h *DijkstraHeap) Pop() any {
//	v := h.nodes[len(h.nodes)-1]
//	h.nodes = h.nodes[:len(h.nodes)-1]
//	return v
//}
//
//var _ heap.Interface = (*DijkstraHeap)(nil)
//
//func Dijkstra(g [][]int, source int) []int {
//	n := len(g)
//	if len(g[0]) != n {
//		panic("g is not a valid graph")
//	}
//
//	distances := make([]int, n)
//	unvisited := new(DijkstraHeap)
//	heap.Push(unvisited, DijkstraNode{Index: 0, Distance: 0})
//	for i := 1; i < n; i++ {
//		heap.Push(unvisited, DijkstraNode{Index: i, Distance: -1})
//	}
//
//	for unvisited.Len() > 0 {
//		current := unvisited.Pop().(DijkstraNode)
//		for i, distance := range g[current.Index] {
//			if distance == 0 {
//				continue
//			}
//			for j, node := range unvisited {
//				if node
//			}
//		}
//	}
//}
