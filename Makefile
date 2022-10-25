PROG = prog

.PHONY: init
init:
	go mod tidy
	go mod vendor

.PHONY: build
build: init
	go build -o ./$(PROG)

.PHONY: run
run: build
	./$(PROG)

.PHONY: clean
	rm ./$(PROG)

.PHONY: test
test: init
	go test ./...