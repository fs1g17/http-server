package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fs1g17/http-server/cache"
	"github.com/fs1g17/http-server/instagram"
	"github.com/fs1g17/http-server/service"
)

const (
	accessToken         = "IGQVJXcUdSNDdPOTdHamJBSDJodDdUSXlsbmpCNVZAFYlZABcTlUWU1kWWZAUS0p1ZAFJTVmJpc095bjRfZAHpJTW9zMC10NUVTWk9YV3dQUDRqY0ZAIV2ZAsanF3QmtVdUo1X0tKdHFfM0ZALMFJSZAUVxcTNsSwZDZD"
	userId              = "5626681774019951"
	addr                = ":9090"
	cacheExpirationTime = time.Hour
)

func main() {
	instaURL := fmt.Sprintf("https://graph.instagram.com/%s/media?access_token=%s", userId, accessToken)
	fetcher := instagram.NewFetcher(instaURL, accessToken)
	cache := cache.New[[]*instagram.Item](fetcher, cacheExpirationTime)
	service := service.New(cache)
	if err := service.Server().Start(addr); err != nil {
		log.Fatalln(err)
	}
}
