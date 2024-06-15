package cache

type LRUCacheNode struct {
	Key   int
	Value int
	Prev  *LRUCacheNode
	Next  *LRUCacheNode
}

type LRUCache struct {
	keyToNode map[int]*LRUCacheNode
	head      *LRUCacheNode // point to the latest node
	tail      *LRUCacheNode // point to the oldest node
	capacity  int
}

func NewLRUCache(capacity int) *LRUCache {
	c := &LRUCache{
		keyToNode: make(map[int]*LRUCacheNode, capacity),
		head:      nil,
		tail:      nil,
		capacity:  capacity,
	}
	return c
}

func (c *LRUCache) Get(key int) int {
	node := c.keyToNode[key]
	if node == nil {
		return -1
	}
	c.bringToFront(node)
	return node.Value
}

func (c *LRUCache) Put(key int, value int) {
	node := c.keyToNode[key]
	if node != nil {
		node.Value = value
		c.bringToFront(node)
		return
	}

	node = &LRUCacheNode{
		Key:   key,
		Value: value,
	}
	c.keyToNode[key] = node

	if c.head != nil {
		currentHead := c.head
		c.head = node
		node.Next = currentHead
		currentHead.Prev = node
	} else {
		c.head = node
		c.tail = node
	}

	if len(c.keyToNode) > c.capacity {
		delete(c.keyToNode, c.tail.Key)
		c.tail.Prev.Next = nil
		c.tail = c.tail.Prev
	}
}

func (c *LRUCache) bringToFront(node *LRUCacheNode) {
	if c.head == node {
		return
	}

	left := node.Prev
	right := node.Next
	left.Next = right
	if right != nil {
		right.Prev = left
	} else if c.tail == node {
		c.tail = left
	}

	currentHead := c.head
	c.head = node
	node.Next = currentHead
	currentHead.Prev = node
	node.Prev = nil
}
