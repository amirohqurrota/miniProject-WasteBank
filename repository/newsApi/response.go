package newsApi

import (
	"math/rand"
	"time"
	newsApi "wastebank-ca/bussines/newsApi"
)

// d9601c590a1c4c1bb36cc321d487d0a9

type ArticlesData struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

type ResponseNews struct {
	Status       string         `json:"status"`
	TotalResults int            `json:"totalResults"`
	Articles     []ArticlesData `json:"articles"`
}

func (resp *ResponseNews) toDomain() newsApi.Domain {
	n := rand.Intn(len(resp.Articles))
	return newsApi.Domain{
		Source:      resp.Articles[n].Source.Name,
		Author:      resp.Articles[n].Author,
		Title:       resp.Articles[n].Title,
		Description: resp.Articles[n].Description,
		Content:     resp.Articles[n].Content,
		Url:         resp.Articles[n].URL,
	}
}
