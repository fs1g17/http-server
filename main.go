package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

type InstaFeed struct {
	Data []struct {
		ID string `json:"id"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

type InstaItem struct {
	MediaURL     string `json:"media_url"`
	Permalink    string `json:"permalink"`
	ThumbnailURL string `json:"thumbnail_url"`
	MediaType    string `json:"media_type"`
	ID           string `json:"id"`
}

const accessToken = "IGQVJXcUdSNDdPOTdHamJBSDJodDdUSXlsbmpCNVZAFYlZABcTlUWU1kWWZAUS0p1ZAFJTVmJpc095bjRfZAHpJTW9zMC10NUVTWk9YV3dQUDRqY0ZAIV2ZAsanF3QmtVdUo1X0tKdHFfM0ZALMFJSZAUVxcTNsSwZDZD"
const userId = "5626681774019951"

var instaUrl = fmt.Sprintf("https://graph.instagram.com/%s/media?access_token=%s", userId, accessToken)

func main() {
	http.HandleFunc("/", getInstaFeed)

	err := http.ListenAndServe(":9090", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getInstaFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	feed, err := fetchInstaFeed(w)
	if err != nil {
		fmt.Fprintf(w, "error: %s", err)
		return
	}

	data := feed.Data

	var instaItems []InstaItem
	for i := 0; i < 9 || i == len(data); i++ {
		instaItem, err := fetchInstaItem(data[i].ID, w)
		if err != nil {
			fmt.Fprintf(w, "error: %s", err)
			return
		}

		instaItems = append(instaItems, *instaItem)
	}

	err = json.NewEncoder(w).Encode(instaItems)
	if err != nil {
		fmt.Fprintf(w, "oops %s", err)
		return
	}

}

func fetchInstaFeed(w http.ResponseWriter) (*InstaFeed, error) {
	resp, err := http.Get(instaUrl)
	if err != nil {
		fmt.Fprintf(w, "error making http request: %s", err)
		return nil, fmt.Errorf("error making http request")
	}

	var feed InstaFeed
	if err := json.NewDecoder(resp.Body).Decode(&feed); err != nil {
		fmt.Fprintf(w, "error decoding insta feed body: %s", err)
		return nil, fmt.Errorf("error decoding insta feed")
	}

	return &feed, nil
}

func fetchInstaItem(id string, w http.ResponseWriter) (*InstaItem, error) {
	mediaUrl := fmt.Sprintf("https://graph.instagram.com/%s?access_token=%s&fields=media_url,permalink,thumbnail_url,media_type", id, accessToken)

	resp, err := http.Get(mediaUrl)
	if err != nil {
		fmt.Fprintf(w, "error getting media item  with id: %s, error: %s", id, err)
		return nil, fmt.Errorf("error getting media item")
	}

	var instaItem InstaItem
	if err := json.NewDecoder(resp.Body).Decode(&instaItem); err != nil {
		fmt.Fprintf(w, "error deconding media item with id: %s, error: %s", id, err)
		return nil, fmt.Errorf("error decoding media item")
	}

	return &instaItem, nil
}
