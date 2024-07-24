# Youtube thumbnail service

Микросервис для получения первью видео по ссылке на видео.

Используемые технологии:
- Redis (Для хранилища кэша)
- Docker (для запуска сервиса)
- gRPC (Для вызова функций другими сервисами или клинетами )
- google.golang.org/api (для обращение к API youtube)



## Getting Started

Для запуска сервиса необходимо предварительно:
- Зарегистрировать приложение в Google Cloud Platform: [Документация](https://developers.google.com/workspace/guides/create-project)
- Создать сервисный аккаунт и его секретный ключ: [Документация](https://developers.google.com/workspace/guides/create-credentials)
- Добавить секретный ключ в директорию secrets
- Добавить .env файл в директорию с проектом и заполнить его данными из .env.example, указав `YOUTUBE_CONFIG_PATH=secrets/your_credentials_file.yaml`
- Файл `your_credentials_file.yaml` заполнить в форме `api_key:"your key"`
- Опционально, настроить `congig/config.yaml` под себя

Для запуска сервиса без интеграции с Google Drive достаточно заполнить .env файл,
оставив переменную `GOOGLE_DRIVE_JSON_FILE_PATH` пустой
## Usage


Запустить сервис можно с помощью команды `make compose-up`

Для запуска тестов необходимо выполнить команду `make test`, для запуска тестов с покрытием `make cover` и `make cover-html` для получения отчёта в html формате


