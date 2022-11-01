package main

import (
	"encoding/json"
	"io"
	"log"

	"main/data"
	"main/discord"
	"main/monitor"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/joho/godotenv"
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
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	defer discord.Timer("main")()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(3),
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
//add begugging memory leaks and impore code
//add proxies
//add loggers to errors
//Mapping upcode
//pass client and data by reference to avoid copy

//----------DEBUGGING----------------
//go build -gcflags="-m" main.go

//----------README----------------
//add readme with all the commands to run the program
//write a software description
