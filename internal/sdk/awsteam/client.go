package awsteam

import (
	"github.com/hasura/go-graphql-client"
)

type Client struct {
	GraphEndpoint string
	GraphClient   *graphql.Client
	Config        *Config
}
