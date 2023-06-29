# Golang-API

This project use packages gorm fiber and use database is MySQL. 

![](https://drive.google.com/uc?id=11wXLcPIqGoljJuyVU2qwolvZscMTblVq)

### Usage
```golang
// command run api
go run main.go
```
[Document Golang](https://go.dev/doc/)

## Gorm
GORM is an ORM library that handles multiple databases such as MySQL, PostgreSQL, and SQLite.

### Installation
```golang
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

### Import
```golang
import (
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)
```

### Connect Mysql
 ```golang
// username = root, ip = 127.0.0.1:3306, database name = test
dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```
[Document Gorm](https://gorm.io/docs/)

## Fiber
Fiber is a modern web framework for the Go programming language. It aims to provide a fast, flexible, and efficient way to build web applications and APIs. Fiber is inspired by Express.js, a popular web framework for JavaScript.

### Installation
```golang
go get github.com/gofiber/fiber/v2
```

### Import
```golang
import (
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)
```

### Config
 ```golang
// Default config
app := fiber.New()

// Custom config
app.Use(cors.New(cors.Config{
	AllowCredentials: true,
}))

// set folder public is static
app.Static("/static", "./public")

// set port
app.Listen((":8080"))
```
[Document Fiber](https://docs.gofiber.io/)

## Mysql
![](https://drive.google.com/uc?id=100mSkPFRLxovYDAHOUUSGVGaDkuUf0rN)

