package main

import (
	"github.com/vite-va/go-ecommerce-backend-api/internal/initialize"
)

func main() {
	// r := router.NewRouter()
	// r.Run(":8002") // listen and server on 0.0.0.0:8080 (for windows "localhost:8080")
	initialize.Run()
}
