# Golang Loker

API backend rest Admin Loker Service using Gin Framework and Gorm ORM Database.

## Command

- ### App Lifecyle

```sh
$ go run main.go || make go run
```

- ### Docker Lifecycle

```sh
$ docker-compose up -d --build
```
## Endpoint

| **Nama**        | **Route**                  | **Method** |
| --------------- | -------------------------- | ---------- |
| **User**        |                            |            |
|                 | */api/v1/auth/ping*        | *GET*      |
|                 | */api/v1/auth/register*    | *POST*     |
|                 | */api/v1/auth/login*        | *POST*    |