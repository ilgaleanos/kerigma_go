FROM golang:1.14-alpine AS builder

WORKDIR /kerigma

COPY ./src /kerigma
RUN go mod download
RUN go build -o app -ldflags="-s -w" ./server


FROM alpine:3.12
EXPOSE 8080

RUN apk update && apk add --no-cache ca-certificates tzdata

WORKDIR /home/kerigma
COPY ./src/.env.prod /home/kerigma/.env
COPY ./src/keys /home/kerigma/keys
COPY --from=builder /kerigma/app /home/kerigma/app

CMD ["./app"]
