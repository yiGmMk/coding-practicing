package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http server start,listen on 8888")
	_ = http.ListenAndServe(":8888", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello"))
	}))
}
