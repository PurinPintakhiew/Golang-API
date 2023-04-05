# Golang-API

### Usage
```golang
  # command run
  go run main.go
```

## Gorm
### Connect mysql by gorm

 ```golang
   /* username = root
      ip = tcp
      database = test */
      
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```
[Document Gorm](https://gorm.io/docs/)

## Fiber

[Document Fiber](https://docs.gofiber.io/)

## Mysql
