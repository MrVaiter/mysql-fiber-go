# Використовуємо офіційний образ Go для стадії збірки
FROM golang:1.22.2 as build

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

# Використовуємо офіційний образ golang для стадії запуску
FROM golang:1.22.2-alpine 

# Встановлюємо робочий каталог в контейнері
WORKDIR /app

# Копіюємо бінарник з стадії збірки
COPY --from=build ./app/main ./main

# Відкриваємо порт 3001
EXPOSE 3001

# Команда для запуску додатку
CMD ["./app", "/main"]