package instagram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Fetcher struct {
	instaURL    string
	accessToken string
}

func NewFetcher(instaURL string, accessToken string) *Fetcher {
	return &Fetcher{
		instaURL:    instaURL,
		accessToken: accessToken,
	}
}

func (f *Fetcher) Fetch() (items []*Item, err error) {
	feed, err := f.feed()
	if err != nil {
		return
	}

	for i := 0; i < 9 && i < len(feed.Data); i++ {
		item, err := f.item(feed.Data[i].ID)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return
}

func (f *Fetcher) feed() (feed *Feed, err error) {
	resp, err := http.Get(f.instaURL)
	if err != nil {
		return nil, fmt.Errorf("error making HTTP request: %s", err)
	}

	feed = new(Feed)
	if err := json.NewDecoder(resp.Body).Decode(feed); err != nil {
		return nil, fmt.Errorf("error decoding insta feed: %s", err)
	}

	return
}

func (f *Fetcher) item(id string) (instaItem *Item, err error) {
	mediaUrl := fmt.Sprintf(
		"https://graph.instagram.com/%s?access_token=%s&fields=media_url,permalink,thumbnail_url,media_type",
		id,
		f.accessToken,
	)

	resp, err := http.Get(mediaUrl)
	if err != nil {
		return nil, fmt.Errorf("error getting media item with id %s: %s", id, err)
	}

	instaItem = new(Item)
	if err := json.NewDecoder(resp.Body).Decode(instaItem); err != nil {
		return nil, fmt.Errorf("error decoding media item with id %s: %s", id, err)
	}

	return
}
