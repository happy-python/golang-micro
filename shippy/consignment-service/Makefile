protoc:
	protoc -I. --go_out=plugins=micro:. proto/consignment/*.proto

build:
	GOOS=linux GOARCH=amd64 go build
