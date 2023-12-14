package awsteam

import (
	"context"

	"github.com/brittandeyoung/terraform-provider-awsteam/internal/sdk/awsteam/setting"
	"github.com/hasura/go-graphql-client"
)

type Client struct {
	GraphEndpoint string
	GraphClient   *graphql.Client
	Config        *Config
}

func (c *Client) SettingClient(ctx context.Context) *setting.Client {
	return &setting.Client{
		GraphClient: c.GraphClient,
	}
}
