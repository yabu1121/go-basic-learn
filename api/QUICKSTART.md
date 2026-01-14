# 🚀 クイックスタートガイド

このガイドでは、Gin + GORM API開発環境を最速でセットアップして、最初のAPIを動かすまでの手順を説明します。

## ⏱️ 所要時間: 約10分

---

## ステップ1: 前提条件の確認

以下がインストールされていることを確認してください:

- ✅ **Docker Desktop** - [ダウンロード](https://www.docker.com/products/docker-desktop)
- ✅ **Go 1.21以上** - [ダウンロード](https://golang.org/dl/)

確認コマンド:
```bash
docker --version
go version
```

---

## ステップ2: プロジェクトディレクトリに移動

```bash
cd c:\Users\nhaya\Desktop\code\prog\golearn\api
```

---

## ステップ3: 環境変数ファイルの作成

```bash
# .env.exampleをコピー
cp .env.example .env

# または手動で作成
# .envファイルに以下を記述:
# DB_HOST=localhost
# DB_PORT=5432
# DB_USER=gouser
# DB_PASSWORD=gopass
# DB_NAME=golearn_db
# APP_PORT=8080
```

---

## ステップ4: Dockerコンテナの起動

```bash
# PostgreSQLコンテナを起動
docker-compose up -d postgres

# 起動確認（healthyになるまで待つ）
docker-compose ps
```

**期待される出力**:
```
NAME                  STATUS              PORTS
golearn_postgres      Up (healthy)        0.0.0.0:5432->5432/tcp
```

---

## ステップ5: Go依存関係のインストール

```bash
# 依存パッケージをダウンロード
go mod download

# または
go mod tidy
```

---

## ステップ6: 最初のドリルを実行

### ドリル01: ヘルスチェックAPI

```bash
# ドリル01のディレクトリに移動
cd drills/level1/01_health_check

# 解答例を実行
go run answer.go
```

別のターミナルで:
```bash
# APIをテスト
curl http://localhost:8080/health
```

**期待される出力**:
```json
{"status":"ok","message":"API is running"}
```

成功です！🎉

---

## ステップ7: データベース接続を試す

### ドリル02: データベース接続

```bash
# Ctrl+C で前のプログラムを停止

# ドリル02のディレクトリに移動
cd ../02_database_connection

# 解答例を実行
go run answer.go
```

別のターミナルで:
```bash
# データベース接続確認
curl http://localhost:8080/db/ping
```

**期待される出力**:
```json
{
  "status":"ok",
  "message":"Database connection successful",
  "database":"golearn_db"
}
```

---

## ステップ8: ユーザー作成APIを試す

### ドリル05: ユーザー作成API

```bash
# Ctrl+C で前のプログラムを停止

# ドリル05のディレクトリに移動
cd ../../level2/05_create_user

# 解答例を実行
go run answer.go
```

別のターミナルで:
```bash
# ユーザーを作成
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "山田太郎",
    "email": "yamada@example.com",
    "age": 25,
    "password": "password123",
    "phone": "09012345678"
  }'
```

**期待される出力**:
```json
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
```

---

## 🎯 次のステップ

おめでとうございます！環境構築が完了しました。

### 学習を続ける

1. **README.mdを読む**: `api/README.md` で全体像を把握
2. **DRILLS.mdを確認**: `api/DRILLS.md` で全30問の問題を確認
3. **順番に進める**: `drills/level1/` から順番に進める

### 推奨学習フロー

```
レベル1（基礎）
  ↓
レベル2（CRUD）
  ↓
レベル3（中級）
  ↓
レベル4（リレーション）
  ↓
レベル5（応用）
  ↓
レベル6（実践）
```

---

## 🛠️ よく使うコマンド

### Docker関連

```bash
# コンテナ起動
docker-compose up -d

# コンテナ停止
docker-compose down

# ログ確認
docker-compose logs -f postgres

# PostgreSQLに接続
docker exec -it golearn_postgres psql -U gouser -d golearn_db

# データベースをリセット
docker-compose down -v
docker-compose up -d
```

### Go関連

```bash
# プログラム実行
go run main.go

# ビルド
go build -o app main.go

# 依存関係の整理
go mod tidy

# 依存関係の更新
go get -u ./...
```

### API テスト

```bash
# GET リクエスト
curl http://localhost:8080/health

# POST リクエスト
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"太郎","email":"taro@example.com","age":25,"password":"pass1234"}'

# PUT リクエスト
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"太郎（更新）"}'

# DELETE リクエスト
curl -X DELETE http://localhost:8080/users/1
```

---

## ❓ トラブルシューティング

### 問題1: ポート8080が使用中

**エラー**: `address already in use`

**解決策**:
```bash
# Windowsの場合
netstat -ano | findstr :8080
taskkill /PID <プロセスID> /F

# または別のポートを使用
# .envファイルで APP_PORT=8081 に変更
```

### 問題2: PostgreSQLに接続できない

**エラー**: `connection refused`

**解決策**:
```bash
# コンテナの状態確認
docker-compose ps

# コンテナが起動していない場合
docker-compose up -d postgres

# ログ確認
docker-compose logs postgres
```

### 問題3: Go依存関係のエラー

**エラー**: `package not found`

**解決策**:
```bash
# 依存関係を再インストール
go mod download
go mod tidy

# キャッシュをクリア
go clean -modcache
go mod download
```

### 問題4: データベースが作成されていない

**エラー**: `database "golearn_db" does not exist`

**解決策**:
```bash
# PostgreSQLコンテナに接続
docker exec -it golearn_postgres psql -U gouser

# データベースを作成
CREATE DATABASE golearn_db;

# 確認
\l

# 終了
\q
```

---

## 📚 参考リンク

- [Gin公式ドキュメント](https://gin-gonic.com/docs/)
- [GORM公式ドキュメント](https://gorm.io/docs/)
- [PostgreSQL公式ドキュメント](https://www.postgresql.org/docs/)
- [Docker公式ドキュメント](https://docs.docker.com/)

---

## 💬 学習のコツ

1. **手を動かす**: コードを読むだけでなく、必ず自分で書く
2. **エラーを恐れない**: エラーは学習の機会
3. **小さく始める**: 一度に全部理解しようとしない
4. **繰り返す**: 同じドリルを何度も解く
5. **応用する**: 学んだことを自分のプロジェクトに活かす

---

## 🎉 準備完了！

環境構築が完了しました。さあ、API開発の学習を始めましょう！

次は `README.md` を読んで、全体像を把握してください。

Happy Coding! 🚀
