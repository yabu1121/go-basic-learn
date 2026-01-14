package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

// ヘルスチェックハンドラー
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "API is running",
	})
}

// ウェルカムメッセージハンドラー（チャレンジ課題1）
func welcome(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Welcome to Gin + GORM API Learning!",
		"time":    time.Now().Format(time.RFC3339),
	})
}

// バージョン情報ハンドラー（チャレンジ課題3）
func version(c *gin.Context) {
	c.JSON(200, gin.H{
		"version": "1.0.0",
		"go":      "1.21",
		"gin":     gin.Version,
	})
}

func main() {
	// Ginルーターを作成（Logger, Recoveryミドルウェア付き）
	r := gin.Default()

	// ルート定義
	r.GET("/health", healthCheck)
	r.GET("/", welcome)
	r.GET("/version", version)

	// サーバー起動
	r.Run(":8080")
}

/*
【解説】

1. **gin.Default() vs gin.New()**
   - Default(): Logger と Recovery ミドルウェアが自動で設定される
   - New(): ミドルウェアなしのルーター

2. **ハンドラー関数**
   - 引数は必ず `*gin.Context`
   - c.JSON() でJSONレスポンスを返す
   - 第1引数はHTTPステータスコード

3. **gin.H**
   - map[string]interface{} のエイリアス
   - JSONレスポンスを簡潔に書ける

4. **ルーティング**
   - r.GET(path, handler) でGETリクエストを処理
   - 他にも POST, PUT, DELETE, PATCH などがある

5. **サーバー起動**
   - r.Run(":8080") でポート8080で起動
   - 引数省略時は :8080 がデフォルト

【テスト方法】

# 基本
curl http://localhost:8080/health

# チャレンジ課題
curl http://localhost:8080/
curl http://localhost:8080/version

【期待される出力】

# /health
{"status":"ok","message":"API is running"}

# /
{"message":"Welcome to Gin + GORM API Learning!","time":"2024-01-14T16:35:00+09:00"}

# /version
{"version":"1.0.0","go":"1.21","gin":"v1.9.1"}
*/
