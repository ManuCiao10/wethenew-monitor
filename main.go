package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/ManuCiao10/wethenew-monitor/data"
	"github.com/ManuCiao10/wethenew-monitor/discord"
	"github.com/ManuCiao10/wethenew-monitor/monitor"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/corpix/uarand"
	"github.com/joho/godotenv"
	"github.com/struCoder/pidusage"
	
)

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	CreateLogFile()
	log.Print("Starting monitor...")
}

func GetProducts() data.Info {
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(30),
		tls_client.WithClientProfile(tls_client.Chrome_105),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithProxyUrl(discord.GetProxy()),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Println("Error creating client: ", err)
	}
	url := "https://api-sell.wethenew.com/sell-nows?skip=0&take=50"
	req, _ := http.NewRequest("GET", url, nil)
	user_agent := uarand.GetRandom()
	req.Header = http.Header{
		"Accept":          {"application/json, text/plain, */*"},
		"accept-language": {"it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5"},
		"user-agent":      {user_agent},
		http.HeaderOrderKey: {
			"Accept",
			"accept-language",
			"user-agent",
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request: ", err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	log.Print("Response status: ", resp.Status)
	var result data.Info
	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Error unmarshalling json: ", err)
	}
	return result
}

func CreateLogFile() {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
}

func main() {
	var pid = os.Getpid()
	sysInfo, _ := pidusage.GetStat(pid)

	log.Printf("[+] CPU: %v%%\n", sysInfo.CPU)
	products := GetProducts()
	monitor.MonitorProducts(products)
}

/*
********IMPROVEMENT********
-Caching response
-Autorestart monioring after crash

*/

