build:
	@echo Building Library Management System Backend Service...
	@go mod tidy
	@echo finished getting dependencies
	@GOOS=linux GOARCH=amd64 go build -o cmd/