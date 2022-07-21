package base

import (
	"log"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	// since 和 util都有正负之分
	now := time.Now()
	time.Sleep(time.Second * 2)

	duration2s := time.Since(now)
	log.Println("duration2s:", duration2s.Milliseconds())

	durationAdd4s := time.Since(now.Add(time.Second * 4))
	log.Println("duration4s:", durationAdd4s.Milliseconds())

	duration2sUtil := time.Until(now)
	log.Println("duration2sUtil:", duration2sUtil.Milliseconds())
}
