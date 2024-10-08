# 安裝

* 以下命令可以建議一些執行的sh檔跟設定檔
* 根據作業系統重新命名 ./Makefile.[*] 為 Makefile

```shell!
make setup
```

# 執行

```shell!
make air
```

# 資料庫遷移

> 使用[golang-migrate](https://github.com/golang-migrate/migrate)做資料庫遷移及做資料表版控

```shell!
make migration
```

# 更新日誌檔

* 可以產生change log

```shell!
make changeLog
```

---

## 目錄結構

```
├── LICENSE
├── fmp_dev.sql
├── Makefile
├── Makefile.example.linux
├── Makefile.example.windows
├── README.md
├── air.example.linux
├── air.example.windows
├── api
├── build
│   └── circleci
│       └── config.yml
├── cmd
├── config
│   ├── config.go.example
│   └── debug_config.go
├── deploy
├── docs
├── go.mod
├── go.sum
├── http
├── internal
│   ├── entity
│   │   ├── db
│   ├── interactor
│   │   ├── constants
│   │   │   └── constants.go
│   │   ├── helpers
│   │   ├── manager
│   │   ├── models
│   │   │   ├── page
│   │   │   │   └── page.go
│   │   │   ├── section
│   │   │   │   └── section.go
│   │   │   └── special
│   │   │       └── backend.go
│   │   ├── service
│   │   └── util
│   │       ├── aes.go
│   │       ├── base.go
│   │       ├── code
│   │       │   └── status.go
│   │       ├── connect
│   │       │   ├── mysql.go
│   │       │   └── postgres.go
│   │       ├── jwe.go
│   │       ├── log
│   │       │   └── log.go
│   │       ├── page.go
│   │       ├── time.go
│   │       ├── util.go
│   │       └── uuid.go
│   ├── presenter
│   └── router
│       ├── middleware
│       └── router.go
├── main.go
├── migrations
└── tools
    ├── autoMigrate
    │   └── main.go
    ├── log
    │   └── run.go
    └── testData
        └── main.go
```

### LICENSE

* 授權檔案 MIT License

### fmp_dev.sql

* 資料庫備份檔

### Makefile.example.linux / Makefile.example.windows

* 根據作業系統重新命名 ./Makefile.[*] 為 Makefile
* 詳細記錄了,所有能夠使用的命令集

### air.example.linux / air.example.windows

* 熱加載設定檔

### /api

* API DOC 置放

### /build

* CI/CD 設定檔

### /cmd

* 本專案的主要應用程式

### /config

* 設定檔

### /deploy

* 被編譯過後的檔案

### /docs

* golang doc

### go.mod

* go mod檔

### /http

* restful api 測試文件

### /internal

* 私有應用程式和函式庫的程式碼,是你不希望其他人在其應用程式或函式庫中匯入的程式碼.

### /internal/entity

* 對應資料庫的CRUD

### /internal/entity/db

* 對應資料表結構檔

### /internal/interactor

* 可共用或不可共用的函式庫的程式碼

### /internal/interactor/constants

* 置放常數資料夾

### /internal/interactor/helpers

* 置放一些共用的manager

### /internal/interactor/manager

* 置放交互調度程式

### /internal/interactor/models

* 置放共用結構檔

### /internal/interactor/models/page

* 置放有關於分頁的結構檔

### /internal/interactor/models/section

* 置放有關於時間的結構檔

### /internal/interactor/models/special

* 置放有關於後端共用結構檔

### /internal/interactor/service

* 置放有關於所有的服務的程式碼

### /internal/interactor/util

* 置放一些小工具,實用程序

### /internal/interactor/util/code

* 置放錯誤代碼或是回應代碼

### /internal/interactor/util/connect

* 置放對應資料庫的連線

### /internal/presenter

* 對應前端第一接觸的地方,API文件註解的地方,驗證輸入的地方

### /internal/router

* 置放路由設定

### main.go

* main func

### /migrations

* 放置資料庫的SQL檔案

### /tools

* 可以置放一些小工具

# License

Vodka is released under the MIT license.