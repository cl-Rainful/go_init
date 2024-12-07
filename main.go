package main

import (
	router "TestDemo/routes"
	"TestDemo/utils"
)

func main() {

	router.InitRouter().Run(utils.HttpPort)
}
