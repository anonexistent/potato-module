# Dockerfile
FROM golang:latest

# Создаем директорию для приложения
WORKDIR /app

# Копируем все файлы в контейнер
COPY . .

# Скачиваем зависимости и собираем приложение
RUN go mod download
RUN go build -o main .

# Указываем команду для запуска приложения
CMD ["/app/main"]