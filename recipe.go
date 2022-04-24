package rakuten

import (
	"encoding/json"
	"fmt"
	"github.com/northfieldzz/rakutengo/rakuten/recipe/category"
	"github.com/northfieldzz/rakutengo/rakuten/recipe/ranking"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const (
	categoriesPath = "Recipe/CategoryList/20170426"
	rankingPath    = "Recipe/CategoryRanking/20170426"
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

func (r *Rakuten) FetchCategories() (result *category.Result, err error) {
	r.SetQuery(categoriesPath, map[string]string{"formatVersion": "2"})
	client := &http.Client{}
	var request *http.Request
	if request, err = http.NewRequest(
		"GET",
		r.url.String(),
		strings.NewReader(url.Values{}.Encode()),
	); err != nil {
		return nil, err
	}
	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return nil, err
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	var body []byte
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}
	if response.StatusCode == 200 {
		//sample, _ := ioutil.ReadAll(response.Body)
		var d category.Data
		if err := json.Unmarshal(body, &d); err != nil {
			return nil, err
		}
		return &category.Result{
			Data:  &d,
			Error: nil,
		}, nil
	} else {
		var d ErrorResponse
		if err := json.Unmarshal(body, &d); err != nil {
			return nil, err
		}
		return &category.Result{
			Data:  nil,
			Error: &d,
		}, nil
	}
}

func (r *Rakuten) FetchRecipeRanking(categoryId string) (result *ranking.Result, err error) {
	r.SetQuery(rankingPath, map[string]string{
		"formatVersion": "2",
		"categoryId":    categoryId,
	})
	client := &http.Client{}
	var request *http.Request
	if request, err = http.NewRequest(
		"GET",
		r.url.String(),
		strings.NewReader(url.Values{}.Encode()),
	); err != nil {
		return nil, err
	}
	var response *http.Response
	if response, err = client.Do(request); err != nil {
		return nil, err
	}
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(response.Body)
	var body []byte
	if body, err = ioutil.ReadAll(response.Body); err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		var d ranking.Data
		if err := json.Unmarshal(body, &d); err != nil {
			return nil, err
		}
		return &ranking.Result{
			Data:  &d,
			Error: nil,
		}, nil
	} else {
		var d ErrorResponse
		if err := json.Unmarshal(body, &d); err != nil {
			return nil, err
		}
		return &ranking.Result{
			Data:  nil,
			Error: &d,
		}, nil
	}
}
