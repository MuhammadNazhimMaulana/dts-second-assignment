package main

import (
	"fmt"

	_ "Assignment_2/database"
	"Assignment_2/routers"
)

func main() {
	fmt.Println("Start Server")
	routers.GetRouter().Run(":8000")
}
