MAKEFILE_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: lint
lint:
	docker run -t --rm -v $(MAKEFILE_DIR):/app -w /app/src golangci/golangci-lint:v2.3.0 golangci-lint run
