package util

import (
	"context"
	"runtime/debug"

	"test/logger"
)

// 仅仅是对函数本身做Panic Recover，自身不会启动协程，需要在协程中调用
func SafeGo(ctx context.Context, fun func()) {
	defer func(ctx context.Context) {
		if err := recover(); err != nil {
			stack := string(debug.Stack())
			logger.Ex(ctx, "SafeGo", "Goroutine Recover: %+v, stack is %s", err, stack)
		}
	}(ctx)
	fun()
}
