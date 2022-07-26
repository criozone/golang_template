# Wss v2


###### Локальная разработка
- Скопировать **docker-compose.override.yml.example** в **docker-compose.override.yml** 
и добавить в него нужные правки по необходимости. Можно переопределить любые
значения заданные в основном docker-compose файле. Например, можно заменть 
 значение переменной окружения LOCAL_HOST на свой выбор
- Добавить в /etc/hosts значение из переменной VIRTUAL_HOST
- Из корня проекта запустить **./docker-start**. Если после старта необходимо перестроить 
конкретный сервис можно запустить: **./docker-start [service-name]** (например: 
**./docker-start wss-go**). 
- Тестовая страница будет доступна по корневому пути по хосту из VIRTUAL_HOST.
По умолчанию: **http://wss-v2.local**

###### Переменные окружения:
- APP_WSS_ADDR - ip адресс который будет слушать сервер. По умолчанию ""
(пусто, тоесть 0.0.0.0) 
- APP_WSS_PORT - порт который будет слушать сервер. По умолчанию 8080
- APP_SENTRY_DSN - TODO: find out and describe
- APP_REDIS_HOST - адресс по которому нужно стучатся в редис
- APP_REDIS_PORT - порт  по которому нужно стучатся в редис
- APP_REDIS_PREFIX - TODO: find out and describe
- SERVICE_NAME - идентификатор сервиса. Используется при логировании.
- DEVELOPMENT - dev mode. Включить для локальной разработки
- VIRTUAL_HOST - используется при локальной разработки чтоб достучатся к приложению через
сервис wss-nginx. Также используется для создания самоподписного сертификата для локальной 
разработки
- REQUEST_ID_HEADER_NAME - название заголовка для идентификатора запроса через все 
микросервисы. По умолчанию: X-Request-Id
- PARENT_SPAN_ID_HEADER_NAME - название  заголовка для идентификатора родительского 
запроса
- MODULE_PATH - идентификатор (путь) модуля Голанг. (Используется в случае
необходимости запуска go mod init)
- HOST_UID=${HOST_UID} - идентификатор локального пользователя. Используется в 
локальной разработке чтоб бы избежать проблемы с правами
- HOST_GID=${HOST_GID} - идентификатор группы локального пользователя. Используется в
  локальной разработке чтоб бы избежать проблемы с правами