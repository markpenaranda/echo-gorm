package main

import (
	"github.com/markpenaranda/echo-gorm/db"
	"github.com/markpenaranda/echo-gorm/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
