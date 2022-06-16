package main

import "testing"

func TestMrc(t *testing.T) {
	files := []string{"./pg-being_ernest.txt"}
	mapReduce(files)
}
