package awsteam

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hasura/go-graphql-client"
	"golang.org/x/oauth2"
)

// The Oath2 token.
type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

// A Config provides service configuration for service clients.
type Config struct {
	// The Oath2 client id
	ClientId string

	// The Oath2 client secret
	ClientSecret string

	// The Graph Client the SDK's API clients will use to invoke Graph requests.
	GraphClient *graphql.Client

	// The graph endpoint where aws team is deployed
	GraphEndpoint string

	// The HTTPClient the SDK's API clients will use to invoke Graph requests.
	HTTPClient *http.Client

	// The Oath2 token to be used for Bearer Authentication
	Token *Token

	// The Oath2 endpoint for getting a token
	TokenEndpoint string
}

func (config *Config) Build() {
	// Configure the AWS TEAM client
	// First we need to get a token from the oath endpoint
	authPayload := strings.NewReader(fmt.Sprintf(`grant_type=client_credentials&client_id=%s&client_secret=%s`, config.ClientId, config.ClientSecret))

	authClient := &http.Client{}
	authReq, err := http.NewRequest("POST", config.TokenEndpoint, authPayload)

	if err != nil {
		panic(err)
	}

	authReq.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := authClient.Do(authReq)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	token := &Token{}

	err = json.Unmarshal(body, token)

	if err != nil {
		panic(err)
	}

	// Initiate clients and save token
	config.GraphClient = &graphql.Client{}
	config.HTTPClient = &http.Client{}
	config.Token = token
}

func (config *Config) NewClient(ctx context.Context) *Client {
	// Returns a configured client
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token.AccessToken},
	)

	config.HTTPClient = oauth2.NewClient(ctx, src)
	config.GraphClient = graphql.NewClient(config.GraphEndpoint, config.HTTPClient)

	client := &Client{
		Config:        config,
		GraphClient:   config.GraphClient,
		GraphEndpoint: config.GraphEndpoint,
	}

	return client
}
