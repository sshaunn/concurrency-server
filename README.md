install:
go mod download
go mod tidy

run:
export ENVIORNMENT=local
go run ./cmd/api/main.go
