// Tools for working with the GeekMarket feature of BoardGameGeek

package bgg

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type SearchConfig struct {
	ItemsPerPage int
	NumItems int
}

type SearchItem struct {
	Href string
	Id string
	ObjectId string
	Subtype string
	PrimaryName string
	NameId string
	YearPublished int
	OrdTitle string
	RepImageId int
	ObjectType string
	Name string
	SortIndex string
	Type string
	ImageUrl string
}

type SearchResponseBody struct {
	Config SearchConfig
	Items []SearchItem
}


func createSearchUrl(s string) string {
	url := "https://boardgamegeek.com/api/market/products/search?ajax=1&marketdomain=boardgame&q="
	url += s
	return url
}

func SearchForGame(s string) []SearchItem {
	url := createSearchUrl(s)

	// make an http request to search for games
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	// get json as bytes from response
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// unmarshall data to a struct
	var body SearchResponseBody
	json.Unmarshal(b, &body)
	return body.Items
}
