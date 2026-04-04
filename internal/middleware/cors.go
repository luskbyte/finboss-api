package middleware

import (
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	origins := getAllowedOrigins()

	config := cors.DefaultConfig()
	config.AllowOrigins = origins
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	return cors.New(config)
}

func getAllowedOrigins() []string {
	if val := os.Getenv("ALLOWED_ORIGINS"); val != "" {
		return strings.Split(val, ",")
	}
	return []string{"http://localhost:5173", "http://localhost:3000"}
}
