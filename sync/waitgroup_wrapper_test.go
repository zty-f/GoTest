package sync

import (
	"context"
	"fmt"
	"sync"
	"testing"
)

func TestWrap(t *testing.T) {
	var (
		res1 string
		err1 error
		res2 string
		err2 error
	)
	wg := WaitGroupWrapper{}
	ctx := context.Background()
	wg.Wrap(ctx, func() {
		res1, err1 = method1()
	})
	wg.Wrap(ctx, func() {
		res2, err2 = method2()
	})
	wg.Wait()
	if err1 != nil {

	}
	if err2 != nil {

	}
	fmt.Println(res1)
	fmt.Println(res2)
}

func TestWrapParam(t *testing.T) {
	var (
		res int
		err error
	)
	arr := []int{1, 2}
	wg := WaitGroupWrapper{}
	m := sync.Map{}
	ctx := context.Background()
	for _, val := range arr {
		newVal := val //注意需要重新赋值
		wg.Wrap(ctx, func() {
			res, err = method3(newVal)
			if err == nil {
				m.Store(newVal, res)
			}
		})
	}
	wg.Wait()
	for _, val := range arr {
		if result, ok := m.Load(val); ok {
			fmt.Println(result)
		}
	}
}

func method1() (string, error) {
	return "a", nil
}
func method2() (string, error) {
	return "b", nil
}

func method3(a int) (int, error) {
	return a, nil
}
