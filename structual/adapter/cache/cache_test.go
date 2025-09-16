package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewDefaultCache()
	setandget(t, c)
}

func setandget(t *testing.T, c Cache) {
	err := c.Put("a", "zhangsan")
	if err != nil {
		t.Fatal(err)
	}
	value, err := c.Get("a")
	if err != nil {
		t.Fatal(err)
	}
	if value != "zhangsan" {
		t.Errorf("got %s", value)
	}
}

func TestRedisCache(t *testing.T) {
	var rdb = redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 1})
	res, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("ping 出错：", err)
	}
	fmt.Println(res)
	c := NewRedisCacheAdapter(rdb)
	setandget(t, c)
}
