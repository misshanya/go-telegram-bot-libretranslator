**Telegram-бот переводчик**
========================
Основан на [LibreTranslate](https://github.com/LibreTranslate/LibreTranslate). 

Использован фреймворк [Golang Telegram Bot](https://github.com/go-telegram/bot).

Инструкция по сборке
-----------------------
#### Требуемые компоненты:
* Go (рекомендуется использовать последнюю версию)
* Make
* [Goose](https://github.com/pressly/goose)

#### Процесс сборки:
```shell
make build
```

#### Миграции базы данных

Выполнить миграцию

```shell
make migrate-up
```

Откатить миграцию

```shell
make migrate-down
```

#### Использование:
```shell
./bin/bot
```

Доступные команды
-----------------
* */start* - приветствие и создание клавиатуры
* */translate текст* - перевести текст

Описание аргументов
-------------
* --mlog - логировать сообщения при помощи middleware