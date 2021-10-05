package main

import (
	"book-api/config"
	"book-api/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}