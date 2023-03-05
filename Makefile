.PHONY: run
run:
	docker compose -f docker-compose.yml up --build

.PHONY: dev
dev:
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up --build

.PHONY: native-build
native-build:
	go build -o ./main ./server/cmd/main.go

.PHONY: native-run
native-run: native-build
	./main
