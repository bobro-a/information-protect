APP_NAME=information-protect
BIN_DIR=bin
OUT=$(BIN_DIR)/$(APP_NAME)

.PHONY: all build clean run

all: build

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(OUT) ./cmd/$(APP_NAME)

clean:
	rm -rf $(BIN_DIR)

run: build
	$(OUT)

files:
	@mkdir -p in
	@echo "123456789" > in/num1.txt
	@echo "987654321" > in/num2.txt
	@echo "Files num1.txt and num2.txt created and filled."

