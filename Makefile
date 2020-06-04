.PHONY: build plugins
ENCODERS = helloencode-encoder openapi-encoder
DECODERS = hellodecode-decoder psr15api-decoder

PROJECT_DIR = $(shell pwd)
BUILD_DIR = build
GO_BUILD = go build -o ../../$(BUILD_DIR)
GO_BUILD_PLUGIN = go build -o ../../../../$(BUILD_DIR) -buildmode=plugin
ENCODER_PLUGIN_PATH = internal/plugins/encoder/
DECODER_PLUGIN_PATH = internal/plugins/decoder/
DOCKER_RUN_BASH = docker-compose run --rm golang bash -c

build:
	@echo "## Building code-generator .."
	@$(DOCKER_RUN_BASH) "cd cmd/code-generator;${GO_BUILD}"

#encoders
%-encoder:
	@ENCODER=$(firstword $(subst -,  ,$@)); \
	echo "## Building encoder plugin $$ENCODER ..."; \
	$(DOCKER_RUN_BASH) "cd $(ENCODER_PLUGIN_PATH)$$ENCODER; $(GO_BUILD_PLUGIN)"

encoders: $(ENCODERS)
	@echo "# Encoders builds done"

#decoders
%-decoder:
	@DECODER=$(firstword $(subst -,  ,$@)); \
	echo "## Building decoder plugin $$DECODER ..."; \
	$(DOCKER_RUN_BASH) "cd $(DECODER_PLUGIN_PATH)$$DECODER; $(GO_BUILD_PLUGIN)"

decoders: $(DECODERS)
	@echo "# Decoders builds done"

all: build encoders decoders
	@echo "# All builds done"