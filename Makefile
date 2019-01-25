# Copyright © 2019 State Street Bank and Trust Company.  All rights reserved
#
# SPDX-License-Identifier: Apache-2.0

GO_CMD		?= go
DEP_CMD		?= dep
LINT_CMD	?= gometalinter

BIN_DIR := $(CURDIR)/bin

all: clean dep install

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: dep
dep:
	$(DEP_CMD) ensure

.PHONY: lint
lint:
	$(LINT_CMD) --disable=gocyclo --disable=gas --deadline=120s --exclude=vendor ./...

.PHONY: test
test:
	$(GO_CMD) test -cover ./...

.PHONY: install
install:
	GOBIN=$(BIN_DIR) $(GO_CMD) install ./cmd/...