package data

type Info struct {
	Results []struct {
		SellNows    []struct {
			ID    int    `json:"id"`
		} `json:"sellNows"`
	} `json:"results"`
}