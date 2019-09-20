protoc:
	protoc -I. --go_out=plugins=micro:. proto/user/*.proto

build:
	GOOS=linux GOARCH=amd64 go build