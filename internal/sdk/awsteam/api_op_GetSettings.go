package awsteam

import (
	"context"
	"encoding/json"
	"fmt"
)

type GetSettingsInput struct {
	Id *string
}

type GetSettingsOutput struct {
	Setting *Settings `json:"getSettings"`
}

func (client *Client) GetSettings(ctx context.Context, in *GetSettingsInput) (*GetSettingsOutput, error) {
	out := &GetSettingsOutput{}
	var id string

	if in.Id != nil {
		id = *in.Id
	} else {
		id = "settings"
	}

	q := fmt.Sprintf(`query GetSettings {
		getSettings(id: "%s") {
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
