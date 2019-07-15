package main

import (
	. "urlshorner/migration"
	"urlshorner/routes"
)

func main() {
	Init()

	routes.AllRoutes()
}
