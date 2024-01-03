package provider

import (
	"context"
	"fmt"

	"github.com/brittandeyoung/terraform-provider-awsteam/internal/sdk/awsteam"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &SettingsResource{}
var _ resource.ResourceWithImportState = &SettingsResource{}

func NewSettingsResource() resource.Resource {
	return &SettingsResource{}
}

// SettingsResource defines the resource implementation.
type SettingsResource struct {
	client *awsteam.Client
}

// SettingsResourceModel describes the data source data model.
type SettingsResourceModel struct {
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

func (r *SettingsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_settings"
}

func (r *SettingsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Allows configuration of the settings within an AWS TEAM deployment.",

		Attributes: map[string]schema.Attribute{
			"approval": schema.BoolAttribute{
				MarkdownDescription: "If disabled, approval will not be required for all elevated access requests. If enabled, approval requirement is managed in eligibility policy configuration.",
				Optional:            true,
				Default:             booldefault.StaticBool(false),
				Computed:            true,
			},
			"comments": schema.BoolAttribute{
				MarkdownDescription: "Determines if comment field is mandatory for all elevated access requests.",
				Optional:            true,
				Default:             booldefault.StaticBool(false),
				Computed:            true,
			},
			"id": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The settings identifier",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"duration": schema.Int64Attribute{
				MarkdownDescription: "Default maximum request duration in hours.",
				Required:            true,
			},
			"expiry": schema.Int64Attribute{
				MarkdownDescription: "Number of time in hours before an unapproved TEAM request expires.",
				Required:            true,
			},
			"ses_notifications_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable sending notifications via Amazon SES.",
				Optional:            true,
				Default:             booldefault.StaticBool(false),
				Computed:            true,
			},
			"ses_source_arn": schema.StringAttribute{
				MarkdownDescription: "ARN of a verified SES identity in another AWS account. Must be configured to authorize sending mail from the TEAM account.",
				Optional:            true,
				Default:             stringdefault.StaticString(""),
				Computed:            true,
			},
			"ses_source_email": schema.StringAttribute{
				MarkdownDescription: "Email address to send notifications from. Must be verified in SES.",
				Optional:            true,
				Default:             stringdefault.StaticString(""),
				Computed:            true,
			},
			"sns_notifications_enabled": schema.BoolAttribute{
				MarkdownDescription: "Send notifications via Amazon SNS. Once enabled, create a subscription to the SNS topic (TeamNotifications-main) in the TEAM account.",
				Optional:            true,
				Default:             booldefault.StaticBool(false),
				Computed:            true,
			},
			"slack_notifications_enabled": schema.BoolAttribute{
				MarkdownDescription: "Enable to send notifications directly to users in Slack via a Slack bot app.",
				Optional:            true,
				Default:             booldefault.StaticBool(false),
				Computed:            true,
			},
			"slack_token": schema.StringAttribute{
				MarkdownDescription: "Slack OAuth token associated with the installed app.",
				Optional:            true,
				Default:             stringdefault.StaticString(""),
				Computed:            true,
			},
			"team_admin_group": schema.StringAttribute{
				MarkdownDescription: "Group of users responsible for managing TEAM administrative configurations",
				Required:            true,
			},
			"team_auditor_group": schema.StringAttribute{
				MarkdownDescription: "Group of users allowed to audit TEAM elevated access requests",
				Required:            true,
			},
			"ticket_no": schema.BoolAttribute{
				MarkdownDescription: "Determines if ticket number field is mandatory for elevated access requests",
				Optional:            true,
				Default:             booldefault.StaticBool(false),
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

func (r *SettingsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*awsteam.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *awsteam.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *SettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data SettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	in := &awsteam.CreateSettingsInput{
		TeamAdminGroup:            data.TeamAdminGroup.ValueStringPointer(),
		TeamAuditorGroup:          data.TeamAuditorGroup.ValueStringPointer(),
		Approval:                  data.Approval.ValueBoolPointer(),
		Comments:                  data.Comments.ValueBoolPointer(),
		SesNotificationsEnabled:   data.SesNotificationsEnabled.ValueBoolPointer(),
		SnsNotificationsEnabled:   data.SnsNotificationsEnabled.ValueBoolPointer(),
		SlackNotificationsEnabled: data.SlackNotificationsEnabled.ValueBoolPointer(),
		TicketNo:                  data.TicketNo.ValueBoolPointer(),
		Duration:                  data.Duration.ValueInt64Pointer(),
		Expiry:                    data.Expiry.ValueInt64Pointer(),
	}

	if !data.SesSourceArn.IsNull() {
		in.SesSourceArn = data.SesSourceArn.ValueStringPointer()
	}

	if !data.SesSourceEmail.IsNull() {
		in.SesSourceEmail = data.SesSourceEmail.ValueStringPointer()
	}

	if !data.SlackToken.IsNull() {
		in.SlackToken = data.SlackToken.ValueStringPointer()
	}

	out, err := r.client.CreateSettings(ctx, in)

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to create settings, got error: %s", err))
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

	tflog.Trace(ctx, "created settings resource")

	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *SettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data SettingsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	in := &awsteam.GetSettingsInput{}

	out, err := r.client.GetSettings(ctx, in)

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read settings, got error: %s", err))
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

func (r *SettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var config, plan, state SettingsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	updateRequired := false

	in := &awsteam.UpdateSettingsInput{
		TeamAdminGroup:            config.TeamAdminGroup.ValueStringPointer(),
		TeamAuditorGroup:          config.TeamAuditorGroup.ValueStringPointer(),
		Approval:                  config.Approval.ValueBoolPointer(),
		Comments:                  config.Comments.ValueBoolPointer(),
		SesNotificationsEnabled:   config.SesNotificationsEnabled.ValueBoolPointer(),
		SnsNotificationsEnabled:   config.SnsNotificationsEnabled.ValueBoolPointer(),
		SlackNotificationsEnabled: config.SlackNotificationsEnabled.ValueBoolPointer(),
		TicketNo:                  config.TicketNo.ValueBoolPointer(),
		Duration:                  config.Duration.ValueInt64Pointer(),
		Expiry:                    config.Expiry.ValueInt64Pointer(),
	}

	if !plan.SesSourceArn.IsUnknown() && !state.SesSourceArn.Equal(plan.SesSourceArn) {
		updateRequired = true
		in.SesSourceArn = config.SesSourceArn.ValueStringPointer()
	}

	if !plan.SesSourceEmail.IsUnknown() && !state.SesSourceEmail.Equal(plan.SesSourceEmail) {
		updateRequired = true
		in.SesSourceEmail = config.SesSourceEmail.ValueStringPointer()
	}

	if !plan.SlackToken.IsUnknown() && !state.SlackToken.Equal(plan.SlackToken) {
		updateRequired = true
		in.SlackToken = config.SlackToken.ValueStringPointer()
	}

	if updateRequired {

		out, err := r.client.UpdateSettings(ctx, in)

		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read settings, got error: %s", err))
			return
		}

		settings := out.Setting
		state.Id = types.StringPointerValue(settings.Id)
		state.Approval = types.BoolPointerValue(settings.Approval)
		state.Comments = types.BoolPointerValue(settings.Comments)
		state.CreatedAt = types.StringPointerValue(settings.CreatedAt)
		state.Duration = types.Int64PointerValue(settings.Duration)
		state.ModifiedBy = types.StringPointerValue(settings.ModifiedBy)
		state.Expiry = types.Int64PointerValue(settings.Expiry)
		state.SesNotificationsEnabled = types.BoolPointerValue(settings.SesNotificationsEnabled)
		state.SesSourceArn = types.StringPointerValue(settings.SesSourceArn)
		state.SesSourceEmail = types.StringPointerValue(settings.SesSourceEmail)
		state.SlackNotificationsEnabled = types.BoolPointerValue(settings.SlackNotificationsEnabled)
		state.SlackToken = types.StringPointerValue(settings.SlackToken)
		state.SnsNotificationsEnabled = types.BoolPointerValue(settings.SnsNotificationsEnabled)
		state.TeamAdminGroup = types.StringPointerValue(settings.TeamAdminGroup)
		state.TeamAuditorGroup = types.StringPointerValue(settings.TeamAuditorGroup)
		state.TicketNo = types.BoolPointerValue(settings.TicketNo)
		state.UpdatedAt = types.StringPointerValue(settings.UpdatedAt)

		tflog.Trace(ctx, "updated settings resource")

	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SettingsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data SettingsResourceModel

	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	in := &awsteam.DeleteSettingsInput{}

	_, err := r.client.DeleteSettings(ctx, in)

	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read settings, got error: %s", err))
		return
	}
}

func (r *SettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
