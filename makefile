#
# Simple build and lint environment
#

.SUFFIXES	:
.SUFFIXES	:	.go

#
# print colored output
RESET_COLOR    := \033[0m
make_std_color := \033[3$1m      # defined for 1 through 7
make_color     := \033[38;5;$1m  # defined for 1 through 255
OK_COLOR       := $(strip $(call make_std_color,2))
WRN_COLOR      := $(strip $(call make_std_color,3))
ERR_COLOR      := $(strip $(call make_std_color,1))
STD_COLOR      := $(strip $(call make_color,8))

COLOR_OUTPUT = 2>&1 |                                        \
    while IFS='' read -r line; do                            \
        if  [[ $$line == FAIL* ]]; then                      \
            echo -e "$(ERR_COLOR)$${line}$(RESETCOLOR)";     \
        elif [[ $$line == *:[\ ]FAIL:* ]]; then              \
            echo -e "$(ERR_COLOR)$${line}$(RESETCOLOR)";     \
        elif [[ $$line == [\-][\-][\-][\ ]FAIL:* ]]; then    \
            echo -e "$(ERR_COLOR)$${line}$(RESETCOLOR)";     \
        elif [[ $$line == WARN* ]]; then                     \
            echo -e "$(WRN_COLOR)$${line}$(RESET_COLOR)";    \
        elif [[ $$line == PASS ]]; then                       \
            echo -e "$(OK_COLOR)$${line}$(RESET_COLOR)";     \
        elif [[ $$line == [\-][\-][\-][\ ]PASS:* ]]; then    \
            echo -e "$(OK_COLOR)$${line}$(RESETCOLOR)";     \
        elif [[ $$line == ok* ]]; then                       \
            echo -e "$(OK_COLOR)$${line}$(RESET_COLOR)";     \
        else                                                 \
            echo -e "$(STD_COLOR)$${line}$(RESET_COLOR)";    \
        fi;                                                  \
    done; exit $${PIPESTATUS[0]};

.DEFAULT: $(help)

BUILD_HASH   := $(shell git rev-parse --short HEAD)
BUILD_NUM   := $(shell expr `cat .git/buildnum 2>/dev/null` + 1 >.git/buildnum && cat .git/buildnum)
BUILD_DATE  := $(shell date +'%Y-%m-%d.%H:%M:%S')
BUILD_TAG  := $(BUILD_HASH).$(BUILD_NUM)
BUILD_VERSION := $(BUILD_HASH).$(BUILD_NUM) ($(BUILD_DATE))

SHELL           := /bin/bash
#GO              := /usr/local/go/bin/go
GO              := go
#GO_PREFIX       := GOCACHE=off
GO_PREFIX       := 
GO_FLAGS        := -v
#GO_FLAGS			:= -v gcflags=\"-m\"
GO_LDFLAGS      := -ldflags="-X github.com/raibru/pktfmt/cmd.buildTag=$(BUILD_TAG) -X github.com/raibru/pktfmt/cmd.buildDate=$(BUILD_DATE)"
BIN_FILE         := pktfmt
MAIN_FILE       := $(BIN_FILE).go
TEST_FILES      := $(wildcard *_test.go)
TEST_FILES      += $(wildcard test/*.go)
SRCS            := $(filter-out $(wildcard *_test.go), $(wildcard *.go))
#SRCS_TEST       := $(wildcard *_test.go)
SRCS_TEST       := ./...
BIN_DIR         := ./bin
BIN_DIR_DEV     := $(BIN_DIR)/dev
RUNTIME_DEV     := ./runtime/develop

CLEAN_FILES 	  :=                    \
									tags                \
									$(wildcard ./tmp/*) \
									./$(BIN_DIR_DEV)/*

help:
	-@echo "Makefile with following options (make <option>):"
	-@echo "	clean"
	-@echo "	tdd"
	-@echo "	test"
	-@echo "	test-cover"
	-@echo "	test-cache"
	-@echo "	test-verbose"
	-@echo "	test-coverage"
	-@echo "	ctags"
	-@echo "	build (curent os)"
	-@echo "	build-windows"
	-@echo "	build-linux"
	-@echo "	run"
	-@echo "    (*) not implemented"
	-@echo ""

print:
	-@echo "SRCS       ==> [$(SRCS)]"
	-@echo "SRCS_TEST  ==> [$(SRCS_TEST)]"
	-@echo "TEST_FILES ==> [$(TEST_FILES)]"

.PHONY: all deploy-all clean tdd test test-cover test-cache test-verbose test-trace-view test-coverage
all: test build
deploy-all: clean test build build-windows build-linux deploy-dev

clean:
	$(GO) clean
	rm -f $(CLEAN_FILES)

tdd:
	@$(GO) test ./... $(COLOR_OUTPUT)

test: $(TEST_FILES) $(SRCS)
	@$(GO) test ./...

test-cover: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -cover ./...

test-cache: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test $(SRCS_TEST) $(COLOR_OUTPUT)

test-verbose: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -v $(SRCS_TEST) $(COLOR_OUTPUT)

test-trace-view: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -trace ./tmp/trace.out $(COLOR_OUTPUT)
	@$(GO) tool trace ./tmp/trace.out

test-coverage: $(TEST_FILES) $(SRCS)
	@$(GO_PREFIX) $(GO) test -coverprofile=./tmp/coverage.out ./... $(COLOR_OUTPUT)

run:
	$(GO) run $(GO_FLAGS) -o $(BIN_DIR_DEV)/$(BIN_FILE) $(MAIN_FILE)

build:
	$(GO) build $(GO_FLAGS) $(GO_LDFLAGS) -o $(BIN_DIR_DEV)/$(BIN_FILE) $(MAIN_FILE)

build-windows:
	GOOS=windows GOARCH=amd64 $(GO) build $(GO_FLAGS) -o $(BIN_DIR)/windows/amd64/$(BIN_FILE).exe $(MAIN_FILE)

build-linux:
	GOOS=linux GOARCH=amd64 $(GO) build $(GO_FLAGS) -o $(BIN_DIR)/linux/amd64/$(BIN_FILE) $(MAIN_FILE)

deploy-dev:
	cp $(BIN_DIR_DEV)/$(BIN_FILE) $(RUNTIME_DEV)/$(BIN_FILE)

ctags:
	ctags -RV .

# EOF