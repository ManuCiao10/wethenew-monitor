package data

//PRODUCT DATA STRUCTURE//

type Info struct {
	Results []struct {
		ID int `json:"id"`
	} `json:"results"`
}

//PRODUCT DATA STRUCTURE//

type ID struct {
	Results []struct {
		ID          int    `json:"id"`
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

//WEBHOOK MESSAGE STRUCTURE//

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	IconURL string `json:"icon_url"`
}

type Fields struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}
type Thumbnail struct {
	URL string `json:"url"`
}

type Image struct {
	URL string `json:"url"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url"`
}
type Embeds struct {
	Author      Author    `json:"author"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
	Description string    `json:"description"`
	Color       int       `json:"color"`
	Fields      []Fields  `json:"fields"`
	Thumbnail   Thumbnail `json:"thumbnail"`
	Image       Image     `json:"image"`
	Footer      Footer    `json:"footer"`
}
type Top struct {
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Content   string   `json:"content"`
	Embeds    []Embeds `json:"embeds"`
}
