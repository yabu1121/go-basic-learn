# ドリル02: データベース接続

## 🎯 目標
GORMを使ってPostgreSQLデータベースに接続します。

## 📝 要件

以下の仕様でデータベース接続機能を実装してください:

### 接続情報
- **ホスト**: localhost
- **ポート**: 5432
- **ユーザー**: gouser
- **パスワード**: gopass
- **データベース名**: golearn_db

### エンドポイント
1. **GET /health** - ヘルスチェック（前回の続き）
2. **GET /db/ping** - データベース接続確認

### レスポンス例

#### `/db/ping` 成功時
```json
{
  "status": "ok",
  "message": "Database connection successful",
  "database": "golearn_db"
}
```

#### `/db/ping` 失敗時
```json
{
  "status": "error",
  "message": "Database connection failed",
  "error": "エラーメッセージ"
}
```

## 🧪 テスト方法

```bash
# 1. PostgreSQLコンテナが起動していることを確認
docker-compose ps

# 2. アプリケーション起動
go run main.go

# 3. データベース接続確認
curl http://localhost:8080/db/ping
```

## 💡 学習ポイント

- GORMの基本的な使い方
- PostgreSQLドライバーの使用方法
- DSN（Data Source Name）の構築
- データベース接続のエラーハンドリング
- グローバル変数でのDB接続管理

## 🚀 チャレンジ課題

1. 環境変数から接続情報を読み込む（godotenvを使用）
2. 接続プールの設定を追加する
3. データベースの統計情報を返すエンドポイントを追加

## ⚠️ 注意点

- データベース接続は1回だけ行い、グローバル変数で管理します
- 接続失敗時はエラーログを出力してください
- 本番環境では接続情報をハードコードしないこと

## 📚 次のステップ

このドリルをクリアしたら、次は `03_model_definition` に進みましょう！
