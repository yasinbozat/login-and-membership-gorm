package main

import (
	"github.com/gofiber/fiber/v2"
)

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
  )
  

  dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
