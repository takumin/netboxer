NAME    := netboxer
VERSION := $(shell gobump show -r)
COMMIT  := $(shell git rev-parse HEAD)
DATE    := $(shell date --utc --rfc-3339=seconds)

LDFLAGS_NAME    := -X "main.name=$(NAME)"
LDFLAGS_VERSION := -X "main.version=$(VERSION)"
LDFLAGS_COMMIT  := -X "main.commit=$(COMMIT)"
LDFLAGS_DATE    := -X "main.date=$(DATE)"
LDFLAGS_CONST   := $(LDFLAGS_NAME) $(LDFLAGS_VERSION) $(LDFLAGS_COMMIT) $(LDFLAGS_DATE)
LDFLAGS         := -ldflags '-s -w $(LDFLAGS_CONST) -extldflags "-static"'

SRCS := $(shell find $(CURDIR) -type f -name '*.go')

.PHONY: default
default: test

.PHONY: run
run: $(SRCS)
	go run $(CURDIR)/cmd/$(NAME)

.PHONY: build
build: $(CURDIR)/bin/$(NAME)

.PHONY: lint
lint:
	golint -set_exit_status ./...

.PHONY: test
test: lint
	go test -v

.PHONY: clean
clean:
	rm -rf $(CURDIR)/bin
	rm -rf $(CURDIR)/dist

$(CURDIR)/bin/$(NAME): $(SRCS)
	go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o $(CURDIR)/bin/$(NAME) $(CURDIR)/cmd/$(NAME)
