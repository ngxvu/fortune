package model

type DataCatCrawled struct {
	Data struct {
		CategoryList []CategoryList `json:"category_list"`
	} `json:"data"`
}

type CategoryList struct {
	Catid       int           `json:"catid"`
	ParentCatid int           `json:"parent_catid"`
	Name        string        `json:"name"`
	DisplayName string        `json:"display_name"`
	Image       string        `json:"image"`
	Level       int           `json:"level"`
	Children    []interface{} `json:"children"`
}
