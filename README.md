# Golang-API

This project use packages gorm fiber and use database is MySQL. 

<img src="https://drive.google.com/uc?id=11wXLcPIqGoljJuyVU2qwolvZscMTblVq" />

### Usage
```golang
// command run
go run main.go
```

## Gorm
### Connect Mysql

 ```golang
// username = root, ip = tcp, database = test
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
```
[Document Fiber](https://docs.gofiber.io/)

## Mysql
![](https://labs.mysql.com/common/logos/mysql-logo.svg?v2)

