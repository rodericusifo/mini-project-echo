start:
	@echo "  >  Starting Program..."
	wire wire/wire.go
	go run cmd/main.go -env=dev