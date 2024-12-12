FROM golang:1.22 AS builder

WORKDIR /api-go-rest

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:latest

COPY --from=builder /api-go-rest/main .

EXPOSE 8000

CMD ["./main"]
