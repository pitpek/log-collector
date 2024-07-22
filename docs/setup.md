# Git

## Комманды для работы с git:

### Откатить последний коммит и сохранить изменения в рабочем каталоге
```
git reset --mixed HEAD~1
```

# Docker

## Комманды для работы с docker:

### Запуск контейнера в фоновом режиме
```
docker-compose up -d
```

### Просмотр логов в реальном времени
```
docker-compose logs -f
```

### Пересоздать контейнер и запустить его
```
docker-compose up --build    из директории /derloyments
docker-compose -f deployments/docker-compose.yml up --build    из корня проекта      
```

### Проверка статуса и портов
```
docker-compose ps
```

### Проверка логов Kafka и Zookeeper
```
docker-compose logs -f kafka
docker-compose logs -f zookeeper
```

### Bridge
```
sudo apt-get update -y
sudo apt-get install bridge-utils
```

### Создать файл миграций 
```
migrate create -ext sql -dir scripts/migrations -seq init
```
### Выполнить миграции
```
migrate -database postgres://postgres:postgres@localhost:5432/log_collector?sslmode=disable -path scripts/migrations up
```

### Посмотреть таблицу
```
docker-compose exec postgres psql -U postgres -d log_collector -c 'SELECT * FROM users;'
```

### Запуск линтера
```
golint ./...
```

### Подключиться к clickhouse через docker
```
docker exec -it clickhouse clickhouse-client -h localhost --user clickhouse --password clickhouse
```

### Открыть таблицу в clickhouse
```
use log_collector
```

### Необходимо запустить docker чтобы выполнить clickhouse_test
```
docker run -d --name clickhouse_test -p 9000:9000 -p 8123:8123 yandex/clickhouse-server
```