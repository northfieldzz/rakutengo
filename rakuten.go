package rakuten

import (
	"net/url"
)

func New() *Rakuten {
	return &Rakuten{
		url: &url.URL{
			Scheme: "https",
			Host:   "app.rakuten.co.jp",
		},
	}
}

type Rakuten struct {
	url           *url.URL
	recipesPath   string
	applicationID string
}
