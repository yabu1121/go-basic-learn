# ヒント: ヘルスチェックAPI

## 🔍 実装の流れ

### 1. 必要なパッケージのインポート

```go
import (
    "github.com/gin-gonic/gin"
)
```

### 2. Ginルーターの作成

Ginでは `gin.Default()` を使うと、Logger と Recovery ミドルウェアが自動で設定されます。

```go
r := gin.Default()
```

**ポイント**: 
- `gin.New()` を使うとミドルウェアなしのルーターが作成されます
- 開発時は `gin.Default()` が便利です

### 3. ルートの定義

GETリクエストを処理するには `r.GET()` を使います。

```go
r.GET("/path", func(c *gin.Context) {
    // ハンドラー関数
})
```

**ポイント**:
- 第1引数: パス（文字列）
- 第2引数: ハンドラー関数
- `c *gin.Context` は現在のリクエストコンテキスト

### 4. JSONレスポンスの返却

`c.JSON()` メソッドを使います。

```go
c.JSON(200, gin.H{
    "key": "value",
})
```

**ポイント**:
- 第1引数: HTTPステータスコード
- 第2引数: レスポンスボディ（構造体またはgin.H）
- `gin.H` は `map[string]interface{}` のエイリアス

### 5. サーバーの起動

```go
r.Run(":8080")
```

**ポイント**:
- デフォルトは `:8080`
- 引数を省略すると `:8080` で起動します

## 📖 よくあるエラーと対処法

### エラー1: `package gin is not in GOROOT`

**原因**: Ginがインストールされていない

**対処法**:
```bash
go mod download
# または
go get github.com/gin-gonic/gin
```

### エラー2: `address already in use`

**原因**: ポート8080が既に使用されている

**対処法**:
```bash
# Windowsの場合
netstat -ano | findstr :8080
taskkill /PID <プロセスID> /F

# または別のポートを使用
r.Run(":8081")
```

## 🎨 コードの構造例

```go
package main

import (
    // インポート
)

func main() {
    // 1. ルーター作成
    
    // 2. ルート定義
    
    // 3. サーバー起動
}
```

## 💡 追加のヒント

### gin.H の使い方

```go
// これらは同じ意味です
gin.H{"key": "value"}
map[string]interface{}{"key": "value"}
```

### ハンドラー関数を分離する

```go
func healthCheck(c *gin.Context) {
    c.JSON(200, gin.H{
        "status": "ok",
    })
}

func main() {
    r := gin.Default()
    r.GET("/health", healthCheck)
    r.Run(":8080")
}
```

### 複数のルートを定義する

```go
r.GET("/health", healthCheck)
r.GET("/", welcome)
r.GET("/version", version)
```

## 🔗 参考リンク

- [Gin Quick Start](https://gin-gonic.com/docs/quickstart/)
- [Gin Context](https://pkg.go.dev/github.com/gin-gonic/gin#Context)
