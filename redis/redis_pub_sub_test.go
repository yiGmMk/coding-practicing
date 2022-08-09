package redis

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestPubSub(t *testing.T) {
	r := NewRedisInmemory(t)
	log.Println(r.Addr)

	topic := "redis_pub_sub"
	ctx := context.Background()

	go func() {
		client := redis.NewClient(&redis.Options{
			Addr:         r.Addr,
			Password:     r.Pass,
			DB:           defaultDatabase,
			MaxRetries:   maxRetries,
			MinIdleConns: idleConns,
		})
		i := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				if err := client.Publish(ctx, topic, strconv.Itoa(i)+"."+
					time.Now().Format("2006-01-02 15:03:04")+"hello").Err(); err != nil {
					log.Println(err)
				}
				time.Sleep(time.Second * 2)
				i++
			}
		}
	}()

	go func() {
		client := redis.NewClient(&redis.Options{
			Addr:         r.Addr,
			Password:     r.Pass,
			DB:           defaultDatabase,
			MaxRetries:   maxRetries,
			MinIdleConns: idleConns,
		})
		sub := client.Subscribe(ctx, topic)
		channel := sub.Channel()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-channel:
				fmt.Printf("c1:%s,%+v\n", msg.Payload, msg)
			}
		}
	}()
	go func() {
		client := redis.NewClient(&redis.Options{
			Addr:         r.Addr,
			Password:     r.Pass,
			DB:           defaultDatabase,
			MaxRetries:   maxRetries,
			MinIdleConns: idleConns,
		})
		sub := client.Subscribe(ctx, topic)
		channel := sub.Channel()
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-channel:
				fmt.Printf("c2:%s,%+v\n", msg.Payload, msg)
			}
		}
	}()
	time.Sleep(time.Second * 60 * 1)
}
