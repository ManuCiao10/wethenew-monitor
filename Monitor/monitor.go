package monitor

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ManuCiao10/wethenew-monitor/data"
	"github.com/ManuCiao10/wethenew-monitor/discord"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/patrickmn/go-cache"
)

func Time() string {
	date := time.Now().Format("15:04:05")
	time := time.Now().UnixNano() / int64(time.Millisecond)
	time_final := fmt.Sprintf("%s.%d", date, time%1000)
	return time_final
}

func MonitorProducts(class data.Info) {
	c := cache.New(cache.NoExpiration, cache.NoExpiration)
	for _, v := range class.Results {
		c.Set(fmt.Sprintf("%d", v.ID), v.ID, cache.NoExpiration)
	}
	// c.Delete(fmt.Sprintf("%d",1691)) For testing
	fmt.Print("Cache: ", c.Items())
	url := "https://api-sell.wethenew.com/sell-nows?skip=0&take=50"
	for {
		options := []tls_client.HttpClientOption{
			tls_client.WithTimeout(30),
			tls_client.WithClientProfile(tls_client.Chrome_105),
			tls_client.WithNotFollowRedirects(),
			tls_client.WithProxyUrl(discord.GetProxy()),
		}

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
		// time.Sleep(time.Duration(3) * time.Second)
		resp, err := client.Do(req)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("[%s] Status code: <|%d|> \n", Time(), resp.StatusCode)
		body, _ := io.ReadAll(resp.Body)
		_ = resp.Body.Close()
		var new_id data.ID
		if err := json.Unmarshal(body, &new_id); err != nil {
			log.Println(err)
		}
		for idx, v := range new_id.Results {
			if _, found := c.Get(fmt.Sprintf("%d", v.ID)); !found {

				c.Set(fmt.Sprintf("%d", v.ID), v.ID, cache.NoExpiration)
				discord.Webhook(new_id, idx)
			}
		}
	}
}
