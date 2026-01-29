package main

import (
	"encoding/json"
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
	http.HandleFunc("/api/user", userHandler)
	http.HandleFunc("/api/users", usersHandler)
	http.HandleFunc("/api/user/create", createHandler)

	// logライブラリを用いてport8081番を用いてhttpリクエストを待ち受ける。
	// 第一引数、ポートを指定、第二引数nilを渡すとhttp.DefaultServeMuxを使用する。
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>hello</h1>")
}

// w: レスポンスを書き込むためのインターフェース
// r: クライアントからのリクエスト情報が格納された構造体。
// w: write, r: read, http.ResponseWriter, http.Request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// ?name=taro の場合、taroを拾えるようにする。コマンド、
	// read, URLはURL, Queryは取得、Getは ～～ = の ~~ を指定している。
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	// この書き方だとドキュメントとして実施される
	fmt.Fprintf(w, "Hello, %s!", name)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(w, "Current Time: %s", now)
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:    1,
		Name:  "太郎",
		Email: "taro@example.com",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{
		{ID: 1, Name: "太郎", Email: "taro@example.com"},
		{ID: 2, Name: "次郎", Email: "jiro@example.com"},
		{ID: 3, Name: "三郎", Email: "saburo@example.com"},
	}

	resp := Response{
		Success: true,
		Message: "ユーザー一覧取得成功",
		Data:    users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}


type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

func createHandler (w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if req.Name == "" || req.Email == "" {
		resp := Response{
			Success: false,
			Message: "名前とメールアドレスは必須です",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	newUser := User{
		ID: 100,
		Name: req.Name,
		Email: req.Email,
	}
	resp := Response{
		Success: true,
		Message: "ユーザー作成成功",
		Data: newUser,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}