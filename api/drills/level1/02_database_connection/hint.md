# ヒント: データベース接続

## 🔍 実装の流れ

### 1. 必要なパッケージのインポート

```go
import (
    "fmt"
    "log"
    
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)
```

### 2. DSN（Data Source Name）の構築

PostgreSQLの接続文字列を作成します。

```go
dsn := "host=localhost user=gouser password=gopass dbname=golearn_db port=5432 sslmode=disable"
```

**フォーマット**:
```
host=ホスト user=ユーザー名 password=パスワード dbname=データベース名 port=ポート sslmode=モード
```

**ポイント**:
- `sslmode=disable` はローカル開発用（本番では `require` を使用）
- スペース区切りで各パラメータを指定

### 3. データベース接続

```go
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
if err != nil {
    log.Fatal("Failed to connect to database:", err)
}
```

**ポイント**:
- `gorm.Open()` の第1引数はドライバー
- 第2引数は設定オプション
- エラーハンドリングは必須

### 4. グローバル変数での管理

```go
var DB *gorm.DB

func initDB() {
    var err error
    dsn := "..."
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    initDB()
    // ...
}
```

**ポイント**:
- グローバル変数 `DB` を定義
- 初期化関数で接続を確立
- `main()` の最初で呼び出す

### 5. 接続確認エンドポイント

```go
func dbPing(c *gin.Context) {
    sqlDB, err := DB.DB()
    if err != nil {
        c.JSON(500, gin.H{
            "status": "error",
            "error": err.Error(),
        })
        return
    }
    
    if err := sqlDB.Ping(); err != nil {
        c.JSON(500, gin.H{
            "status": "error",
            "error": err.Error(),
        })
        return
    }
    
    c.JSON(200, gin.H{
        "status": "ok",
        "message": "Database connection successful",
    })
}
```

**ポイント**:
- `DB.DB()` で `*sql.DB` を取得
- `Ping()` で接続確認
- エラー時は500ステータスを返す

## 📖 よくあるエラーと対処法

### エラー1: `dial tcp [::1]:5432: connect: connection refused`

**原因**: PostgreSQLが起動していない

**対処法**:
```bash
docker-compose up -d postgres
docker-compose ps
```

### エラー2: `password authentication failed`

**原因**: 認証情報が間違っている

**対処法**:
- DSNの user, password を確認
- docker-compose.yml の環境変数を確認

### エラー3: `database "golearn_db" does not exist`

**原因**: データベースが作成されていない

**対処法**:
```bash
docker exec -it golearn_postgres psql -U gouser -c "CREATE DATABASE golearn_db;"
```

## 🎨 環境変数を使う方法（チャレンジ課題1）

### 1. godotenvのインストール

```bash
go get github.com/joho/godotenv
```

### 2. .envファイルの作成

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=gouser
DB_PASSWORD=gopass
DB_NAME=golearn_db
```

### 3. 環境変数の読み込み

```go
import (
    "os"
    "github.com/joho/godotenv"
)

func initDB() {
    // .envファイルを読み込む
    godotenv.Load()
    
    // 環境変数から取得
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        host, user, password, dbname, port,
    )
    
    // ...
}
```

## 💡 接続プールの設定（チャレンジ課題2）

```go
sqlDB, err := DB.DB()
if err != nil {
    log.Fatal(err)
}

// 最大オープン接続数
sqlDB.SetMaxOpenConns(100)

// 最大アイドル接続数
sqlDB.SetMaxIdleConns(10)

// 接続の最大ライフタイム
sqlDB.SetConnMaxLifetime(time.Hour)
```

## 🔗 参考リンク

- [GORM - Connecting to Database](https://gorm.io/docs/connecting_to_the_database.html)
- [PostgreSQL Driver](https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL)
- [godotenv](https://github.com/joho/godotenv)
