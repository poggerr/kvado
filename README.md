markdown
# KvadoTest

Это сервис на языке Golang, который позволяет выполнять поиск книг по автору и поиск авторов по книге, используя базу данных MySQL.

## Требования

- Golang
- MySQL
- Docker (для запуска базы данных с тестовыми данными)

## Установка и настройка

1. Склонируйте репозиторий:
   ```shell
   git clone https://github.com/your/repo.git

2. Установите зависимости:
   shell
   go mod download


3. Создайте базу данных MySQL и настройте параметры подключения в файле config.go.

4. Запустите миграции базы данных:
   shell
   go run migrate.go


## Запуск сервиса

1. Запустите базу данных MySQL с помощью Docker:
   shell
   docker-compose up -d


2. Запустите сервис:
   shell
   go run main.go


## Использование

### GRPC-запросы

Вы можете использовать любой клиент GRPC, чтобы отправлять запросы к сервису. Примеры запросов приведены ниже:

1. Запрос на поиск книг по автору:
   protobuf
   message AuthorRequest {
   int32 author_id = 1;
   }

   message Book {
   int32 book_id = 1;
   string book_name = 2;
   }

   message BookResponse {
   repeated Book books = 1;
   }

   service YourService {
   rpc FindBooksByAuthor(AuthorRequest) returns (BookResponse) {}
   }


2. Запрос на поиск авторов по книге:
   protobuf
   message BookRequest {
   string book_name = 1;
   }

   message Author {
   int32 author_id = 1;
   string author_name = 2;
   }

   message AuthorResponse {
   repeated Author authors = 1;
   }

   service YourService {
   rpc FindAuthorsByBook(BookRequest) returns (AuthorResponse) {}
   }


### REST API

Сервис также предоставляет REST API для выполнения запросов. Примеры эндпоинтов приведены ниже:

1. Запрос на поиск книг по автору:
   http
   GET /books?author_id=1


2. Запрос на поиск авторов по книге:
   http
   GET /authors?book_name=Book 1


## Тестирование

Вы можете запустить unit-тесты с помощью следующей команды:
shell
go test -v ./...


## Вклад

Если вы нашли ошибку или хотите внести улучшения, пожалуйста, создайте issue или отправьте pull request.

## Лицензия

Этот проект лицензирован под MIT License - подробности см. в файле [LICENSE](LICENSE.txt).

