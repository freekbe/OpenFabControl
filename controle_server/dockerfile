FROM golang:latest AS builder
WORKDIR /app
COPY go.mod ./
# Ensure dependencies and go.sum are populated inside the builder
RUN go mod download && go mod tidy && go get github.com/lib/pq@v1.10.0 && go get golang.org/x/crypto@v0.46.0
COPY . .
# Build a static Linux binary to avoid runtime loader issues on Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /app/main /app/main
# Ensure the binary is executable
RUN chmod +x /app/main
EXPOSE 3000
CMD ["/app/main"]
