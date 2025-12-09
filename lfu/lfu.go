package lfu

import (
	"algorithm/double_linked_list"
	"sync/atomic"
)

type LFU[K comparable, V any] struct {
	entries     map[K]V
	minFreq     atomic.Int32
	keyToFreq   map[K]int
	keyToNode   map[K]*double_linked_list.DNode[K]
	freqToDLink map[int]*double_linked_list.DoubleLinkedList[K]
	capacity    int
}

func NewLFU[K comparable, V any](capacity int) *LFU[K, V] {
	return &LFU[K, V]{
		entries:     make(map[K]V),
		keyToFreq:   make(map[K]int),
		keyToNode:   make(map[K]*double_linked_list.DNode[K]),
		freqToDLink: make(map[int]*double_linked_list.DoubleLinkedList[K]),
		capacity:    capacity,
	}
}

func (l *LFU[K, V]) Get(key K) (V, bool) {
	value, ok := l.entries[key]
	if !ok {
		return value, false
	}

	currentFreq := l.keyToFreq[key]
	dl := l.freqToDLink[currentFreq]
	node := l.keyToNode[key]
	dl.Remove(node)

	if dl.Size() == 0 {
		l.compareAndSwap(currentFreq, currentFreq+1)
	}

	freq := currentFreq + 1
	nowDL, ok := l.freqToDLink[freq]
	if !ok {
		nowDL = double_linked_list.NewDoubleLinkedList[K]()
		l.freqToDLink[freq] = nowDL
	}

	l.keyToFreq[key] = freq
	newNode := nowDL.Append(key)
	l.keyToNode[key] = newNode

	return value, true
}

func (l *LFU[K, V]) Put(key K, value V) error {
	if _, ok := l.entries[key]; ok {
		l.Get(key)
		l.entries[key] = value

		return nil
	}

	// 当容量达到上限时，需要淘汰元素
	if len(l.entries) >= l.capacity {
		min := l.getMinFreq()
		minDl, ok := l.freqToDLink[min]
		if !ok {
			// 如果最小频率对应的链表不存在，说明LFU可能有问题
			return nil
		}

		// 从最小频率链表中移除头部元素（最早访问的）
		deleteKey, err := minDl.RemoveHead()
		if err != nil {
			return err
		}

		// 如果移除后链表为空，删除该频率的映射
		if minDl.Size() == 0 {
			delete(l.freqToDLink, min)
		}

		// 从所有映射中删除该键
		delete(l.entries, deleteKey)
		delete(l.keyToNode, deleteKey)
		delete(l.keyToFreq, deleteKey)
	}

	// 创建或获取频率为1的链表
	dl, ok := l.freqToDLink[1]
	if !ok {
		dl = double_linked_list.NewDoubleLinkedList[K]()
		l.freqToDLink[1] = dl
	}

	// 将新键添加到频率为1的链表中
	node := dl.Append(key)

	// 更新所有映射
	l.keyToNode[key] = node
	l.keyToFreq[key] = 1
	l.entries[key] = value

	// 由于是新键，最小频率一定是1
	l.updateMinFreq(1)

	return nil
}

func (l *LFU[K, V]) getMinFreq() int {
	return int(l.minFreq.Load())
}

func (l *LFU[K, V]) updateMinFreq(freq int) {
	l.minFreq.Store(int32(freq))
}

func (l *LFU[K, V]) compareAndSwap(old, new int) {
	l.minFreq.CompareAndSwap(int32(old), int32(new))
}
