package singlelinkedlist_test

import (
	"testing"

	"algorithm/singlelinkedlist"

	"github.com/stretchr/testify/suite"
)

type SingleLinkedListTestSuite struct {
	suite.Suite
	list *singlelinkedlist.List[int]
}

// SetupTest 在每个测试前运行
func (s *SingleLinkedListTestSuite) SetupTest() {
	s.list = singlelinkedlist.New[int]()
}

// TestNewList 测试创建新链表
func (s *SingleLinkedListTestSuite) TestNewList() {
	s.NotNil(s.list)
	s.Nil(s.list.Head)
	s.Equal(0, s.list.Size)
}

// TestHeadInsert 测试头部插入
func (s *SingleLinkedListTestSuite) TestHeadInsert() {
	// 插入一个元素
	s.list.HeadInsert(1)
	s.NotNil(s.list.Head)
	s.Equal(1, s.list.Head.Value)
	s.Equal(1, s.list.Size)

	// 再插入一个元素
	s.list.HeadInsert(2)
	s.Equal(2, s.list.Head.Value)
	s.Equal(1, s.list.Head.Next.Value)
	s.Equal(2, s.list.Size)

	// 再插入一个元素
	s.list.HeadInsert(3)
	s.Equal(3, s.list.Head.Value)
	s.Equal(2, s.list.Head.Next.Value)
	s.Equal(1, s.list.Head.Next.Next.Value)
	s.Equal(3, s.list.Size)
}

// TestAppend 测试尾部添加
func (s *SingleLinkedListTestSuite) TestAppend() {
	// 追加一个元素
	s.list.Append(1)
	s.NotNil(s.list.Head)
	s.Equal(1, s.list.Head.Value)
	s.Equal(1, s.list.Size)

	// 再追加一个元素
	s.list.Append(2)
	s.Equal(1, s.list.Head.Value)
	s.Equal(2, s.list.Head.Next.Value)
	s.Equal(2, s.list.Size)

	// 再追加一个元素
	s.list.Append(3)
	s.Equal(1, s.list.Head.Value)
	s.Equal(2, s.list.Head.Next.Value)
	s.Equal(3, s.list.Head.Next.Next.Value)
	s.Equal(3, s.list.Size)
}

// TestMixedOperations 测试混合操作
func (s *SingleLinkedListTestSuite) TestMixedOperations() {
	// 头部插入
	s.list.HeadInsert(1) // Size: 1
	s.list.HeadInsert(2) // Size: 2

	// 尾部添加
	s.list.Append(3) // Size: 3
	s.list.Append(4) // Size: 4

	// 头部插入
	s.list.HeadInsert(0) // Size: 5

	// 验证链表结构: 0 -> 2 -> 1 -> 3 -> 4
	s.Equal(0, s.list.Head.Value)
	s.Equal(2, s.list.Head.Next.Value)
	s.Equal(1, s.list.Head.Next.Next.Value)
	s.Equal(3, s.list.Head.Next.Next.Next.Value)
	s.Equal(4, s.list.Head.Next.Next.Next.Next.Value)
	s.Equal(5, s.list.Size)
}

// 运行测试套件
func TestSingleLinkedListTestSuite(t *testing.T) {
	suite.Run(t, new(SingleLinkedListTestSuite))
}
