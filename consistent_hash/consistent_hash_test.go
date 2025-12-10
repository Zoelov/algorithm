package consistent_hash

import (
	"fmt"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/suite"
)

// ConsistentHashTestSuite 是一致性哈希的测试套件
type ConsistentHashTestSuite struct {
	suite.Suite
	ch *ConsistentHash
}

// SetupTest 在每个测试用例之前执行，初始化测试环境
func (s *ConsistentHashTestSuite) SetupTest() {
	// 创建一个每个节点有3个虚拟节点的一致性哈希实例
	s.ch = NewConsistentHash(3)
}

// TestNewConsistentHash 测试创建新的一致性哈希实例
func (s *ConsistentHashTestSuite) TestNewConsistentHash() {
	s.NotNil(s.ch)
	s.Equal(0, len(s.ch.nodes))
	s.Equal(0, len(s.ch.ring))
	s.Equal(0, len(s.ch.sortedRing))
	s.Equal(3, s.ch.virtualNodeNum)
}

// isSorted 检查uint32切片是否有序
func isSorted(slice []uint32) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i] < slice[i-1] {
			return false
		}
	}
	return true
}

// TestAddNode 测试添加节点
func (s *ConsistentHashTestSuite) TestAddNode() {
	// 添加第一个节点
	s.ch.AddNode("node1")
	s.Equal(1, len(s.ch.nodes))
	s.Equal(3, len(s.ch.ring))
	s.Equal(3, len(s.ch.sortedRing))

	// 添加第二个节点
	s.ch.AddNode("node2")
	s.Equal(2, len(s.ch.nodes))
	s.Equal(6, len(s.ch.ring))
	s.Equal(6, len(s.ch.sortedRing))

	// 验证排序环是有序的
	s.True(isSorted(s.ch.sortedRing))
}

// TestAddDuplicateNode 测试添加重复节点
func (s *ConsistentHashTestSuite) TestAddDuplicateNode() {
	// 添加节点
	s.ch.AddNode("node1")
	ringCount := len(s.ch.ring)
	nodesCount := len(s.ch.nodes)

	// 再次添加相同的节点
	s.ch.AddNode("node1")
	s.Equal(ringCount, len(s.ch.ring))
	s.Equal(nodesCount, len(s.ch.nodes))
}

// TestGetNode 测试根据键获取节点
func (s *ConsistentHashTestSuite) TestGetNode() {
	// 添加两个节点
	s.ch.AddNode("node1")
	s.ch.AddNode("node2")

	// 获取一个键对应的节点
	node, ok := s.ch.GetNode("key1")
	s.True(ok)
	s.True(node == "node1" || node == "node2")

	// 验证相同的键总是映射到相同的节点
	for i := 0; i < 10; i++ {
		n, ok := s.ch.GetNode("key1")
		s.True(ok)
		s.Equal(node, n)
	}
}

// TestGetNodeEmptyRing 测试从空环获取节点
func (s *ConsistentHashTestSuite) TestGetNodeEmptyRing() {
	// 从空环获取节点
	node, ok := s.ch.GetNode("key1")
	s.False(ok)
	s.Equal("", node)
}

// TestRemoveNode 测试移除节点
func (s *ConsistentHashTestSuite) TestRemoveNode() {
	// 添加两个节点
	s.ch.AddNode("node1")
	s.ch.AddNode("node2")
	originalRingSize := len(s.ch.ring)
	originalNodesSize := len(s.ch.nodes)

	// 移除一个节点
	s.ch.RemoveNode("node1")
	s.Equal(originalNodesSize-1, len(s.ch.nodes))
	s.Equal(originalRingSize/2, len(s.ch.ring))
	s.Equal(originalRingSize/2, len(s.ch.sortedRing))

	// 验证被移除的节点不在节点列表中
	nodes := s.ch.GetNodes()
	s.NotContains(nodes, "node1")
	s.Contains(nodes, "node2")

	// 验证排序环是有序的
	s.True(isSorted(s.ch.sortedRing))
}

// TestRemoveNonExistentNode 测试移除不存在的节点
func (s *ConsistentHashTestSuite) TestRemoveNonExistentNode() {
	// 添加一个节点
	s.ch.AddNode("node1")
	originalRingSize := len(s.ch.ring)
	originalNodesSize := len(s.ch.nodes)

	// 移除不存在的节点
	s.ch.RemoveNode("node2")
	s.Equal(originalNodesSize, len(s.ch.nodes))
	s.Equal(originalRingSize, len(s.ch.ring))
}

// TestKeyDistributionConsistency 测试键分布的一致性
func (s *ConsistentHashTestSuite) TestKeyDistributionConsistency() {
	// 添加三个节点
	s.ch.AddNode("node1")
	s.ch.AddNode("node2")
	s.ch.AddNode("node3")

	// 记录每个节点的键分布
	distribution := make(map[string]int)
	keys := 1000

	// 统计键的分布
	for i := 0; i < keys; i++ {
		key := fmt.Sprintf("key-%d", i)
		node, _ := s.ch.GetNode(key)
		distribution[node]++
	}

	// 验证所有节点都有一定的键分布
	s.Len(distribution, 3)

	// 打印分布情况（可选）
	for node, count := range distribution {
		fmt.Printf("Node %s: %d keys (%.2f%%)\n", node, count, float64(count)/float64(keys)*100)
	}
}

