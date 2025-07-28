
all: run

run:
	@go run cmd/mkgo/main.go

build:
	@go build -o mkgo cmd/mkgo/main.go 
