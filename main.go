package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/data"
	"main/discord"
	"main/monitor"

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

func Login_init(client tls_client.HttpClient) {
	url :=  "https://api-sell.wethenew.com/sell-nows?skip=0&take=50"
	req, err := http.NewRequest("GET",url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "it-IT,it;q=0.9,en-US;q=0.8,en;q=0.7,de;q=0.6,fr;q=0.5")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var result data.Info
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to the go struct pointer
		log.Panic("Can not unmarshal JSON")
	}
	monitor.Check(result) // check if there is a new product
	discord.Webhook(result)
	//---CREATE A LOOP TO GET ALL THE SELL NOWS----//
	//---If is new send a webhook----//
	// fmt.Println(PrettyPrint(result))
	// for _, rec := range result.Results {
	// 	fmt.Println(rec.Name)
	// }

}

func main() {
	defer discord.Timer("main")()
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

}

//----------IMPROVEMENT----------------
//add begugging memory leaks and impore code
//add proxies
//save data and create a for loop
//check cache to do not do more requests
//check if cookies expired or try to do the login
//add loggers to errors
//add rotare user-agent
//imporve headers clean look better

//add an array with all the actual ID numebres and check if there is a new one => send webhook
//check long polling to get the new products

//----------DEBUGGING----------------
//go build -gcflags="-m" main.go

//----------README----------------
//add readme with all the commands to run the program

//----------TODO----------------
//1.function to save all the data
//2.polling until new products
//3.check if there is a new product
//4.send webhook


//----------NOTES----------------
//