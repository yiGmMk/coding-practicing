package redis

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestStream(t *testing.T) {
	r := NewRedisInmemory(t)

	ctx := context.Background()
	id := 0
	fileName := "./info.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return
	}
	defer func() {
		file.Close()
		_ = os.Remove(fileName)
	}()
	go func() {
		client := NewClientFromOptions(t, r)
		for range time.Tick(time.Second * 2) {
			err := client.XAdd(ctx, &redis.XAddArgs{
				Stream: _stream,
				MaxLen: 10,
				Values: map[string]interface{}{
					"id": id,
				},
			}).Err()
			if err != nil {
				t.Error(err)
				continue
			}
			id++
			log.Println("produce:", id)

			g, err := client.XInfoGroups(ctx, _stream).Result()
			if err != nil {
				log.Println(err)
				continue
			}
			_, err = file.WriteString(fmt.Sprintf("%+v\n", g))
			if err != nil {
				log.Println("write xinfo", err)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	go func() {
		client := NewClientFromOptions(t, r)
		res, err := client.XGroupCreate(ctx, _stream, _group, "0").Result()
		if err != nil {
			if err.Error() != "BUSYGROUP Consumer Group name already exists" {
				log.Println("创建group失败", err, res)
				return
			}
		}

		consumer, err := GenConsumerName(1)
		if err != nil {
			log.Println("获取消费者失败", err)
			return
		}
		log.Println("consumer:", consumer)
		for range time.Tick(time.Second * 4) {
			streams, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    _group,
				Consumer: consumer,
				Streams:  []string{_stream, ">"},
				Count:    1,
			}).Result()
			if err != nil {
				log.Println("读取stream失败", streams, err)
				continue
			}
			log.Println("consumer:", consumer, ",consume:", streams)

			for _, stream := range streams {
				if stream.Stream != _stream {
					continue
				}
				for _, v := range stream.Messages {
					val, err := client.XAck(ctx, _stream, _group, v.ID).Result()
					if err != nil {
						log.Println("ack error", err, val)
					}
				}
			}
		}
	}()

	go func() {
		client := NewClientFromOptions(t, r)
		res, err := client.XGroupCreate(ctx, _stream, _group, "0").Result()
		if err != nil {
			if err.Error() != "BUSYGROUP Consumer Group name already exists" {
				log.Println("创建group失败", err, res)
				return
			}
			log.Println("创建group失败", err, res)
			return
		}

		consumer, err := GenConsumerName(2)
		if err != nil {
			log.Println("获取消费者失败", err)
			return
		}
		log.Println("consumer:", consumer)
		for range time.Tick(time.Second * 8) {
			streams, err := client.XReadGroup(ctx, &redis.XReadGroupArgs{
				Group:    _group,
				Consumer: consumer,
				Streams:  []string{_stream, ">"},
				Count:    1,
				NoAck:    true,
			}).Result()
			if err != nil {
				log.Println("读取stream失败", streams, err)
				continue
			}
			log.Println("consumer:", consumer, ",consume:", streams)

			for _, stream := range streams {
				if stream.Stream != _stream {
					continue
				}
				for _, v := range stream.Messages {
					val, err := client.XAck(ctx, _stream, _group, v.ID).Result()
					if err != nil {
						log.Println("ack error", err, val)
					}
				}
			}
		}
	}()
	time.Sleep(time.Second * 20)
}
