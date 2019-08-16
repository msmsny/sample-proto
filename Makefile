PROTO_FILES := \
	$(shell find ./sample -type f -name '*.proto') \
	vendor/github.com/protocolbuffers/protobuf/src \
	vendor/github.com/mwitkow/go-proto-validators/validator.proto

.PHONY: build
build: $(PROTO_FILES) build-go build-doc

build-go: $(PROTO_FILES)
	protoc \
		--proto_path=. \
		--proto_path=vendor \
		--proto_path=vendor/github.com/protocolbuffers/protobuf/src \
		--go_out=plugins=grpc:./ \
		--govalidators_out=. \
		sample/*.proto

build-doc: $(PROTO_FILES)
	mkdir -p sample/doc
	protoc \
		--proto_path=. \
		--proto_path=vendor \
		--proto_path=vendor/github.com/protocolbuffers/protobuf/src \
		--doc_out=sample/doc \
		--doc_opt=markdown,sample.md \
		sample/*.proto

build-protoset: $(PROTO_FILES) vendor/github.com/googleapis/googleapis/google/rpc/error_details.proto
	protoc \
		--proto_path=. \
		--proto_path=vendor \
		--proto_path=vendor/github.com/protocolbuffers/protobuf/src \
		--descriptor_set_out=sample/sample.protoset \
		--include_imports \
		sample/*.proto \
		vendor/github.com/googleapis/googleapis/google/rpc/error_details.proto

.PHONY: clean
clean:
	find ./sample -name '*.pb.go' -type f | xargs rm -f
	find ./sample -name '*.md' -type f | xargs rm -f
	find ./sample -name '*.protoset' -type f | xargs rm -f
