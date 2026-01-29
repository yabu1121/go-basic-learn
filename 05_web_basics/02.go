package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	
	http.HandleFunc("/", greet)
	http.HandleFunc("/without-middleware", handleNonMiddleware)
	// with-middlewareは内側の処理から行われていく。
	// アクセス権限を持たせるときに利用をする。
	http.Handle("/with-middleware",
		loggingMiddleware(
			authMiddleware(
				http.HandlerFunc(protectedHandler),
			),
		),
	)

	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func handleNonMiddleware (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ミドルウェアなし")
}