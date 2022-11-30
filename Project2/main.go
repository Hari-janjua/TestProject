package main

import (
	"Project2/Routes"
)

func main() {
	router := Routes.SetupRouter()
	router.Run()
}
