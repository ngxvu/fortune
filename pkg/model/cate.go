package model

type DataCatCrawled struct {
	Data struct {
		CategoryList []CategoryList `json:"category_list"`
	} `json:"data"`
}

type CategoryList struct {
	CatID           int    `json:"catid"`
	ParentCatID     int    `json:"parent_catid"`
	Name            string `json:"name"`
	DisplayName     string `json:"display_name"`
	Image           string `json:"image"`
	UnselectedImage string `json:"unselected_image"`
	SelectedImage   string `json:"selected_image"`
	Level           int    `json:"level"`
	Children        []struct {
		CatID              int         `json:"catid"`
		ParentCatID        int         `json:"parent_catid"`
		Name               string      `json:"name"`
		DisplayName        string      `json:"display_name"`
		Image              string      `json:"image"`
		UnselectedImage    interface{} `json:"unselected_image"`
		SelectedImage      interface{} `json:"selected_image"`
		Level              int         `json:"level"`
		Children           interface{} `json:"children"`
		BlockBuyerPlatform interface{} `json:"block_buyer_platform"`
	} `json:"children"`
	BlockBuyerPlatform interface{} `json:"block_buyer_platform"`
}
