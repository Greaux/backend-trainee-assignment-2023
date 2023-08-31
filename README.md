# Тестовое задание AvitoTech

## Инструкция по установке/запуску
### На ПК должно быть установлено:
1) Git (https://git-scm.com/downloads)
2) Docker (https://www.docker.com/products/docker-desktop/)
### Первая установка:
Если на ПК имеется Git и Docker, можно приступать в первому запуску:
1) Сначала необходимо скопировать репозиторий
- `git clone https://github.com/Greaux/backend-trainee-assignment-2023`
3) Перейти в директорию программы:
- `cd backend-trainee-assignment-2023`
4) Сборка контейнера при помощи Docker-compose
- `docker-compose build`
6) Запуск Docker контейнера
- `docker-compose up` 

### Команды после первой установки:
1) Завершение работы контейнера:
- `docker-compose down`
2) Запуск контейнера:
- `docker-compose up`
3) Пересборка контейнера
- `docker-compose up --build`

## Использованные фреймворки

1) Фреймворк для работы с запросами:
Gofiber - https://github.com/gofiber/fiber/v2

2) Библиотека для работы с БД
GORM - https://gorm.io
 
## Работа с API

### Создание юзера
POST /users
Добавление user в БД.

Тело запроса:
- `username` - уникальное имя для нового пользователя 
- `id` - уникальный ID пользователя (не обязательно)

Тело ответа:
- `ID` - ID который принадлежит новому пользователю
- `Username` - Имя которое присвоено новому пользователю
- `Segments` - Список сегментов которому принадлежит пользователь (на данном этапе null)

![Пример запроса (POSTMAN)](https://raw.githubusercontent.com/Greaux/backend-trainee-assignment-2023/main/Screenshots/%D0%A1%D0%BE%D0%B7%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%D0%AE%D0%B7%D0%B5%D1%80%D0%B0.png)
### Создание сегмента:
POST /segments
Добавление сегмента в БД.

Тело запроса: 
- `Name` - уникальное имя для нового сегмента

Тело ответа:
- `ID` - ID который принадлежит новому сегменту
- `Name` - Имя которое присвоено новому сегменту
- `Users` - Список пользователей которыи принадлежит сегмент (на данном этапе null)

![Пример запроса (POSTMAN)](https://raw.githubusercontent.com/Greaux/backend-trainee-assignment-2023/main/Screenshots/%D0%A1%D0%BE%D0%B7%D0%B4%D0%B0%D0%BD%D0%B8%D0%B5%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D0%B0.png)
### Удаление сегмента:
DELETE /segments
Удаление сегмента в БД.

Тело запроса: 
- `Name` - Имя сегмента который нужно удалить

Тело ответа:
- `ID` - ID который пренадлежал сегменту
- `Name` - Имя которое принадлежало сегменту
- `Users` - null

![Пример запроса (POSTMAN)](https://raw.githubusercontent.com/Greaux/backend-trainee-assignment-2023/main/Screenshots/%D0%A3%D0%B4%D0%B0%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D0%B0.png)
### Добавление пользователя в сегменты:
POST /editUserSegments
Добавление пользователя в сегмент в БД.

Тело запроса:
- `userid` - ID пользователя которого нужно добавить
- `segments` - список сегментов в которые его нужно добавить через запятую (например SEGMENT1,SEGMENT2)

Тело ответа:
- `ID` - ID который принадлежит пользователю
- `Username` - Имя пользователя
- `segments` - Список сегментов которые были добавлены пользователю в данном запросе

![Пример запроса (POSTMAN)](https://raw.githubusercontent.com/Greaux/backend-trainee-assignment-2023/main/Screenshots/%D0%94%D0%BE%D0%B1%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5%20%D1%8E%D0%B7%D0%B5%D1%80%D0%B0%20%D0%B2%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D1%8B.png)
### Удаление пользователя из сегментов:
DELETE /editUserSegments
Удаление пользователя из сегментов в БД.

Тело запроса: 
- `userid` - ID пользователя которого нужно добавить
- `segments` - список сегментов из которых его нужно удалить через запятую (например SEGMENT1,SEGMENT2)

Тело ответа:
- `ID` - ID который принадлежит пользователю
- `Username` - Имя пользователя
- `segments` - если null, значит пользователь удален из данных сегментов, либо не был в них
		
![Пример запроса (POSTMAN)](https://raw.githubusercontent.com/Greaux/backend-trainee-assignment-2023/main/Screenshots/%D0%A3%D0%B4%D0%B0%D0%BB%D0%B5%D0%BD%D0%B8%D1%8F%20%D1%83%20%D1%8E%D0%B7%D0%B5%D1%80%D0%B0%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D0%BE%D0%B2.png)
### Получение списка сегментов для пользователя по ID:
GET /user/:id
Получение списка сегментов из БД

ID пользователя передается напрямую в URL

Тело ответа:
- `ID` - ID сегмента
- `Name` - Имя сегмента
- `Users` - null

В случае если у пользователя нет сегментов ответ - record not found

![Пример запроса (POSTMAN)](https://raw.githubusercontent.com/Greaux/backend-trainee-assignment-2023/main/Screenshots/%D0%9F%D0%BE%D0%BB%D1%83%D1%87%D0%B5%D0%BD%D0%B8%D0%B5%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D0%BE%D0%B2%20%D1%8E%D0%B7%D0%B5%D1%80%D0%B0%20%D0%BF%D0%BE%20ID.png)
### Получение списка сегментов для пользователя по username:
GET /username/:Username
Получение списка сегментов из БД

* username пользователя передается напрямую в URL

Тело ответа:
- `ID` - ID сегмента
- `Name` - Имя сегмента
- `Users` - null
		
![Пример запроса (POSTMAN)](https://raw.githubusercontent.com/Greaux/backend-trainee-assignment-2023/main/Screenshots/%D0%9F%D0%BE%D0%BB%D1%83%D1%87%D0%B5%D0%BD%D0%B8%D0%B5%20%D1%81%D0%B5%D0%B3%D0%BC%D0%B5%D0%BD%D1%82%D0%BE%D0%B2%20%D1%8E%D0%B7%D0%B5%D1%80%D0%B0%20%D0%BF%D0%BE%20UserName.png)
