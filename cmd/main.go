package main

import (
	"fmt"

	"github.com/ResetPlease/Babito/api/router"
)

func main() {
	fmt.Println("Babito service init")
	r := router.SetupRouter()
	r.Run(":8080")
}
