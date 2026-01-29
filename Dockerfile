# ベースイメージとしてGo 1.23 (Alpine Linux)を使用
# go.modでは1.21以上となっていたため、より新しい安定版を選択
FROM golang:1.23-alpine

# 作業ディレクトリを設定
WORKDIR /app

# 必要なツールとビルド依存関係をインストール
# git: go mod downloadに必要
# build-base: CGO (SQLiteなど) に必要
# curl: デバッグ用
RUN apk add --no-cache git build-base curl

# go.modとおそらく存在するgo.sumをコピー
COPY go.mod go.sum* ./

# 依存関係をダウンロード
RUN go mod download

# ソースコードをコピー
COPY . .

# コンテナがすぐに終了しないように待機させる
# 実際にはdocker composeでバインドマウントして使うことを想定
CMD ["tail", "-f", "/dev/null"]
