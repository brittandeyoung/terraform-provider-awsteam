package awsteam

type Approver struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Approvers  []string `json:"approvers"`
	GroupIds   []string `json:"groupIds"`
	TicketNo   string   `json:"ticketNo"`
	ModifiedBy string   `json:"modifiedBy"`
	CreatedAt  string   `json:"createdAt"`
	UpdatedAt  string   `json:"updatedAt"`
}
