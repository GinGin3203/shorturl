# url_shortener

## Запуск

```shell
$ docker-compose -f docker-compose.local.yaml up
```

## Пример работы

В примерах используется тулза [rs/curlie](https://github.com/rs/curlie).

### Создание новой сокращенной ссылки

```shell
$ curlie -v POST localhost:8080/api/v1/shortener url="https://www.youtube.com/watch?v=dQw4w9WgXcQ" 

*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
POST /api/v1/shortener HTTP/1.1
Host: localhost:8080
User-Agent: curl/7.74.0
Content-Type: application/json
Accept: application/json, */*
Content-Length: 53

{
    "url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
}


* upload completely sent off: 53 out of 53 bytes
* Mark bundle as not supporting multiuse
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 25 Feb 2022 03:23:08 GMT
Content-Length: 34

* Connection #0 to host localhost left intact
{
    "short_url": "https://o.co/AAAAB"
}
```

### Получение оригинальной ссылки

```shell
$ curlie "localhost:8080/api/v1/shortener?url=https://o.co/AAAAB"                                 

HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 25 Feb 2022 03:26:11 GMT
Content-Length: 62

{
    "original_url": "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
}
```
