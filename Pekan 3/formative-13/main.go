package main

import (
	"formative-13/routers"
)

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}
