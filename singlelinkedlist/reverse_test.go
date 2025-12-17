package singlelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ReverseTestSuite struct {
	suite.Suite
}

func (s *ReverseTestSuite) SetupTest() {
}

func (s *ReverseTestSuite) TestReverseList_Empty() {
	var head *Node[int]
	result := ReverseList(head)
	s.Nil(result)
}

func (s *ReverseTestSuite) TestReverseList_Single() {
	head := &Node[int]{Value: 1}
	result := ReverseList(head)
	s.NotNil(result)
	s.Equal(1, result.Value)
	s.Nil(result.Next)
}

func (s *ReverseTestSuite) TestReverseList_Multiple() {
	// 1->2->3->4->5
	list := New[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)

	head := list.Head
	result := ReverseList(head)

	// 5->4->3->2->1
	s.NotNil(result)
	s.Equal(5, result.Value)
	s.Equal(4, result.Next.Value)
	s.Equal(3, result.Next.Next.Value)
	s.Equal(2, result.Next.Next.Next.Value)
	s.Equal(1, result.Next.Next.Next.Next.Value)
	s.Nil(result.Next.Next.Next.Next.Next)
}

func TestReverseTestSuite(t *testing.T) {
	suite.Run(t, new(ReverseTestSuite))
}
