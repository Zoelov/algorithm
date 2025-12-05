package linked_list

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteLinkedList struct {
	suite.Suite
	link *LinkedList[int]
	data []int
}

func (s *SuiteLinkedList) SetupTest() {
	s.link = NewLinkedList[int](nil)
	s.data = make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		s.data = append(s.data, i)
	}
}

func (s *SuiteLinkedList) TearDownSuite() {

}

func (s *SuiteLinkedList) TestBuild() {
	s.link.Build(s.data)
	s.link.Print()
	all := s.link.Get()

	s.Equal(s.data, all, "Build should build linked list correctly")
}

func (s *SuiteLinkedList) TestAppend() {
	s.link.Build(s.data)
	s.link.Append(20)
	all := s.link.Get()
	expected := make([]int, 0)
	expected = append(expected, s.data...)
	expected = append(expected, 20)
	s.Equal(expected, all)
}

func (s *SuiteLinkedList) TestHeadInsert() {
	l := NewLinkedList[int](nil)
	l.HeadInsert([]int{10, 20, 30})
	all := l.Get()
	s.Equal([]int{30, 20, 10}, all)
}

func TestLinkedList(t *testing.T) {
	suite.Run(t, &SuiteLinkedList{})
}
