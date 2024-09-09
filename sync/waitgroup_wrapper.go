package sync

import (
	"context"
	"sync"
	"test/sync/util"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (wg *WaitGroupWrapper) Wrap(ctx context.Context, f func()) {
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		util.SafeGo(ctx, f)
	}(ctx)
}
