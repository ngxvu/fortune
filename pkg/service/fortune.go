package service

import (
	"encoding/json"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/model"
	"net/http"
)

type FortuneService struct {
	Client *http.Client
}

func NewFortuneService(client *http.Client) FortuneInterface {
	return &FortuneService{Client: client}
}

type FortuneInterface interface {
	ProcessURLsParentCate(client *http.Client, urls model.Data) (model.DataCatCrawled, error)
	ProcessURLsShop(client *http.Client, urls model.Data) (model.DataShopCrawled, error)
}

func (s *FortuneService) ProcessURLsParentCate(client *http.Client, urls model.Data) (model.DataCatCrawled, error) {
	method := "GET"
	var bodies []model.DataCatCrawled
	for _, urls := range urls.APIs {
		req, err := http.NewRequest(method, urls.URL, nil)
		if err != nil {
			return model.DataCatCrawled{}, err
		}
		res, err := client.Do(req)
		if err != nil {
			return model.DataCatCrawled{}, err
		}
		defer res.Body.Close()

		var body model.DataCatCrawled
		err = json.NewDecoder(res.Body).Decode(&body)
		if err != nil {
			return model.DataCatCrawled{}, err
		}
		bodies = append(bodies, body)
	}
	rs := model.DataCatCrawled{}
	for _, v := range bodies {
		for _, u := range v.Data.CategoryList {
			rs.Data.CategoryList = append(rs.Data.CategoryList, u)
		}
	}
	return rs, nil
}

func (s *FortuneService) ProcessURLsShop(client *http.Client, urls model.Data) (model.DataShopCrawled, error) {
	method := "GET"
	var bodies []model.DataShopCrawled
	for _, urls := range urls.APIs {
		req, err := http.NewRequest(method, urls.URL, nil)
		if err != nil {
			return model.DataShopCrawled{}, err
		}
		res, err := client.Do(req)
		if err != nil {
			return model.DataShopCrawled{}, err
		}
		defer res.Body.Close()

		var body model.DataShopCrawled
		err = json.NewDecoder(res.Body).Decode(&body)
		if err != nil {
			return model.DataShopCrawled{}, err
		}
		bodies = append(bodies, body)
	}
	rs := model.DataShopCrawled{}

	for i := 0; i < len(bodies); i++ {
		if len(bodies) > 0 {
			rs.Data.Total += bodies[i].Data.Total
		}
	}

	for _, v := range bodies {
		for _, u := range v.Data.OfficialShops {
			rs.Data.OfficialShops = append(rs.Data.OfficialShops, u)
		}
	}
	return rs, nil
}
