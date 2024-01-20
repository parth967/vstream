curl -s  "http://127.0.0.1:8080/d"

echo "restarting server..."

go run cmd/main.go
