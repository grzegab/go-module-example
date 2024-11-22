.DEFAULT_GOAL := help

help:
	@echo "make build - build docker compose"
	@echo "make dev - run docker compose"
	@echo "make bash - exec into service container"
	@echo "make clean - stop & remove docker tmp files"
.PHONY:help

build:
	docker compose down
	docker compose build --no-cache
.PHONY:build

dev:
	docker compose up -d
.PHONY:dev

bash:
	docker compose exec -it ex_school bash
.PHONY:bash

tests:
.PHONY:tests

check:tests
.PHONY:check

clean:
	docker compose down
	rm -rf docker/sql/data/*
.PHONY:clean

test:
	go test ./...
.PHONY:test