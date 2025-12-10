package consistent_hash

import (
	"fmt"
	"hash/crc32"
	"slices"
	"sort"
	"sync"
)

type ConsistentHash struct {
	sync.RWMutex
	sortedRing     []uint32
	ring           map[uint32]string
	hashFunc       func([]byte) uint32
	virtualNodeNum int
	nodes          map[string]struct{}
}

func NewConsistentHash(virtualNodes int) *ConsistentHash {
	return &ConsistentHash{
		hashFunc:       crc32.ChecksumIEEE,
		virtualNodeNum: virtualNodes,
		ring:           make(map[uint32]string),
		sortedRing:     make([]uint32, 0),
		nodes:          make(map[string]struct{}),
	}
}

func (c *ConsistentHash) AddNode(node string) {
	c.Lock()
	defer c.Unlock()

	_, ok := c.nodes[node]
	if ok {
		return
	}

	c.nodes[node] = struct{}{}

	for i := range c.virtualNodeNum {
		virtualKey := fmt.Sprintf("%s#%d", node, i)
		hash := c.hashFunc([]byte(virtualKey))

		for {
			_, ok := c.ring[hash]
			if !ok {
				break
			}
			virtualKey := fmt.Sprintf("%s#%d-%d", node, i, hash)
			hash = c.hashFunc([]byte(virtualKey))
		}

		c.ring[hash] = node
		c.sortedRing = append(c.sortedRing, hash)
	}

	slices.Sort(c.sortedRing)
}

func (c *ConsistentHash) RemoveNode(node string) {
	c.Lock()
	defer c.Unlock()

	if _, exists := c.nodes[node]; !exists {
		return
	}

	delete(c.nodes, node)

	newRing := make(map[uint32]string)
	newSortedRing := make([]uint32, 0, len(c.ring))

	for hash, nodeName := range c.ring {
		if nodeName != node {
			newRing[hash] = nodeName
			newSortedRing = append(newSortedRing, hash)
		}
	}

	c.ring = newRing
	c.sortedRing = newSortedRing

	slices.Sort(c.sortedRing)
}

func (c *ConsistentHash) GetNode(key string) (string, bool) {
	c.RLock()
	defer c.RUnlock()

	if len(c.nodes) == 0 {
		return "", false
	}

	hash := c.hashFunc([]byte(key))

	idx := sort.Search(len(c.sortedRing), func(i int) bool {
		return c.sortedRing[i] >= hash
	})

	if idx == len(c.sortedRing) {
		idx = 0
	}

	return c.ring[c.sortedRing[idx]], true
}

func (c *ConsistentHash) GetNodes() []string {
	c.RLock()
	defer c.RUnlock()

	nodes := make([]string, 0, len(c.nodes))
	for node := range c.nodes {
		nodes = append(nodes, node)
	}

	sort.Strings(nodes)
	return nodes
}
