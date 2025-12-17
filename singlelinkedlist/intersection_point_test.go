package singlelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// 辅助函数：创建一个有环的链表
func createCircularList(data []int, pos int) *List[int] {
	list := New[int]()
	var cycleNode *Node[int]

	for i, val := range data {
		list.Append(val)
		// 找到要作为环起点的节点
		if i == pos {
			// 遍历到当前节点
			current := list.Head
			for j := 0; j < i; j++ {
				current = current.Next
			}
			cycleNode = current
		}
	}

	// 找到尾节点并将其指向环的起始节点
	if cycleNode != nil {
		tail := list.Head
		for tail.Next != nil {
			tail = tail.Next
		}
		tail.Next = cycleNode
	}

	return list
}

// IntersectionTestSuite 测试链表相交相关函数的测试套件
type IntersectionTestSuite struct {
	suite.Suite
}

// SetupTest 在每个测试用例前执行
func (suite *IntersectionTestSuite) SetupTest() {
	// 测试前的初始化逻辑
}

// TestGetIntersectionNode_BothEmptyLists 测试两个空链表的情况
func (suite *IntersectionTestSuite) TestGetIntersectionNode_BothEmptyLists() {
	list1 := New[int]()
	list2 := New[int]()

	result := GetIntersectionNode(list1, list2)
	suite.Nil(result, "两个空链表不应该相交")
}

// TestGetIntersectionNode_OneEmptyList 测试一个空链表和一个非空链表的情况
func (suite *IntersectionTestSuite) TestGetIntersectionNode_OneEmptyList() {
	list1 := New[int]()
	list2 := New[int]()
	list2.Append(1)
	list2.Append(2)
	list2.Append(3)

	result1 := GetIntersectionNode(list1, list2)
	suite.Nil(result1, "一个空链表和一个非空链表不应该相交")

	result2 := GetIntersectionNode(list2, list1)
	suite.Nil(result2, "一个非空链表和一个空链表不应该相交")
}

// TestGetIntersectionNode_NoIntersection 测试两个不相交的链表
func (suite *IntersectionTestSuite) TestGetIntersectionNode_NoIntersection() {
	list1 := New[int]()
	list1.Append(1)
	list1.Append(2)
	list1.Append(3)
	list1.Append(4)

	list2 := New[int]()
	list2.Append(5)
	list2.Append(6)
	list2.Append(7)

	result := GetIntersectionNode(list1, list2)
	suite.Nil(result, "两个不相交的链表应该返回nil")
}

// TestGetIntersectionNode_IntersectionAtHead 测试两个链表在头部相交
func (suite *IntersectionTestSuite) TestGetIntersectionNode_IntersectionAtHead() {
	// 创建一个公共部分
	common := &Node[int]{Value: 2}
	common.Next = &Node[int]{Value: 3}
	common.Next.Next = &Node[int]{Value: 4}

	// 第一个链表：1 -> 公共部分
	list1 := New[int]()
	list1.Head = &Node[int]{Value: 1}
	list1.Head.Next = common
	list1.Size = 4

	// 第二个链表：直接指向公共部分的头部
	list2 := New[int]()
	list2.Head = common
	list2.Size = 3

	result := GetIntersectionNode(list1, list2)
	suite.NotNil(result, "两个相交的链表应该返回相交节点")
	suite.Equal(2, result.Value, "相交节点的值应该是2")
}

// TestGetIntersectionNode_IntersectionAtMiddle 测试两个链表在中间节点相交
func (suite *IntersectionTestSuite) TestGetIntersectionNode_IntersectionAtMiddle() {
	// 创建一个公共部分
	common := &Node[int]{Value: 4}
	common.Next = &Node[int]{Value: 5}
	common.Next.Next = &Node[int]{Value: 6}

	// 第一个链表：1 -> 2 -> 3 -> 公共部分
	list1 := New[int]()
	list1.Head = &Node[int]{Value: 1}
	list1.Head.Next = &Node[int]{Value: 2}
	list1.Head.Next.Next = &Node[int]{Value: 3}
	list1.Head.Next.Next.Next = common
	list1.Size = 6

	// 第二个链表：7 -> 8 -> 公共部分
	list2 := New[int]()
	list2.Head = &Node[int]{Value: 7}
	list2.Head.Next = &Node[int]{Value: 8}
	list2.Head.Next.Next = common
	list2.Size = 5

	result := GetIntersectionNode(list1, list2)
	suite.NotNil(result, "两个相交的链表应该返回相交节点")
	suite.Equal(4, result.Value, "相交节点的值应该是4")
}

