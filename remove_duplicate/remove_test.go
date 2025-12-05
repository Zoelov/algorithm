package remove_duplicate

import (
	"algorithm/linked_list"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SuiteRemoveDuplicate struct {
	suite.Suite
}

func (s *SuiteRemoveDuplicate) SetupTestSuite() {

}

func (s *SuiteRemoveDuplicate) TearDownSuite() {

}

func (s *SuiteRemoveDuplicate) Test() {
	link := linked_list.NewLinkedList[int](nil)
	data := []int{1, 1, 2, 3, 4, 4, 5, 5, 6, 6, 6, 7, 8, 9, 10, 10}
	link.Build(data)

	head := link.Head()
	link.Print()

	head = RemoveDuplicateElement(head)
	link.Print()

	all := linked_list.NewLinkedList(head).Get()

	s.Equal([]int{2, 3, 7, 8, 9}, all)
}

func TestRemoveDuplicate(t *testing.T) {
	suite.Run(t, new(SuiteRemoveDuplicate))
}
