.PHONY: test build run docker-test docker-build docker-run ci clean check-ports

# Определяем команду docker compose
DOCKER_COMPOSE := $(shell if command -v docker-compose >/dev/null 2>&1; then echo "docker-compose"; else echo "docker compose"; fi)

# Тестирование
test:
	go test -v -count=1 ./internal/tests/...

# Локальная сборка
build:
	go build -o main .

# Локальный запуск
run: build
	./main

# Очистка
clean:
	rm -f main
	docker system prune -f

# Docker операции
docker-test:
	docker build --target builder -t url-shortener-test .
	docker run --rm url-shortener-test go test -v -count=1 ./internal/tests/...

docker-build: docker-test
	docker build -t url-shortener .

docker-run: check-ports docker-build
	$(DOCKER_COMPOSE) up -d

docker-stop:
	$(DOCKER_COMPOSE) down

# CI/CD пайплайн
ci: test docker-build

# Развертывание
deploy: ci docker-run

check-ports:
	@echo "Checking if ports are available..."
	@$(DOCKER_COMPOSE) down 2>/dev/null || true
	@if lsof -i :8081 > /dev/null; then \
		echo "Port 8081 is in use. Stopping service..."; \
		sudo lsof -ti :8081 | xargs kill -9; \
	fi
	@if lsof -i :5433 > /dev/null; then \
		echo "Port 5433 is in use. Stopping service..."; \
		sudo lsof -ti :5433 | xargs kill -9; \
	fi