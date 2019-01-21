all: build

check:
	go mod tidy -v && go mod vendor -v && go mod verify

generate:
	go generate ./...

build:
	go build -v -o ./service .

clean:
	rm -rf ./service
