package monitor

import (
	"io"
	"log"
	"main/data"

	// "main/discord"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

type ID struct {
	New_ID    int    `json:"id"`
}

func MonitorProducts(products data.Info, client tls_client.HttpClient) {
	//do another request an dsave ONLY the ID shoes
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
	var ID_new data.Info


	
	//lopp each ID and check if is not in the struct
	//if is not in the struct send a webhook
	//add the new ID to the struct
	// fmt.Println(products)

	
}



// func FindIndex(s []int, e int) int {
// 	for i, a := range s {
// 		if a == e {
// 			return i
// 		}
// 	}
// 	return -1
// }

// func MonitorProducts(class data.Info) {
// 	Inventory := make([]int, 50)
// 	for _, product := range class.Results {
// 		for _, sellNow := range product.SellNows {
// 			fmt.Println(sellNow.ID)
// 			index := FindIndex(Inventory,sellNow.ID)
// 			if index != -1 {
// 				Inventory[index] = sellNow.ID
// 				fmt.Println("New product")
// 				discord.Webhook()
// 			}

			
// 		}
// 	}

// }