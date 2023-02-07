package jzoffer

import (
	"reflect"
	"testing"
)

var _testcase = []struct {
	name, time, expect []string
}{
	{
		name: []string{
			"leslie", "leslie", "leslie",
			"clare", "clare", "clare", "clare"},
		time: []string{
			"13:00", "13:20", "14:00",
			"18:00", "18:51", "19:30", "19:49"},
		expect: []string{"clare", "leslie"},
	},
	{
		name:   []string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
		time:   []string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
		expect: []string{"daniel"},
	},
	{
		name:   []string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
		time:   []string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"},
		expect: []string{"bob"},
	},

	{
		name: []string{
			"a", "a", "a", "a",
			"b", "b", "b", "b", "b", "b",
			"c", "c", "c", "c"},
		time: []string{
			"01:35", "08:43", "20:49", "00:01",
			"17:44", "02:50", "18:48", "22:27", "14:12", "18:00",
			"12:38", "20:40", "03:59", "22:24"},
	},
}

func Test1604(t *testing.T) {
	for i, v := range _testcase {
		got := alertNames(v.name, v.time)
		if !reflect.DeepEqual(got, v.expect) {
			t.Errorf("[%d],got:%s,want:%s", i, got, v.expect)
		}
	}
}
