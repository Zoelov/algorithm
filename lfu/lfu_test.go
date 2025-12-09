package lfu

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// LFUTestSuite 是LFU缓存的测试套件
type LFUTestSuite struct {
	suite.Suite
	lfu *LFU[string, int]
}

// SetupTest 在每个测试用例之前执行，初始化测试环境
func (s *LFUTestSuite) SetupTest() {
	// 创建一个容量为3的LFU缓存
	s.lfu = NewLFU[string, int](3)
}

// TestNewLFU 测试创建新的LFU缓存
func (s *LFUTestSuite) TestNewLFU() {
	s.NotNil(s.lfu)
	s.Equal(3, s.lfu.capacity)
}

// TestPutAndGet 测试基本的Put和Get操作
func (s *LFUTestSuite) TestPutAndGet() {
	// 添加元素
	err := s.lfu.Put("key1", 1)
	s.NoError(err)

	// 获取元素
	val, ok := s.lfu.Get("key1")
	s.True(ok)
	s.Equal(1, val)

	// 获取不存在的元素
	val, ok = s.lfu.Get("key2")
	s.False(ok)
	s.Equal(0, val)

	err = s.lfu.Put("key2", 2)
	s.Nil(err)

	val, ok = s.lfu.Get("key2")
	s.True(ok)
	s.Equal(2, val)
	s.Equal(2, s.lfu.keyToFreq["key2"])


	err = s.lfu.Put("key2", 3)
	s.Nil(err)

	s.Equal(3, s.lfu.keyToFreq["key2"])
}

// TestUpdateValue 测试更新已有元素的值
func (s *LFUTestSuite) TestUpdateValue() {
	// 添加元素
	err := s.lfu.Put("key1", 1)
	s.NoError(err)

	// 更新元素值
	err = s.lfu.Put("key1", 10)
	s.NoError(err)

	// 验证更新后的值
	val, ok := s.lfu.Get("key1")
	s.True(ok)
	s.Equal(10, val)
}


// TestFrequencyUpdate 测试频率更新机制
func (s *LFUTestSuite) TestFrequencyUpdate() {
	// 添加元素
	err := s.lfu.Put("key1", 1)
	s.NoError(err)

	// 第一次访问
	_, ok := s.lfu.Get("key1")
	s.True(ok)
	s.Equal(2, s.lfu.keyToFreq["key1"])

	// 第二次访问
	_, ok = s.lfu.Get("key1")
	s.True(ok)
	s.Equal(3, s.lfu.keyToFreq["key1"])

	// 添加第二个元素并访问
	err = s.lfu.Put("key2", 2)
	s.NoError(err)
	s.Equal(1, s.lfu.keyToFreq["key2"])

	_, ok = s.lfu.Get("key2")
	s.True(ok)
	s.Equal(2, s.lfu.keyToFreq["key2"])
}

// TestEviction 测试淘汰机制
func (s *LFUTestSuite) TestEviction() {
	// 添加超过容量的元素
	err := s.lfu.Put("key1", 1)
	s.NoError(err)
	err = s.lfu.Put("key2", 2)
	s.NoError(err)
	err = s.lfu.Put("key3", 3)
	s.NoError(err)
	err = s.lfu.Put("key4", 4)
	s.NoError(err)

	// 应该淘汰频率最低的元素
	val, ok := s.lfu.Get("key1")
	s.False(ok)
	s.Equal(0, val)

	// 其他元素应该仍然存在
	val, ok = s.lfu.Get("key2")
	s.True(ok)
	s.Equal(2, val)

	val, ok = s.lfu.Get("key3")
	s.True(ok)
	s.Equal(3, val)

	val, ok = s.lfu.Get("key4")
	s.True(ok)
	s.Equal(4, val)
}

// TestEvictionWithFrequency 测试基于频率的淘汰
func (s *LFUTestSuite) TestEvictionWithFrequency() {
	// 添加元素
	err := s.lfu.Put("key1", 1)
	s.NoError(err)
	err = s.lfu.Put("key2", 2)
	s.NoError(err)
	err = s.lfu.Put("key3", 3)
	s.NoError(err)

	// 增加key1和key2的频率
	_, ok := s.lfu.Get("key1")
	s.True(ok)
	_, ok = s.lfu.Get("key1")
	s.True(ok)
	_, ok = s.lfu.Get("key2")
	s.True(ok)

	// 此时频率: key1=3, key2=2, key3=1
	s.Equal(3, s.lfu.keyToFreq["key1"])
	s.Equal(2, s.lfu.keyToFreq["key2"])
	s.Equal(1, s.lfu.keyToFreq["key3"])

	// 添加新元素，应该淘汰频率最低的key3
	err = s.lfu.Put("key4", 4)
	s.NoError(err)

	// key3应该被淘汰
	val, ok := s.lfu.Get("key3")
	s.False(ok)
	s.Equal(0, val)

	// key1, key2, key4应该存在
	val, ok = s.lfu.Get("key1")
	s.True(ok)
	s.Equal(1, val)

	val, ok = s.lfu.Get("key2")
	s.True(ok)
	s.Equal(2, val)

	val, ok = s.lfu.Get("key4")
	s.True(ok)
	s.Equal(4, val)
}

