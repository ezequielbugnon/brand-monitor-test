package frameworks

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber"
)

type mockController struct{}

func (m *mockController) PostFile(c *fiber.Ctx) {
	// Implementar lógica de prueba para PostFile si es necesario
}

func (m *mockController) GetData(c *fiber.Ctx) {
	// Implementar lógica de prueba para GetData si es necesario
}

func TestFiberFrameworkConnection(t *testing.T) {
	app := fiber.New()
	controller := &mockController{}
	framework := NewApp(controller, app)

	go func() {
		// Ejecutar el servidor en segundo plano
		framework.Conection()
	}()
	defer framework.app.Shutdown()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:3000")
	if err != nil {
		t.Fatalf("Error haciendo una solicitud GET: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Se esperaba un código de estado 404 , pero se obtuvo %d", resp.StatusCode)
	}
}

func TestFiberFrameworkControllers(t *testing.T) {
	app := fiber.New()
	controller := &mockController{}
	framework := NewApp(controller, app)

	framework.Controllers()

	reqFile := httptest.NewRequest("POST", "/api/v1/file", nil)
	respFile, err := app.Test(reqFile)
	if err != nil {
		t.Fatalf("Error haciendo una solicitud POST a /api/v1/file: %v", err)
	}
	defer respFile.Body.Close()

	if respFile.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un código de estado 200 OK para /api/v1/file, pero se obtuvo %d", respFile.StatusCode)
	}

	reqData := httptest.NewRequest("GET", "/api/v1/data", nil)
	respData, err := app.Test(reqData)
	if err != nil {
		t.Fatalf("Error haciendo una solicitud GET a /api/v1/data: %v", err)
	}
	defer respData.Body.Close()

	if respData.StatusCode != http.StatusOK {
		t.Errorf("Se esperaba un código de estado 200 OK para /api/v1/data, pero se obtuvo %d", respData.StatusCode)
	}
}
