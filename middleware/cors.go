package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CorsMiddleware() fiber.Handler {
	allowedOrigins := os.Getenv("CORS_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "https://frontend-manbig-phi.vercel.app,http://localhost:5173"
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowedOrigins, // Frontend tetap + Railway domain untuk Swagger
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowCredentials: true, // Kembali ke true untuk frontend
	})
}
