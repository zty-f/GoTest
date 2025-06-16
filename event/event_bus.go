package event

var defaultEventBus = New()

func Subscribe(topic string, fn interface{}) error {
	return defaultEventBus.Subscribe(topic, fn)
}

// 相同 topic 的异步逻辑放在同步逻辑前边，可以避免同步影响异步的执行。
func SubscribeAsync(topic string, fn interface{}) error {
	return defaultEventBus.SubscribeAsync(topic, fn, false)
}

func Publish(topic string, args ...interface{}) {
	defaultEventBus.Publish(topic, args...)
}

func PublishWithPool(pool Pool, topic string, args ...interface{}) {
	defaultEventBus.PublishWithPool(pool, topic, args...)
}
