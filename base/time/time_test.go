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

// 1.19 new function
func Test119(t *testing.T) {
	// returns the start and end times of the time zone in effect at a given time.
	st, ed := time.Now().ZoneBounds()
	log.Println(
		"st:", st.Format("2006-01-02 15:04:05"),
		"ed:", ed.Format("2006-01-02 15:04:05"))

	d := ed.Sub(st).Abs()
	log.Println("d:", d.Hours())
}
