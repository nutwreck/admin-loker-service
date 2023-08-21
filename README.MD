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

- ### Swagger

```sh
$ go get -u github.com/swaggo/swag/cmd/swag
$ $HOME/go/bin/swag init
```
## Endpoint
[open swagger](/docs/index.html)

| **Nama**        | **Route**                  | **Method** |
| --------------- | -------------------------- | ---------- |
| **User**        |                            |            |
|                 | */api/v1/auth/ping*        | *GET*      |
|                 | */api/v1/auth/register*    | *POST*     |
|                 | */api/v1/auth/login*        | *POST*    |

struktur API nantinya ada
Admin Loker Service => internal API untuk Superadmin / Kebutuhan internal
Company Loker Service => API untuk user type perusahaan
User Loker Service => API untuk user pencari lowongan

hirego.id

chmod +x run.sh

./run.sh

https://www.kaggle.com/datasets/anggagewor/data-wilayah-republic-indonesia?resource=download