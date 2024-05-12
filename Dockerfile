#ベースイメージを指定
#バイナリファイルをビルド
FROM golang:1.22-alpine3.19 AS builder

#コンテナ内のカレントディレクトリ
#無ければ作成される
WORKDIR /app

#[コピー元] [コピー先]
COPY . .

RUN go build -o main main.go

FROM  alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .


EXPOSE 8080
CMD ["/app/main"]
