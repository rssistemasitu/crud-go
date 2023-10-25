package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rssistemasitu/crud-go/internal/controller/routes"
)

// func main() {
// 	err := godotenv.Load("configs/.env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	print(os.Getenv("TEST"))
// }

func main() {
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
