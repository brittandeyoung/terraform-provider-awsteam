package awsteam

import (
	"context"
	"encoding/json"
	"fmt"
)

type DeleteSettingsInput struct {
	Id *string
}

type DeleteSettingsOutput struct {
	Setting *Settings `json:"deleteSettings"`
}

func (client *Client) DeleteSettings(ctx context.Context, in *DeleteSettingsInput) (*DeleteSettingsOutput, error) {
	out := &DeleteSettingsOutput{}
	var id string

	if in.Id != nil {
		id = *in.Id
	} else {
		id = "settings"
	}

	q := fmt.Sprintf(`mutation DeleteSettings {
		deleteSettings(input: { id: "%s" }) {
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
			id
		}
	}	
	`, id)

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
