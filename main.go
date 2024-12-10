package main

import (
	"TestDemo/model"
	router "TestDemo/routes"
	"TestDemo/utils"
)

func main() {
	model.InitDB()

	router.InitRouter().Run(utils.HttpPort)
}
