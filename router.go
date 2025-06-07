package gcpoauthexample

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode) // Set Gin mode: debug / release / test
	r := gin.Default()
	registerGcpOAuthRoutes(r)
	// You can add Swagger UI for API documentation here.

	return r
}

func registerGcpOAuthRoutes(r *gin.Engine) {
	r.GET("/auth/gcp", GCPOAuth)
	r.GET("/auth/gcp/callback", GCPOAuthCallback)
}
