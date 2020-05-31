.PHONY: build

build:
	docker-compose run --rm golang bash -c "cd cmd/code-generator;go build -o ../../build"