// TestCapacityOne 测试容量为1的特殊情况
func (s *LFUTestSuite) TestCapacityOne() {
	lfu := NewLFU[string, int](1)

	// 添加第一个元素
	err := lfu.Put("key1", 1)
	s.NoError(err)

	// 添加第二个元素，第一个应该被淘汰
	err = lfu.Put("key2", 2)
	s.NoError(err)

	// 验证第一个元素被淘汰
	val, ok := lfu.Get("key1")
	s.False(ok)
	s.Equal(0, val)

	// 验证第二个元素存在
	val, ok = lfu.Get("key2")
	s.True(ok)
	s.Equal(2, val)

	// 访问第二个元素，增加其频率
	val, ok = lfu.Get("key2")
	s.True(ok)
	s.Equal(2, val)

	// 添加第三个元素，应该淘汰频率更低的（虽然只有一个元素）
	err = lfu.Put("key3", 3)
	s.NoError(err)

	// 验证第二个元素被淘汰
	val, ok = lfu.Get("key2")
	s.False(ok)
	s.Equal(0, val)

	// 验证第三个元素存在
	val, ok = lfu.Get("key3")
	s.True(ok)
	s.Equal(3, val)
}

// TestComplexOperations 测试复杂操作组合
func (s *LFUTestSuite) TestComplexOperations() {
	// 创建容量为4的LFU
	lfu := NewLFU[string, int](4)

	// 执行一系列操作
	err := lfu.Put("key1", 1)
	s.NoError(err)
	err = lfu.Put("key2", 2)
	s.NoError(err)
	err = lfu.Put("key3", 3)
	s.NoError(err)
	err = lfu.Put("key4", 4)
	s.NoError(err)

	// 增加key1和key2的频率
	_, ok := lfu.Get("key1")
	s.True(ok)
	_, ok = lfu.Get("key1")
	s.True(ok)
	_, ok = lfu.Get("key2")
	s.True(ok)

	// 此时频率: key1=3, key2=2, key3=1, key4=1

	// 添加新元素，应该淘汰key3或key4
	err = lfu.Put("key5", 5)
	s.NoError(err)

	// 验证其中一个低频元素被淘汰
	_, ok1 := lfu.Get("key3")
	_, ok2 := lfu.Get("key4")
	// 至少有一个应该被淘汰
	s.True(!ok1 || !ok2)

	// 验证其他元素存在
	_, ok = lfu.Get("key1")
	s.True(ok)

	_, ok = lfu.Get("key2")
	s.True(ok)

	_, ok = lfu.Get("key5")
	s.True(ok)

	// 继续操作
	err = lfu.Put("key6", 6)
	s.NoError(err)

	// 此时应该淘汰剩下的那个低频元素（key3或key4）
	_, ok3 := lfu.Get("key3")
	_, ok4 := lfu.Get("key4")
	// 两个都应该被淘汰
	s.False(ok3 && ok4)
}

// TestGetNonExistentKey 测试获取不存在的键
func (s *LFUTestSuite) TestGetNonExistentKey() {
	val, ok := s.lfu.Get("non-existent-key")
	s.False(ok)
	s.Equal(0, val)
}

// TestPutDuplicateKey 测试插入重复的键
func (s *LFUTestSuite) TestPutDuplicateKey() {
	// 第一次插入
	err := s.lfu.Put("key1", 1)
	s.NoError(err)
	s.Equal(1, len(s.lfu.entries))

	// 第二次插入相同的键
	err = s.lfu.Put("key1", 10)
	s.NoError(err)
	// 容量不应该增加
	s.Equal(1, len(s.lfu.entries))

	// 验证值被更新
	val, ok := s.lfu.Get("key1")
	s.True(ok)
	s.Equal(10, val)
}

// TestEvictionWhenFrequencyListEmpty 测试频率列表为空时的淘汰
func (s *LFUTestSuite) TestEvictionWhenFrequencyListEmpty() {
	// 添加三个元素
	err := s.lfu.Put("key1", 1)
	s.NoError(err)
	err = s.lfu.Put("key2", 2)
	s.NoError(err)
	err = s.lfu.Put("key3", 3)
	s.NoError(err)

	// 访问所有三个元素，使它们的频率变为2
	_, ok := s.lfu.Get("key1")
	s.True(ok)
	_, ok = s.lfu.Get("key2")
	s.True(ok)
	_, ok = s.lfu.Get("key3")
	s.True(ok)

	// 此时频率为1的列表应该为空
	s.Equal(0, s.lfu.freqToDLink[1].Size())

	// 添加第四个元素，应该淘汰其中一个元素
	err = s.lfu.Put("key4", 4)
	s.NoError(err)

	// 验证缓存中只有4个元素
	s.Equal(3, len(s.lfu.entries))
}

// TestLRU 运行所有LFU测试
func TestLFU(t *testing.T) {
	suite.Run(t, new(LFUTestSuite))
}
