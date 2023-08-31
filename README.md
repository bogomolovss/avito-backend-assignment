# Тестовое задание для стажёра Backend
## Сервис динамического сегментирования пользователей

### Как запустить проект

Скопируйте проект с github.
Запустите docker-compose.yml.
```
sudo docker-compose up --build
```
Зайдите через docker-compose bash в контейнер базы данных.
```
sudo docker-compose exec postgres bash
```
Создайте базу данных (также консоль попросит ввести пароль для бд : postgres).
```
createdb -h localhost -p 5432 -U postgres avito_db --password
```
Остановите docker-compose и выполните команду
```
sudo docker-compose up
```

## Примеры API
Все методы API описаны в avito_user_segment_postman.json. Этот файл можно импортировать в Postman.

### Метод создания сегмента. Принимает slug(название сегмента)
```
 [POST]http://localhost:8080/api/segments
{
    "title": "AVITO_MESSAGE_100"
}
```

### Метод удаления сегмента. Принимает slug (название) сегмента.
```
 [DELETE]http://localhost:8080/api/segments/AVITO_DISCOUNT_50
```

### Метод добавления пользователя в сегмент. Принимает список slug (названий) сегментов которые нужно добавить пользователю, список slug (названий) сегментов которые нужно удалить у пользователя, id пользователя.
(параметр в url = id пользователя)
```
 [POST]http://localhost:8080/api/user_segments/1
[
    {"title": "AVITO_DISCOUNT_90"},
    {"title": "AVITO_MESSAGE_100"}
]
```
В качестве ответа получаем информацию о пользователе, которому добавили сегменты и **добавленные** сегменты.
Пример:
```
{
    "data": {
        "ID": 1,
        "CreatedAt": "2023-08-31T00:32:49.196113+03:00",
        "UpdatedAt": "2023-08-31T13:26:47.095080371+03:00",
        "DeletedAt": null,
        "username": "Diman",
        "Segments": [
            {
                "ID": 5,
                "CreatedAt": "2023-08-31T13:08:46.894525+03:00",
                "UpdatedAt": "2023-08-31T13:08:46.894525+03:00",
                "DeletedAt": null,
                "title": "AVITO_DISCOUNT_90",
                "Users": null
            },
            {
                "ID": 7,
                "CreatedAt": "2023-08-31T13:26:43.537493+03:00",
                "UpdatedAt": "2023-08-31T13:26:43.537493+03:00",
                "DeletedAt": null,
                "title": "AVITO_MESSAGE_100",
                "Users": null
            }
        ]
    }
}
```
### Метод получения активных сегментов пользователя. Принимает на вход id пользователя. (Имеет две версии)
```
 [GET]http://localhost:8080/api/user_segments/1
```
Пример ответа API:
```
{
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2023-08-31T13:09:38.377072Z",
            "UpdatedAt": "2023-08-31T13:09:38.377072Z",
            "DeletedAt": null,
            "title": "AVITO_MESSAGE_100",
            "Users": null
        },
        {
            "ID": 2,
            "CreatedAt": "2023-08-31T13:58:59.865234Z",
            "UpdatedAt": "2023-08-31T13:58:59.865234Z",
            "DeletedAt": null,
            "title": "AVITO_DISCOUNT",
            "Users": null
        },
        {
            "ID": 3,
            "CreatedAt": "2023-08-31T13:59:05.910368Z",
            "UpdatedAt": "2023-08-31T13:59:05.910368Z",
            "DeletedAt": null,
            "title": "AVITO_VOICE_MESSAGE",
            "Users": null
        }
    ]
}
```

Или если вызывать упрощенную версию (выводит только названия сегментов пользователя)
```
 [GET]http://localhost:8080/api/user_segments_lite/1
```
Пример ответа:
```
{
    "data": [
        {
            "title": "AVITO_MESSAGE_100"
        },
        {
            "title": "AVITO_DISCOUNT"
        },
        {
            "title": "AVITO_VOICE_MESSAGE"
        }
    ]
}
```

Такие методы как добавление/удаление/получение пользователя/сегмента, списка пользователей/сегментов описаны в
**avito_user_segment_postman.json**.
