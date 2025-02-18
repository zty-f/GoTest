package util

import (
	"context"
	"test/logger"
	"time"
)

func AsyncSetNxWithRetry(ctx context.Context, maxRetries int, delay time.Duration, key string, value string, expiration time.Duration, setNx func(ctx context.Context, key string, value string, expiration time.Duration) (bool, error)) {
	go func() {
		for retries := 1; retries <= maxRetries; retries++ {
			_, err := setNx(ctx, key, value, expiration)
			if err != nil {
				logger.Ex(ctx, "AsyncSetNxWithRetry", "setNx redis cache failed,key[%v],retryTimes:[%v],err[%v]", key, retries, err)
				time.Sleep(delay) // 等待一段时间再重试
				continue
			}
			break
		}
	}()
}

func AsyncHMSetWithRetry(ctx context.Context, maxRetries int, delay time.Duration, key string, values map[string]interface{}, HMSet func(ctx context.Context, key string, values map[string]interface{}) bool) {
	go func() {
		for retries := 1; retries <= maxRetries; retries++ {
			success := HMSet(ctx, key, values)
			if !success {
				logger.Ex(ctx, "AsyncHMSetWithRetry", "HMSet redis cache failed,key[%v],values[%+v],retryTimes:[%v]", key, values, retries)
				time.Sleep(delay) // 等待一段时间再重试
				continue
			}
			break
		}
	}()
}

func AsyncIncrByWithRetry(ctx context.Context, maxRetries int, delay time.Duration, key string, inc int64, incrBy func(ctx context.Context, key string, inc int64) (res int64, err error)) {
	go func() {
		for retries := 1; retries <= maxRetries; retries++ {
			_, err := incrBy(ctx, key, inc)
			if err != nil {
				logger.Ex(ctx, "AsyncIncrByWithRetry", "incrBy redis cache failed,key[%v],retryTimes:[%v],err[%v]", key, retries, err)
				time.Sleep(delay) // 等待一段时间再重试
				continue
			}
			break
		}
	}()
}

func AsyncSendKafkaNewTermActivityTaskWithRetry(ctx context.Context, maxRetries int, delay time.Duration, stuId int64, growthValue int, sendKafkaNewTermActivityTask func(ctx context.Context, stuId int64, growthValue int) error) {
	go func() {
		for retries := 1; retries <= maxRetries; retries++ {
			err := sendKafkaNewTermActivityTask(ctx, stuId, growthValue)
			if err != nil {
				logger.Ex(ctx, "AsyncSendKafkaNewTermActivityTaskWithRetry", "sendKafkaNewTermActivityTask failed,stuId[%d],growthValue[%d],retryTimes:[%v],err[%v]", stuId, growthValue, retries, err)
				time.Sleep(delay) // 等待一段时间再重试
				continue
			}
			break
		}
	}()
}

func AsyncSendKafkaOctoberActivityTaskWithRetry(ctx context.Context, maxRetries int, delay time.Duration, stuId int64, roundNo, finishNum int, SendKafkaOctoberActivityTask func(ctx context.Context, stuId int64, roundNo, finishNum int) error) {
	go func() {
		for retries := 1; retries <= maxRetries; retries++ {
			err := SendKafkaOctoberActivityTask(ctx, stuId, roundNo, finishNum)
			if err != nil {
				logger.Ex(ctx, "AsyncSendKafkaOctoberActivityTaskWithRetry", "sendKafkaNewTermActivityTask failed,stuId[%d],roundNo[%d],finishNum:[%d],retryTimes:[%v],err[%v]", stuId, roundNo, finishNum, retries, err)
				time.Sleep(delay) // 等待一段时间再重试
				continue
			}
			break
		}
	}()
}
