# goenv: Go の環境を構築
FROM golang:1.21.0-alpine as goenv

RUN apk update

WORKDIR /app

# 共通リソースを配置
COPY go.mod .
COPY go.sum .
RUN go mod download

# サーバーコードを配置
COPY server server/

# builder: サーバーを構築
FROM goenv as builder

# ビルドタグを指定
ARG build

# DI をビルドタグに合わせて再設定
RUN go run github.com/google/wire/cmd/wire \
    gen -tags "$build" ./.../injection/...

# サーバーをビルド
RUN CGO_ENABLED=0 GOOS=linux go build -v -o /app/server ./server/cmd

# run: 実行環境の構築
FROM gcr.io/distroless/static-debian10 as run

COPY --from=builder /app/server /server
COPY --from=builder /app/etc/ /etc/

CMD ["/server"]

# air: デバッグ実行環境を準備
FROM goenv as dev

# ホットスワップライブラリを導入
RUN go install github.com/cosmtrek/air@latest

# ビルドタグを指定
ARG build
ENV BUILD $build

# DI をビルドタグに合わせて再設定
# (ローカルも変更になるので注意)
CMD go run github.com/google/wire/cmd/wire \
    gen -tags "$BUILD" ./.../injection/... && \
    air -c ./server/.air.toml