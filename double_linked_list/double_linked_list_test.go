package double_linked_list

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// DoubleLinkedListTestSuite 双向链表测试套件
type DoubleLinkedListTestSuite struct {
	suite.Suite
	list *DoubleLinkedList[int]
}

// SetupTest 每个测试前的设置
func (suite *DoubleLinkedListTestSuite) SetupTest() {
	suite.list = NewDoubleLinkedList[int]()
}

// TestNewDoubleLinkedList 测试创建新链表
func (suite *DoubleLinkedListTestSuite) TestNewDoubleLinkedList() {
	suite.NotNil(suite.list)
	suite.NotNil(suite.list.dummpyHead)
	suite.NotNil(suite.list.dummpyTail)
	suite.Equal(0, suite.list.Size())
	suite.Equal(suite.list.dummpyTail, suite.list.dummpyHead.Next)
	suite.Equal(suite.list.dummpyHead, suite.list.dummpyTail.Pre)
}

// TestPrepend 测试在头部添加元素
func (suite *DoubleLinkedListTestSuite) TestPrepend() {
	// 测试添加一个元素
	suite.list.Prepend(1)
	suite.Equal(1, suite.list.Size())
	suite.Equal(1, suite.list.dummpyHead.Next.Val)
	suite.Equal(suite.list.dummpyTail, suite.list.dummpyHead.Next.Next)
	suite.Equal(suite.list.dummpyHead, suite.list.dummpyHead.Next.Pre)

	// 测试添加多个元素
	suite.list.Prepend(2)
	suite.Equal(2, suite.list.Size())
	suite.Equal(2, suite.list.dummpyHead.Next.Val)
	suite.Equal(1, suite.list.dummpyHead.Next.Next.Val)
	suite.Equal(suite.list.dummpyHead.Next, suite.list.dummpyHead.Next.Next.Pre)
}

// TestAppend 测试在尾部添加元素
func (suite *DoubleLinkedListTestSuite) TestAppend() {
	// 测试添加一个元素
	suite.list.Append(1)
	suite.Equal(1, suite.list.Size())
	suite.Equal(1, suite.list.dummpyTail.Pre.Val)
	suite.Equal(suite.list.dummpyHead, suite.list.dummpyTail.Pre.Pre)
	suite.Equal(suite.list.dummpyTail, suite.list.dummpyTail.Pre.Next)

	// 测试添加多个元素
	suite.list.Append(2)
	suite.Equal(2, suite.list.Size())
	suite.Equal(2, suite.list.dummpyTail.Pre.Val)
	suite.Equal(1, suite.list.dummpyTail.Pre.Pre.Val)
	suite.Equal(suite.list.dummpyTail.Pre, suite.list.dummpyTail.Pre.Pre.Next)
}

// TestSize 测试获取链表大小
func (suite *DoubleLinkedListTestSuite) TestSize() {
	suite.Equal(0, suite.list.Size())

	suite.list.Append(1)
	suite.Equal(1, suite.list.Size())

	suite.list.Prepend(2)
	suite.Equal(2, suite.list.Size())

	suite.list.Append(3)
	suite.Equal(3, suite.list.Size())
}

// TestInsertAfter 测试在指定节点后插入元素
func (suite *DoubleLinkedListTestSuite) TestInsertAfter() {
	// 准备数据
	suite.list.Append(1)
	suite.list.Append(3)
	middle := suite.list.dummpyHead.Next

	// 测试在有效节点后插入
	err := suite.list.InsertAfter(middle, 2)
	suite.NoError(err)
	suite.Equal(3, suite.list.Size())
	suite.Equal(1, middle.Val)
	suite.Equal(2, middle.Next.Val)
	suite.Equal(3, middle.Next.Next.Val)
	suite.Equal(middle, middle.Next.Pre)
	suite.Equal(middle.Next, middle.Next.Next.Pre)

	// 测试在nil节点后插入
	err = suite.list.InsertAfter(nil, 4)
	suite.Error(err)

	// 测试在dummy tail后插入
	err = suite.list.InsertAfter(suite.list.dummpyTail, 4)
	suite.Error(err)
}

