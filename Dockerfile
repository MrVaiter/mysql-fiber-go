# Використовуємо офіційний образ Go
FROM golang:1.22.2

# Встановлюємо робочий каталог в контейнері
WORKDIR /app

# Копіюємо go mod та sum файли
COPY go.mod go.sum ./

# Завантажуємо всі залежності
RUN go mod download

# Копіюємо вихідний код в контейнер
COPY . .

# Збираємо додаток
RUN go build ./cmd/mysql-controller/main.go 
# Відкриваємо порт 3001
EXPOSE 3001

# Команда для запуску додатку
CMD ["./main"]