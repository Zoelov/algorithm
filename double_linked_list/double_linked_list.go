// Package double_linked_list 实现了一个泛型双向链表
//
// 该双向链表使用虚拟头尾节点（dummy head/tail）来简化边界条件处理，
// 支持在O(1)时间内进行头/尾节点的插入和删除操作，
// 以及在指定节点前后进行插入和删除操作。
package double_linked_list

import "errors"

// DNode 双向链表节点
type DNode[T any] struct {
	Val  T         // 节点存储的值
	Pre  *DNode[T] // 指向前一个节点的指针
	Next *DNode[T] // 指向后一个节点的指针
}

// DoubleLinkedList 双向链表结构体
type DoubleLinkedList[T any] struct {
	dummpyHead *DNode[T] // 虚拟头节点，不存储实际数据
	dummpyTail *DNode[T] // 虚拟尾节点，不存储实际数据
	size       int       // 链表中实际元素的数量
}

// NewDoubleLinkedList 创建一个新的空双向链表
//
// 返回值：
//
//	*DoubleLinkedList[T]: 指向新创建的双向链表的指针
func NewDoubleLinkedList[T any]() *DoubleLinkedList[T] {
	// 创建虚拟头节点和虚拟尾节点
	head := &DNode[T]{}
	tail := &DNode[T]{}

	// 初始化虚拟节点的指针
	head.Next = tail
	tail.Pre = head

	// 返回新创建的链表
	return &DoubleLinkedList[T]{
		dummpyHead: head,
		dummpyTail: tail,
	}
}

// Prepend 在链表头部添加一个新元素
//
// 参数：
//
//	val: 要添加的元素值
func (d *DoubleLinkedList[T]) Prepend(val T) {
	// 获取当前的第一个实际节点
	first := d.dummpyHead.Next

	// 创建新节点，指向虚拟头节点和第一个实际节点
	node := &DNode[T]{
		Val:  val,
		Next: first,
		Pre:  d.dummpyHead,
	}

	// 更新虚拟头节点和第一个实际节点的指针
	d.dummpyHead.Next = node
	first.Pre = node

	// 链表大小加1
	d.size++
}

// Append 在链表尾部添加一个新元素
//
// 参数：
//
//	val: 要添加的元素值
func (d *DoubleLinkedList[T]) Append(val T) *DNode[T] {
	// 获取当前的最后一个实际节点
	tail := d.dummpyTail.Pre

	// 创建新节点，指向最后一个实际节点和虚拟尾节点
	node := &DNode[T]{
		Val:  val,
		Next: d.dummpyTail,
		Pre:  tail,
	}

	// 更新最后一个实际节点和虚拟尾节点的指针
	tail.Next = node
	d.dummpyTail.Pre = node

	// 链表大小加1
	d.size++

	return node
}

// Size 返回链表中实际元素的数量
//
// 返回值：
//
//	int: 链表的大小
func (d *DoubleLinkedList[T]) Size() int {
	return d.size
}

// InsertAfter 在指定节点之后插入一个新元素
//
// 参数：
//
//	node: 要在其后插入新元素的节点
//	val: 要添加的元素值
//
// 返回值：
//
//	error: 如果node为nil或node是虚拟尾节点，则返回错误
func (d *DoubleLinkedList[T]) InsertAfter(node *DNode[T], val T) error {
	// 检查node是否有效
	if node == nil || node == d.dummpyTail {
		return errors.New("node is invalid")
	}

	// 获取node的下一个节点
	next := node.Next

	// 创建新节点
	newNode := &DNode[T]{
		Val:  val,
		Next: next,
		Pre:  node,
	}

	// 更新指针
	node.Next = newNode
	next.Pre = newNode

	// 链表大小加1
	d.size++
	return nil
}

// InsertBefore 在指定节点之前插入一个新元素
//
// 参数：
//
//	node: 要在其前插入新元素的节点
//	val: 要添加的元素值
//
// 返回值：
//
//	error: 如果node为nil或node是虚拟头节点，则返回错误
func (d *DoubleLinkedList[T]) InsertBefore(node *DNode[T], val T) error {
	// 检查node是否有效
	if node == nil || node == d.dummpyHead {
		return errors.New("node is invalid")
	}

	// 获取node的前一个节点
	pre := node.Pre

	// 创建新节点
	newNode := &DNode[T]{
		Val:  val,
		Next: node,
		Pre:  pre,
	}

	// 更新指针
	pre.Next = newNode
	node.Pre = newNode

	// 链表大小加1
	d.size++

	return nil
}

// RemoveHead 删除链表的第一个实际元素
//
// 返回值：
//
//	T: 被删除的元素值
//	error: 如果链表为空，则返回错误
func (d *DoubleLinkedList[T]) RemoveHead() (T, error) {
	var zero T

	// 检查链表是否为空
	if d.size == 0 {
		return zero, errors.New("list is empty")
	}

	// 获取当前的头节点
	head := d.dummpyHead.Next

	// 更新虚拟头节点和头节点的下一个节点的指针
	d.dummpyHead.Next = head.Next
	head.Next.Pre = d.dummpyHead

	// 断开被删除节点的指针，防止内存泄漏
	head.Pre = nil
	head.Next = nil

	// 链表大小减1
	d.size--

	return head.Val, nil
}

// RemoveTail 删除链表的最后一个实际元素
//
// 返回值：
//
//	T: 被删除的元素值
//	error: 如果链表为空，则返回错误
func (d *DoubleLinkedList[T]) RemoveTail() (T, error) {
	var zero T

	// 检查链表是否为空
	if d.size == 0 {
		return zero, errors.New("list is empty")
	}

	// 获取当前的尾节点
	tail := d.dummpyTail.Pre

	// 更新虚拟尾节点和尾节点的前一个节点的指针
	d.dummpyTail.Pre = tail.Pre
	tail.Pre.Next = d.dummpyTail

	// 断开被删除节点的指针，防止内存泄漏
	tail.Next = nil
	tail.Pre = nil

	// 链表大小减1
	d.size--

	return tail.Val, nil
}

// Remove 删除指定的节点
//
// 参数：
//
//	node: 要删除的节点
//
// 返回值：
//
//	error: 如果node为nil或node是虚拟头/尾节点，则返回错误
func (d *DoubleLinkedList[T]) Remove(node *DNode[T]) error {
	// 检查node是否有效
	if node == nil || node == d.dummpyHead || node == d.dummpyTail {
		return errors.New("node is invalid")
	}

	// 获取node的前后节点
	pre := node.Pre
	next := node.Next

	// 更新前后节点的指针，绕过要删除的节点
	pre.Next = next
	next.Pre = pre

	// 断开被删除节点的指针，防止内存泄漏
	node.Pre = nil
	node.Next = nil

	// 链表大小减1
	d.size--

	return nil
}

// MoveToTail 将指定节点移动到链表尾部
//
// 参数：
//
//	node: 要移动的节点
//
// 返回值：
//
//	error: 如果node为nil或node是虚拟头/尾节点，则返回错误
func (d *DoubleLinkedList[T]) MoveToTail(node *DNode[T]) error {
	// 检查node是否有效
	if node == nil || node == d.dummpyHead || node == d.dummpyTail {
		return errors.New("node is invalid")
	}

	// 首先从链表中删除该节点
	if err := d.Remove(node); err != nil {
		return err
	}

	// 将节点添加到链表尾部
	tail := d.dummpyTail.Pre
	tail.Next = node
	node.Next = d.dummpyTail
	node.Pre = tail
	d.dummpyTail.Pre = node

	// 链表大小加1（因为之前删除时减了1）
	d.size++

	return nil
}