// TestInsertBefore 测试在指定节点前插入元素
func (suite *DoubleLinkedListTestSuite) TestInsertBefore() {
	// 准备数据
	suite.list.Append(1)
	suite.list.Append(3)
	middle := suite.list.dummpyTail.Pre

	// 测试在有效节点前插入
	err := suite.list.InsertBefore(middle, 2)
	suite.NoError(err)
	suite.Equal(3, suite.list.Size())
	suite.Equal(1, middle.Pre.Pre.Val)
	suite.Equal(2, middle.Pre.Val)
	suite.Equal(3, middle.Val)
	suite.Equal(middle.Pre.Pre, middle.Pre.Pre.Pre.Next)
	suite.Equal(middle.Pre, middle.Pre.Next.Pre)

	// 测试在nil节点前插入
	err = suite.list.InsertBefore(nil, 4)
	suite.Error(err)

	// 测试在dummy head前插入
	err = suite.list.InsertBefore(suite.list.dummpyHead, 4)
	suite.Error(err)
}

// TestRemoveHead 测试删除头元素
func (suite *DoubleLinkedListTestSuite) TestRemoveHead() {
	// 测试从空链表删除
	val, err := suite.list.RemoveHead()
	suite.Error(err)
	suite.Zero(val)

	// 准备数据
	suite.list.Append(1)
	suite.list.Append(2)
	suite.list.Append(3)

	// 测试删除头元素
	val, err = suite.list.RemoveHead()
	suite.NoError(err)
	suite.Equal(1, val)
	suite.Equal(2, suite.list.Size())
	suite.Equal(2, suite.list.dummpyHead.Next.Val)

	// 测试删除所有元素
	val, err = suite.list.RemoveHead()
	suite.NoError(err)
	suite.Equal(2, val)
	suite.Equal(1, suite.list.Size())

	val, err = suite.list.RemoveHead()
	suite.NoError(err)
	suite.Equal(3, val)
	suite.Equal(0, suite.list.Size())
	suite.Equal(suite.list.dummpyTail, suite.list.dummpyHead.Next)
	suite.Equal(suite.list.dummpyHead, suite.list.dummpyTail.Pre)
}

// TestRemoveTail 测试删除尾元素
func (suite *DoubleLinkedListTestSuite) TestRemoveTail() {
	// 测试从空链表删除
	val, err := suite.list.RemoveTail()
	suite.Error(err)
	suite.Zero(val)

	// 准备数据
	suite.list.Append(1)
	suite.list.Append(2)
	suite.list.Append(3)

	// 测试删除尾元素
	val, err = suite.list.RemoveTail()
	suite.NoError(err)
	suite.Equal(3, val)
	suite.Equal(2, suite.list.Size())
	suite.Equal(2, suite.list.dummpyTail.Pre.Val)

	// 测试删除所有元素
	val, err = suite.list.RemoveTail()
	suite.NoError(err)
	suite.Equal(2, val)
	suite.Equal(1, suite.list.Size())

	val, err = suite.list.RemoveTail()
	suite.NoError(err)
	suite.Equal(1, val)
	suite.Equal(0, suite.list.Size())
	suite.Equal(suite.list.dummpyTail, suite.list.dummpyHead.Next)
	suite.Equal(suite.list.dummpyHead, suite.list.dummpyTail.Pre)
}

