.PHONY: build
build:
	go build -o ./main ./server/cmd/main.go

.PHONY: run
run: build
	./main