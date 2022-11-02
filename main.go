package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"main/data"
	"main/discord"
	"main/monitor"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/joho/godotenv"
	"github.com/struCoder/pidusage"
	"go.uber.org/zap"
)

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
}

func GetProducts(client tls_client.HttpClient) data.Info {
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
	var result data.Info
	if err := json.Unmarshal(body, &result); err != nil {
		log.Panic("Can not unmarshal JSON")
	}
	return result

}

func main() {
	var pid = os.Getpid()
	sysInfo, err := pidusage.GetStat(pid)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("CPU: %v%%\n", sysInfo.CPU)
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	defer discord.Timer("main")()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(20),
		tls_client.WithClientProfile(tls_client.Chrome_105),
		tls_client.WithInsecureSkipVerify(),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		logger.Fatal("Can not create client")
	}
	products := GetProducts(client)
	monitor.MonitorProducts(products, client)
	
}

//----------IMPROVEMENT----------------
//restart monioring after crash
//add begugging memory leaks and improve code
//run in a server / docker
//send an error webhook message if the response != 200

//----------DEBUGGING----------------
//go build -gcflags="-m" main.go

//----------README----------------
//write a software description
