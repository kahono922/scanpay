package main

import(
	"log"
	"github.com/gofiber/fiber/v2"
	_"github.com/go-sql-driver/mysql"
)

func main(){
	app := fiber.New()

	app.Get("/",func(c *fiber.Ctx)error{
		return c.SendString("Hello")
	})
	log.Fatal(app.Listen(":3000"))
}

