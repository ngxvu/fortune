package service

import (
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/model"
	"io/ioutil"
	"net/http"
)

type FortuneService struct {
	Client *http.Client
}

func NewFortuneService(client *http.Client) FortuneInterface {
	return &FortuneService{Client: client}
}

type FortuneInterface interface {
	ProcessURLsCate(client *http.Client, urls model.Data) (string, error)
}

func (s *FortuneService) ProcessURLsCate(client *http.Client, urls model.Data) (string, error) {
	method := "GET"
	var bodies [][]byte
	for _, urls := range urls.APIs {
		req, err := http.NewRequest(method, urls.URL, nil)
		if err != nil {
			return "", err
		}
		res, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		bodies = append(bodies, body)
	}

	var result string
	for _, body := range bodies {
		result += string(body)
	}
	return result, nil
}
