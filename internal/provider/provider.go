// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/brittandeyoung/terraform-provider-awsteam/internal/envvar"
	"github.com/brittandeyoung/terraform-provider-awsteam/internal/sdk/awsteam"
	"github.com/brittandeyoung/terraform-provider-awsteam/internal/service/setting"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

type AWSTEAMClient struct {
	Client        *awsteam.Client
	Config        *awsteam.Config
	Token         *awsteam.Token
	GraphEndpoint string
}

// Ensure AWSTEAMProvider satisfies various provider interfaces.
var _ provider.Provider = &AWSTEAMProvider{}

// AWSTEAMProvider defines the provider implementation.
type AWSTEAMProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// AWSTEAMProviderModel describes the provider data model.
type AWSTEAMProviderModel struct {
	ClientId      types.String `tfsdk:"client_id"`
	ClientSecret  types.String `tfsdk:"client_secret"`
	GraphEndpoint types.String `tfsdk:"graph_endpoint"`
	TokenEndpoint types.String `tfsdk:"token_endpoint"`
}

func (p *AWSTEAMProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "awsteam"
	resp.Version = p.version
}

func (p *AWSTEAMProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"client_id": schema.StringAttribute{
				MarkdownDescription: "The client id for authenticating to the oauth2 token endpoint.",
				Optional:            true,
			},
			"client_secret": schema.StringAttribute{
				MarkdownDescription: "The client secret for authenticating to the oauth2 token endpoint.",
				Optional:            true,
			},
			"graph_endpoint": schema.StringAttribute{
				MarkdownDescription: "The graph endpoint for the AWS TEAM deployment.",
				Optional:            true,
			},
			"token_endpoint": schema.StringAttribute{
				MarkdownDescription: "The token endpoint for the oath2 authenticator for AWS TEAMS.",
				Optional:            true,
			},
		},
	}
}

func (p *AWSTEAMProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data AWSTEAMProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	clientId := fieldOrEnvVar(data.ClientId, "client_id", envvar.AWSTEAMClientId, resp)
	clientSecret := fieldOrEnvVar(data.ClientId, "client_secret", envvar.AWSTEAMClientSecret, resp)
	graphEndpoint := fieldOrEnvVar(data.ClientId, "graph_endpoint", envvar.AWSTEAMGraphEndpoint, resp)
	TokenEndpoint := fieldOrEnvVar(data.ClientId, "token_endpoint", envvar.AWSTEAMTokenEndpoint, resp)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configure the AWS TEAM client
	// First we need to get a token from the oath endpoint

	config := &awsteam.Config{
		ClientId:      clientId,
		ClientSecret:  clientSecret,
		GraphEndpoint: graphEndpoint,
		TokenEndpoint: TokenEndpoint,
	}

	config.Build()

	// Now we can create the graph client with the authentication token
	// meta := &awsteam.Client{
	// 	Client:        config.NewClient(ctx),
	// 	Config:        config,
	// 	GraphEndpoint: config.GraphEndpoint,
	// 	Token:         config.Token,
	// }
	meta := config.NewClient(ctx)

	resp.DataSourceData = meta
	resp.ResourceData = meta
}

func (p *AWSTEAMProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *AWSTEAMProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		setting.NewSettingSettingsDataSource,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AWSTEAMProvider{
			version: version,
		}
	}
}

func fieldOrEnvVar(field basetypes.StringValue, fieldName string, envvarName string, resp *provider.ConfigureResponse) string {
	var value string
	if field.IsNull() {
		value = os.Getenv(envvarName)
		if value == "" {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Providing a value for %s is required. This can also be handled by setting the %s environment variable.", fieldName, envvarName))
		}
	} else {
		value = field.String()
	}
	return value
}
