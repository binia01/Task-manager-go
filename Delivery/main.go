package main

import (
	routers "task-manager-go/Delivery/router"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
