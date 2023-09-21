package frameworks

import (
	"github.com/ezequiel-bugnon/brandmonitor/controllers"
	"github.com/gofiber/fiber"
)

type FiberFramework struct {
	controller controllers.Controllers
	app        *fiber.App
}

func NewApp(controller controllers.Controllers, app *fiber.App) *FiberFramework {
	return &FiberFramework{
		controller: controller,
		app:        app,
	}
}

func (f *FiberFramework) Conection() {
	f.app.Listen(":3000")
}

func (f *FiberFramework) Controllers() {

	f.app.Post("/api/v1/file", func(c *fiber.Ctx) {
		f.controller.PostFile(c)
	})

	f.app.Get("/api/v1/data", func(c *fiber.Ctx) {
		f.controller.GetData(c)
	})

}
