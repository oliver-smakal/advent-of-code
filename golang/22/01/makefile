BINARY_NAME=main.out
 
all: build test
 
build:
	go build -o ${BINARY_NAME}
 
test:
	go test -v ./...
	cat ./tests/01 | ./${BINARY_NAME}| diff -y tests/01res -;echo
 
run:
	go build -o ${BINARY_NAME}
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}
