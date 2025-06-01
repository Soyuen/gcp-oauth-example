package main

import (
	gcpoauthexample "gcp-oauth-example"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_REDIRECT_URL") // e.g.  http://localhost:8000/auth/gcp/callback

	if clientID == "" || clientSecret == "" || redirectURL == "" {
		log.Fatal("Missing Google OAuth configuration")
	}

	// 初始化 Google OAuth config
	gcpoauthexample.InitGoogleOAuthConfig(clientID, clientSecret, redirectURL)
	r := gcpoauthexample.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	if err := r.Run("localhost:" + port); err != nil && err != http.ErrServerClosed {
		log.Fatalf("failed to run server: %v", err)
	}
}
