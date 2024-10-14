# Запуск серверов:

### Пересборка образов:
- `Linux:` GOOS=linux docker compose -f docker-compose.dev.yml build
- `Windows:` GOOS=windows docker compose -f docker-compose.dev.yml build

### Непосредственно запуск:
- `Любая OS:` docker compose -f docker-compose.dev.yml up

### Остановка серверов:
- `Любая OS:` docker compose -f docker-compose.dev.yml down или Ctrl+C