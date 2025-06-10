# WorkerPool-Vk

## Обзор
WorkerPool-Vk — это простая реализация пула воркеров на языке Go. Она позволяет динамически добавлять и удалять воркеров, которые обрабатывают задачи из общего входного канала.

## Структура проекта
```
.
├── .gitignore
├── cmd/
│   └── main.go <- пример реализации
├── go.mod
├── README.md
└── workerpool/
    ├── pool.go <- реализация пула
    └── worker.go <- реалиизация воркеров
```

## Использование
Запустите приложение с помощью команды:
```sh
git clone git@github.com:stupidcabbage/WorkerPool-Vk.git
cd WorkerPool-Vk
go run cmd/main.go
```
