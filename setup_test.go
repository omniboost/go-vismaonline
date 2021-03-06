package vismaonline_test

import (
	"context"
	"log"
	"net/url"
	"os"
	"testing"

	vismaonline "github.com/omniboost/go-vismaonline"
	"golang.org/x/oauth2"
)

var (
	client *vismaonline.Client
)

func TestMain(m *testing.M) {
	baseURL := os.Getenv("BASE_URL")
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	refreshToken := os.Getenv("REFRESH_TOKEN")
	tokenURL := os.Getenv("TOKEN_URL")
	debug := os.Getenv("DEBUG")

	oauthConfig := vismaonline.NewOauth2Config()
	oauthConfig.ClientID = clientID
	oauthConfig.ClientSecret = clientSecret

	// set alternative token url
	if tokenURL != "" {
		oauthConfig.Endpoint.TokenURL = tokenURL
	}

	token := &oauth2.Token{
		RefreshToken: refreshToken,
	}

	// get http client with automatic oauth logic
	httpClient := oauthConfig.Client(context.Background(), token)

	client = vismaonline.NewClient(httpClient)
	if debug != "" {
		client.SetDebug(true)
	}

	if baseURL != "" {
		u, err := url.Parse(baseURL)
		if err != nil {
			log.Fatal(err)
		}
		client.SetBaseURL(*u)
	}

	client.SetDisallowUnknownFields(true)
	m.Run()
}
