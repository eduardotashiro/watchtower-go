FROM golang:1.26-alpine as builder

WORKDIR /app

COPY . . 

RUN go build -o watchtower-go .

FROM alpine:latest

COPY --from=builder /app/watchtower-go ./watchtower-go

CMD ["./watchtower-go"]

