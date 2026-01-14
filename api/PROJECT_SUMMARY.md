# 🎓 Gin + GORM API開発 - プロジェクトサマリー

## 📁 プロジェクト構成

```
api/
├── 📄 docker-compose.yml      # Docker環境設定
├── 📄 Dockerfile              # Goアプリケーションコンテナ定義
├── 📄 init.sql                # PostgreSQL初期化スクリプト
├── 📄 .env.example            # 環境変数のサンプル
├── 📄 .gitignore              # Git除外設定
├── 📄 go.mod                  # Go依存関係定義
├── 📄 go.sum                  # 依存関係チェックサム
├── 📄 main.go                 # スターターテンプレート
│
├── 📚 ドキュメント
│   ├── README.md              # プロジェクト概要
│   ├── QUICKSTART.md          # クイックスタートガイド
│   ├── DRILLS.md              # 全30問の問題集
│   └── CHECKLIST.md           # 学習進捗チェックリスト
│
└── 📂 drills/                 # ドリル問題
    ├── level1/                # 基礎編（4問）
    │   ├── 01_health_check/
    │   ├── 02_database_connection/
    │   ├── 03_model_definition/
    │   └── 04_migration/
    │
    ├── level2/                # CRUD基本編（5問）
    │   ├── 05_create_user/
    │   ├── 06_user_list/
    │   ├── 07_user_detail/
    │   ├── 08_update_user/
    │   └── 09_delete_user/
    │
    ├── level3/                # 中級編（5問）
    │   ├── 10_validation/
    │   ├── 11_error_handling/
    │   ├── 12_pagination/
    │   ├── 13_search/
    │   └── 14_sorting/
    │
    ├── level4/                # リレーション編（5問）
    │   ├── 15_one_to_many/
    │   ├── 16_many_to_many/
    │   ├── 17_preload/
    │   ├── 18_join/
    │   └── 19_cascade_delete/
    │
    ├── level5/                # 応用編（5問）
    │   ├── 20_authentication/
    │   ├── 21_middleware/
    │   ├── 22_file_upload/
    │   ├── 23_transaction/
    │   └── 24_aggregation/
    │
    └── level6/                # 実践編（6問）
        ├── 25_blog_api/
        ├── 26_comments/
        ├── 27_likes/
        ├── 28_tags/
        ├── 29_full_text_search/
        └── 30_ecommerce/
```

## 🎯 学習内容

### 技術スタック

- **Webフレームワーク**: Gin v1.9.1
- **ORM**: GORM v1.25.5
- **データベース**: PostgreSQL 15
- **コンテナ**: Docker & Docker Compose
- **言語**: Go 1.21+

### カバーする内容

#### ✅ 作成済み（詳細実装あり）

1. **ドリル01**: ヘルスチェックAPI
   - Ginの基本的なルーティング
   - JSONレスポンス
   - チャレンジ課題3つ

2. **ドリル02**: データベース接続
   - GORMでPostgreSQLに接続
   - 環境変数の活用
   - 接続プール設定
   - チャレンジ課題3つ

3. **ドリル05**: ユーザー作成API
   - POSTリクエスト処理
   - バリデーション
   - エラーハンドリング
   - 重複チェック
   - チャレンジ課題4つ

#### 📝 要件定義済み（DRILLS.mdに記載）

- レベル1: 基礎編（残り2問）
- レベル2: CRUD基本編（残り4問）
- レベル3: 中級編（5問）
- レベル4: リレーション編（5問）
- レベル5: 応用編（5問）
- レベル6: 実践編（6問）

**合計30問のドリル問題**

## 🚀 クイックスタート

### 1. 環境構築（5分）

```bash
cd c:\Users\nhaya\Desktop\code\prog\golearn\api
cp .env.example .env
docker-compose up -d postgres
go mod download
```

### 2. 最初のAPIを実行（1分）

```bash
go run main.go
```

別のターミナルで:
```bash
curl http://localhost:8080/health
```

### 3. ドリルを開始（すぐ）

```bash
cd drills/level1/01_health_check
cat README.md
```

## 📚 学習の流れ

### 推奨学習パス