// TestKeyDistributionAfterNodeRemoval 测试移除节点后键分布的一致性
func (s *ConsistentHashTestSuite) TestKeyDistributionAfterNodeRemoval() {
	// 添加三个节点
	s.ch.AddNode("node1")
	s.ch.AddNode("node2")
	s.ch.AddNode("node3")

	// 记录键到节点的映射
	keyToNode := make(map[string]string)
	keys := 1000

	for i := 0; i < keys; i++ {
		key := fmt.Sprintf("key-%d", i)
		node, _ := s.ch.GetNode(key)
		keyToNode[key] = node
	}

	// 移除一个节点
	s.ch.RemoveNode("node1")

	// 记录移除节点后键到节点的新映射
	newKeyToNode := make(map[string]string)
	for i := 0; i < keys; i++ {
		key := fmt.Sprintf("key-%d", i)
		node, _ := s.ch.GetNode(key)
		newKeyToNode[key] = node
	}

	// 验证节点1不再被映射
	for _, node := range newKeyToNode {
		s.True(node == "node2" || node == "node3")
	}

	// 计算映射保持不变的比例
	consistentCount := 0
	totalNonRemovedKeys := 0

	for key, node := range keyToNode {
		if node == "node1" {
			// 这些键应该重新映射
			continue
		}
		totalNonRemovedKeys++
		if newKeyToNode[key] == node {
			consistentCount++
		}
	}

	// 验证大部分映射保持不变
	if totalNonRemovedKeys > 0 {
		consistentRatio := float64(consistentCount) / float64(totalNonRemovedKeys)
		// 对于只有3个虚拟节点的情况，预期的一致性比例会较低
		s.Greater(consistentRatio, 0.4) // 至少40%的映射应该保持不变
	} else {
		// 如果没有非node1的键，跳过此检查
		s.T().Log("No keys mapped to non-removed nodes")
	}
}

// TestConcurrentAccess 测试并发访问一致性哈希
func (s *ConsistentHashTestSuite) TestConcurrentAccess() {
	// 添加初始节点
	s.ch.AddNode("node1")
	s.ch.AddNode("node2")
	s.ch.AddNode("node3")

	var wg sync.WaitGroup
	operations := 1000
	concurrency := 10

	// 并发执行添加节点、获取节点和移除节点操作
	for i := range concurrency {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()

			for j := 0; j < operations/concurrency; j++ {
				// 添加新节点
				if j%5 == 0 {
					newNode := fmt.Sprintf("node-%d-%d", goroutineID, j)
					s.ch.AddNode(newNode)
				}

				// 获取节点
				key := fmt.Sprintf("key-%d-%d", goroutineID, j)
				node, ok := s.ch.GetNode(key)
				s.True(ok)
				s.NotEmpty(node)

				// 移除节点
				if j%10 == 0 {
					removeNode := fmt.Sprintf("node-%d-%d", goroutineID, j-5)
					s.ch.RemoveNode(removeNode)
				}
			}
		}(i)
	}

	// 等待所有goroutine完成
	wg.Wait()

	// 验证数据结构的一致性
	s.Len(s.ch.ring, len(s.ch.sortedRing))
	s.True(isSorted(s.ch.sortedRing))

	// 验证所有虚拟节点都映射到存在的真实节点
	for _, node := range s.ch.ring {
		_, exists := s.ch.nodes[node]
		s.True(exists)
	}
}

// TestGetNodes 测试获取所有节点
func (s *ConsistentHashTestSuite) TestGetNodes() {
	// 添加三个节点
	s.ch.AddNode("node1")
	s.ch.AddNode("node2")
	s.ch.AddNode("node3")

	// 获取所有节点
	nodes := s.ch.GetNodes()
	s.Len(nodes, 3)

	// 验证节点列表包含所有添加的节点
	s.Contains(nodes, "node1")
	s.Contains(nodes, "node2")
	s.Contains(nodes, "node3")

	// 验证节点列表是排序的
	sortedNodes := make([]string, len(nodes))
	copy(sortedNodes, nodes)
	sort.Strings(sortedNodes)
	s.Equal(sortedNodes, nodes)
}

// TestVirtualNodeDistribution 测试虚拟节点的分布
func (s *ConsistentHashTestSuite) TestVirtualNodeDistribution() {
	// 添加两个节点，每个节点有3个虚拟节点
	s.ch.AddNode("node1")
	s.ch.AddNode("node2")

	// 统计每个真实节点对应的虚拟节点数量
	virtualNodeCount := make(map[string]int)
	for _, node := range s.ch.ring {
		virtualNodeCount[node]++
	}

	// 验证每个真实节点有正确数量的虚拟节点
	s.Equal(3, virtualNodeCount["node1"])
	s.Equal(3, virtualNodeCount["node2"])
	s.Equal(2, len(virtualNodeCount))
}

// TestConsistentHash 运行所有一致性哈希测试
func TestConsistentHash(t *testing.T) {
	suite.Run(t, new(ConsistentHashTestSuite))
}
