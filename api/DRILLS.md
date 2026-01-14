# 🎯 Gin + GORM ドリル問題集 - 全問題一覧

このファイルには、全30問のドリル問題の詳細な要件が記載されています。

---

## 📚 レベル1: 基礎編

### ドリル01: ヘルスチェックAPI ✅
**ディレクトリ**: `level1/01_health_check/`

**学習内容**: Ginの基本的なルーティング、JSONレスポンス

**要件**:
- GET `/health` エンドポイントを作成
- `{"status": "ok", "message": "API is running"}` を返す

---

### ドリル02: データベース接続 ✅
**ディレクトリ**: `level1/02_database_connection/`

**学習内容**: GORMでPostgreSQLに接続

**要件**:
- PostgreSQLに接続
- GET `/db/ping` でデータベース接続確認
- 環境変数から接続情報を読み込む

---

### ドリル03: モデル定義
**ディレクトリ**: `level1/03_model_definition/`

**学習内容**: GORMのモデル構造体の作成

**要件**:
- Userモデルを定義（ID, Name, Email, Age, CreatedAt, UpdatedAt）
- gormタグとjsonタグを適切に設定
- GET `/models/user` でモデル情報を返す

**ヒント**:
```go
type User struct {
    ID        uint      `gorm:"primaryKey" json:"id"`
    Name      string    `gorm:"size:100;not null" json:"name"`
    Email     string    `gorm:"size:100;unique;not null" json:"email"`
    Age       int       `gorm:"not null" json:"age"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

---

### ドリル04: マイグレーション
**ディレクトリ**: `level1/04_migration/`

**学習内容**: AutoMigrateでテーブル作成

**要件**:
- `DB.AutoMigrate(&User{})` でテーブル作成
- GET `/migrate` でマイグレーション実行
- マイグレーション結果を返す

**ヒント**:
```go
if err := DB.AutoMigrate(&User{}); err != nil {
    // エラー処理
}
```

---

## 📚 レベル2: CRUD基本編

### ドリル05: ユーザー作成API ✅
**ディレクトリ**: `level2/05_create_user/`

**学習内容**: POST リクエストとCreate操作

**要件**:
- POST `/users` でユーザー作成
- バリデーション（required, email, gte=0）
- 重複メールアドレスのチェック
- 201 Created を返す

---

### ドリル06: ユーザー一覧取得
**ディレクトリ**: `level2/06_user_list/`

**学習内容**: GET リクエストとFind操作

**要件**:
- GET `/users` でユーザー一覧を取得
- 全ユーザーを配列で返す
- 空の場合は空配列を返す

**ヒント**:
```go
var users []User
DB.Find(&users)
c.JSON(200, gin.H{"data": users})
```

---

### ドリル07: ユーザー詳細取得
**ディレクトリ**: `level2/07_user_detail/`

**学習内容**: パスパラメータとFirst操作

**要件**:
- GET `/users/:id` でユーザー詳細を取得
- 存在しない場合は404を返す

**ヒント**:
```go
id := c.Param("id")
var user User
if err := DB.First(&user, id).Error; err != nil {
    c.JSON(404, gin.H{"error": "User not found"})
    return
}
```

---

### ドリル08: ユーザー更新API
**ディレクトリ**: `level2/08_update_user/`

**学習内容**: PUT リクエストとUpdate操作

**要件**:
- PUT `/users/:id` でユーザー情報を更新
- 部分更新に対応
- 存在しない場合は404を返す

**ヒント**:
```go
var user User
if err := DB.First(&user, id).Error; err != nil {
    c.JSON(404, gin.H{"error": "User not found"})
    return
}
DB.Model(&user).Updates(req)
```

---

### ドリル09: ユーザー削除API
**ディレクトリ**: `level2/09_delete_user/`

**学習内容**: DELETE リクエストとDelete操作

**要件**:
- DELETE `/users/:id` でユーザーを削除
- ソフトデリート（DeletedAt）を使用
- 204 No Content を返す

**ヒント**:
```go
if err := DB.Delete(&User{}, id).Error; err != nil {
    c.JSON(500, gin.H{"error": err.Error()})
    return
}
c.Status(204)
```

---

## 📚 レベル3: 中級編

### ドリル10: バリデーション
**ディレクトリ**: `level3/10_validation/`

**学習内容**: binding タグでの入力検証

**要件**:
- 複雑なバリデーションルールを実装
- カスタムバリデーターを作成
- エラーメッセージをカスタマイズ

**バリデーション例**:
- 名前: 必須、1-100文字
- メール: 必須、メール形式
- 年齢: 0-150
- パスワード: 必須、8文字以上、英数字含む

---

### ドリル11: エラーハンドリング
**ディレクトリ**: `level3/11_error_handling/`

**学習内容**: カスタムエラーレスポンス

**要件**:
- エラーハンドリングミドルウェアを作成
- 統一されたエラーレスポンス形式
- エラーログの記録

