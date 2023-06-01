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

### Connect Mysql
 ```golang
// username = root, ip = 127.0.0.1:3306, database name = test
dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```
[Document Gorm](https://gorm.io/docs/)

## Fiber

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

