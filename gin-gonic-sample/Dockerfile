FROM golang:1.21.3-alpine as builder
COPY  etc/ssl/ /etc/ssl/certs/
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/main ./cmd/app
ENTRYPOINT ["/out/main"]