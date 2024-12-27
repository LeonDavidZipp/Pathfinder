NAME := solver

all:
	go build -o $(NAME) src/main.go

clean:
	rm -f $(NAME)