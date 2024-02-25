build:
	go build -o bin/mindfck

run: build
	./bin/mindfck

start:
	go run .\server.go