package main

import (
	"api-crypto/db"
	"api-crypto/route"
)

func main() {
	db.Init()
	e := route.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
