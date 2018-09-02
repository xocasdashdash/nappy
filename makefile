all:
	CGO_ENABLED=0 go build -o nappy
run:
	go run nappy.go 1s