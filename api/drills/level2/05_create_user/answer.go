package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// グローバルなデータベース接続
var DB *gorm.DB

// Userモデル
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Email     string         `gorm:"size:100;unique;not null" json:"email"`
	Age       int            `gorm:"not null" json:"age"`
	Password  string         `gorm:"size:255;not null" json:"-"` // チャレンジ課題1
	Phone     string         `gorm:"size:20" json:"phone"`       // チャレンジ課題2
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// ユーザー作成リクエスト
type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=1,max=100"`
	Email    string `json:"email" binding:"required,email"`
	Age      int    `json:"age" binding:"required,gte=0,lte=150"`
	Password string `json:"password" binding:"required,min=8"` // チャレンジ課題1
	Phone    string `json:"phone" binding:"omitempty,len=11"`  // チャレンジ課題2
}

// データベース接続の初期化
func initDB() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "gouser")
	password := getEnv("DB_PASSWORD", "gopass")
	dbname := getEnv("DB_NAME", "golearn_db")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		host, user, password, dbname, port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// マイグレーション
	if err := DB.AutoMigrate(&User{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database connection established and migrated successfully")
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// ユーザー作成ハンドラー
func createUser(c *gin.Context) {
	var req CreateUserRequest

	// JSONバインディングとバリデーション
	if err := c.ShouldBindJSON(&req); err != nil {
		// バリデーションエラーの詳細を取得
		var errors []string
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				errors = append(errors, formatValidationError(e))
			}
		} else {
			errors = append(errors, err.Error())
		}

		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Validation failed",
			"errors":  errors,
		})
		return
	}

	// メールアドレスの重複チェック（チャレンジ課題4）
	var existingUser User
	if err := DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		c.JSON(409, gin.H{
			"status":  "error",
			"message": "Email already exists",
		})
		return
	}

	// Userモデルの作成
	user := User{
		Name:     req.Name,
		Email:    req.Email,
		Age:      req.Age,
		Password: req.Password, // 本番ではハッシュ化が必要
		Phone:    req.Phone,
	}

	// データベースに保存
	if err := DB.Create(&user).Error; err != nil {
		// 重複エラーのチェック
		if strings.Contains(err.Error(), "duplicate key") {
			c.JSON(409, gin.H{
				"status":  "error",
				"message": "Email already exists",
			})
			return
		}

		c.JSON(500, gin.H{
			"status":  "error",
			"message": "Failed to create user",
			"error":   err.Error(),
		})
		return
	}

	// 日本時間に変換（チャレンジ課題3）
	loc, _ := time.LoadLocation("Asia/Tokyo")
	user.CreatedAt = user.CreatedAt.In(loc)
	user.UpdatedAt = user.UpdatedAt.In(loc)

	// 成功レスポンス
	c.JSON(201, gin.H{
		"status":  "success",
		"message": "User created successfully",
		"data":    user,
	})
}

// バリデーションエラーのフォーマット
func formatValidationError(e validator.FieldError) string {
	field := e.Field()
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email address", field)
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", field, e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", field, e.Param())
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to %s", field, e.Param())
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to %s", field, e.Param())
	case "len":
		return fmt.Sprintf("%s must be exactly %s characters", field, e.Param())
	default:
		return fmt.Sprintf("%s is invalid", field)
	}
}

// ヘルスチェックハンドラー
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "API is running",
	})
}

func main() {
	// データベース接続
	initDB()

	// Ginルーター
	r := gin.Default()

	// ルート定義
	r.GET("/health", healthCheck)
	r.POST("/users", createUser)

	// サーバー起動
	port := getEnv("APP_PORT", "8080")
	log.Printf("Server starting on port %s...", port)
	r.Run(":" + port)
}

/*
【解説】

1. **Userモデル**
   - gormタグでデータベース制約を定義
   - jsonタグでJSONフィールド名を指定
   - json:"-" でJSONに含めない（パスワード）
   - DeletedAt でソフトデリート対応

2. **CreateUserRequest**
   - bindingタグでバリデーションルールを定義
   - required: 必須
   - email: メールアドレス形式
   - gte/lte: 数値範囲
   - min/max: 文字列長
   - omitempty: 省略可能

3. **c.ShouldBindJSON()**
   - JSONをGoの構造体にバインド
   - バリデーションも同時に実行
   - エラー時は400を返す

4. **DB.Create()**
   - GORMでレコードを作成
   - 成功時はIDが自動設定される
   - CreatedAt/UpdatedAtも自動設定

5. **エラーハンドリング**
   - バリデーションエラー: 400
   - 重複エラー: 409
   - データベースエラー: 500

6. **バリデーションエラーのフォーマット**
   - validator.ValidationErrors を解析
   - ユーザーフレンドリーなメッセージに変換

7. **重複チェック**
   - DB.Where().First() で既存レコード検索
   - err == nil なら既に存在する
   - 409 Conflict を返す

【テスト方法】

# 正常なユーザー作成
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "山田太郎",
    "email": "yamada@example.com",
    "age": 25,
    "password": "password123",
    "phone": "09012345678"
  }'

# バリデーションエラー（名前が空）
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "",
    "email": "test@example.com",
    "age": 25,
    "password": "password123"
  }'

# バリデーションエラー（メール形式が不正）
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "佐藤花子",
    "email": "invalid-email",
    "age": 30,
    "password": "password123"
  }'

# バリデーションエラー（年齢が範囲外）
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "鈴木一郎",
    "email": "suzuki@example.com",
    "age": -5,
    "password": "password123"
  }'

# 重複エラー（同じメールアドレス）
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "田中次郎",
    "email": "yamada@example.com",
    "age": 28,
    "password": "password123"
  }'

【期待される出力】

# 成功時
{
  "status": "success",
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "山田太郎",
    "email": "yamada@example.com",
    "age": 25,
    "phone": "09012345678",
    "created_at": "2024-01-14T16:35:00+09:00",
    "updated_at": "2024-01-14T16:35:00+09:00"
  }
}

# バリデーションエラー時
{
  "status": "error",
  "message": "Validation failed",
  "errors": [
    "Name is required",
    "Email must be a valid email address"
  ]
}

# 重複エラー時
{
  "status": "error",
  "message": "Email already exists"
}
*/
