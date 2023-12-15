package provider

import (
	"context"
	"fmt"

	"github.com/brittandeyoung/terraform-provider-awsteam/internal/sdk/awsteam"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &SettingsDataSource{}

func NewSettingsDataSource() datasource.DataSource {
	return &SettingsDataSource{}
}

// SettingsDataSource defines the data source implementation.
type SettingsDataSource struct {
	client *awsteam.Client
}

// SettingsDataSourceModel describes the data source data model.
type SettingsDataSourceModel struct {
	Approval                  types.Bool   `tfsdk:"approval"`
	Comments                  types.Bool   `tfsdk:"comments"`
	Id                        types.String `tfsdk:"id"`
	Duration                  types.Int64  `tfsdk:"duration"`
	Expiry                    types.Int64  `tfsdk:"expiry"`
	SesNotificationsEnabled   types.Bool   `tfsdk:"ses_notifications_enabled"`
	SnsNotificationsEnabled   types.Bool   `tfsdk:"sns_notifications_enabled"`
	SlackNotificationsEnabled types.Bool   `tfsdk:"slack_notifications_enabled"`
	SesSourceEmail            types.String `tfsdk:"ses_source_email"`
	SesSourceArn              types.String `tfsdk:"ses_source_arn"`
	SlackToken                types.String `tfsdk:"slack_token"`
	TeamAdminGroup            types.String `tfsdk:"team_admin_group"`
	TeamAuditorGroup          types.String `tfsdk:"team_auditor_group"`
	TicketNo                  types.Bool   `tfsdk:"ticket_no"`
	ModifiedBy                types.String `tfsdk:"modified_by"`
	CreatedAt                 types.String `tfsdk:"created_at"`
	UpdatedAt                 types.String `tfsdk:"updated_at"`
}

func (d *SettingsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings"
}

func (d *SettingsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "Provides a data source for AWS TEAM Settings",

		Attributes: map[string]schema.Attribute{
			"approval": schema.BoolAttribute{
				MarkdownDescription: "If disabled, approval will not be required for all elevated access requests. If enabled, approval requirement is managed in eligibility policy configuration.",
				Computed:            true,
			},
			"comments": schema.BoolAttribute{
				MarkdownDescription: "Determines if comment field is mandatory for all elevated access requests.",
				Computed:            true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Settings identifier.",
				Computed:            true,
			},
			"duration": schema.Int64Attribute{
				MarkdownDescription: "Default maximum request duration in hours.",
				Computed:            true,
			},
			"expiry": schema.Int64Attribute{
				MarkdownDescription: "Number of time in hours before an unapproved TEAM request expires.",
				Computed:            true,
			},
			"ses_notifications_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable sending notifications via Amazon SES.",
				Computed:            true,
			},
			"ses_source_arn": schema.StringAttribute{
				MarkdownDescription: "ARN of a verified SES identity in another AWS account. Must be configured to authorize sending mail from the TEAM account.",
				Computed:            true,
			},
			"ses_source_email": schema.StringAttribute{
				MarkdownDescription: "Email address to send notifications from. Must be verified in SES.",
				Computed:            true,
			},
			"sns_notifications_enabled": schema.BoolAttribute{
				MarkdownDescription: "Send notifications via Amazon SNS. Once enabled, create a subscription to the SNS topic (TeamNotifications-main) in the TEAM account.",
				Computed:            true,
			},
			"slack_notifications_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable to send notifications directly to users in Slack via a Slack bot app.",
				Computed:            true,
			},
			"slack_token": schema.StringAttribute{
				MarkdownDescription: "Slack OAuth token associated with the installed app.",
				Computed:            true,
			},
			"team_admin_group": schema.StringAttribute{
				MarkdownDescription: "Group of users responsible for managing TEAM administrative configurations",
				Computed:            true,
			},
			"team_auditor_group": schema.StringAttribute{
				MarkdownDescription: "Group of users allowed to audit TEAM elevated access requests",
				Computed:            true,
			},
			"ticket_no": schema.BoolAttribute{
				MarkdownDescription: "Determines if ticket number field is mandatory for elevated access requests",
				Computed:            true,
			},
			"modified_by": schema.StringAttribute{
				MarkdownDescription: "The user to last modify the settings",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				MarkdownDescription: "The date and time that the setting was created",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				MarkdownDescription: "The date and time of the last time the settings were updated",
				Computed:            true,
			},
		},
	}
}

func (d *SettingsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*awsteam.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *awsteam.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *SettingsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data SettingsDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	in := &awsteam.GetSettingsInput{}

	out, err := d.client.GetSettings(ctx, in)

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read settings, got error: %s", err))
		return
	}

	if out.Setting == nil {
		resp.Diagnostics.AddError("Client Error", "Settings does not exist")
		return
	}

	settings := out.Setting
	data.Id = types.StringPointerValue(settings.Id)
	data.Approval = types.BoolPointerValue(settings.Approval)
	data.Comments = types.BoolPointerValue(settings.Comments)
	data.CreatedAt = types.StringPointerValue(settings.CreatedAt)
	data.Duration = types.Int64PointerValue(settings.Duration)
	data.ModifiedBy = types.StringPointerValue(settings.ModifiedBy)
	data.Expiry = types.Int64PointerValue(settings.Expiry)
	data.SesNotificationsEnabled = types.BoolPointerValue(settings.SesNotificationsEnabled)
	data.SesSourceArn = types.StringPointerValue(settings.SesSourceArn)
	data.SesSourceEmail = types.StringPointerValue(settings.SesSourceEmail)
	data.SlackNotificationsEnabled = types.BoolPointerValue(settings.SlackNotificationsEnabled)
	data.SlackToken = types.StringPointerValue(settings.SlackToken)
	data.SnsNotificationsEnabled = types.BoolPointerValue(settings.SnsNotificationsEnabled)
	data.TeamAdminGroup = types.StringPointerValue(settings.TeamAdminGroup)
	data.TeamAuditorGroup = types.StringPointerValue(settings.TeamAuditorGroup)
	data.TicketNo = types.BoolPointerValue(settings.TicketNo)
	data.UpdatedAt = types.StringPointerValue(settings.UpdatedAt)

	tflog.Trace(ctx, "read settings resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
