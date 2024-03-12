# task-list

Restful task list API 

## 環境
- Go 版本: 1.18.2
- Gin 框架版本: 1.9.1

## URL

- URL需加上版本號 /v1

## 授權
- 對task做CRUD都需要進行授權驗證。
- 可呼叫 `/apikey` 取得API KEY。
- 每個API KEY只有一分鐘的效期，並且僅供一次使用。

## 專案結構
```
project-root/
│
├── config/
│   ├── (configuration files)
│
├── controllers/
│   ├── (controllers for handling HTTP requests)
│
├── middleware/
│   ├── (middleware for request/response processing)
│
├── models/
│   ├── (unities models)
│
├── repositories/
│   ├── (repositories for interacting with the database)
│
├── routers/
│   ├── (router setup and endpoint definitions)
│
├── services/
│   ├── (business logic services)
│
├── utils/
│   ├── (utility functions and helpers)
│
├── main.go
```
