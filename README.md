
## go-calc

# Структура проекта

---
```go
go-calc  
│  
├── cmd  
│   └── main.go  
│  
├── internal  
│   └── application  
│       ├── application.go  
│       └── application_test.go  
│  
├── pkg  
│   └── calculation  
│       ├── calculation.go  
│       └── calculation_test.go  
│  
├── go.mod  
└── README.md
```
---
### Простой калькулятор на Go

Простой калькулятор на языке программирования Go, который может выполнять арифметические операции с арабскими числами.

### Файлы Конфигурации

Для смены порта необходимо изменить переменную *config.Addr* в методе *application*:
```go
func ConfigFromEnv() *Config {
    config := new(Config)  
    config.Addr = os.Getenv("PORT")  
    if config.Addr == "" {  
        config.Addr = "8080"  
    }  
    return config  
}
```
### Предварительные требования
Перед началом работы с проектом необходимо убедиться, что на вашем компьютере установлены следующие инструменты и компоненты:  
- Go (версия 1.23.0 или выше): Язык программирования Go необходим для разработки и запуска серверной части приложения. Вы можете скачать его с официального сайта [go.dev](https://go.dev/).
- Git: Система контроля версий, необходимая для клонирования проекта из репозитория на GitHub. Скачать Git можно с сайта [git-scm.com](https://git-scm.com/).
Убедитесь, что все эти компоненты установлены и правильно настроены на вашем компьютере перед продолжением работы с проектом.
## Начало работы  
#### Копирование проекта с GitHub  
Для начала работы с проектом необходимо клонировать репозиторий на локальный компьютер, используя следующую команду в терминале:
```go
git clone https://github.com/neandrson/go-calc.git  
```
После клонирования репозитория, перейдите в папку проекта для выполнения последующих команд:
```go
cd go-calc
```
#### Установка зависимостей  
go.mod: Определяет зависимости проекта модуля.  
Чтобы установить все зависимости проекта, перейдите в директорию проекта, откройте в ней терминал и выполните команду:  
```go
go mod tidy
```
Эта команда скачает и установит все необходимые зависимости, указанные в файле go.mod.
#### Инструкция по запуску  
Для запуска приложения go-calc необходимо выполнить команду:  
```go
go run ./cmd/main.go
```
После успешного запуска сервера в терминале появится сообщение: "Starting server on port 8080...".
#### Описание методов  
##### ***Метод application***  
Метод *application* управляет распределением задач и предоставляет API для взаимодействия с *calculation*.  
##### ***Метод calculation***  
Метод *calculation* обрабатывает вычислительные задачи, отправленные *application*. Он предоставляет API для приема и выполнения вычислений.
#### Запуск вычисления  
Для отправки вычислительной задачи на сервер калькулятора, используйте приложение ***curl*** в командной строке ОС Windows с запросом:
```go
curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"2+2*2\" }"
```
Для передачи выражения используется POST запрос.  
В запросе должны использоваться арифметические выражения с арабскими числами.  
#### Формат ввода  
Калькулятор поддерживает следующие операторы: ***"+", "-", "\*", "/"***, а также приоритеты вычисления с поддержкой ***"(", ")"***.
#### Запуск тестов  
Запуск теста из папки модуля используется:
```go
go test -v
```
Запуск всех тестов используйте:
```go
go test -v ./...
```
### Примеры ответов  
#### Успешный результат
**Запрос:**
```go
{  
   "expression": "2+2*2"
}
```
**Ответ:**
```go
{
   "result": "6.000000"
}
```

#### Неверный запрос
**Запрос:**
```go
curl --location "http://localhost:8080/api/v1/calculate" --header "Content-Type: application/json" --data "{ \"expression\": \"2+2+\" }"
```
**Ответ:**
```go
{
   "error": "invalid expression"
}
```
Всего программа может возвращать 6 видов ошибок:

<table>
	<thead>
		<tr>
			<th>№</th>
			<th>Пример тела запроса/Метод запроса</th>
			<th>Ответ сервера</th>
			<th>Описание ошибки</th>
		</tr>
	</thead>
	<tbody>
		<tr>
			<td>1</td>
			<td>{"expression": "2+abc"}/POST</td>
			<td>invalid character</td>
			<td>В выражении присутствуют недопустимые символы.</td>
		</tr>
		<tr>
			<td>2</td>
			<td>{"expression": "2+3)"}/POST</td>
			<td>unmatched parentheses</td>
			<td>В выражении присутствует открывающая или закрывающая скобка без пары.</td>
		</tr>
		<tr>
			<td>3</td>
			<td>{"expression": "2/0"}/POST</td>
			<td>division by zero</td>
			<td>Попытка деления на ноль.</td>
		</tr>
		<tr>
			<td>4</td>
			<td>{"expression": "2+2"}/GET</td>
			<td>method not allowed</td>
			<td>Неверный HTTP-метод (должен быть POST).</td>
		</tr>
		<tr>
			<td>5</td>
			<td>{"express"/POST</td>
			<td>expression is not valid</td>
			<td>Неверный формат тела запроса.</td>
		</tr>
		<tr>
			<td>6</td>
			<td>{"expression": "2+"}/POST</td>
			<td>invalid expression</td>
			<td>Некорректное выражение.</td>
		</tr>
	</tbody>
</table>

Контакт для связи в телеграме @Neandrs
