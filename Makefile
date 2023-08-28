BINARY_NAME=turbo
OUTPUT_DIR=bin

build:
	go build -o ${OUTPUT_DIR}/${BINARY_NAME} main.go
	# GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}.exe main.go
	# GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}.mac main.go
	# GOARCH=arm64 GOOS=linux go build -o ${BINARY_NAME}.arm main.go
	# GOARCH=arm64 GOOS=windows go build -o ${BINARY_NAME}.arm.exe main.go
	# GOARCH=arm64 GOOS=darwin go build -o ${BINARY_NAME}.arm.mac main.go

run: build
	./${OUTPUT_DIR}/${BINARY_NAME}

clean:
	go clean
	rm ${OUTPUT_DIR}/${BINARY_NAME}

test:
	go test ./...

test_coverage:
	go test ./... -coverprofile=coverage.out
