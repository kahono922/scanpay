package routerHandler

import(
	"github.com/gofiber/fiber/v2"
	_"github.com/gofiber/template/html"
)

func Home(c *fiber.Ctx) error{
  return c.Render("home",&fiber.Map{})
}
