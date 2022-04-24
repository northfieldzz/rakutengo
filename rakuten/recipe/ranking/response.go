package ranking

import rakuten "github.com/northfieldzz/rakutengo"

type Recipe struct {
	Id             string   `json:"recipeId"`
	Title          string   `json:"recipeTitle"`
	Description    string   `json:"recipeDescription"`
	Url            string   `json:"recipeUrl"`
	Cost           string   `json:"recipeCost"`
	Indication     string   `json:"recipeIndication"`
	Material       []string `json:"recipeMaterial"`
	PublishDay     string   `json:"recipePublishDay"`
	NickName       string   `json:"nickname"`
	Rank           string   `json:"rank"`
	PickUp         int64    `json:"pickup"`
	Shop           int64    `json:"shop"`
	FoodImageUrl   string   `json:"foodImageUrl"`
	SmallImageUrl  string   `json:"smallImageUrl"`
	MediumImageUrl string   `json:"mediumImageUrl"`
}

type Data struct {
	Result []Recipe `json:"result"`
}

type Result struct {
	Data  *Data
	Error *rakuten.ErrorResponse
}
