package main

import (
	"gweb/core"
	"gweb/routes"
)

func main() {
	routes.DefinedRouter()
	err := core.Server(":8080", routes.Router()).ListenAndServe()
	if err != nil {
		panic(err)
	}
}
