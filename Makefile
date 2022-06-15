

build:
	go build -o bin/pokemon-server ./cmd/pokemon/m/v1/main.go

start_server:
	go run ./cmd/pokemon/m/v1/main.go

e2e_test:
	go test -v ./tests/e2e/

all_unit_and_integration_tests:
	go test -v ./internal/helpers

test: e2e_test  all_unit_and_integration_tests


compile:
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/pokemon-server-main-darwin-amd64 ./cmd/pokemon/m/v1/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/pokemon-server-main-freebsd-386 ./cmd/pokemon/m/v1/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/pokemon-server-main-linux-amd64 ./cmd/pokemon/m/v1/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/pokemon-server-main-windows-amd64 ./cmd/pokemon/m/v1/main.go

swagger:
	swagger generate spec -o ./api/swagger.yaml --scan-models

api_doc:
		@echo "Ensure you have go-swagger installed if not installed already"
		@echo "To install go-swaggger: run [brew tap go-swagger/go-swagger]"
		@echo "Then run: run [brew install go-swagger]"
		swagger serve ./api/swagger.yaml