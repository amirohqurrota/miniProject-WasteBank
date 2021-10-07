package newsApi

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	newsapi "wastebank-ca/bussines/newsApi"
)

const APIKEY = "d9601c590a1c4c1bb36cc321d487d0a9"

var tag = []string{
	"waste-sorting",
	"recycling-waste",
	"sorting-waste",
	"organic-waste",
	"waste-management",
	"global-warming",
	"climate-change",
	"sustainability-waste",
	"go-green",
	"green-sustainability",
	"recycling-green",
}

type NewsApi struct {
	httpClient http.Client
}

func NewNewsApi() newsapi.Repository {
	return &NewsApi{
		httpClient: http.Client{},
	}
}

func (newsApi *NewsApi) GetNews() (newsapi.Domain, error) {
	keyword := tag[rand.Intn(len(tag))]
	// fmt.Println(keyword)
	response, err := http.Get(fmt.Sprintf("https://newsapi.org/v2/everything?q=%s&apiKey=%s", keyword, APIKEY))
	if err != nil {
		return newsapi.Domain{}, err
	}
	result := ResponseNews{}
	err = json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		return newsapi.Domain{}, err
	}

	return result.toDomain(), nil

}
