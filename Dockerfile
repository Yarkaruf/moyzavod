FROM node as web

WORKDIR /web
COPY ./public ./src ./package.json ./package-lock.json ./rollup.config.js ./
RUN npm ci
RUN npm run build

FROM golang:1.18 AS builder

WORKDIR /go/src/app/
COPY ./server/go.mod ./server/go.sum ./
RUN go mod download
COPY ./server .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o moyzavod . && chmod +x moyzavod

FROM alpine:3.15

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /go/src/app/moyzavod ./moyzavod
COPY --from=web /web/public ./public

ENTRYPOINT ["./moyzavod"]

EXPOSE 80
