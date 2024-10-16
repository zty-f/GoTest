package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestZScore(t *testing.T) {
	ctx := context.Background()
	rd.ZAdd(ctx, "zset", &redis.Z{Score: 1, Member: "one"})
	result, err := rd.ZScore(ctx, "zset", "one").Result()
	fmt.Println(result, "---", err) // 1 --- <nil>
	result, err = rd.ZScore(ctx, "zset", "two").Result()
	fmt.Println(result, "---", err) // 0 --- redis: nil

	x, err := rd.ZRank(ctx, "zset234", "two").Result()
	fmt.Println(x, "---", err) // 0 --- redis: nil
}
