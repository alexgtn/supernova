FROM golang:1.18-alpine3.15 as builder

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN apk add git

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" .

FROM scratch

WORKDIR /app

COPY --from=builder /app/supernova .

CMD ["/app/supernova", "main"]