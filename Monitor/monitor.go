package monitor

import (
	"encoding/json"
	"io"
	"log"
	"main/data"
	"main/discord"
	// "fmt"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

func SaveSlice(class data.Info) []int {
	var slice []int

	for _, v := range class.Results {
		slice = append(slice, v.ID)
	}
	return slice
}

func SaveSliceTest(class data.Info) []int {
	var slice []int

	for _, v := range class.Results {
		if v.ID != 275 {
			slice = append(slice, v.ID)
		}
	}
	return slice
}

func Contains(s []int, id int) bool {
	for _, v := range s {
		if v == id {
			return true
		}
	}
	return false
}

func MonitorProducts(class data.Info, client tls_client.HttpClient) {
	Slice := SaveSliceTest(class) //TRY TO USE THE NEW_ID TO ADD THE FIRST TIME ALL THE PRODUCTS AND AFTER USEA WHILE LOOP OR A TIMER OUT FOR REQUEST
	url := "https://api-sell.wethenew.com/sell-nows?skip=0&take=50"
	for {
		duration := time.Duration(4)*time.Second
		time.Sleep(duration)
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
		
		body, _ := io.ReadAll(resp.Body)
		var new_id data.ID
		if err := json.Unmarshal(body, &new_id); err != nil {
			log.Fatal(err)
		}
		for idx, v := range new_id.Results {
			if !Contains(Slice, v.ID) {
				Slice = append(Slice, v.ID)
				discord.Webhook(new_id, idx)
				continue
			}

		}
		
	}
}
