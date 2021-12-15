package router

import(
	"github.com/gofiber/fiber/v2"
	"github.com/kahono922/scanpay/routerHandler"
)

func SetUpRoutes(app *fiber.App){
	app.Get("/",routerHandler.Home)
}
