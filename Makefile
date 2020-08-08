test:
	go mod tidy
	go test
	golangci-lint run