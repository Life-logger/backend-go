package main

import (
	"io"
	"log"
	"os"

	"github.com/joho/godotenv"
	"taskbuddy.io/taskbuddy/server/controller"
	"taskbuddy.io/taskbuddy/server/router"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("empty .env " + err.Error())
	}

	//config.InitializeDB()

	multiWriter := io.MultiWriter(os.Stdout)
	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	app := router.Initialize(multiWriter)

	loginController := controller.LoginController{}
	app.Get("/callback", loginController.Login)
	// Start the server on port 3000

	log.Fatal(app.Listen(":22250"))
}
