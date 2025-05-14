lint:
	goimports -d .
	golangci-lint run -v

format:
	goimports -w .
	golangci-lint run --fix
	golangci-lint fmt

test:
	go test -coverprofile=coverage.out -timeout 10m -count=1 ./...

coverage: ## run coverage
	go tool cover -func=coverage.out

doc:
	go run main.go docs
