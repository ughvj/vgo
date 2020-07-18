package main

import (
	"vgo/route"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := route.Init()
	router.Logger.Fatal(router.Start(":8000"))
}