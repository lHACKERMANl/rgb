.PHONY: build run clean deps

APP_NAME = renderApp
BUILD_DIR = build

# Сборка приложения
build:
	go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/renderApp/

# Запуск приложения
run:
	go run ./cmd/renderApp/

# Установка зависимостей
deps:
	go mod download
	go mod tidy

# Очистка
clean:
	rm -rf $(BUILD_DIR)
	go clean

# Инициализация проекта
init: deps
	mkdir -p $(BUILD_DIR)

# Все команды
all: clean init build
