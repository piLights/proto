.PHONY: proto lint

proto:
	@protoc --go_out=plugins=grpc:. *.proto

lint:
	@protoc --lint_out=. *.proto
