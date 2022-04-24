package recipe

type CategoryResponse struct {
	Result struct {
		Large  []Category `json:"large"`
		Medium []Category `json:"medium"`
		Small  []Category `json:"small"`
	} `json:"result"`
}

type RankingResponse struct {
	Result []Recipe `json:"result"`
}
