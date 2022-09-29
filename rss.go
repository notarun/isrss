package main

import "encoding/xml"

// https://validator.w3.org/feed/docs/rss2.html
type RSS struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel RSSChannel `xml:"channel"`
}

type RSSChannel struct {
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	LastBuildDate string    `xml:"lastBuildDate"`
	Items         []RSSItem `xml:"item"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PublishDate string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Description CDATA  `xml:"description"`
}

type CDATA struct {
	Value string `xml:",cdata"`
}

func NewRSS(response *InshortsNewsResponse, title string) *RSS {
	rss := RSS{
		Version: "2.0",
		Channel: RSSChannel{
			Title:         title,
			Link:          "https://www.inshorts.com/",
			Description:   "Inshorts RSS Feed",
			LastBuildDate: response.GetLastNewsDate(),
		},
	}

	for _, item := range response.Data.NewsList {
		rssItem := RSSItem{
			Title:       item.NewsObject.Title,
			Link:        item.NewsObject.URL,
			Description: CDATA{item.NewsObject.GetMarkupContent()},
			PublishDate: item.NewsObject.GetCreatedAt(),
			GUID:        item.NewsObject.URL,
		}
		rss.Channel.Items = append(rss.Channel.Items, rssItem)
	}

	return &rss
}
