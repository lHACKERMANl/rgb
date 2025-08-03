# RGB 2D Render Engine

Минимальный 2D рендер движок на Go с использованием OpenGL.

## Требования

- Go 1.21+
- OpenGL 4.5+
- GLFW библиотеки

### Установка зависимостей

#### Linux (Ubuntu/Debian):
```bash
sudo apt-get update
sudo apt-get install libgl1-mesa-dev libxrandr-dev libxcursor-dev libxinerama-dev libxi-dev libglfw3-dev build-essential
```

#### macOS:
```bash
brew install glfw
xcode-select --install
```

#### Windows:
- Установите TDM-GCC или Visual Studio
- Скачайте GLFW pre-compiled binaries

## Сборка и запуск

```bash
# Инициализация зависимостей
go mod tidy

# Сборка
go build -o renderApp ./cmd/renderApp/

# Запуск
./renderApp
```

Или напрямую:
```bash
go run ./cmd/renderApp/
```

## Использование

При запуске откроется окно 800x600 с оранжевым треугольником.
- Нажмите ESC для выхода

## Архитектура

- `WindowService` - управление окном и событиями
- `ShaderService` - компиляция и управление шейдерами
- `RenderService` - рендеринг фигур через OpenGL
- `glPackage` - низкоуровневая работа с OpenGL

## Создание фигур

```go
triangle := &glPackage.Shape{
    Vertices: []float32{
        -0.5, -0.5,  
         0.5, -0.5,  
         0.0,  0.5,  
    },
    Indices: []uint32{0, 1, 2},
}
```
