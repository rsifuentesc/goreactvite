package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
    app := fiber.New()

    // Configura CORS para permitir solicitudes desde cualquier origen
    app.Use(cors.New(cors.Config{
        AllowOrigins: "*",
    }))

    // Sirve archivos estáticos desde la carpeta ./goreactvite/dist
    app.Static("/", "./goreactvite/dist")

    // Maneja la ruta /users
    app.Get("/users", func(c *fiber.Ctx) error {
        return c.JSON(&fiber.Map{
            "data": "Usuarios desde el backend",
        })
    })

    // Ejecuta la aplicación Fiber en un hilo separado para evitar bloqueos
    go app.Listen(":8080")

    fmt.Println("Server on port 8080")
}
