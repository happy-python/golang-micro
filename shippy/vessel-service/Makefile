protoc:
	protoc -I. --go_out=plugins=micro:. proto/vessel/*.proto

build:
	GOOS=linux GOARCH=amd64 go build