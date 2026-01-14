# ドリル05: ユーザー作成API

## 🎯 目標
POSTリクエストを処理して、データベースに新しいユーザーを作成します。

## 📝 要件

### エンドポイント
- **メソッド**: POST
- **パス**: `/users`
- **リクエストボディ**:
  ```json
  {
    "name": "山田太郎",
    "email": "yamada@example.com",
    "age": 25
  }
  ```

### レスポンス

#### 成功時（201 Created）
```json
{
  "status": "success",
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "山田太郎",
    "email": "yamada@example.com",
    "age": 25,
    "created_at": "2024-01-14T16:35:00+09:00",
    "updated_at": "2024-01-14T16:35:00+09:00"
  }
}
```

#### バリデーションエラー時（400 Bad Request）
```json
{
  "status": "error",
  "message": "Validation failed",
  "errors": [
    "Name is required",
    "Email must be valid"
  ]
}
```

#### データベースエラー時（500 Internal Server Error）
```json
{
  "status": "error",
  "message": "Failed to create user",
  "error": "エラーメッセージ"
}
```

## 🧪 テスト方法

```bash
# ユーザー作成
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "山田太郎",
    "email": "yamada@example.com",
    "age": 25
  }'

# バリデーションエラーのテスト
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "",
    "email": "invalid-email"
  }'

# 重複メールアドレスのテスト
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "佐藤花子",
    "email": "yamada@example.com"
  }'
```

## 💡 学習ポイント

- POSTリクエストの処理方法
- JSONリクエストボディのバインディング
- バリデーションタグの使用
- GORMのCreate操作
- 適切なHTTPステータスコードの返却
- エラーハンドリング

## 🚀 チャレンジ課題

1. パスワードフィールドを追加（ハッシュ化は不要）
2. 電話番号フィールドを追加（正規表現バリデーション）
3. 作成日時を日本時間で返す
4. 同じメールアドレスの重複登録を防ぐ

## ⚠️ 注意点

- メールアドレスは一意である必要があります
- 年齢は0以上である必要があります
- 名前は必須項目です
- リクエストボディが不正な場合は400エラーを返します

## 📚 次のステップ

このドリルをクリアしたら、次は `06_user_list` に進みましょう！