// TestGetIntersectionNode_IntersectionAtTail 测试两个链表在尾节点相交
func (suite *IntersectionTestSuite) TestGetIntersectionNode_IntersectionAtTail() {
	// 创建一个公共的尾节点
	commonTail := &Node[int]{Value: 5}

	// 第一个链表：1 -> 2 -> 3 -> 4 -> 公共尾节点
	list1 := New[int]()
	list1.Head = &Node[int]{Value: 1}
	list1.Head.Next = &Node[int]{Value: 2}
	list1.Head.Next.Next = &Node[int]{Value: 3}
	list1.Head.Next.Next.Next = &Node[int]{Value: 4}
	list1.Head.Next.Next.Next.Next = commonTail
	list1.Size = 5

	// 第二个链表：6 -> 7 -> 公共尾节点
	list2 := New[int]()
	list2.Head = &Node[int]{Value: 6}
	list2.Head.Next = &Node[int]{Value: 7}
	list2.Head.Next.Next = commonTail
	list2.Size = 3

	result := GetIntersectionNode(list1, list2)
	suite.NotNil(result, "两个相交的链表应该返回相交节点")
	suite.Equal(5, result.Value, "相交节点的值应该是5")
}

// TestGetIntersectionNode_OneListIsSubset 测试一个链表是另一个链表的子链表
func (suite *IntersectionTestSuite) TestGetIntersectionNode_OneListIsSubset() {
	// 创建主链表
	mainList := New[int]()
	mainList.Append(1)
	mainList.Append(2)
	mainList.Append(3)
	mainList.Append(4)
	mainList.Append(5)

	// 创建子链表，指向主链表的中间节点
	subList := New[int]()
	subList.Head = mainList.Head.Next.Next // 指向值为3的节点
	subList.Size = 3                       // 3, 4, 5

	result := GetIntersectionNode(mainList, subList)
	suite.NotNil(result, "当一个链表是另一个链表的子链表时应该返回相交节点")
	suite.Equal(3, result.Value, "相交节点的值应该是3")
}

// TestIntersectionTestSuite 运行测试套件
func TestIntersectionTestSuite(t *testing.T) {
	suite.Run(t, new(IntersectionTestSuite))
}

// CircleTestSuite 测试链表有环检测相关函数的测试套件
type CircleTestSuite struct {
	suite.Suite
}

// SetupTest 在每个测试用例前执行
func (suite *CircleTestSuite) SetupTest() {
	// 测试前的初始化逻辑
}

// TestHasCircle_EmptyList 测试空链表的情况
func (suite *CircleTestSuite) TestHasCircle_EmptyList() {
	list := New[int]()
	result := HasCircle(list)
	suite.False(result, "空链表不应该有环")
}

// TestHasCircle_SingleNode 测试只有一个节点的链表（无环）
func (suite *CircleTestSuite) TestHasCircle_SingleNode() {
	list := New[int]()
	list.Append(1)
	result := HasCircle(list)
	suite.False(result, "只有一个节点的链表不应该有环")
}

// TestHasCircle_NoCircle 测试有多个节点的链表（无环）
func (suite *CircleTestSuite) TestHasCircle_NoCircle() {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	result := HasCircle(list)
	suite.False(result, "有多个节点的无环链表不应该有环")
}

// TestHasCircle_CircleAtHead 测试有环的链表（环在头部）
func (suite *CircleTestSuite) TestHasCircle_CircleAtHead() {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// 创建环：尾节点指向头节点
	tail := list.Head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = list.Head

	result := HasCircle(list)
	suite.True(result, "环在头部的链表应该检测到环")
}

// TestHasCircle_CircleAtMiddle 测试有环的链表（环在中间）
func (suite *CircleTestSuite) TestHasCircle_CircleAtMiddle() {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	// 创建环：尾节点指向中间节点（值为3的节点）
	tail := list.Head
	var middle *Node[int]
	for tail.Next != nil {
		if tail.Value == 3 {
			middle = tail
		}
		tail = tail.Next
	}
	tail.Next = middle

	result := HasCircle(list)
	suite.True(result, "环在中间的链表应该检测到环")
}

// TestHasCircle_CircleAtTail 测试有环的链表（环在尾部，自环）
func (suite *CircleTestSuite) TestHasCircle_CircleAtTail() {
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// 创建环：尾节点指向自己
	tail := list.Head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = tail

	result := HasCircle(list)
	suite.True(result, "环在尾部（自环）的链表应该检测到环")
}

// TestCircleTestSuite 运行测试套件
func TestCircleTestSuite(t *testing.T) {
	suite.Run(t, new(CircleTestSuite))
}
