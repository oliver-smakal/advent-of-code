BINARY_NAME=main.out

ifndef VERBOSE
.SILENT:
endif
 
all: build test
 
build:
	go build -o ${BINARY_NAME}
 
test:
	go test -v ./...
	./test.sh | ./${BINARY_NAME}
 
run:
	go build -o ${BINARY_NAME}
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}
