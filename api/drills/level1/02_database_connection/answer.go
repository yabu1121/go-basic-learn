package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// グローバルなデータベース接続
var DB *gorm.DB

// データベース接続の初期化
func initDB() {
	// 環境変数の読み込み（チャレンジ課題1）
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default values")
	}

	// 環境変数から接続情報を取得（デフォルト値あり）
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "gouser")
	password := getEnv("DB_PASSWORD", "gopass")
	dbname := getEnv("DB_NAME", "golearn_db")

	// DSN（Data Source Name）の構築
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, user, password, dbname, port,
	)

	// データベース接続
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// ログレベルの設定（開発時は詳細ログを表示）
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// 接続プールの設定（チャレンジ課題2）
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	// 最大オープン接続数
	sqlDB.SetMaxOpenConns(100)

	// 最大アイドル接続数
	sqlDB.SetMaxIdleConns(10)

	// 接続の最大ライフタイム
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connection established successfully")
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
	// *sql.DBインスタンスを取得
	sqlDB, err := DB.DB()
	if err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to get database instance",
			"error":   err.Error(),
		})
		return
	}

	// 接続確認
	if err := sqlDB.Ping(); err != nil {
		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Database connection failed",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"status":   "ok",
		"message":  "Database connection successful",
		"database": getEnv("DB_NAME", "golearn_db"),
	})
}

// データベース統計情報ハンドラー（チャレンジ課題3）
func dbStats(c *gin.Context) {
	sqlDB, err := DB.DB()
	if err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	stats := sqlDB.Stats()
	c.JSON(200, gin.H{
		"status": "ok",
		"stats": gin.H{
			"max_open_connections": stats.MaxOpenConnections,
			"open_connections":     stats.OpenConnections,
			"in_use":               stats.InUse,
			"idle":                 stats.Idle,
			"wait_count":           stats.WaitCount,
			"wait_duration":        stats.WaitDuration.String(),
			"max_idle_closed":      stats.MaxIdleClosed,
			"max_lifetime_closed":  stats.MaxLifetimeClosed,
		},
	})
}

func main() {
	// データベース接続の初期化
	initDB()

	// Ginルーターの作成
	r := gin.Default()

	// ルート定義
	r.GET("/health", healthCheck)
	r.GET("/db/ping", dbPing)
	r.GET("/db/stats", dbStats) // チャレンジ課題3

	// サーバー起動
	port := getEnv("APP_PORT", "8080")
	r.Run(":" + port)
}

/*
【解説】

1. **DSNの構築**
   - PostgreSQLの接続文字列を作成
   - host, user, password, dbname, port を指定
   - sslmode=disable はローカル開発用

2. **gorm.Open()**
   - 第1引数: ドライバー（postgres.Open(dsn)）
   - 第2引数: 設定オプション（&gorm.Config{}）
   - 戻り値: *gorm.DB と error

3. **グローバル変数 DB**
   - アプリケーション全体で使用するため
   - 初期化関数で接続を確立
   - 各ハンドラーから参照可能

4. **DB.DB()**
   - GORMのDBから標準の*sql.DBを取得
   - Ping()やStats()などの操作に必要

5. **接続プール設定**
   - SetMaxOpenConns: 最大同時接続数
   - SetMaxIdleConns: アイドル接続の最大数
   - SetConnMaxLifetime: 接続の最大生存時間

6. **環境変数の活用**
   - godotenvで.envファイルを読み込み
   - os.Getenv()で環境変数を取得
   - デフォルト値を設定して柔軟性を確保

【テスト方法】

# データベース接続確認
curl http://localhost:8080/db/ping

# データベース統計情報
curl http://localhost:8080/db/stats

【期待される出力】

# /db/ping
{
  "status": "ok",
  "message": "Database connection successful",
  "database": "golearn_db"
}

# /db/stats
{
  "status": "ok",
  "stats": {
    "max_open_connections": 100,
    "open_connections": 1,
    "in_use": 0,
    "idle": 1,
    ...
  }
}

【トラブルシューティング】

1. 接続エラーが出る場合
   - PostgreSQLコンテナが起動しているか確認
   - docker-compose ps で確認
   - docker-compose logs postgres でログ確認

2. 認証エラーが出る場合
   - .envファイルの内容を確認
   - docker-compose.ymlの環境変数と一致しているか確認

3. データベースが見つからない場合
   - docker exec -it golearn_postgres psql -U gouser -l
   - でデータベース一覧を確認
*/
