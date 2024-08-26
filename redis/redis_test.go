package cache

import (
	"codeup.aliyun.com/61e54b0e0bb300d827e1ae27/backend/golib/logger"
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestRedisContextDeadlineExceeded(t *testing.T) {
	// 创建一个Redis客户端
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis服务器的地址和端口
		Password: "",               // Redis服务器的密码，如果没有设置密码则为空字符串
		DB:       0,                // Redis数据库的索引
	})

	// 创建一个上下文对象并设置超时时间为1秒   指定时间内没有完成操作就会报错context deadline exceeded
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	// 执行GET操作，并传入上下文对象
	value, err := client.Get(ctx, "b").Result()
	if err == redis.Nil {
		fmt.Println("key does not exist")
	} else if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("value:", value)
	}

	// 处理超时错误
	if errors.Is(err, context.DeadlineExceeded) {
		fmt.Println("Redis context deadline exceeded")
	} else if err != nil {
		fmt.Println("error:", err)
	}
}

func TestHGet(t *testing.T) {
	ctx := context.Background()
	v, err := rd.HGet(ctx, "yyy", "t").Result()
	if err != nil {
		if err != redis.Nil {
			logger.Ex(ctx, "redisHGetFailed", "%s", err.Error())
		}
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestHSetNx(t *testing.T) {
	ctx := context.Background()
	v, err := rd.HSetNX(ctx, "a", "c", "2").Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestHLen(t *testing.T) {
	ctx := context.Background()
	v, err := rd.HLen(ctx, "4444444").Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestExists(t *testing.T) {
	ctx := context.Background()
	v, err := rd.Exists(ctx, "zset").Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestSetNx(t *testing.T) {
	ctx := context.Background()
	v, err := rd.SetNX(ctx, "t", "c", 30*time.Second).Result()
	fmt.Println("-------")
	fmt.Println(v)
	fmt.Println(err)
}

func TestExpireAt(t *testing.T) {
	ctx := context.Background()
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	x := rand.Int63()
	v, err := rd.ExpireAt(ctx, "zset", time.Now().Add(time.Duration(x)*time.Second)).Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func TestExpireNX(t *testing.T) {
	ctx := context.Background()
	// redis 7.0以上版本才支持
	v, err := rd.ExpireNX(ctx, "list2", 120*time.Second).Result()
	if err != nil {
		fmt.Println(err, v)
		return
	}
	fmt.Println("-------")
	fmt.Println(v)
}

func Test取余(t *testing.T) {
	fmt.Println(2100051684 / 8 % 3)
}

// GetDayEndTime 获取一天的结束
func GetDayEndTime(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func TestRand(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	randSec := rand.Intn(10)
	fmt.Println(randSec)
	fmt.Println(GetDayEndTime(time.Now()).Add(time.Duration(randSec) * time.Second).Sub(time.Now()))
}

func TestTime(t *testing.T) {
	fmt.Println(time.Now().Add(-6 * time.Second).Before(time.Now()))
	fmt.Println(time.Now().Add(-6 * time.Second).After(time.Now()))

	lastFinishTime, _ := time.ParseInLocation("2006-01-02", "2024-04-09", time.Local)
	fmt.Println(lastFinishTime)
	fmt.Println(time.Now().Sub(lastFinishTime) < 7*24*time.Hour)
}

type User struct {
	m map[int]string
}

func TestMap(t *testing.T) {
	u := &User{}
	fmt.Println(u.m)        // map[]
	fmt.Println(u.m == nil) // true
	s, ok := u.m[1]
	fmt.Println(s, ok) // false
	u.m[1] = "1"       // 空指针异常
	var m map[int]string
	fmt.Println(m) // map[]
	s, ok = m[1]
	fmt.Println(s, ok) // false
	m[1] = "1"         // 空指针异常
	// 总结 当map为nil的时候，取值会返回零值，不出错，赋值会报空指针异常
}

func TestGet(t *testing.T) {
	result, err := rd.Get(context.Background(), "key").Result()
	fmt.Println(result)
	fmt.Println(err)
}

func TestMultiSet(t *testing.T) {
	ctx := context.Background()
	keys := []string{"key1", "key2", "key3", "key4"}
	result := make(map[string]string)

	pipeline := rd.Pipeline()

	redisResultMap := make(map[string]*redis.StatusCmd)
	for _, key := range keys {
		redisResultMap[key] = pipeline.Set(ctx, key, 1, 10*time.Minute)
	}

	_, err := pipeline.Exec(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}

	for key, redisResult := range redisResultMap {
		result[key] = redisResult.Val()
		fmt.Println(key, result[key])
		fmt.Println(redisResult.Result())
	}

	err = pipeline.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func TestMultiGet(t *testing.T) {
	ctx := context.Background()
	keys := []string{"key1", "key2", "key3", "key4", "key5"}
	result := make(map[string]string)

	pipeline := rd.Pipeline()

	redisResultMap := make(map[string]*redis.StringCmd)
	for _, key := range keys {
		redisResultMap[key] = pipeline.Get(ctx, key)
	}

	_, err := pipeline.Exec(ctx)
	if err != nil && !errors.Is(err, redis.Nil) {
		fmt.Println(err)
		return
	}

	for key, redisResult := range redisResultMap {
		result[key] = redisResult.Val()
		fmt.Println(redisResult.Val())
		fmt.Println(redisResult.Result())
	}

	err = pipeline.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}

func TestHMGet(t *testing.T) {
	ctx := context.Background()
	rd.HSet(ctx, "hash", "key1", "value1", "key2", "value2", "key3", "value3", "key4", "value4", "key5", "value5")
	keys := []string{"key1", "key2", "key223"}
	result, err := rd.HMGet(ctx, "hash", keys...).Result()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result[0])
	fmt.Println(cast.ToInt(result[0]))
	fmt.Printf("%+v\n", result)
}
