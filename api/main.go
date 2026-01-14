package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// グローバルなデータベース接続
var DB *gorm.DB

// データベース接続の初期化
func initDB() {
	// 環境変数の読み込み
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	// 接続情報
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "gouser")
	password := getEnv("DB_PASSWORD", "gopass")
	dbname := getEnv("DB_NAME", "golearn_db")

	// DSN構築
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, user, password, dbname, port,
	)

	// データベース接続
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("✅ Database connection established successfully")
}

// 環境変数を取得（デフォルト値付き）
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// ヘルスチェックハンドラー
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "API is running",
	})
}

// データベース接続確認ハンドラー
func dbPing(c *gin.Context) {
	sqlDB, err := DB.DB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	if err := sqlDB.Ping(); err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":   "ok",
		"message":  "Database connection successful",
		"database": getEnv("DB_NAME", "golearn_db"),
	})
}

func main() {
	// データベース接続
	initDB()

	// Ginルーター
	r := gin.Default()

	// ルート定義
	r.GET("/health", healthCheck)
	r.GET("/db/ping", dbPing)

	// TODO: ここにあなたのAPIエンドポイントを追加してください
	// 例:
	// r.POST("/users", createUser)
	// r.GET("/users", getUsers)
	// r.GET("/users/:id", getUser)
	// r.PUT("/users/:id", updateUser)
	// r.DELETE("/users/:id", deleteUser)

	// サーバー起動
	port := getEnv("APP_PORT", "8080")
	log.Printf("🚀 Server starting on port %s...", port)
	log.Printf("📝 Try: curl http://localhost:%s/health", port)
	r.Run(":" + port)
}

/*
【このファイルについて】

これはGin + GORMのスターターテンプレートです。
基本的なデータベース接続とヘルスチェックAPIが実装されています。

【使い方】

1. 環境構築
   docker-compose up -d postgres
   cp .env.example .env

2. 実行
   go run main.go

3. テスト
   curl http://localhost:8080/health
   curl http://localhost:8080/db/ping

4. 開発
   main()関数にあなたのAPIエンドポイントを追加してください

【学習の進め方】

1. drills/level1/ から順番に進める
2. 各ドリルのREADME.mdを読む
3. 自分で実装してみる
4. hint.mdを参考にする
5. answer.goで答え合わせ

詳しくは README.md と QUICKSTART.md を参照してください。

Happy Coding! 🚀
*/
