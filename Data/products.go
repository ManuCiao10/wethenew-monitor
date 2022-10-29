package data

type Info struct {
	Results []struct {
		Name        string `json:"name"`
		Image       string `json:"image"`
		ProductType string `json:"productType"`
		SellNows    []struct {
			ID    int    `json:"id"`
			Size  string `json:"size"`
			Price int    `json:"price"`
		} `json:"sellNows"`
	} `json:"results"`
}