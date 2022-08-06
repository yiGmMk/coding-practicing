package timingwheel

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

const (
	key = "seconds:%d"
)

func Key(seconds int) string {
	return fmt.Sprintf(key, seconds)
}

func NewRedis() (string, error) {
	// run redis in memory
	r, err := miniredis.Run()
	if err != nil {
		return "", err
	}
	redisAddress := r.Addr()
	return redisAddress, nil
}

func worker(id int, redisAddress string) {
	log.Println("-----------------worker start-------------------")
	index := 0
	r := redis.New(redisAddress)
	ticker := time.NewTicker(time.Second)
	for now := range ticker.C {
		vs, err := r.Smembers(Key(now.Second()))
		if err != nil {
			log.Println(err)
			continue
		} else {
			_, _ = r.Del(Key(now.Second()))
		}
		log.Println(id, ".worker,second:", now.Second(), "deal with:", vs, " count:", len(vs))
		index++
		if index == 100 {
			break
		}
	}
}

func producer(id int, redisAddress string) {
	log.Println("-----------------producer start-------------------")
	r := redis.New(redisAddress)

	for now := range time.Tick(time.Second) {
		if now.Second()%id == 0 {
			_, _ = r.Sadd(Key(now.Add(time.Second).Second()), id)
			log.Println(id, ".producer,second:", now.Add(time.Second).Second(), "add:", id)
		}
	}
}

/*
producer 每次把当前需要处理的数据加入下一秒的队列中

	worker   逐秒处理队列中的数据
*/
func TestTime(t *testing.T) {
	redisAddress, err := NewRedis()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("redis address:", redisAddress)

	log.Println("add second:", time.Date(2006, 1, 2, 3, 4, 59, 1, time.Local).Add(time.Second).Second())
	go worker(100, redisAddress)
	go producer(1, redisAddress)
	go producer(3, redisAddress)
	go producer(5, redisAddress)
	go producer(7, redisAddress)
	go producer(11, redisAddress)
	go producer(15, redisAddress)

	now := time.Now()
	for {
		if time.Now().After(now.Add(time.Minute * 2)) {
			break
		}
	}
}
