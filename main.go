package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "os"
	// "strings"

	// "github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Login struct {
	Token    string `json:"token"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Info struct {
	Name     string `json:"name"`
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
	Price    int    `json:"price"`
	Url      string `json:"url"`
}

func init() {
	err := godotenv.Load("config/.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Login_init() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://sell.wethenew.com/sell-now", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "sell.wethenew.com")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
	req.Header.Set("cache-control", "no-cache")
	req.Header.Set("cookie", "_ga=GA1.1.2092679579.1666843583; euconsent-v2=CPhgeQAPhgeQAAHABBENCmCgAP_AAAEAAAqIAAAAAAAA.f_gAACAAAAAA; _cs_c=1; didomi_token=eyJ1c2VyX2lkIjoiMTg0MTcxZTItZTI3OC02MzM5LTgxYzgtNzVlOWNiNzUwYTk4IiwiY3JlYXRlZCI6IjIwMjItMTAtMjdUMDQ6MDY6MjkuNzUzWiIsInVwZGF0ZWQiOiIyMDIyLTEwLTI3VDA0OjA2OjI5Ljc1M1oiLCJ2ZXJzaW9uIjoyLCJ2ZW5kb3JzIjp7ImVuYWJsZWQiOlsiYzpjb250ZW50c3F1YXJlIiwiYzpnb29nbGVhbmEtUkVVeXFXRHciXX0sInZlbmRvcnNfbGkiOnsiZW5hYmxlZCI6WyJjOmdvb2dsZWFuYS1SRVV5cVdEdyJdfX0=; _gid=GA1.2.445192787.1666843838; _cs_id=eda9db4c-4504-ac8b-f794-d826b8156749.1666843588.1.1666843838.1666843588.1.1701007588298; slrspc_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYW51ZWxlLmFyZGluZ2hpQGdtYWlsLmNvbSIsImZpcnN0bmFtZSI6ImVtYW51ZWxlIiwibGFzdG5hbWUiOiJhcmRpbmdoaSIsImlhdCI6MTY2Njg0Mzg1NywiZXhwIjoxNjcyMDI3ODU3fQ.0ZqIsxYCTiu8n44Nw9M9OFFdF1huH1shBItHsFUlU1g; slrspc_refresh_token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYW51ZWxlLmFyZGluZ2hpQGdtYWlsLmNvbSIsImZpcnN0bmFtZSI6ImVtYW51ZWxlIiwibGFzdG5hbWUiOiJhcmRpbmdoaSIsImlhdCI6MTY2Njg0Mzg1NywiZXhwIjoxNjY3NDQ4NjU3fQ.b8a4lLqMRdtTA9QTwT1dk8IMZQK2rPy7GC56_GWYWCQ; ABTasty=uid=j9y47r4wpkgk69zn&fst=1666843582502&pst=-1&cst=1666843582502&ns=1&pvt=4&pvis=4&th=; _ga_FTYGFSRM68=GS1.1.1666843583.1.1.1666844496.0.0.0")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("sec-ch-ua", `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"macOS"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", bodyText)
	fmt.Println(resp.StatusCode)

}

func main() {
	Login_init()
	// app := fiber.New()
	// login := Login{
	// 	Token:    os.Getenv("TOKEN"),
	// 	Email:    os.Getenv("EMAIL"),
	// 	Password: os.Getenv("PASSWORD"),
	// }

}

//add proxies
