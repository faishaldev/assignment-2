// @title Orders API
// @version 1.0.0
// @description Simple Orders API
// @termsOfService http://swagger.io/terms
// @contact.name Musthafa Faishal
// @contact.email musthafafaisha@gmail.com
// @license.name Apache 2.0
// license.url http://www.apache.org/licenses/LICENSE-
// @host localhost:8080
// @BasePath /

package main

import (
	db "assignment-2/models"
	"assignment-2/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	db.Init()

	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file: ", e)
	}

	router := routers.SetupRouter()
	port := os.Getenv("SERVER_PORT")

	if len(os.Args) > 1 {
		reqPort := os.Args[1]

		if reqPort != "" {
			port = reqPort
		}
	}

	if port == "" {
		port = "8080"
	}

	type Job interface {
		Run()
	}

	router.Run(":" + port)
}
