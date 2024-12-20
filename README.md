
# go-calc

Структура проекта

go-calc
│
├── cmd
│ ├── main.go
│ └── main_test.go
│
├── internal
│ └── application
│    ├── application.go
│    └── application_test.go
│
├── pkg
│ └── calculation
│    ├── calculation.go
│    └── calculation_test.go
│
├── go.mod
└── README.md

Простой калькулятор на Go

Простой калькулятор на языке программирования Go, который может выполнять арифметические операции с арабскими числами.

Файлы Конфигурации
go.mod: Определяют зависимости проекта модуля.

Предварительные требования
Перед началом работы с проектом необходимо убедиться, что на вашем компьютере установлены следующие инструменты и компоненты:

- Go (версия 1.23.0 или выше): Язык программирования Go необходим для разработки и запуска серверной части приложения. Вы можете скачать его с официального сайта golang.org.
- Git: Система контроля версий, необходимая для клонирования проекта из репозитория на GitHub. Скачать Git можно с сайта git-scm.com.
Убедитесь, что все эти компоненты установлены и правильно настроены на вашем компьютере перед продолжением работы с проектом.

Начало работы
Копирование проекта с GitHub
Для начала работы с проектом необходимо клонировать репозиторий на локальный компьютер, используя следующую команду в терминале:

git clone https://github.com/neandrson/go-calc.git

После клонирования репозитория, перейдите в папку проекта для выполнения последующих команд.

Установка зависимостей
Чтобы установить все зависимости проекта, перейдите в директорию проекта, откройте в ней терминал и выполните команду: go mod tidy

Эта команда скачает и установит все необходимые зависимости, указанные в файле go.mod.

Инструкция по запуску
Для запуска приложения go-calc необходимо выполнить следующие шаги:

Сервер go-calc - запуск
Для запуска сервера go-calc откройте новое терминальное окно в директории go-calc и выполните команду: go run ./cmd/main.go
После успешного запуска сервера в терминале появится сообщение: "Starting server on port 8080...".

Описание методов
Описание метода application
Сервер application управляет распределением задач и предоставляет API для взаимодействия с calculation. Ниже приведены основные методы API и примеры их использования с помощью curl.

Описание метода calculation
Сервер калькулятора обрабатывают вычислительные задачи, отправленные application. Он предоставляют API для приема и выполнения вычислений.

Запуск вычисления
Для отправки вычислительной задачи на сервер калькулятора, используйте следующий curl запрос:

curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{
    \"expression\": \"2+2*2\"
    }"

Пример ответа сервера:

{
    "result": "6.000000"
}

Для передачи выражения используется POST запрос

В запросе могут использоваться арифметические выражения в формате "число оператор число...".

Формат ввода

Калькулятор поддерживает следующие операторы: "+", "-", "*", "/".

Выражения могут быть введены в формате арабских чисел.

Ограничения
Калькулятор принимает числа от 1 до 9 включительно.
Результатом операции деления является вещественное число с точностью до 6 разряда.

Примеры ответов

{
    "expression": "2+2*2"
}

{
    "result": "6.000000"
}

и ответ сервера кодом 200, если выражение вычислено успешно

{
    "expression": "2+2*2a"
}

{
    "error": "Expression is not valid"
}

и ответ сервера кодом 422, если входные данные не соответствуют требованиям приложения
