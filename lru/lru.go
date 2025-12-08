package lru

import (
	"algorithm/double_linked_list"
)

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRU[K comparable, V any] struct {
	cache    map[K]*double_linked_list.DNode[entry[K, V]]
	list     *double_linked_list.DoubleLinkedList[entry[K, V]]
	capacity int
}

func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	return &LRU[K, V]{
		cache:    make(map[K]*double_linked_list.DNode[entry[K, V]]),
		list:     double_linked_list.NewDoubleLinkedList[entry[K, V]](),
		capacity: capacity,
	}
}

func (l *LRU[K, V]) Get(key K) (V, bool) {
	var zero V
	node, ok := l.cache[key]
	if !ok {
		return zero, false
	}

	err := l.list.MoveToTail(node)
	if err != nil {
		return zero, false
	}

	return node.Val.value, true
}

func (l *LRU[K, V]) Put(key K, value V) {
	node, ok := l.cache[key]
	if ok {
		node.Val = entry[K, V]{key, value}
		l.list.MoveToTail(node)
		return
	}

	newNode := l.list.Append(entry[K, V]{key, value})
	l.cache[key] = newNode

	if l.list.Size() > l.capacity {
		removed, err := l.list.RemoveHead()
		if err == nil {
			key := removed.key
			delete(l.cache, key)
		}
	}
}

func (l *LRU[K, V]) Remove(key K) bool {
	node, ok := l.cache[key]
	if !ok {
		return false
	}

	if err := l.list.Remove(node); err != nil {
		return false
	}

	delete(l.cache, key)
	return true
}

func (l *LRU[K, V]) Size() int {
	return l.list.Size()
}
