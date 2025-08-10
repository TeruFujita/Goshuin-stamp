# Go バックエンド用 Dockerfile
FROM golang:1.21-alpine AS builder

# 作業ディレクトリの設定
WORKDIR /app

# 必要なパッケージのインストール
RUN apk add --no-cache git

# Go modules のコピー
COPY go.mod go.sum ./

# 依存関係のダウンロード
RUN go mod download

# ソースコードのコピー
COPY . .

# バイナリのビルド
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/server

# 実行用イメージ
FROM alpine:latest

# 必要なパッケージのインストール
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# ビルドしたバイナリをコピー
COPY --from=builder /app/main .

# ポートの公開
EXPOSE 8080

# アプリケーションの実行
CMD ["./main"]
