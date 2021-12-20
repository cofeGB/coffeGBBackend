## CoffeGB ##

CoffeGB - это сервис для ... (TODO: описание).

В текущем репозитории ведется разработка бекэнда сервиса.

## Начало работы ##

Для ознакомления можно воспользоваться [демо-сервером](https://coffegb.herokuapp.com/api).

Спецификацию на API вы можете найти [тут]().

## Установка ##

Самый простой путь для локальной установки - это использование docker.

Для этого необходимо:

1) установить git, doker-ce, docker-compose на вашу машину
1) клонировать репозиторий
    ```
    git clone https://github.com/cofeGB/coffeGBBackend.git
    cd coffeGBBackend
    ```
1) создать в корне проекта файл ```.env``` с переменными окружения:
    ```
    COMPOSE_PROJECT_NAME=coffebg_dev
    VERSION=0.1.0

    PG_DATABASE=coffebg_dev
    # имя и пароль необходимо сменить
    PG_USER=coffe
    PG_PASS=coffe
    ```
1) скопировать манифест docker-compose:
    ```
    cp docker-compose.dev.yml docker-compose.yml
    ```
1) запустить проект
    ```
    docker-compose up -d
    ```

API-сервер доступен по адресу: [http://<IP машины с докером>:8123/api](http://localhost:8123/api)


## Версионирование ##

Мы используем [SemVer](http://semver.org/). Доступные версии можно увидеть  [в тегах репозитория](https://github.com/cofeGB/coffeGBBackend/tags). 
