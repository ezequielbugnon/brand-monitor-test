package frameworks

import (
	"github.com/ezequiel-bugnon/brandmonitor/controllers"
	"github.com/gofiber/fiber"
)

type FiberFramework struct {
	controller controllers.Controllers
	app        *fiber.App
}

func NewApp(controller controllers.Controllers) *FiberFramework {
	return &FiberFramework{
		controller: controller,
		app:        fiber.New(),
	}
}

func (f *FiberFramework) Conection() {
	f.app.Listen(":3000")
}

func (f *FiberFramework) Controllers() {

	f.app.Get("/", func(c *fiber.Ctx) {
		f.controller.PostFile(c)
	})

	f.app.Get("/dardo", func(c *fiber.Ctx) {
		f.controller.GetData(c)
	})

}
