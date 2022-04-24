package rakuten

import (
	"encoding/json"
	"fmt"
	"github.com/northfieldzz/rakutengo/rakuten/recipe"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (r *Rakuten) SetQuery(path string, values map[string]string) {
	u := r.url
	u.Path = fmt.Sprintf("services/api/%s", path)
	q := u.Query()
	q.Set("applicationId", os.Getenv("RAKUTEN_APPLICATION_ID"))
	for k, v := range values {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()
}

func (r *Rakuten) FetchCategories() {
	path := "Recipe/CategoryList/20170426"
	r.SetQuery(path, map[string]string{"formatVersion": "2"})
	client := &http.Client{}
	request, err := http.NewRequest(
		"GET",
		r.url.String(),
		strings.NewReader(url.Values{}.Encode()),
	)
	response, err := client.Do(request)
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	if response.StatusCode != 200 {
		//sample, _ := ioutil.ReadAll(response.Body)
		return
	}
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var d recipe.Category
	err = json.Unmarshal(body, &d)
	if err != nil {
		return
	}
}

func (r *Rakuten) FetchRecipeRanking(categoryId string) {
	path := "Recipe/CategoryRanking/20170426"
	u := r.url
	u.Path = fmt.Sprintf("services/api/%s", path)
	q := u.Query()
	q.Set("applicationId", os.Getenv("RAKUTEN_APPLICATION_ID"))
	q.Set("formatVersion", "2")
	u.RawQuery = q.Encode()
	client := &http.Client{}
	request, err := http.NewRequest(
		"GET",
		u.String(),
		strings.NewReader(url.Values{}.Encode()),
	)
	response, err := client.Do(request)
	if response.StatusCode != 200 {
		return
	}
	if err != nil {
		return
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	var d recipe.RankingResponse
	err = json.Unmarshal(body, &d)
	if err != nil {
		return
	}
}
