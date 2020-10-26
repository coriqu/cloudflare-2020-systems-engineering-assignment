hello:
	echo "Hello"
build:
	go build hello.go
run:
	go run hello.go
all: hello build