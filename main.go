package main

import (
	"github.com/VituSuperMEg/take-tree-go/config"
	"github.com/VituSuperMEg/take-tree-go/routes"
)

func main() {
	config.InitDB()
	routes.InitApi()
}
