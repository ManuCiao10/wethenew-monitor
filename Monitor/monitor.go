package monitor

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/data"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

// type ID struct {
// 	New_ID    int    `json:"id"`
// }

type ID struct {
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

func SaveSlice(class data.Info) []int{
	var slice []int

	for _, v := range class.Results {
		for _, v := range v.SellNows {
			slice = append(slice, v.ID)
		}
	}
	return slice
}

func MonitorProducts(class data.Info, client tls_client.HttpClient) {
	url := "https://api-sell.wethenew.com/sell-nows?skip=0&take=50"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header = http.Header{
		"Accept":          {"application/json, text/plain, */*"},
		"accept-language": {"it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5"},
		"user-agent":      {"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36"},
		http.HeaderOrderKey: {
			"Accept",
			"accept-language",
			"user-agent",
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var new_id ID
	if err := json.Unmarshal(body, &new_id); err != nil {
		log.Fatal(err)
	}
	Slice := SaveSlice(class) //TRY TO USE THE NEW_ID TO ADD THE FIRST TIME ALL THE PRODUCTS AND AFTER USEA WHILE LOOP OR A TIMER OUT FOR REQUEST
	for _, v := range new_id.Results {
		for _, v := range v.SellNows {
			//check if v.ID is already Slice
			//if not send a message to discord
			fmt.Println(v.ID)
			
			
		}
	}
}

