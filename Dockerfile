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
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main ./cmd/mysql-controller/main.go 

# Використовуємо офіційний образ golang для стадії запуску
FROM alpine:latest  

# Встановлюємо робочий каталог в контейнері
WORKDIR /app

# Копіюємо бінарник з стадії збірки
COPY --from=build /app/main .

# Відкриваємо порт 3001
EXPOSE 3001

# Команда для запуску додатку
CMD ["/app/main"]