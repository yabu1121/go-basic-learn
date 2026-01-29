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
	http.HandleFunc("/api/user/search", searchHandler)

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


// request(postメソッドで送信するオブジェクトの型を指定する)
type CreateUserRequest struct {
	Name string `json:"name"`
	Email string `json:"email"`
}

// createhandle関数を宣言これはhttp, handlefuncで呼び出す関数。
// w: クライアントへ送るレスポンスを書き込むための道具
// r: クライアントから届いたリクエストの情報が入っている。
func createHandler (w http.ResponseWriter, r *http.Request) {
	// r.Methodは普通にrequestのメソッドがある、それを指定できるのがhttp.methodPostなど
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// var req CreateUserRequestというのはとりあえず宣言しておく。
	// さいしょはから　。
	// jsonのbodyを読み取ってgoの構造体reqに変換して流し込む。
	// defer: その関数が終了する直前に、特定の処理を必ず実行させる
	var req CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// 慣習的にはリソースを弾いたらすぐに書くのが一般的
	defer r.Body.Close()

	//　バリデーション 
	if req.Name == "" || req.Email == "" {
		resp := Response{
			Success: false,
			Message: "名前とメールアドレスは必須です",
		}

		// headerに記述する
		// Content-Type:application/jsonである。
		// つまりjson形式で記述されていることをヘッダーに示す。
		// 最悪これがなくても一応処理はできる。
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(resp)
		return
	}

	// 
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


func searchHandler (w http.ResponseWriter, r *http.Request) {
	// クエリパラメータを用いたデータ取得
	query := r.URL.Query().Get("q")
	page := r.URL.Query().Get("page")

	if query == "" {
		http.Error(w, "検索キーワードが必要です", http.StatusBadRequest)
		return
	}

	if page == "" {
		page = "1"
	}

	resp := Response{
		Success: true,
		Message: "検索成功",
		Data: map[string]string{
			"query": query,
			"page": page,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}