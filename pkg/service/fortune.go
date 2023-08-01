package service

import (
	"encoding/json"
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
	ProcessURLsParentCate(client *http.Client, urls model.Data) (model.DataCatCrawled, error)
	ProcessURLsShop(client *http.Client, urls model.Data) (model.DataShopCrawled, error)
	ProcessURLsItem(client *http.Client, urls model.Data) (string, error)
	ProcessURLsItem1(client *http.Client, urls model.Data) ([]model.DataItemCrawled, error)
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

func (s *FortuneService) ProcessURLsItem(client *http.Client, urls model.Data) (string, error) {
	method := "GET"
	var rs string
	for _, urls := range urls.APIs {
		req, err := http.NewRequest(method, urls.URL, nil)
		if err != nil {
			return "", err
		}
		req.Header.Add("af-ac-enc-dat", "null")
		req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188")
		req.Header.Add("Cookie", "REC_T_ID=3d3b6607-2518-11ee-810b-2cea7fa63dbf; SPC_EC=eWtHVlB4VDQxQkpwdmcxdgYVajAeQAi2dVmwP3pnLHj4gI+2rbcddXqIpW9OqjieK8VB22aA8WD9fqaT726IZa/N8ubp8DPIBZgarTZpzqhukwiJ1PVJk+tjDZ4jpb0S1tHEz8uRF5RB3Kd0m7SnLdH8Bf9uecqu2wDyCTEHUAI=; SPC_F=VOhuVbKnQstddpDzrzY0ycTU0joRtPFL; SPC_R_T_ID=7NI3u4BqV0cnKF4InjzK63ImxL2D7+Mx+N65/scaxcs8+5/VrJlYm3Cn/GM3HlNdLUtUrqLxKcv534gHTGO+Ol0scjFW7rHCE3VfLyX1nWKdWxhH09ExtcY7X2l7wLd6J/xfG8y38jPpP3y+dlT3ypx/v7pKQAMGl1FPfVmrRvc=; SPC_R_T_IV=SHk2TWRSU2hSM05PZUVVUQ==; SPC_SI=HsHIZAAAAABMUGZlQ3Zzbhj6AAAAAAAAVjl5ZnZBQ0k=; SPC_ST=.Y0lRREk4eTNWZFM4ZWtiYjsRXvw2wghsy7JXn6Q+3DenBcMD5qglHdfEjeaeClhM0YCsyMfbQqFqkWPOF/rGgFdxdysuQJR0Gq/FWNYtvKTiNVN62pri9tcgCr4ORSeML1lVEPJRCImSLvDVoSYyw3n8H7pQRIStv+PEgLD1KFYGFzFW2ip8rEDJ9DD4GpLhEHaiU5i67iybn5Kge37zfA==; SPC_T_ID=7NI3u4BqV0cnKF4InjzK63ImxL2D7+Mx+N65/scaxcs8+5/VrJlYm3Cn/GM3HlNdLUtUrqLxKcv534gHTGO+Ol0scjFW7rHCE3VfLyX1nWKdWxhH09ExtcY7X2l7wLd6J/xfG8y38jPpP3y+dlT3ypx/v7pKQAMGl1FPfVmrRvc=; SPC_T_IV=SHk2TWRSU2hSM05PZUVVUQ==; SPC_U=62219631")
		res, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		rs = string(body)
	}
	return rs, nil
}

func (s *FortuneService) ProcessURLsItem1(client *http.Client, urls model.Data) ([]model.DataItemCrawled, error) {
	method := "GET"
	var bodies []model.DataItemCrawled
	for _, urls := range urls.APIs {
		req, err := http.NewRequest(method, urls.URL, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("af-ac-enc-dat", "null")
		req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.188")
		req.Header.Add("Cookie", "REC_T_ID=3d3b6607-2518-11ee-810b-2cea7fa63dbf; SPC_EC=eWtHVlB4VDQxQkpwdmcxdgYVajAeQAi2dVmwP3pnLHj4gI+2rbcddXqIpW9OqjieK8VB22aA8WD9fqaT726IZa/N8ubp8DPIBZgarTZpzqhukwiJ1PVJk+tjDZ4jpb0S1tHEz8uRF5RB3Kd0m7SnLdH8Bf9uecqu2wDyCTEHUAI=; SPC_F=VOhuVbKnQstddpDzrzY0ycTU0joRtPFL; SPC_R_T_ID=7NI3u4BqV0cnKF4InjzK63ImxL2D7+Mx+N65/scaxcs8+5/VrJlYm3Cn/GM3HlNdLUtUrqLxKcv534gHTGO+Ol0scjFW7rHCE3VfLyX1nWKdWxhH09ExtcY7X2l7wLd6J/xfG8y38jPpP3y+dlT3ypx/v7pKQAMGl1FPfVmrRvc=; SPC_R_T_IV=SHk2TWRSU2hSM05PZUVVUQ==; SPC_SI=HsHIZAAAAABMUGZlQ3Zzbhj6AAAAAAAAVjl5ZnZBQ0k=; SPC_ST=.Y0lRREk4eTNWZFM4ZWtiYjsRXvw2wghsy7JXn6Q+3DenBcMD5qglHdfEjeaeClhM0YCsyMfbQqFqkWPOF/rGgFdxdysuQJR0Gq/FWNYtvKTiNVN62pri9tcgCr4ORSeML1lVEPJRCImSLvDVoSYyw3n8H7pQRIStv+PEgLD1KFYGFzFW2ip8rEDJ9DD4GpLhEHaiU5i67iybn5Kge37zfA==; SPC_T_ID=7NI3u4BqV0cnKF4InjzK63ImxL2D7+Mx+N65/scaxcs8+5/VrJlYm3Cn/GM3HlNdLUtUrqLxKcv534gHTGO+Ol0scjFW7rHCE3VfLyX1nWKdWxhH09ExtcY7X2l7wLd6J/xfG8y38jPpP3y+dlT3ypx/v7pKQAMGl1FPfVmrRvc=; SPC_T_IV=SHk2TWRSU2hSM05PZUVVUQ==; SPC_U=62219631")
		res, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()

		var body model.DataItemCrawled
		err = json.NewDecoder(res.Body).Decode(&body)
		if err != nil {
			return nil, err
		}
		body.Data.ProductPrice.Price.SingleValue /= 100000
		bodies = append(bodies, body)
	}
	return bodies, nil
}
