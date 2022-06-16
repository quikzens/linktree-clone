package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Port               string
	ServerAddress      string
	SecretKey          string
	DBSource           string
	DBName             string
	GoogleClientID     string
	GoogleClientSecret string
	GoogleRedirectUrl  string
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Couldn't loading .env file")
	}

	Port = os.Getenv("PORT")
	ServerAddress = os.Getenv("SERVER_ADDRESS")
	SecretKey = os.Getenv("SECRET_KEY")
	DBSource = os.Getenv("DB_SOURCE")
	GoogleClientID = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
	GoogleRedirectUrl = os.Getenv("GOOGLE_REDIRECT_URL")
}
