package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"time"

	"github.com/ManuCiao10/wethenew-monitor/data"

)

const (
	Image_URL = "https://cdn.discordapp.com/attachments/965899789021642752/965899835570016286/DBFF8755-874B-4436-B79A-0C02DDBBEBBA.jpg"
)

var cookieJar, _ = cookiejar.New(nil)

var client = &http.Client{
	Jar: cookieJar,
}

func Time() string {
	date := time.Now().Format("15:04:05")
	time := time.Now().UnixNano() / int64(time.Millisecond)
	time_final := fmt.Sprintf("%s.%d", date, time%1000)
	return time_final
}

func Webhook(new_id data.ID, idx int) {
	var webhookURL = os.Getenv("DISCORD_WEBHOOK_URL_TEST")
	n_size := len(new_id.Results[idx].SellNows)
	var fields []data.Fields
	for i := 0; i < n_size; i++ {
		fields = append(fields, data.Fields{
			Name:   "Payout",
			Value:  "[" + strconv.Itoa(new_id.Results[idx].SellNows[i].Price) + " €" + " | " + new_id.Results[idx].SellNows[i].Size + "]" + "(" + "https://sell.wethenew.com/sell-now/" + strconv.Itoa(new_id.Results[idx].SellNows[i].ID) + "?holding-Lab" + ")",
			Inline: true,
		})
	}
	payload := &data.Top{
		Username:  "Wethenew Monitor",
		AvatarURL: Image_URL,
		Content:   "",
		Embeds: []data.Embeds{
			{
				Title: new_id.Results[idx].Name,
				// Description: "Sell Now",
				Color:  1999236,
				Fields: fields,
				Thumbnail: data.Thumbnail{
					URL: new_id.Results[idx].Image,
				},
				Footer: data.Footer{
					IconURL: Image_URL,
					Text:    "Wethenew | Holding-Lab " + Time(),
				},
			},
		},
	}
	payloadBuf := new(bytes.Buffer)
	_ = json.NewEncoder(payloadBuf).Encode(payload)

	if webhookURL == "" {
		fmt.Println("SET DISCORD_WEBHOOK_URL ENV VAR")
	}
	SendWebhook, err := http.NewRequest("POST", webhookURL, payloadBuf)
	if err != nil {
		fmt.Println(err)
	}
	SendWebhook.Header.Set("content-type", "application/json")

	sendWebhookRes, err := client.Do(SendWebhook)
	if err != nil {
		fmt.Print(err)
	}
	if sendWebhookRes.StatusCode != 204 {
		fmt.Printf("Webhook failed with status %d\n", sendWebhookRes.StatusCode)
	}
	defer sendWebhookRes.Body.Close()
}