FROM golang:1.11-alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/api

# Prepare the final image
FROM scratch
WORKDIR /app
COPY --from=builder /app/api .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
CMD ["./api"]
