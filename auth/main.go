package main

import (
	"auth/route"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := route.Init()
	router.Logger.Fatal(router.Start(":2434"))
}