dev:
	go run main.go

build:
	go build -o bin/backend main.go

clean:
	rm -rf bin

run:
	./bin/backend