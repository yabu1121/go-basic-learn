# ヒント: ユーザー作成API

## 🔍 実装の流れ

### 1. Userモデルの定義

```go
type User struct {
    ID        uint           `gorm:"primaryKey" json:"id"`
    Name      string         `gorm:"size:100;not null" json:"name"`
    Email     string         `gorm:"size:100;unique;not null" json:"email"`
    Age       int            `gorm:"not null" json:"age"`
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
```

**ポイント**:
- `gorm` タグ: データベースの制約
- `json` タグ: JSONのフィールド名
- `DeletedAt`: ソフトデリート用（削除フラグ）

### 2. リクエスト用の構造体

```go
type CreateUserRequest struct {
    Name  string `json:"name" binding:"required"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"required,gte=0"`
}
```

**bindingタグの種類**:
- `required`: 必須
- `email`: メールアドレス形式
- `gte=0`: 0以上
- `min=3`: 最小長3
- `max=100`: 最大長100

### 3. POSTハンドラーの実装

```go
func createUser(c *gin.Context) {
    var req CreateUserRequest
    
    // 1. JSONをバインド
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{
            "status": "error",
            "message": "Validation failed",
            "error": err.Error(),
        })
        return
    }
    
    // 2. モデルを作成
    user := User{
        Name:  req.Name,
        Email: req.Email,
        Age:   req.Age,
    }
    
    // 3. データベースに保存
    if err := DB.Create(&user).Error; err != nil {
        c.JSON(500, gin.H{
            "status": "error",
            "message": "Failed to create user",
            "error": err.Error(),
        })
        return
    }
    
    // 4. 成功レスポンス
    c.JSON(201, gin.H{
        "status": "success",
        "message": "User created successfully",
        "data": user,
    })
}
```

### 4. ルーティング

```go
r.POST("/users", createUser)
```

## 📖 バリデーションの詳細

### よく使うバリデーションタグ

```go
type Example struct {
    // 必須
    Field1 string `binding:"required"`
    
    // メールアドレス
    Field2 string `binding:"email"`
    
    // 数値範囲
    Field3 int `binding:"gte=0,lte=100"`  // 0以上100以下
    
    // 文字列長
    Field4 string `binding:"min=3,max=20"`
    
    // 複数条件
    Field5 string `binding:"required,email"`
}
```

### カスタムエラーメッセージ

```go
if err := c.ShouldBindJSON(&req); err != nil {
    var errors []string
    for _, e := range err.(validator.ValidationErrors) {
        errors = append(errors, fmt.Sprintf("%s is invalid", e.Field()))
    }
    c.JSON(400, gin.H{
        "status": "error",
        "errors": errors,
    })
    return
}
```

## 💡 GORMのCreate操作

### 基本的な使い方

```go
// 単一レコード作成
user := User{Name: "太郎", Email: "taro@example.com"}
result := DB.Create(&user)

// エラーチェック
if result.Error != nil {
    // エラー処理
}

// 作成されたレコードのIDを取得
fmt.Println(user.ID)  // 自動で設定される
```

### 複数レコード作成

```go
users := []User{
    {Name: "太郎", Email: "taro@example.com"},
    {Name: "花子", Email: "hanako@example.com"},
}
DB.Create(&users)
```

### 特定のフィールドのみ保存

```go
DB.Select("Name", "Email").Create(&user)
```

## 🎨 エラーハンドリングのパターン

### パターン1: シンプル

```go
if err := DB.Create(&user).Error; err != nil {
    c.JSON(500, gin.H{"error": err.Error()})
    return
}
```

### パターン2: 詳細

```go
result := DB.Create(&user)
if result.Error != nil {
    // 重複エラーの判定
    if strings.Contains(result.Error.Error(), "duplicate key") {
        c.JSON(409, gin.H{
            "status": "error",
            "message": "Email already exists",
        })
        return
    }
    
    c.JSON(500, gin.H{
        "status": "error",
        "message": "Database error",
        "error": result.Error.Error(),
    })
    return
}
```

## 🚀 チャレンジ課題のヒント

### 課題1: パスワードフィールド

```go
type User struct {
    // ...
    Password string `gorm:"size:255;not null" json:"-"`  // JSONには含めない
}

type CreateUserRequest struct {
    // ...
    Password string `json:"password" binding:"required,min=8"`
}
```

### 課題2: 電話番号バリデーション

```go
type CreateUserRequest struct {
    // ...
    Phone string `json:"phone" binding:"required,e164"`  // E.164形式
    // または
    Phone string `json:"phone" binding:"required,len=11"`  // 11桁
}
```

### 課題3: 日本時間

```go
loc, _ := time.LoadLocation("Asia/Tokyo")
user.CreatedAt = user.CreatedAt.In(loc)
```

### 課題4: 重複チェック

```go
// 既存ユーザーをチェック
var existingUser User
if err := DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
    c.JSON(409, gin.H{
        "status": "error",
        "message": "Email already exists",
    })
    return
}
```

## 🔗 参考リンク

- [Gin Binding](https://gin-gonic.com/docs/examples/binding-and-validation/)
- [GORM Create](https://gorm.io/docs/create.html)
- [Validator](https://github.com/go-playground/validator)