// TestRemove 测试删除指定节点
func (suite *DoubleLinkedListTestSuite) TestRemove() {
	// 测试删除nil节点
	err := suite.list.Remove(nil)
	suite.Error(err)

	// 测试删除dummy节点
	err = suite.list.Remove(suite.list.dummpyHead)
	suite.Error(err)
	err = suite.list.Remove(suite.list.dummpyTail)
	suite.Error(err)

	// 准备数据
	suite.list.Append(1)
	suite.list.Append(2)
	suite.list.Append(3)
	middle := suite.list.dummpyHead.Next.Next

	// 测试删除中间节点
	err = suite.list.Remove(middle)
	suite.NoError(err)
	suite.Equal(2, suite.list.Size())
	suite.Equal(1, suite.list.dummpyHead.Next.Val)
	suite.Equal(3, suite.list.dummpyHead.Next.Next.Val)
	suite.Equal(suite.list.dummpyHead.Next, suite.list.dummpyHead.Next.Next.Pre)

	// 测试删除头节点
	head := suite.list.dummpyHead.Next
	err = suite.list.Remove(head)
	suite.NoError(err)
	suite.Equal(1, suite.list.Size())
	suite.Equal(3, suite.list.dummpyHead.Next.Val)

	// 测试删除尾节点
	tail := suite.list.dummpyTail.Pre
	err = suite.list.Remove(tail)
	suite.NoError(err)
	suite.Equal(0, suite.list.Size())
}

// TestMoveToTail 测试将节点移动到尾部
func (suite *DoubleLinkedListTestSuite) TestMoveToTail() {
	// 测试移动nil节点
	err := suite.list.MoveToTail(nil)
	suite.Error(err)

	// 测试移动dummy节点
	err = suite.list.MoveToTail(suite.list.dummpyHead)
	suite.Error(err)
	err = suite.list.MoveToTail(suite.list.dummpyTail)
	suite.Error(err)

	// 准备数据
	suite.list.Append(1)
	suite.list.Append(2)
	suite.list.Append(3)
	head := suite.list.dummpyHead.Next

	// 测试移动头节点到尾部
	err = suite.list.MoveToTail(head)
	suite.NoError(err)
	suite.Equal(3, suite.list.Size())
	suite.Equal(2, suite.list.dummpyHead.Next.Val)
	suite.Equal(3, suite.list.dummpyHead.Next.Next.Val)
	suite.Equal(1, suite.list.dummpyTail.Pre.Val)
	suite.Equal(suite.list.dummpyHead.Next.Next, suite.list.dummpyTail.Pre.Pre)

	// 准备新数据
	suite.list = NewDoubleLinkedList[int]()
	suite.list.Append(1)
	suite.list.Append(2)
	suite.list.Append(3)
	suite.list.Append(4)
	middle := suite.list.dummpyHead.Next.Next

	// 测试移动中间节点到尾部
	err = suite.list.MoveToTail(middle)
	suite.NoError(err)
	suite.Equal(4, suite.list.Size())
	suite.Equal(1, suite.list.dummpyHead.Next.Val)
	suite.Equal(3, suite.list.dummpyHead.Next.Next.Val)
	suite.Equal(4, suite.list.dummpyHead.Next.Next.Next.Val)
	suite.Equal(2, suite.list.dummpyTail.Pre.Val)
	suite.Equal(suite.list.dummpyHead.Next.Next.Next, suite.list.dummpyTail.Pre.Pre)

	// 测试移动尾节点到尾部（应该不变）
	tail := suite.list.dummpyTail.Pre
	err = suite.list.MoveToTail(tail)
	suite.NoError(err)
	suite.Equal(4, suite.list.Size())
	suite.Equal(2, suite.list.dummpyTail.Pre.Val)
}

// TestComplexOperations 测试复杂操作组合
func (suite *DoubleLinkedListTestSuite) TestComplexOperations() {
	// 测试多种操作组合
	suite.list.Append(1)
	suite.list.Append(2)
	suite.list.Prepend(0)
	suite.list.InsertAfter(suite.list.dummpyHead.Next, -1)
	suite.list.InsertBefore(suite.list.dummpyTail.Pre, 3)

	suite.Equal(5, suite.list.Size())

	// 验证链表结构
	current := suite.list.dummpyHead.Next
	suite.Equal(0, current.Val)
	current = current.Next
	suite.Equal(-1, current.Val)
	current = current.Next
	suite.Equal(1, current.Val)
	current = current.Next
	suite.Equal(3, current.Val)
	current = current.Next
	suite.Equal(2, current.Val)
}

// 运行测试套件
func TestDoubleLinkedList(t *testing.T) {
	suite.Run(t, new(DoubleLinkedListTestSuite))
}
