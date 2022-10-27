package main

import (
	// "fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
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
	


func main() {
	// app := fiber.New()
	// login := Login{
	// 	Token:    os.Getenv("TOKEN"),
	// 	Email:    os.Getenv("EMAIL"),
	// 	Password: os.Getenv("PASSWORD"),
	// }
	

	


}

//add proxies
//add request
