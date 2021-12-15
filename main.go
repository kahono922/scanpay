package main

import(
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/kahono922/scanpay/database"
	"github.com/gofiber/template/html"
	"github.com/kahono922/scanpay/router"
)

func main(){
  engine := html.New("./views",".html")
	app := fiber.New(fiber.Config{
	  Views: engine,
	})
	router.SetUpRoutes(app)
	
	database.Connect()
	log.Fatal(app.Listen(":3000"))
}

