package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"text/template"
	"time"
)

type NewsCategories string

const (
	ALL_NEWS      NewsCategories = "all_news"
	TOP_NEWS      NewsCategories = "top_stories"
	TRENDING_NEWS NewsCategories = "trending"
)

const baseUrl = "https://inshorts.com/api/en/news"

var inshortsClient = http.Client{
	Timeout: 10 * time.Second,
}

type InshortsNewsResponse struct {
	Data InshortsNewsData
}

type InshortsNewsData struct {
	MinNewsId string         `json:"min_news_id"`
	NewsList  []InshortsNews `json:"news_list"`
}

type InshortsNews struct {
	Type       string
	HashId     string             `json:"hash_id"`
	NewsObject InshortsNewsObject `json:"news_obj"`
}

type InshortsNewsObject struct {
	HashId    string `json:"hash_id"`
	Title     string
	Content   string
	Image     string `json:"image_url"`
	Author    string `json:"author_name"`
	URL       string `json:"source_url"`
	CreatedAt int64  `json:"created_at"`
}

func (r InshortsNewsResponse) GetLastNewsDate() string {
	var date int64
	if len(r.Data.NewsList) > 0 {
		date = r.Data.NewsList[0].NewsObject.CreatedAt / 1000
	} else {
		date = time.Now().Unix()
	}
	return time.Unix(date, 0).Format(time.RFC822)
}

func (o InshortsNewsObject) GetCreatedAt() string {
	return time.Unix(o.CreatedAt/1000, 0).Format(time.RFC822)
}

func (o InshortsNewsObject) GetMarkupContent() string {
	t, _ := template.New("content").Parse(`
	<img src="{{.Image}}" alt="{{.Title}}" style="max-width: 100%; height: auto;" />
	<p>{{.Content}}</p>
	`)

	var b bytes.Buffer
	t.Execute(&b, o)
	return b.String()
}

func GetResults(category NewsCategories) (*InshortsNewsResponse, error) {
	queryParams := "?include_card_data=true"

	switch category {
	case ALL_NEWS, TRENDING_NEWS, TOP_NEWS:
		queryParams += "&category=" + string(category)
	default:
		queryParams += "&category=" + string(ALL_NEWS)
	}

	resp, err := inshortsClient.Get(baseUrl + queryParams)
	if err != nil {
		return nil, errors.New("Error getting result from inshorts")
	}
	defer resp.Body.Close()

	var parsed InshortsNewsResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&parsed)

	return &parsed, nil
}