**エラーレスポンス形式**:
```json
{
  "status": "error",
  "code": "USER_NOT_FOUND",
  "message": "User not found",
  "timestamp": "2024-01-14T16:35:00+09:00"
}
```

---

### ドリル12: ページネーション
**ディレクトリ**: `level3/12_pagination/`

**学習内容**: Limit/Offsetでのページング

**要件**:
- GET `/users?page=1&limit=10` でページング
- デフォルト: page=1, limit=20
- 総件数、総ページ数を返す

**レスポンス例**:
```json
{
  "data": [...],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 100,
    "total_pages": 10
  }
}
```

**ヒント**:
```go
page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
offset := (page - 1) * limit

var users []User
var total int64
DB.Model(&User{}).Count(&total)
DB.Limit(limit).Offset(offset).Find(&users)
```

---

### ドリル13: 検索機能
**ディレクトリ**: `level3/13_search/`

**学習内容**: Where句での条件検索

**要件**:
- GET `/users/search?name=太郎&age_min=20&age_max=30`
- 複数条件での検索
- 部分一致検索（LIKE）

**ヒント**:
```go
query := DB.Model(&User{})
if name := c.Query("name"); name != "" {
    query = query.Where("name LIKE ?", "%"+name+"%")
}
if ageMin := c.Query("age_min"); ageMin != "" {
    query = query.Where("age >= ?", ageMin)
}
```

---

### ドリル14: ソート機能
**ディレクトリ**: `level3/14_sorting/`

**学習内容**: Order句での並び替え

**要件**:
- GET `/users?sort=name&order=asc`
- 複数フィールドでのソート
- デフォルト: created_at desc

**ヒント**:
```go
sort := c.DefaultQuery("sort", "created_at")
order := c.DefaultQuery("order", "desc")
DB.Order(sort + " " + order).Find(&users)
```

---

## 📚 レベル4: リレーション編

### ドリル15: 1対多リレーション
**ディレクトリ**: `level4/15_one_to_many/`

**学習内容**: HasMany/BelongsTo

**要件**:
- Userモデルに複数のPostsを関連付け
- POST `/users/:id/posts` で投稿作成
- GET `/users/:id/posts` で投稿一覧取得

**モデル定義**:
```go
type User struct {
    ID    uint
    Name  string
    Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
    ID      uint
    Title   string
    Content string
    UserID  uint
    User    User `gorm:"foreignKey:UserID"`
}
```

---

### ドリル16: 多対多リレーション
**ディレクトリ**: `level4/16_many_to_many/`

**学習内容**: Many2Many

**要件**:
- UserとRoleの多対多リレーション
- POST `/users/:id/roles` でロール追加
- DELETE `/users/:id/roles/:role_id` でロール削除

**モデル定義**:
```go
type User struct {
    ID    uint
    Name  string
    Roles []Role `gorm:"many2many:user_roles;"`
}

type Role struct {
    ID    uint
    Name  string
    Users []User `gorm:"many2many:user_roles;"`
}
```

---

### ドリル17: Preload
**ディレクトリ**: `level4/17_preload/`

**学習内容**: 関連データの事前読み込み

**要件**:
- GET `/users/:id?include=posts,roles` で関連データを含める
- N+1問題を回避

**ヒント**:
```go
query := DB.Model(&User{})
if strings.Contains(c.Query("include"), "posts") {
    query = query.Preload("Posts")
}
if strings.Contains(c.Query("include"), "roles") {
    query = query.Preload("Roles")
}
query.First(&user, id)
```

---

### ドリル18: Join
**ディレクトリ**: `level4/18_join/`

**学習内容**: テーブル結合クエリ

**要件**:
- GET `/posts?user_name=太郎` でユーザー名で投稿検索
- Joinを使った効率的なクエリ

**ヒント**:
```go
DB.Joins("JOIN users ON users.id = posts.user_id").
   Where("users.name LIKE ?", "%"+userName+"%").
   Find(&posts)
```

---

### ドリル19: カスケード削除
**ディレクトリ**: `level4/19_cascade_delete/`

**学習内容**: OnDelete制約

**要件**:
- ユーザー削除時に関連する投稿も削除
- トランザクションを使用

**ヒント**:
```go
DB.Transaction(func(tx *gorm.DB) error {
    if err := tx.Where("user_id = ?", userID).Delete(&Post{}).Error; err != nil {
        return err
    }
    if err := tx.Delete(&User{}, userID).Error; err != nil {
        return err
    }
    return nil
})
```

---

## 📚 レベル5: 応用編

### ドリル20: 認証API
**ディレクトリ**: `level5/20_authentication/`

**学習内容**: JWT認証の実装

**要件**:
- POST `/auth/register` でユーザー登録
- POST `/auth/login` でログイン（JWTトークン発行）
- パスワードのハッシュ化（bcrypt）

**必要なパッケージ**:
```bash
go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
```

---

### ドリル21: ミドルウェア
**ディレクトリ**: `level5/21_middleware/`

