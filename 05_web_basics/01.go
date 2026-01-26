package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)

	// logライブラリを用いてport8081番を用いてhttpリクエストを待ち受ける。
	// 第一引数、ポートを指定、第二引数nilを渡すとhttp.DefaultServeMuxを使用する。
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homeHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"<h1>hello</h1>")
}