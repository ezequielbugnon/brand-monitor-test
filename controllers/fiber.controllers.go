package controllers

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/ezequiel-bugnon/brandmonitor/dto"
	"github.com/ezequiel-bugnon/brandmonitor/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber"
)

type FiberController struct {
	service services.Service
}

func NewController(service services.Service) *FiberController {
	return &FiberController{
		service,
	}
}

func (f *FiberController) PostFile(c *fiber.Ctx) {
	file, err := c.FormFile("file")
	if err != nil {
		c.Status(http.StatusBadRequest).SendString("Error al leer el archivo CSV")
		return
	}

	uploadedFile, err := file.Open()
	if err != nil {
		c.Status(http.StatusInternalServerError).SendString("Error al abrir el archivo CSV")
		return
	}
	defer uploadedFile.Close()

	csvReader := csv.NewReader(uploadedFile)

	header, err := csvReader.Read()
	if err != nil {
		c.Status(http.StatusBadRequest).SendString("Error al leer el encabezado CSV")
		return
	}

	cleanedHeader := make([]string, len(header))
	for i, columnName := range header {
		cleanedHeader[i] = strings.TrimSpace(columnName)
	}

	expectedHeader := []string{"Indicador1", "Indicador2", "Indicador3"}
	if !sliceEqual(header, expectedHeader) {
		c.Status(http.StatusBadRequest).SendString("El encabezado CSV no tiene el formato esperado")
		return
	}

	lineCounter := 1
	validate := validator.New()

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			c.Status(http.StatusBadRequest).SendString("Error al leer los datos CSV")
			return
		}

		data := dto.FileDto{}
		for i, value := range record {
			intValue, err := strconv.Atoi(value)
			if err != nil {
				c.Status(http.StatusBadRequest).SendString(
					fmt.Sprintf("Error al convertir el valor en la línea %d, columna %d a entero", lineCounter, i+1),
				)
				return
			}

			switch i {
			case 0:
				data.Indicador1 = intValue
			case 1:
				data.Indicador2 = intValue
			case 2:
				data.Indicador3 = intValue
			}
		}

		if err := validate.Struct(data); err != nil {
			c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return
		}

		if err := f.service.PostFile(data); err != nil {
			c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "datos no procesados"})
			return
		}

		log.Printf("Datos procesados: %+v", data)
	}

	c.Status(http.StatusOK).JSON(fiber.Map{"datos": "datos procesados correctamente"})
}

func (f *FiberController) GetData(c *fiber.Ctx) {
	data, err := f.service.GetData()
	if err != nil {
		c.Status(http.StatusNotFound).JSON(fiber.Map{"datos": "no existen datos"})
		return
	}

	c.Status(http.StatusOK).JSON(fiber.Map{"datos": data})
}

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