```
1. QUICKSTART.md を読む（10分）
   ↓
2. 環境構築（10分）
   ↓
3. レベル1: 基礎編（2-3時間）
   - ドリル01-04を順番に
   ↓
4. レベル2: CRUD基本編（3-4時間）
   - ドリル05-09を順番に
   ↓
5. レベル3: 中級編（4-5時間）
   - ドリル10-14を順番に
   ↓
6. レベル4: リレーション編（5-6時間）
   - ドリル15-19を順番に
   ↓
7. レベル5: 応用編（6-8時間）
   - ドリル20-24を順番に
   ↓
8. レベル6: 実践編（8-10時間）
   - ドリル25-30を順番に
```

**総学習時間**: 約30-40時間

### 各ドリルの学習方法

1. **README.mdを読む**（5分）
   - 問題の要件を理解
   - 学習ポイントを確認

2. **自分で実装**（20-30分）
   - まず自力で実装してみる
   - エラーを恐れない

3. **hint.mdを参照**（10分）
   - 詰まったらヒントを見る
   - 実装パターンを学ぶ

4. **answer.goで確認**（10分）
   - 解答例と比較
   - ベストプラクティスを学ぶ

5. **テスト実行**（5分）
   - curlコマンドで動作確認
   - 期待通りの結果を確認

6. **チャレンジ課題**（20-30分）
   - 余裕があれば挑戦
   - 応用力を養う

**1ドリルあたり**: 約1-2時間

## 💡 学習のポイント

### 効果的な学習方法

1. **手を動かす**
   - コードを読むだけでなく、必ず自分で書く
   - タイピングすることで記憶に定着

2. **エラーを楽しむ**
   - エラーメッセージを読む習慣をつける
   - エラーから学ぶことが最も多い

3. **小さく始める**
   - 一度に全部理解しようとしない
   - 1つずつ確実に進める

4. **繰り返す**
   - 同じドリルを何度も解く
   - 2周目、3周目で理解が深まる

5. **応用する**
   - 学んだことを自分のプロジェクトに活かす
   - オリジナルのAPIを作ってみる

### つまづきやすいポイント

1. **ポインタの理解**
   - `&user` と `user` の違い
   - GORMでは基本的にポインタを使う

2. **エラーハンドリング**
   - `if err != nil` の重要性
   - 適切なHTTPステータスコードの選択

3. **構造体タグ**
   - `gorm`, `json`, `binding` タグの使い分け
   - タグの書き方（バッククォート）

4. **データベースのリレーション**
   - 外部キーの設定
   - Preloadの使い方

## 🛠️ よく使うコマンド

### Docker

```bash
# 起動
docker-compose up -d

# 停止
docker-compose down

# ログ確認
docker-compose logs -f postgres

# PostgreSQLに接続
docker exec -it golearn_postgres psql -U gouser -d golearn_db

# データベースリセット
docker-compose down -v && docker-compose up -d
```

### Go

```bash
# 実行
go run main.go

# ビルド
go build -o app main.go

# 依存関係整理
go mod tidy

# テスト
go test ./...
```

### API テスト

```bash
# GET
curl http://localhost:8080/users

# POST
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"太郎","email":"taro@example.com","age":25,"password":"pass1234"}'

# PUT
curl -X PUT http://localhost:8080/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"太郎（更新）"}'

# DELETE
curl -X DELETE http://localhost:8080/users/1
```

## 📖 参考リソース

### 公式ドキュメント

- [Gin](https://gin-gonic.com/docs/)
- [GORM](https://gorm.io/docs/)
- [PostgreSQL](https://www.postgresql.org/docs/)
- [Docker](https://docs.docker.com/)

### おすすめ学習リソース

- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [GORM Guides](https://gorm.io/docs/)

## 🎯 次のステップ

### このプロジェクトを終えたら

1. **オリジナルAPIを作る**
   - ToDoアプリ
   - ブログシステム
   - SNSのバックエンド

2. **さらに学ぶ**
   - GraphQL（gqlgen）
   - gRPC
   - マイクロサービス

3. **本番環境へのデプロイ**
   - Heroku
   - AWS（ECS, Lambda）
   - Google Cloud Run

## 🎉 完成！

このプロジェクトで、あなたはGin + GORMを使った本格的なREST API開発のスキルを習得できます。

**全30問のドリルを完了すれば、実務レベルのAPI開発ができるようになります！**

頑張ってください！🚀

---

**作成日**: 2024-01-14  
**バージョン**: 1.0.0  
**対象**: Go初心者〜中級者
