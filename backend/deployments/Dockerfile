# Используем в качестве базового образа Golang
FROM golang:1.21.2

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем все остальные файлы из текущей директории в контейнер
COPY ../ ./

# Выполняем команду для управления зависимостями модулей Go
RUN go mod tidy

# Собираем ваше приложение
RUN go build -o app cmd/app/main.go

# Указываем команду для запуска контейнера
CMD ["./app"]
