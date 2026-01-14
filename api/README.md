# 🎯 Gin + GORM API開発ドリル

このディレクトリでは、**Gin**（Webフレームワーク）と**GORM**（ORM）を使用したREST API開発を段階的に学習できます。

## 📋 目次

1. [環境構築](#環境構築)
2. [ドリル問題一覧](#ドリル問題一覧)
3. [学習の進め方](#学習の進め方)
4. [ヒント集](#ヒント集)

---

## 🚀 環境構築

### 前提条件
- Docker Desktop がインストールされていること
- Goがインストールされていること（バージョン1.21以上）

### セットアップ手順

```bash
# 1. このディレクトリに移動
cd api

# 2. 環境変数ファイルをコピー
cp .env.example .env

# 3. Dockerコンテナを起動
docker-compose up -d

# 4. Go modulesの依存関係をダウンロード
go mod download

# 5. アプリケーションを実行
go run main.go
```

### 動作確認

```bash
# ヘルスチェック
curl http://localhost:8080/health

# PostgreSQL接続確認
docker exec -it golearn_postgres psql -U gouser -d golearn_db
```

---

## 📚 ドリル問題一覧

### レベル1: 基礎編（drills/level1/）

| 問題 | タイトル | 学習内容 |
|------|----------|----------|
| 01 | ヘルスチェックAPI | Ginの基本的なルーティング |
| 02 | データベース接続 | GORMでPostgreSQLに接続 |
| 03 | モデル定義 | GORMのモデル構造体の作成 |
| 04 | マイグレーション | AutoMigrateでテーブル作成 |

### レベル2: CRUD基本編（drills/level2/）

| 問題 | タイトル | 学習内容 |
|------|----------|----------|
| 05 | ユーザー作成API | POST リクエストとCreate操作 |
| 06 | ユーザー一覧取得 | GET リクエストとFind操作 |
| 07 | ユーザー詳細取得 | パスパラメータとFirst操作 |
| 08 | ユーザー更新API | PUT リクエストとUpdate操作 |
| 09 | ユーザー削除API | DELETE リクエストとDelete操作 |

### レベル3: 中級編（drills/level3/）

| 問題 | タイトル | 学習内容 |
|------|----------|----------|
| 10 | バリデーション | binding タグでの入力検証 |
| 11 | エラーハンドリング | カスタムエラーレスポンス |
| 12 | ページネーション | Limit/Offsetでのページング |
| 13 | 検索機能 | Where句での条件検索 |
| 14 | ソート機能 | Order句での並び替え |

### レベル4: リレーション編（drills/level4/）

| 問題 | タイトル | 学習内容 |
|------|----------|----------|
| 15 | 1対多リレーション | HasMany/BelongsTo |
| 16 | 多対多リレーション | Many2Many |
| 17 | Preload | 関連データの事前読み込み |
| 18 | Join | テーブル結合クエリ |
| 19 | カスケード削除 | OnDelete制約 |

### レベル5: 応用編（drills/level5/）

| 問題 | タイトル | 学習内容 |
|------|----------|----------|
| 20 | 認証API | JWT認証の実装 |
| 21 | ミドルウェア | 認証ミドルウェア |
| 22 | ファイルアップロード | 画像アップロード処理 |
| 23 | トランザクション | DB.Transaction |
| 24 | 集計クエリ | Count/Sum/Avg |

### レベル6: 実践編（drills/level6/）

| 問題 | タイトル | 学習内容 |
|------|----------|----------|
| 25 | ブログAPI | 記事のCRUD |
| 26 | コメント機能 | ネストされたリソース |
| 27 | いいね機能 | 中間テーブルの活用 |
| 28 | タグ機能 | 多対多の実装 |
| 29 | 全文検索 | LIKE検索の最適化 |
| 30 | 総合演習 | ECサイトAPI |

---

## 🎓 学習の進め方

### 1. 問題を読む
各ドリルフォルダには以下のファイルがあります:
- `README.md` - 問題文と要件
- `hint.md` - ヒントと参考情報
- `answer.go` - 解答例（見る前に自分で実装してみましょう！）

### 2. 自分で実装する
`main.go` または指定されたファイルに実装してください。

### 3. テストする
各問題には動作確認用のcurlコマンドが用意されています。

### 4. 解答を確認する
実装後、`answer.go`と比較して学びを深めましょう。

---

## 💡 ヒント集

### Ginの基本パターン

```go
// ルーティング
r := gin.Default()
r.GET("/path", handlerFunc)
r.POST("/path", handlerFunc)
r.PUT("/path/:id", handlerFunc)
r.DELETE("/path/:id", handlerFunc)

// パラメータ取得
id := c.Param("id")           // パスパラメータ
name := c.Query("name")       // クエリパラメータ
var input Input
c.ShouldBindJSON(&input)      // JSONボディ

// レスポンス
c.JSON(200, gin.H{"message": "success"})
c.JSON(400, gin.H{"error": "bad request"})
```

### GORMの基本パターン

```go
// 接続
dsn := "host=localhost user=gouser password=gopass dbname=golearn_db port=5432"
db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// マイグレーション
db.AutoMigrate(&User{})

// Create
db.Create(&user)

// Read
db.Find(&users)                    // 全件取得
db.First(&user, id)                // IDで取得
db.Where("name = ?", name).Find(&users)  // 条件検索

// Update
db.Model(&user).Updates(User{Name: "新しい名前"})

// Delete
db.Delete(&user, id)
```

### モデル定義のパターン

```go
type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `gorm:"size:100;not null" json:"name" binding:"required"`
    Email     string         `gorm:"size:100;unique;not null" json:"email" binding:"required,email"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

---

## 🔧 便利なコマンド

```bash
# Dockerコンテナの状態確認
docker-compose ps

# ログ確認
docker-compose logs -f app

# データベースに接続
docker exec -it golearn_postgres psql -U gouser -d golearn_db

# コンテナの再起動
docker-compose restart

# コンテナの停止と削除
docker-compose down

# データベースも含めて完全削除
docker-compose down -v
```

---

## 📖 参考リンク

- [Gin公式ドキュメント](https://gin-gonic.com/docs/)
- [GORM公式ドキュメント](https://gorm.io/docs/)
- [PostgreSQL公式ドキュメント](https://www.postgresql.org/docs/)

---

## 🎯 次のステップ

1. まずは `drills/level1/01_health_check/` から始めましょう
2. 各レベルを順番に進めることをおすすめします
3. わからない場合は `hint.md` を参照してください
4. 実装後は必ず動作確認をしましょう

頑張ってください！🚀
