package limit

import (
	"sync"
	"time"
)

type TokenBucket struct {
	latestTokenTime time.Time     // 上次放入令牌的时间
	availableToken  int64         // 可用令牌数量
	capacity        int64         // 令牌桶容量
	interval        time.Duration // 放入一个令牌需要的时间周期
	lock            sync.Mutex    // 并发控制
}

// NewTokenBucket 创建令牌桶，capacity表示容量，rate表示1s放入的令牌数量
func NewTokenBucket(capacity int64, rate int64) *TokenBucket {
	// 计算放入一个令牌需要的时间周期
	interval := time.Second / time.Duration(rate)
	return &TokenBucket{
		latestTokenTime: time.Now(),
		availableToken:  capacity,
		capacity:        capacity,
		interval:        interval,
	}
}

// 调整令牌数量，模拟匀速放入令牌的操作
func (tb *TokenBucket) adjust() {
	// 令牌桶已满，不放入令牌
	if tb.availableToken == tb.capacity {
		return
	}
	now := time.Now()
	// 距上次放入令牌经过的时间周期，也就是需要放入的令牌数量
	newTokenCount := int64(now.Sub(tb.latestTokenTime) / tb.interval)
	if newTokenCount == 0 {
		return
	}
	// 放入令牌，并处理令牌桶溢出情况
	tb.availableToken += newTokenCount
	if tb.availableToken > tb.capacity {
		tb.availableToken = tb.capacity
	}
	// 更新放入令牌的时间
	tb.latestTokenTime = now
}

// Take 获取令牌
func (tb *TokenBucket) Take() bool {
	tb.lock.Lock()
	defer tb.lock.Unlock()
	tb.adjust() // 调整令牌桶中最新的令牌数量
	// 若有可用令牌，则允许请求通过
	if tb.availableToken > 0 {
		tb.availableToken -= 1
		return true
	}
	return false
}

func (tb *TokenBucket) TakeN(n int64) bool {
	tb.lock.Lock()
	defer tb.lock.Unlock()
	tb.adjust() // 调整令牌桶中最新的令牌数量
	// 若有足够的可用令牌，则允许请求通过
	if tb.availableToken >= n {
		tb.availableToken -= n
		return true
	}
	return false
}
