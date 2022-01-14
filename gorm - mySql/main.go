package main

import (
  "github.com/gofiber/fiber/v2"
  "gorm.io/gorm"
  "gorm.io/driver/mysql"
)

func main() {
  _, err := gorm.Open(mysql.Open("root:root@/teste"), &gorm.Config{})

  if err != nil {
    panic("could not connect to the database")
  }

  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) error {
    return c.SendString("Running")
  })

  app.Listen(":8080")
}