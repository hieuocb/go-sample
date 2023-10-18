package main

import (
	"gin-gonic-sample/routers"
)

func main() {
	router := routers.AddDemoRouter()
	router.Run(":8080")
}
