package awsteam

type Approver struct {
	Id         *string   `json:"id"`
	Name       *string   `json:"name"`
	Type       *string   `json:"type"`
	Approvers  []*string `json:"approvers"`
	GroupIds   []*string `json:"groupIds"`
	TicketNo   *string   `json:"ticketNo"`
	ModifiedBy *string   `json:"modifiedBy"`
	CreatedAt  *string   `json:"createdAt"`
	UpdatedAt  *string   `json:"updatedAt"`
}

type Settings struct {
	Approval                  *bool   `json:"approval"`
	Comments                  *bool   `json:"comments"`
	Duration                  *int64  `json:"duration,string"`
	Expiry                    *int64  `json:"expiry,string"`
	Id                        *string `json:"id"`
	SesNotificationsEnabled   *bool   `json:"sesNotificationsEnabled"`
	SnsNotificationsEnabled   *bool   `json:"snsNotificationsEnabled"`
	SlackNotificationsEnabled *bool   `json:"slackNotificationsEnabled"`
	SesSourceEmail            *string `json:"sesSourceEmail"`
	SesSourceArn              *string `json:"sesSourceArn"`
	SlackToken                *string `json:"slackToken"`
	TeamAdminGroup            *string `json:"teamAdminGroup"`
	TeamAuditorGroup          *string `json:"teamAuditorGroup"`
	TicketNo                  *bool   `json:"ticketNo"`
	ModifiedBy                *string `json:"modifiedBy"`
	CreatedAt                 *string `json:"createdAt"`
	UpdatedAt                 *string `json:"updatedAt"`
}
