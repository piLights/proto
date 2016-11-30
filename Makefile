.PHONY: proto gateway swagger lint

proto:
	@protoc --go_out=plugins=grpc:. *.proto

lint:
	@protoc --lint_out=. *.proto
