package main

import (
	database "assignment-2/config"
	"assignment-2/helpers"
	"assignment-2/routers"
	"os"
)

func main() {
	database.ConnectDb()

	helpers.LoadEnvChecking()

	var PORT = os.Getenv("PORT")

	routers.StartServer().Run(PORT)
}
