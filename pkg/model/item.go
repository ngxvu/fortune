package model

type DataItemCrawled struct {
	Data struct {
		Item struct {
			ItemID     int64  `json:"item_id"`
			ShopID     int    `json:"shop_id"`
			Title      string `json:"title"`
			Image      string `json:"image"`
			CatID      int    `json:"cat_id"`
			ItemRating struct {
				RatingStar float64 `json:"rating_star"`
			} `json:"item_rating"`
			FeCategories []struct {
				Catid       int    `json:"catid"`
				DisplayName string `json:"display_name"`
			} `json:"fe_categories"`
		} `json:"item"`
		ProductPrice struct {
			Price struct {
				SingleValue int64 `json:"single_value"`
			} `json:"price"`
		} `json:"product_price"`
	} `json:"data"`
}
