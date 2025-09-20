# learning-protobuf-go

## Prerequisites

- Downloaded the protobuf for Golang


## Generate a protobuf file

Had run the following
```sh
protoc --go_out=. person.proto

# on one window
go run main.go

# on another window
go test ./...
```
