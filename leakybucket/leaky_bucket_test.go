package leakybucket

import (
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

func TestLeakyBucket(t *testing.T) {
	// 创建一个容量为10，每秒漏1个水的漏桶
	bucket := New(10, 1)

	// 前10个请求应该都通过
	for i := 0; i < 10; i++ {
		if !bucket.Allow() {
			t.Errorf("第%d个请求应该通过", i+1)
		}
	}

	// 第11个请求应该被拒绝
	if bucket.Allow() {
		t.Error("第11个请求应该被拒绝")
	}

	// 等待1秒后，应该又能通过一个请求
	time.Sleep(1 * time.Second)
	if !bucket.Allow() {
		t.Error("等待1秒后，应该能通过一个请求")
	}
}

func TestLeakyBucket2(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		bucket := New(20, 1)

		for i := 0; i < 20; i++ {
			if !bucket.Allow() {
				t.Errorf("第%d个请求应该通过", i+1)
			}
		}

		if bucket.Allow() {
			t.Error("第21个请求应该被拒绝")
		}

		time.Sleep(1 * time.Second)
		synctest.Wait()

		if !bucket.Allow() {
			t.Error("等待1秒后，应该能通过一个请求")
		}
	})
}

func TestLeakyBucket3(t *testing.T) {
	bucket := New(20, 1)

	synctest.Test(t, func(t *testing.T) {
		var wg sync.WaitGroup
		var errCount int
		for i := 0; i < 22; i++ {
			wg.Go(func() {
				if !bucket.Allow() {
					// t.Errorf("第%d个请求应该通过", i+1)
					errCount++

				}
			})
		}

		if errCount > 0 {
			t.Error("不应该有错误的请求")
		}
		synctest.Wait()

		if errCount != 2 {
			t.Error("应该有2个错误的请求")
		}
		wg.Wait()
	})
}
