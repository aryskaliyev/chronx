package pkg

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/fatih/color"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var tokenStash string

func GetClient() (*http.Client, error) {
	getOS()

	clientID := os.Getenv("CHRONXID")
	clientSecret := os.Getenv("CHRONXSECRET")
	redirectURL := "urn:ietf:wg:oauth:2.0:oob"

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/calendar",
			"openid",
		},
		Endpoint: google.Endpoint,
	}

	token, err := getToken(config)
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), token)

	return client, nil
}

func getToken(config *oauth2.Config) (*oauth2.Token, error) {
	stashedToken, err := getStashedToken()
	if err != nil {
		token, err := getNewToken(config)
		if err != nil {
			return token, err
		}
		stashToken(token)
		return token, nil
	}
	return stashedToken, nil
}

func getStashedToken() (*oauth2.Token, error) {
	tokenB, err := os.ReadFile(tokenStash)
	if err != nil {
		return &oauth2.Token{}, err
	}

	if len(tokenB) == 0 {
		return &oauth2.Token{}, fmt.Errorf("unable to access token")
	}
	var token oauth2.Token
	_ = json.Unmarshal(tokenB, &token)
	return &token, nil
}

func stashToken(token *oauth2.Token) {}

func getNewToken(config *oauth2.Config) (*oauth2.Token, error) {}

func getOS() {}
