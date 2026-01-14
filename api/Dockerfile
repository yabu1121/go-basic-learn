FROM golang:1.21-alpine

# 必要なツールをインストール
RUN apk add --no-cache git

# 作業ディレクトリを設定
WORKDIR /app

# Go modulesの依存関係をコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# ポート8080を公開
EXPOSE 8080

# アプリケーションを実行
CMD ["go", "run", "main.go"]
