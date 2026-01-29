# ベースイメージ
FROM golang:1.23-alpine

# 作業ディレクトリ
WORKDIR /app

# 必要なツール
RUN apk add --no-cache git build-base curl

# airのインストール (バイナリ版)
# go installの失敗を避けるため、インストールスクリプトを使用
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b /go/bin

# go.mod類
COPY go.mod go.sum* ./
RUN go mod download

# ソース
COPY . .

CMD ["tail", "-f", "/dev/null"]
