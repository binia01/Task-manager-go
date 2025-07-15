package main

import (
	"task-manager-go/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
