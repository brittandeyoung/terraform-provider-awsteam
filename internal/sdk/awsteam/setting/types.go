package setting

type GroupSetting struct {
	ID               string `json:"id"`
	TeamAdminGroup   string `json:"teamAdminGroup"`
	TeamAuditorGroup string `json:"teamAuditorGroup"`
}

type Settings struct {
	Approval                  bool   `json:"approval"`
	Comments                  bool   `json:"comments"`
	Duration                  string `json:"duration"`
	Expiry                    string `json:"expiry"`
	ID                        string `json:"id"`
	SesNotificationsEnabled   bool   `json:"sesNotificationsEnabled"`
	SnsNotificationsEnabled   bool   `json:"snsNotificationsEnabled"`
	SlackNotificationsEnabled bool   `json:"slackNotificationsEnabled"`
	SesSourceEmail            string `json:"sesSourceEmail"`
	SesSourceArn              string `json:"sesSourceArn"`
	SlackToken                string `json:"slackToken"`
	TeamAdminGroup            string `json:"teamAdminGroup"`
	TeamAuditorGroup          string `json:"teamAuditorGroup"`
	TicketNo                  bool   `json:"ticketNo"`
	ModifiedBy                string `json:"modifiedBy"`
	CreatedAt                 string `json:"createdAt"`
	UpdatedAt                 string `json:"updatedAt"`
}
