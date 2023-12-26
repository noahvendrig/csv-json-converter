build:
	@go build -o bin/csv-json-converter

run: build
	@./bin/csv-json-converter

test:
	@go test -v ./...