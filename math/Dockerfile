FROM golang:1.25.9-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .

RUN chmod +x ./app

CMD ["./app"] 


