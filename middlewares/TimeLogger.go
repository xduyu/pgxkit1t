package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func TimeLogger(c fiber.Ctx) error {
	start := time.Now()
	log.Printf("Time to start: %v", start)

	err := c.Next()

	stop := time.Since(start)
	log.Printf("Needed time: %v", stop)

	return err
}
