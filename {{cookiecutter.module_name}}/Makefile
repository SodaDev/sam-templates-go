build-{{cookiecutter.module_name}}Function:
	go mod tidy
	go test ./... -cover
	GOOS=linux GOARCH=arm64 go build -o bootstrap {{cookiecutter.module_name}}/cmd/main.go
	cp bootstrap $(ARTIFACTS_DIR)/bootstrap
