package model

type DataShopCrawled struct {
	Data struct {
		Total         int `json:"total"`
		OfficialShops []struct {
			Userid           int    `json:"userid"`
			Username         string `json:"username"`
			Shopid           int    `json:"shopid"`
			ShopName         string `json:"shop_name"`
			Logo             string `json:"logo"`
			LogoPc           string `json:"logo_pc"`
			ShopCollectionID int    `json:"shop_collection_id"`
			Ctime            int    `json:"ctime"`
			BrandLabel       int    `json:"brand_label"`
		} `json:"official_shops"`
	} `json:"data"`
}
