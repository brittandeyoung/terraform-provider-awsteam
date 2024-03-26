package awsteam

import (
	"context"
	"encoding/json"
)

type GetAccountsInput struct{}

type GetAccountsOutput struct {
	Accounts []*Account `json:"getAccounts"`
}

func (client *Client) GetAccounts(ctx context.Context, in *GetAccountsInput) (*GetAccountsOutput, error) {
	out := &GetAccountsOutput{}

	q := `query GetAccounts {
		getAccounts {
			name
			id
		}
	}`

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
