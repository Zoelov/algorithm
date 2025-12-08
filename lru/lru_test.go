package lru

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// LRUTestSuite 是LRU缓存的测试套件
type LRUTestSuite struct {
	suite.Suite
	lru *LRU[string, int]
}

// SetupTest 在每个测试用例之前执行，初始化测试环境
func (s *LRUTestSuite) SetupTest() {
	// 创建一个容量为3的LRU缓存
	s.lru = NewLRU[string, int](3)
}

// TestNewLRU 测试创建新的LRU缓存
func (s *LRUTestSuite) TestNewLRU() {
	s.NotNil(s.lru)
	s.Equal(3, s.lru.capacity)
	s.Equal(0, s.lru.Size())
}

// TestPutAndGet 测试基本的Put和Get操作
func (s *LRUTestSuite) TestPutAndGet() {
	// 添加元素
	s.lru.Put("key1", 1)
	s.Equal(1, s.lru.Size())

	// 获取元素
	val, ok := s.lru.Get("key1")
	s.True(ok)
	s.Equal(1, val)

	// 获取不存在的元素
	val, ok = s.lru.Get("key2")
	s.False(ok)
	s.Equal(0, val)
}

// TestUpdateValue 测试更新已有元素的值
func (s *LRUTestSuite) TestUpdateValue() {
	// 添加元素
	s.lru.Put("key1", 1)

	// 更新元素值
	s.lru.Put("key1", 10)
	s.Equal(1, s.lru.Size())

	// 验证更新后的值
	val, ok := s.lru.Get("key1")
	s.True(ok)
	s.Equal(10, val)
}

// TestLRUEviction 测试LRU缓存的淘汰机制
func (s *LRUTestSuite) TestLRUEviction() {
	// 添加超过容量的元素
	s.lru.Put("key1", 1)
	s.lru.Put("key2", 2)
	s.lru.Put("key3", 3)
	s.lru.Put("key4", 4)

	// 验证容量限制
	s.Equal(3, s.lru.Size())

	// 验证最久未使用的元素被淘汰
	val, ok := s.lru.Get("key1")
	s.False(ok)
	s.Equal(0, val)

	// 验证其他元素仍然存在
	val, ok = s.lru.Get("key2")
	s.True(ok)
	s.Equal(2, val)

	val, ok = s.lru.Get("key3")
	s.True(ok)
	s.Equal(3, val)

	val, ok = s.lru.Get("key4")
	s.True(ok)
	s.Equal(4, val)
}

// TestAccessOrder 测试访问顺序对LRU的影响
func (s *LRUTestSuite) TestAccessOrder() {
	// 添加元素
	s.lru.Put("key1", 1)
	s.lru.Put("key2", 2)
	s.lru.Put("key3", 3)

	// 访问最久未使用的元素
	val, ok := s.lru.Get("key1")
	s.True(ok)
	s.Equal(1, val)

	// 添加新元素，应该淘汰key2而不是key1
	s.lru.Put("key4", 4)

	// 验证key2被淘汰
	val, ok = s.lru.Get("key2")
	s.False(ok)
	s.Equal(0, val)

	// 验证其他元素仍然存在
	val, ok = s.lru.Get("key1")
	s.True(ok)
	s.Equal(1, val)

	val, ok = s.lru.Get("key3")
	s.True(ok)
	s.Equal(3, val)

	val, ok = s.lru.Get("key4")
	s.True(ok)
	s.Equal(4, val)
}

// TestRemove 测试删除元素
func (s *LRUTestSuite) TestRemove() {
	// 添加元素
	s.lru.Put("key1", 1)
	s.lru.Put("key2", 2)

	// 删除不存在的元素
	ok := s.lru.Remove("key3")
	s.False(ok)
	s.Equal(2, s.lru.Size())

	// 删除存在的元素
	ok = s.lru.Remove("key1")
	s.True(ok)
	s.Equal(1, s.lru.Size())

	// 验证元素被删除
	val, ok := s.lru.Get("key1")
	s.False(ok)
	s.Equal(0, val)

	// 验证其他元素仍然存在
	val, ok = s.lru.Get("key2")
	s.True(ok)
	s.Equal(2, val)
}

// TestSize 测试Size方法
func (s *LRUTestSuite) TestSize() {
	// 初始大小为0
	s.Equal(0, s.lru.Size())

	// 添加元素后大小增加
	s.lru.Put("key1", 1)
	s.Equal(1, s.lru.Size())

	s.lru.Put("key2", 2)
	s.Equal(2, s.lru.Size())

	// 删除元素后大小减少
	s.lru.Remove("key1")
	s.Equal(1, s.lru.Size())

	// 超过容量时大小保持不变
	s.lru.Put("key3", 3)
	s.lru.Put("key4", 4)
	s.Equal(3, s.lru.Size())
}

// TestCapacityOne 测试容量为1的特殊情况
func (s *LRUTestSuite) TestCapacityOne() {
	lru := NewLRU[string, int](1)

	// 添加第一个元素
	lru.Put("key1", 1)
	s.Equal(1, lru.Size())

	// 添加第二个元素，第一个应该被淘汰
	lru.Put("key2", 2)
	s.Equal(1, lru.Size())

	// 验证第一个元素被淘汰
	val, ok := lru.Get("key1")
	s.False(ok)
	s.Equal(0, val)

	// 验证第二个元素存在
	val, ok = lru.Get("key2")
	s.True(ok)
	s.Equal(2, val)

	// 更新现有元素，大小不变
	lru.Put("key2", 20)
	s.Equal(1, lru.Size())

	// 验证更新后的值
	val, ok = lru.Get("key2")
	s.True(ok)
	s.Equal(20, val)
}

// TestComplexOperations 测试复杂操作组合
func (s *LRUTestSuite) TestComplexOperations() {
	// 执行一系列操作
	s.lru.Put("key1", 1)
	s.lru.Put("key2", 2)
	s.lru.Put("key3", 3)
	s.lru.Get("key1")    // key1变为最近使用
	s.lru.Put("key4", 4) // 应该淘汰key2
	s.lru.Get("key3")    // key3变为最近使用
	s.lru.Put("key5", 5) // 应该淘汰key1
	s.lru.Remove("key1") // 删除key1（可能已经被淘汰）
	s.lru.Put("key6", 6) // 应该淘汰key4

	// 验证最终状态
	s.Equal(3, s.lru.Size())

	// 验证存在的元素
	val, ok := s.lru.Get("key3")
	s.True(ok)
	s.Equal(3, val)

	val, ok = s.lru.Get("key5")
	s.True(ok)
	s.Equal(5, val)

	val, ok = s.lru.Get("key6")
	s.True(ok)
	s.Equal(6, val)

	// 验证被淘汰的元素
	val, ok = s.lru.Get("key2")
	s.False(ok)

	val, ok = s.lru.Get("key4")
	s.False(ok)

	val, ok = s.lru.Get("key1")
	s.False(ok)
}

// TestLRU 运行所有LRU测试
func TestLRU(t *testing.T) {
	suite.Run(t, new(LRUTestSuite))
}
