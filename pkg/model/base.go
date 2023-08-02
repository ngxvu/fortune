package model

type API struct {
	URL string `json:"url"`
}

type Data struct {
	APIs []API `json:"data"`
}
