package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		log.Println(err)
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
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(resp.StatusCode)
	var result data.Info
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON => RATE_LIMITED", err)
	}
	return result
}

func main() {
	var pid = os.Getpid()
	sysInfo, _ := pidusage.GetStat(pid)
	fmt.Printf("CPU: %v%%\n", sysInfo.CPU)

	products := GetProducts()
	monitor.MonitorProducts(products)

}

//----------IMPROVEMENT----------------
//Save Logs in a file
//restart monioring after crash
//fix docker

//----------DEBUGGING----------------
//go build -gcflags="-m" main.go

//----------README----------------
