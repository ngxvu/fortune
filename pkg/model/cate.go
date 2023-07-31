package model

type API struct {
	URL string `json:"url"`
}

type Data struct {
	APIs []API `json:"data"`
}

type DataCrawled struct {
	Data struct {
		CategoryList []CategoryList `json:"category_list"`
	} `json:"data"`
}

type CategoryList struct {
	CatID              int           `json:"catid"`
	ParentCatID        int           `json:"parent_catid"`
	Name               string        `json:"name"`
	DisplayName        string        `json:"display_name"`
	Image              string        `json:"image"`
	UnselectedImage    string        `json:"unselected_image"`
	SelectedImage      string        `json:"selected_image"`
	Level              int           `json:"level"`
	Children           []interface{} `json:"children"`
	BlockBuyerPlatform interface{}   `json:"block_buyer_platform"`
}
