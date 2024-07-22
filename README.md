# LogCollector

Приложения на Go, которое будет представлять собой систему мониторинга и анализа логов в реальном времени для веб-сервисов. Это приложение использует Kafka, Redis, ClickHouse, Docker, Grafana, Prometheus. Вот основные компоненты и их роль в системе:

### Основные компоненты системы

1. **Kafka**:
   - Kafka будет использоваться для приема и распределения логов в реальном времени. Логи от разных сервисов будут поступать в Kafka, где будут обрабатываться и передаваться в другие компоненты системы.

2. **Redis**:
   - Redis будет использоваться для кэширования и быстрой обработки данных. Например, для хранения состояния текущих обработок логов или временных данных.

3. **ClickHouse**:
   - В ClickHouse будут сохраняться все логи для последующего анализа и отчетов. Это позволит иметь долговременное хранилище с возможностью выполнения сложных SQL-запросов для анализа данных.

4. **Docker**:
   - Все компоненты системы будут упакованы в Docker-контейнеры, что обеспечит простоту развертывания и масштабирования приложения.

5. **Grafana**:
   - Grafana будет использоваться для визуализации данных. Она будет подключена к базе данных ClickHouse и к Prometheus для отображения метрик и логов в реальном времени.

6. **Prometheus**:
   - Prometheus будет использоваться для сбора метрик с разных компонентов системы. Эти метрики будут включать данные о производительности системы, количестве обработанных логов, задержках и т.д.


# Запуск приложения в docker-compose
```
docker-compose -f deployments/docker-compose.yml up -d
```