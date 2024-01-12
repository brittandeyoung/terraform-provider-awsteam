package awsteam

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/smithy-go/ptr"
)

type UpdateSettingsInput struct {
	Approval                  *bool
	Comments                  *bool
	Duration                  *int64
	Expiry                    *int64
	Id                        *string
	SesNotificationsEnabled   *bool
	SnsNotificationsEnabled   *bool
	SlackNotificationsEnabled *bool
	SesSourceEmail            *string
	SesSourceArn              *string
	SlackToken                *string
	TeamAdminGroup            *string
	TeamAuditorGroup          *string
	TicketNo                  *bool
	ModifiedBy                *string
	CreatedAt                 *string
	UpdatedAt                 *string
}

type UpdateSettingsOutput struct {
	Settings *Settings `json:"updateSettings"`
}

func (client *Client) UpdateSettings(ctx context.Context, in *UpdateSettingsInput) (*UpdateSettingsOutput, error) {
	out := &UpdateSettingsOutput{}
	var id string

	if in.Id != nil {
		id = *in.Id
	} else {
		id = "settings"
	}

	q := fmt.Sprintf(`mutation UpdateSettings {
		updateSettings(
			input: {
				id: "%s"
				duration: "%d"
				expiry: "%d"
				comments: %t
				ticketNo: %t
				approval: %t
				modifiedBy: "%s"
				sesNotificationsEnabled: %t
				snsNotificationsEnabled: %t
				slackNotificationsEnabled: %t
				sesSourceEmail: "%s"
				sesSourceArn: "%s"
				slackToken: "%s"
				teamAdminGroup: "%s"
				teamAuditorGroup: "%s"
			}
		) {
			id
			duration
			expiry
			comments
			ticketNo
			approval
			modifiedBy
			sesNotificationsEnabled
			snsNotificationsEnabled
			slackNotificationsEnabled
			sesSourceEmail
			sesSourceArn
			slackToken
			teamAdminGroup
			teamAuditorGroup
			createdAt
			updatedAt
		}
	}
	`, id, ptr.ToInt64(in.Duration),
		ptr.ToInt64(in.Expiry),
		ptr.ToBool(in.Comments),
		ptr.ToBool(in.TicketNo),
		ptr.ToBool(in.Approval),
		ptr.ToString(in.ModifiedBy),
		ptr.ToBool(in.SesNotificationsEnabled),
		ptr.ToBool(in.SnsNotificationsEnabled),
		ptr.ToBool(in.SlackNotificationsEnabled),
		ptr.ToString(in.SesSourceEmail),
		ptr.ToString(in.SesSourceArn),
		ptr.ToString(in.SlackToken),
		ptr.ToString(in.TeamAdminGroup),
		ptr.ToString(in.TeamAuditorGroup),
	)

	raw, err := client.GraphClient.ExecRaw(ctx, q, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(raw, out)

	if err != nil {
		return nil, err
	}

	return out, nil
}
