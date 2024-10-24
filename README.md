# Как работать с сервисами:

### Пересборка образов:
- docker compose -f docker-compose.dev.yml build

### Запуск:
- docker compose -f docker-compose.dev.yml up llm_server database -d
- Небольшая пауза для инициализации DB
- docker compose -f docker-compose.dev.yml up backend -d

### Остановка и удаление:
- docker compose -f docker-compose.dev.yml down