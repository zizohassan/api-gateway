package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB = nil
var err error

func main() {
	g := SetupRouter()
	g.Run(":5050")
	//autotls.Run(g , "site.com")
}