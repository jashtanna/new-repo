package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// RegisterCorsMiddleware sets up CORS  for the Gin router.
// This is important when your frontend and backend are hosted on different domains or ports.
// It ensures the browser allows requests between them.
func RegisterCorsMiddleware(r *gin.Engine) {
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// Middleware function to wrap and apply CORS settings on each incoming request
	r.Use(func(c *gin.Context) {
		httpResponseWriter := c.Writer
		httpRequest := c.Request

		// Apply the CORS rules before continuing to the next middleware/handler
		corsHandler.ServeHTTP(httpResponseWriter, httpRequest, func(w http.ResponseWriter, r *http.Request) {
			c.Next()
		})

		c.Next()
	})
}
