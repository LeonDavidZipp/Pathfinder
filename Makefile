NAME := solver

all: build

build:
	go build -o $(NAME) src/main.go

re: clean build

clean:
	rm -f $(NAME)