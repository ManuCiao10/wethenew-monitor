package monitor

import (
	"encoding/json"
	"io"
	"log"
	// "os"
	"time"

	"github.com/ManuCiao10/wethenew-monitor/data"
	"github.com/ManuCiao10/wethenew-monitor/discord"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
)

func MonitorProducts(class data.Info) {
	Slice := discord.SaveSlice(class)
	url := "https://api-sell.wethenew.com/sell-nows?skip=0&take=50"
	for {
		options := []tls_client.HttpClientOption{
			tls_client.WithTimeout(30),
			tls_client.WithClientProfile(tls_client.Chrome_105),
			tls_client.WithNotFollowRedirects(),
			tls_client.WithProxyUrl(discord.GetProxy()),
		}
		time.Sleep(time.Duration(10) * time.Second)
		client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
		if err != nil {
			log.Print(err)
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println(err)
			continue
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
		time.Sleep(time.Duration(10) * time.Second)
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("Status code: %d\n", resp.StatusCode)

		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		var new_id data.ID
		if err := json.Unmarshal(body, &new_id); err != nil {
			log.Println(err)
		}
		for idx, v := range new_id.Results {
			if !discord.Contains(Slice, v.ID) {
				log.Print("New product found!")
				Slice = append(Slice, v.ID)
				discord.Webhook(new_id, idx)
				continue
			}
		}
	}
}
