package event

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/atomic"
)

func TestEventBus(t *testing.T) {
	t.Skip()
	var cnt atomic.Int32
	sleep1 := func() {
		time.Sleep(1 * time.Second)
		cnt.Add(1)
	}
	topic1, topic2 := "test1", "topic2"
	{
		// case 1: 同步模式下，同一个 topic 中的 不同 publish 互不影响
		bus := New()
		cnt.Store(0)
		bus.Subscribe(topic1, sleep1)

		now := time.Now()
		var wg sync.WaitGroup
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				bus.Publish(topic1)
			}()
		}
		wg.Wait()
		bus.WaitAsync()

		assert.Equal(t, cnt.Load(), int32(3))
		assert.Less(t, time.Since(now), 2*time.Second)
	}
	{
		// case 2: 异步模式下，同一个 topic 中的 不同 publish 互不影响
		bus := New()
		cnt.Store(0)
		bus.SubscribeAsync(topic1, sleep1, false)

		now := time.Now()
		var wg sync.WaitGroup
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				bus.Publish(topic1)
			}()
		}
		wg.Wait()
		bus.WaitAsync()

		assert.Equal(t, cnt.Load(), int32(3))
		assert.Less(t, time.Since(now), 2*time.Second)
	}
	{
		// case 3: 同步模式下，不同 topic 中的 不同 publish 互不影响
		bus := New()
		cnt.Store(0)
		bus.Subscribe(topic1, sleep1)
		bus.Subscribe(topic2, sleep1)

		now := time.Now()
		var wg sync.WaitGroup
		for i := 0; i < 3; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				bus.Publish(topic1)
			}()
			go func() {
				defer wg.Done()
				bus.Publish(topic2)
			}()
		}
		wg.Wait()
		bus.WaitAsync()

		assert.Equal(t, cnt.Load(), int32(6))
		assert.Less(t, time.Since(now), 2*time.Second)
	}

	{
		// case 4: 异步模式下，不同 topic 中的 不同 publish 互不影响
		bus := New()
		cnt.Store(0)
		bus.SubscribeAsync(topic1, sleep1, false)
		bus.SubscribeAsync(topic2, sleep1, false)

		now := time.Now()
		var wg sync.WaitGroup
		for i := 0; i < 3; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				bus.Publish(topic1)
			}()
			go func() {
				defer wg.Done()
				bus.Publish(topic2)
			}()
		}
		wg.Wait()
		bus.WaitAsync()

		assert.Equal(t, cnt.Load(), int32(6))
		assert.Less(t, time.Since(now), 2*time.Second)
	}
	{
		// case 5: 同步异步模式下，相同topic，异步不影响同步。
		bus := New()
		cnt.Store(0)
		bus.SubscribeAsync(topic1, sleep1, false)
		bus.Subscribe(topic1, sleep1)

		now := time.Now()
		var wg sync.WaitGroup
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				bus.Publish(topic1)
			}()
		}
		wg.Wait()
		bus.WaitAsync()

		assert.Equal(t, cnt.Load(), int32(6))
		assert.Less(t, time.Since(now), 2*time.Second)
	}
	{
		// case 6: 同步异步模式下，不同 topic 中的 不同 publish 互不影响
		bus := New()
		cnt.Store(0)
		bus.Subscribe(topic1, sleep1)
		bus.SubscribeAsync(topic2, sleep1, false)

		now := time.Now()
		var wg sync.WaitGroup
		for i := 0; i < 3; i++ {
			wg.Add(2)
			go func() {
				defer wg.Done()
				bus.Publish(topic1)
			}()
			go func() {
				defer wg.Done()
				bus.Publish(topic2)
			}()
		}
		wg.Wait()
		bus.WaitAsync()

		assert.Equal(t, cnt.Load(), int32(6))
		assert.Less(t, time.Since(now), 2*time.Second)
	}
}
