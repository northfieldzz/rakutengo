package category

import rakuten "github.com/northfieldzz/rakutengo"

type Category struct {
	Id   string `json:"categoryId"`
	Name string `json:"categoryName"`
	Url  string `json:"categoryUrl"`
}

type Data struct {
	Result struct {
		Large  []Category `json:"large"`
		Medium []Category `json:"medium"`
		Small  []Category `json:"small"`
	} `json:"result"`
}

type Result struct {
	Data  *Data
	Error *rakuten.ErrorResponse
}
