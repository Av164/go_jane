package models

//import()

type CampaignNew struct{
	//CampaignId string `json:"campaign_id"`
	StartCampaign int `json:"start_campaign"`
	EndCampaign int `json:"end_campaign"`
	MaxImpressions int `json:"max_impressions"`
	Cpm int `json:"cpm"`
	Keywords []string `json:"keywords"`
}