package models

type Addecision struct{
	Keywords[] string `json:"keywords"`
}

type AddecisionRes struct{
	CampaignId int `json:"campaign_id"`
	URL string `json:"impression_url"`

}