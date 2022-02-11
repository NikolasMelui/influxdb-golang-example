.PHONY: clean
clean:
	@echo "Run clean"
	@rm -rf ./bin
	@echo "Successfully cleaned"

APP_BIN = ./bin/influxdb-golang-example

.PHONY: compile
compile:
	@echo "Run compile influxdb-golang-example"
	@env go build -o ${APP_BIN} ./cmd
	@echo "Successfully compiled"


run: compile
	@env ${APP_BIN}

.PHONY: test
test:
	@env go test ./internal/*