## Запуск сервера
### Создать БД postgreSQL

### Создать `.env.dev`

```bash
touch server/.env.dev
```

#### Необходимо обязательно добавить следующие перменные 
```
DB_NAME=DBName
DB_HOST=DBHost
DB_PORT=DBPort
DB_USER=DBUser
DB_PASS=DBPass
DB_SSL_MODE=DBSSL

JWT_EXPIRATION=6000
```
### Установливаем библиотеку для миграций

```bash
cd server
go get -u github.com/golang-migrate/migrate/v4
```
### Запускаем сервер
```bash
go run cmd/main.go
```

### Прогоняем миграции
```bash
make migrate_up
```

## Запуск клиента

### Создать `.env.dev`

```bash
touch client/.env.dev
```

#### Необходимо обязательно добавить следующие перменные
```
SERVER_HOST=http://localhost:8000
```
### Билдим приложение

```bash
cd client
go build .
```
### Используем клиента
```bash
./client [command]
```

### Чтобы узнать какие команды существуют:
```bash
./client --help
```

### Для использования сервера необходимо сначала зарегестрироваться:
```bash
./client auth registration -e MyEmail -p MyPassword 
```
### Затем залогиниться:
```bash
./client auth login -e MyEmail -p MyPassword 
```

### Для загрузки файла на сервер:
```bash
./client upload -f PathToYourFile
```

### Для просмотра загруженных файлов на сервере:
```bash
./client index
```

### Для загрузки файла с сервера:
```bash
./client download -n NameFileOnServer
```
P.S. Загрузку файла на сервер и выгрузку сделал немножко разными способами.
