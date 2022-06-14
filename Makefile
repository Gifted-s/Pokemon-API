

build:
	go build -o bin/pokemon-server ./cmd/pokemon/m/v1/main.go

run:
	go run ./cmd/pokemon/m/v1/main.go

e2e_test:
	go test -v ./tests/e2e/

unit_tests:
	go test -v ./tests/unit_tests/customize_query_params_test && go test -v ./tests/unit_tests/levenshtein_distance_algorithm_test && go test -v ./tests/unit_tests/operator_matcher_test  && go test -v ./tests/unit_tests/parse_csv_test && go test -v ./tests/unit_tests/sort_pokemon_based_on_edit_distance_test

integration_tests:
	go test -v ./tests/integration_tests/compute_edit_distance_test
    
test: e2e_test  unit_tests integration_tests


compile:
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/pokemon-server-main-darwin-amd64 ./cmd/pokemon/m/v1/main.go
	GOOS=freebsd GOARCH=386 go build -o bin/pokemon-server-main-freebsd-386 ./cmd/pokemon/m/v1/main.go
	GOOS=linux GOARCH=amd64 go build -o bin/pokemon-server-main-linux-amd64 ./cmd/pokemon/m/v1/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/pokemon-server-main-windows-amd64 ./cmd/pokemon/m/v1/main.go
