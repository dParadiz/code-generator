.PHONY: build plugins

build:
	docker-compose run --rm golang bash -c "cd cmd/code-generator;go build -o ../../build"

plugins:
	docker-compose run --rm golang bash -c "cd internal/plugins/encoder/helloencode;go build -o ../../../../build -buildmode=plugin"
	docker-compose run --rm golang bash -c "cd internal/plugins/decoder/hellodecode;go build -o ../../../../build -buildmode=plugin"