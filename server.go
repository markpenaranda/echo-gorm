package main

import (
	"echo-gorm/db"
	"echo-gorm/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
