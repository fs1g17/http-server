package instagram

type Feed struct {
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

type Item struct {
	MediaURL     string `json:"media_url"`
	Permalink    string `json:"permalink"`
	ThumbnailURL string `json:"thumbnail_url"`
	MediaType    string `json:"media_type"`
	ID           string `json:"id"`
}
