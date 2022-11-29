package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http server start,listen on 8888")
	http.ListenAndServe(":8888", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	}))
}
