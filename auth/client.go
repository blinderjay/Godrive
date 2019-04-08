package auth

import (
	"net/http"

	"github.com/blinderjay/Godrive/env"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	home, err := env.Home()
	tokFile := home + "/.config/token.json"
	// tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}
