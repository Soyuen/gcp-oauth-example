package gcpoauthexample

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Tags OAuth
// @Summary Start Google OAuth flow
// @Description Redirects the user to Google's OAuth consent page
// @Accept json
// @Produce html
// @Success 302 {string} string "Redirect to Google OAuth consent screen"
// @Failure 500 {object} ErrorCodeResponse "Internal server error"
// @Router /auth/gcp [get]
func GCPOAuth(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	authURL, err := GetOAuth(ctx)
	if err != 0 {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Code: err})
		return
	}

	c.Redirect(http.StatusFound, authURL)
}

// @Tags OAuth
// @Summary Google OAuth callback
// @Description Handle Google OAuth callback
// @Accept json
// @Produce json
// @Param state query string true "OAuth state"
// @Param code query string true "OAuth code"
// @Success 302 {string} string "Redirect to final URL"
// @Failure 400 {object} ErrorCodeResponse "Invalid request"
// @Failure 500 {object} ErrorCodeResponse "Internal server error"
// @Router /auth/gcp/callback [get]
func GCPOAuthCallback(c *gin.Context) {
	state := c.Query("state")
	code := c.Query("code")

	c.JSON(http.StatusOK, gin.H{
		"message": "OAuth Success",
		"code":    code,
		"state":   state,
	})

	// You can use redirect to your frontend app after successful login (e.g., localhost:3000)
	//	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	//  defer cancel()
	// 	redirectURL, err := OAuthCallback(ctx, state, code)
	// if err != 0 {
	// 	c.JSON(http.StatusInternalServerError, ErrorResponse{Code: err})
	// 	return
	// }
	// c.Redirect(http.StatusFound, redirectURL)

}
