package main

import (
	"bufio"
	"io/ioutil"
	// "io/ioutil"
	// "strings"
	// "bytes"
	// "encoding/json"
	"fmt"
	// "io"
	// "io/ioutil"
	"log"
	"math/rand"

	// "net/url"
	"os"
	"sync"
	"time"

	"github.com/ManuCiao10/wethenew-monitor/data"
	"github.com/ManuCiao10/wethenew-monitor/monitor"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/corpix/uarand"
	"github.com/joho/godotenv"
	"github.com/struCoder/pidusage"
	"go.uber.org/zap"
)

var (
	mu sync.Mutex
)

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getProxy() string{
	mu.Lock()
	file, err := os.Open("proxies.txt")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
	_ = file.Close()

	if len(txtlines) == 0{
		panic("Please add proxies to proxies.txt")
	}

	index := rand.Intn(len(txtlines))
	mu.Unlock()
	fmt.Print(txtlines[index])
	return txtlines[index]
}

func GetProducts(client tls_client.HttpClient) data.Info {
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
	fmt.Println(string(body))
	var result data.Info
	// fmt.Println(string(body))
	// if err := json.Unmarshal(body, &result); err != nil {
	// 	log.Panic("Can not unmarshal JSON => RATE_LIMITED", err)
	// }
	return result

}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
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
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(30),
		tls_client.WithClientProfile(tls_client.Chrome_105),
		tls_client.WithInsecureSkipVerify(),
		tls_client.WithProxyUrl("http://hgj3x3cas2:0ef2uixpcu@185.91.205.175:5874"),//

	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		logger.Fatal("Can not create client")
	}
	products := GetProducts(client)
	monitor.MonitorProducts(products, client)

}

//----------IMPROVEMENT----------------
//Save Logs in a file
//fix ID unique size
//restart monioring after crash
//fix docker
//add rotare ISP porxies

//----------DEBUGGING----------------
//go build -gcflags="-m" main.go

//----------README----------------
