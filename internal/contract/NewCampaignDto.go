package contract

//uma struct do typo new campanha

type NewCampaign struct {
	ID      string
	Name    string
	Content string
	Emails  []string
}

type NewCampaignResponse struct {
	ID      string
	Name    string
	Content string
	Status  string
}
