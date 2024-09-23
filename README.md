**Telegram-бот переводчик**
========================
Основан на [LibreTranslate](https://github.com/LibreTranslate/LibreTranslate). 

Использован фреймворк [Golang Telegram Bot](https://github.com/go-telegram/bot).

Инструкция по сборке
-----------------------
#### Требуемые компоненты:
* Go (рекомендуется использовать последнюю версию)

#### Процесс сборки:
```shell
go build ./cmd/bot
```

Доступные команды
-----------------
* */start* - приветствие и создание клавиатуры
* */translate текст* - перевести текст