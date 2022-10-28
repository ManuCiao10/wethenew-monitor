package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/discord"
	"time"

	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/joho/godotenv"
)

type Login struct {
	Token    string `json:"token"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)

}

func init() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("|%v|\n", time.Since(start))
	}
}

func Login_init(client tls_client.HttpClient) {

	req, err := http.NewRequest("GET", "https://api-sell.wethenew.com/sell-nows?skip=0&take=50", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api-sell.wethenew.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
	req.Header.Set("authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYW51ZWxlLmFyZGluZ2hpQGdtYWlsLmNvbSIsImZpcnN0bmFtZSI6ImVtYW51ZWxlIiwibGFzdG5hbWUiOiJhcmRpbmdoaSIsImlhdCI6MTY2Njg0Mzg1NywiZXhwIjoxNjcyMDI3ODU3fQ.0ZqIsxYCTiu8n44Nw9M9OFFdF1huH1shBItHsFUlU1g")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("feature-policy", "microphone 'none'; geolocation 'none'; camera 'none'; payment 'none'; battery 'none'; gyroscope 'none'; accelerometer 'none';")
	req.Header.Set("origin", "https://sell.wethenew.com")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://sell.wethenew.com/")
	req.Header.Set("sec-ch-ua", `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	req.Header.Set("x-xss-protection", "1;mode=block")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result discord.Info
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	discord.Webhook(result)

	// fmt.Println(PrettyPrint(result))
	for _, rec := range result.Results {
		fmt.Println(rec.Name)
	}


}





func main() {
	defer timer("main")()
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeout(7),
		tls_client.WithClientProfile(tls_client.Chrome_105),
		tls_client.WithInsecureSkipVerify(),
		// tls_client.WithNotFollowRedirects(),
		//tls_client.WithProxyUrl("http://user:pass@host:ip"),
		// tls_client.WithCookieJar(cJar), // create cookieJar instance and pass it as argument
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Fatal(err)
	}
	Login_init(client)

	// app := fiber.New()
	// login := Login{
	// 	Token:    os.Getenv("TOKEN"),
	// 	Email:    os.Getenv("EMAIL"),
	// 	Password: os.Getenv("PASSWORD"),
	// }

}

//add begugging memory leaks and impore code
//add proxies
//save data and create a for loop
//check only new data at index 0
//check cache to do not do more requests
//check if cookies expired or try to do the login

//url := "https://sell.wethenew.com/sell-now"
// 	req, err := http.NewRequest("GET", url, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	req.Header.Set("authority", "sell.wethenew.com")
// 	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
// 	req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
// 	req.Header.Set("cache-control", "no-cache")
// 	req.Header.Set("cookie", "euconsent-v2=CPhgeQAPhgeQAAHABBENCmCgAP_AAAEAAAqIAAAAAAAA.f_gAACAAAAAA; _cs_c=1; didomi_token=eyJ1c2VyX2lkIjoiMTg0MTcxZTItZTI3OC02MzM5LTgxYzgtNzVlOWNiNzUwYTk4IiwiY3JlYXRlZCI6IjIwMjItMTAtMjdUMDQ6MDY6MjkuNzUzWiIsInVwZGF0ZWQiOiIyMDIyLTEwLTI3VDA0OjA2OjI5Ljc1M1oiLCJ2ZXJzaW9uIjoyLCJ2ZW5kb3JzIjp7ImVuYWJsZWQiOlsiYzpjb250ZW50c3F1YXJlIiwiYzpnb29nbGVhbmEtUkVVeXFXRHciXX0sInZlbmRvcnNfbGkiOnsiZW5hYmxlZCI6WyJjOmdvb2dsZWFuYS1SRVV5cVdEdyJdfX0=; _gid=GA1.2.445192787.1666843838; _cs_id=eda9db4c-4504-ac8b-f794-d826b8156749.1666843588.1.1666843838.1666843588.1.1701007588298; slrspc_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYW51ZWxlLmFyZGluZ2hpQGdtYWlsLmNvbSIsImZpcnN0bmFtZSI6ImVtYW51ZWxlIiwibGFzdG5hbWUiOiJhcmRpbmdoaSIsImlhdCI6MTY2Njg0Mzg1NywiZXhwIjoxNjcyMDI3ODU3fQ.0ZqIsxYCTiu8n44Nw9M9OFFdF1huH1shBItHsFUlU1g; slrspc_refresh_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYW51ZWxlLmFyZGluZ2hpQGdtYWlsLmNvbSIsImZpcnN0bmFtZSI6ImVtYW51ZWxlIiwibGFzdG5hbWUiOiJhcmRpbmdoaSIsImlhdCI6MTY2Njg0Mzg1NywiZXhwIjoxNjY3NDQ4NjU3fQ.b8a4lLqMRdtTA9QTwT1dk8IMZQK2rPy7GC56_GWYWCQ; ABTasty=uid=j9y47r4wpkgk69zn&fst=1666843582502&pst=-1&cst=1666843582502&ns=1&pvt=4&pvis=4&th=; _ga_FTYGFSRM68=GS1.1.1666843583.1.1.1666844496.0.0.0")
// 	req.Header.Set("pragma", "no-cache")
// 	req.Header.Set("sec-ch-ua", `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`)
// 	req.Header.Set("sec-ch-ua-mobile", "?0")
// 	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
// 	req.Header.Set("sec-fetch-dest", "document")
// 	req.Header.Set("sec-fetch-mode", "navigate")
// 	req.Header.Set("sec-fetch-site", "same-origin")
// 	req.Header.Set("sec-fetch-user", "?1")
// 	req.Header.Set("upgrade-insecure-requests", "1")
// 	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
// 	response, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	doc, _ := goquery.NewDocumentFromReader(response.Body)
// 	// fmt.Println(doc.Html())

// 	fmt.Println(doc.Find("name").Text())
//     title := doc.Find("price").Text()
//     fmt.Println(title)

//----------DEBUGGING----------------
//go build -gcflags="-m" main.go