**学習内容**: 認証ミドルウェア

**要件**:
- JWTトークンの検証ミドルウェア
- 保護されたエンドポイントの作成
- GET `/me` で現在のユーザー情報取得

**ヒント**:
```go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        // トークン検証
        c.Next()
    }
}

r.GET("/me", AuthMiddleware(), getCurrentUser)
```

---

### ドリル22: ファイルアップロード
**ディレクトリ**: `level5/22_file_upload/`

**学習内容**: 画像アップロード処理

**要件**:
- POST `/upload` で画像アップロード
- ファイルサイズ制限（5MB）
- 画像形式チェック（jpg, png, gif）
- ファイル名の生成（UUID）

**ヒント**:
```go
file, _ := c.FormFile("image")
filename := uuid.New().String() + filepath.Ext(file.Filename)
c.SaveUploadedFile(file, "./uploads/"+filename)
```

---

### ドリル23: トランザクション
**ディレクトリ**: `level5/23_transaction/`

**学習内容**: DB.Transaction

**要件**:
- 銀行振込のシミュレーション
- POST `/transfer` で送金処理
- ロールバック処理

**ヒント**:
```go
DB.Transaction(func(tx *gorm.DB) error {
    // 送金元の残高を減らす
    if err := tx.Model(&account1).Update("balance", balance1-amount).Error; err != nil {
        return err
    }
    // 送金先の残高を増やす
    if err := tx.Model(&account2).Update("balance", balance2+amount).Error; err != nil {
        return err
    }
    return nil
})
```

---

### ドリル24: 集計クエリ
**ディレクトリ**: `level5/24_aggregation/`

**学習内容**: Count/Sum/Avg

**要件**:
- GET `/stats/users` でユーザー統計
- 総ユーザー数、平均年齢、年齢分布

**ヒント**:
```go
var total int64
var avgAge float64
DB.Model(&User{}).Count(&total)
DB.Model(&User{}).Select("AVG(age)").Scan(&avgAge)
```

---

## 📚 レベル6: 実践編

### ドリル25: ブログAPI
**ディレクトリ**: `level6/25_blog_api/`

**学習内容**: 記事のCRUD

**要件**:
- 記事の作成、取得、更新、削除
- マークダウン対応
- 下書き/公開ステータス

---

### ドリル26: コメント機能
**ディレクトリ**: `level6/26_comments/`

**学習内容**: ネストされたリソース

**要件**:
- POST `/posts/:id/comments` でコメント作成
- GET `/posts/:id/comments` でコメント一覧
- ネストされたコメント（返信機能）

---

### ドリル27: いいね機能
**ディレクトリ**: `level6/27_likes/`

**学習内容**: 中間テーブルの活用

**要件**:
- POST `/posts/:id/like` でいいね
- DELETE `/posts/:id/like` でいいね解除
- GET `/posts/:id/likes` でいいね一覧

---

### ドリル28: タグ機能
**ディレクトリ**: `level6/28_tags/`

**学習内容**: 多対多の実装

**要件**:
- 記事にタグを追加
- GET `/posts?tags=Go,API` でタグ検索
- GET `/tags` で人気タグ一覧

---

### ドリル29: 全文検索
**ディレクトリ**: `level6/29_full_text_search/`

**学習内容**: LIKE検索の最適化

**要件**:
- GET `/search?q=キーワード` で全文検索
- タイトル、本文、タグから検索
- 検索結果のハイライト

---

### ドリル30: 総合演習（ECサイトAPI）
**ディレクトリ**: `level6/30_ecommerce/`

**学習内容**: これまでの総まとめ

**要件**:
- 商品管理（CRUD）
- カート機能
- 注文処理
- 在庫管理
- ユーザー認証
- 注文履歴

**モデル構成**:
- User（ユーザー）
- Product（商品）
- Category（カテゴリ）
- Cart（カート）
- CartItem（カート商品）
- Order（注文）
- OrderItem（注文商品）

---

## 🎓 学習の進め方

1. **順番に進める**: レベル1から順番に進めることを推奨
2. **自分で実装**: まず自分で実装してみる
3. **ヒントを見る**: 詰まったらhint.mdを参照
4. **解答を確認**: 実装後、answer.goと比較
5. **テストする**: curlコマンドで動作確認
6. **チャレンジ**: 余裕があればチャレンジ課題に挑戦

## 📝 各ドリルの構成

各ドリルディレクトリには以下のファイルがあります:

- `README.md` - 問題文と要件
- `hint.md` - ヒントと実装のポイント
- `answer.go` - 解答例（完全動作するコード）

## 🔗 参考リソース

- [Gin公式ドキュメント](https://gin-gonic.com/docs/)
- [GORM公式ドキュメント](https://gorm.io/docs/)
- [Go by Example](https://gobyexample.com/)
- [PostgreSQL公式ドキュメント](https://www.postgresql.org/docs/)

頑張ってください！🚀
