# proto-sample

## Requirements

* protoc 3.x

generate Go code

* Go
* protoc-gen-go

generate doc

* protoc-gen-doc

## Generate code/documents

### Setup

protoc

```bash
# Mac
$ curl -sL https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-osx-x86_64.zip -o protoc-3.9.1-osx-x86_64.zip
$ unzip protoc-3.9.1-osx-x86_64.zip

# Linux
$ curl -sL https://github.com/protocolbuffers/protobuf/releases/download/v3.9.1/protoc-3.9.1-linux-x86_64.zip -o protoc-3.9.1-linux-x86_64.zip
$ unzip protoc-3.9.1-linux-x86_64.zip
```

```bash
$ sudo cp bin/protoc /usr/local/bin
```

generate Go code

```bash
# basic
$ go get -u -v github.com/golang/protobuf/protoc-gen-go

# grpc_validator
$ go get -u -v github.com/mwitkow/go-proto-validators/protoc-gen-govalidators

```

generate documents

```bash
# Mac
$ curl -sL https://github.com/pseudomuto/protoc-gen-doc/releases/download/v1.3.0/protoc-gen-doc-1.3.0.darwin-amd64.go1.11.2.tar.gz -o protoc-gen-doc-1.3.0.darwin-amd64.go1.11.2.tar.gz
$ tar xvzf protoc-gen-doc-1.3.0.darwin-amd64.go1.11.2.tar.gz
$ sudo cp protoc-gen-doc-1.3.0.darwin-amd64.go1.11.2/protoc-gen-doc /usr/local/bin

# Linux
$ curl -sL https://github.com/pseudomuto/protoc-gen-doc/releases/download/v1.3.0/protoc-gen-doc-1.3.0.linux-amd64.go1.11.2.tar.gz -o protoc-gen-doc-1.3.0.linux-amd64.go1.11.2.tar.gz
$ tar xvzf protoc-gen-doc-1.3.0.linux-amd64.go1.11.2.tar.gz
$ sudo cp protoc-gen-doc-1.3.0.linux-amd64.go1.11.2/protoc-gen-doc /usr/local/bin
```

.proto dependencies

```bash
$ git clone --depth 1 https://github.com/protocolbuffers/protobuf vendor/github.com/protocolbuffers/protobuf
$ git clone --depth 1 https://github.com/mwitkow/go-proto-validators vendor/github.com/mwitkow/go-proto-validators
$ git clone --depth 1 https://github.com/googleapis/googleapis vendor/github.com/googleapis/googleapis
```

## Build

generate Go code

```bash
$ make build-go
```

generate documents

```bash
$ make build-doc
```

generate Go code and documents

```bash
$ make build
```

generate protoset (for [grpcurl](https://github.com/fullstorydev/grpcurl))

```bash
$ make build-protoset
```
