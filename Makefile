APP_NAME=information-protect
BIN_DIR=bin
OUT_DIR=$(BIN_DIR)/out
CONFIG=configs/config-local.yaml
OUT=$(BIN_DIR)/$(APP_NAME)

# дефолтные значения
NUM1 ?= 123456789
NUM2 ?= 987654321

# если передали make files 111 222 — подменим NUM1 и NUM2
ifeq ($(word 2, $(MAKECMDGOALS)),)
else
  NUM1 := $(word 2, $(MAKECMDGOALS))
endif

ifeq ($(word 3, $(MAKECMDGOALS)),)
else
  NUM2 := $(word 3, $(MAKECMDGOALS))
endif

.PHONY: all build clean run files test

all: build

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(OUT) ./cmd/$(APP_NAME)

clean:
	rm -rf $(BIN_DIR)

run: build
	$(OUT) -config=$(CONFIG)

files:
	@mkdir -p in
	@rm -f in/num1.txt in/num2.txt
	@echo "$(NUM1)" > in/num1.txt
	@echo "$(NUM2)" > in/num2.txt
	@echo "Files num1.txt and num2.txt created with values: $(NUM1), $(NUM2)"

test:
	go test ./internal/... -v

%:
	@:
