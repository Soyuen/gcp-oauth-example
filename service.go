package gcpoauthexample

import (
	"context"
	"fmt"

	"gcp-oauth-example/pkg"

	"golang.org/x/oauth2"
)

func GetOAuth(ctx context.Context) (string, int) {
	oauthConfig := GoogleOAuthConfig()
	if oauthConfig == nil {
		fmt.Println(ErrOAuthConfigMissing.Error())
		return "", ErrCodeOAuthConfigMissing // use error code for better consistency, traceability, and security
	}

	state, err := pkg.GenerateSecurityToken()
	if err != nil {
		fmt.Println(ErrGenerateOAuthState.Error(), err)
		return "", ErrCodeGenerateOAuthState
	}
	// If you want to include custom information (e.g. your web user's ID),
	// you can JSON-encode it into the `state` parameter here.

	authURL := oauthConfig.AuthCodeURL(state,
		oauth2.AccessTypeOffline,
		oauth2.SetAuthURLParam("prompt", "consent"),
	)

	return authURL, 0
}

func OAuthCallback(ctx context.Context, code, state string) (string, int) {
	// Step 1: (Optional) Validate the `state` parameter
	// If you encoded custom data (e.g. CSRF token or user ID), decode and verify it here

	// Step 2: Exchange the authorization code (`code`) for an access token
	// token, err := oauthConfig.Exchange(ctx, code)

	// Step 3: (Optional) Store the access token and refresh token securely
	// You can save them in a database, session, or set them in an HTTP-only cookie

	// Step 4: Redirect the user to the original page (e.g. dashboard or homepage)
	// You can include `state` or other context in the redirect URL if needed

	// NOTE: This demo skips token exchange and storage for simplicity

	redirectURL := fmt.Sprintf("https://localhost:3000/oauth/callback?state=%s&code=%s", state, code)
	return redirectURL, 0
}
