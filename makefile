dev:
	go run src/server.go

build:
	go build -o bin/backend src/server.go 

clean:
	rm -rf bin