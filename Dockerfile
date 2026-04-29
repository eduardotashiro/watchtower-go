FROM golang:1.26-alpine as builder

RUN apk add --no-cache git

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o watchtower-go .

FROM alpine:latest

RUN apk add --no-cache chromium xvfb ca-certificates

COPY --from=builder /app/watchtower-go ./watchtower-go

CMD ["./watchtower-go"]