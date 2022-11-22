package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fs1g17/http-server/cache"
	"github.com/fs1g17/http-server/instagram"
	"github.com/fs1g17/http-server/service"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	accessToken := os.Getenv("ACCESS_TOKEN")
	userId := os.Getenv("USER_ID")
	addr := ":9090"
	cacheExpirationTime := time.Hour

	instaURL := fmt.Sprintf("https://graph.instagram.com/%s/media?access_token=%s", userId, accessToken)
	fetcher := instagram.NewFetcher(instaURL, accessToken)
	cache := cache.New[[]*instagram.Item](fetcher, cacheExpirationTime)
	service := service.New(cache)
	if err := service.Server().Start(addr); err != nil {
		log.Fatalln(err)
	}
}
