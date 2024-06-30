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


