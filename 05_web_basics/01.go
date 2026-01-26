package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// ここでhttpのドメイン後のurlで何の関数を実行するか指定する。
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/time", timeHandler)

	// logライブラリを用いてport8081番を用いてhttpリクエストを待ち受ける。
	// 第一引数、ポートを指定、第二引数nilを渡すとhttp.DefaultServeMuxを使用する。
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homeHandler (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"<h1>hello</h1>")
}

// w: レスポンスを書き込むためのインターフェース
// r: クライアントからのリクエスト情報が格納された構造体。
// w: write, r: read, http.ResponseWriter, http.Request
func helloHandler (w http.ResponseWriter, r * http.Request){
	// ?name=taro の場合、taroを拾えるようにする。コマンド、
	// read, URLはURL, Queryは取得、Getは ～～ = の ~~ を指定している。
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	// この書き方だとドキュメントとして実施される
	fmt.Fprintf(w, "Hello, %s!", name)
}

func timeHandler (w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Current Time: %s", now)
}