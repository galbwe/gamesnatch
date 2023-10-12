// scrape currency exchange rates from federalreserve.gov
package currency

import (
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type RawExchangeRate struct {
	// Exchange rate in currency units per US dollar
	Country string
	Currency string
	ExchangeDate string // TODO: convert to a date object
	ScrapeDate time.Time // TODO: convert to a date object
	Rate string
}


// TODO: write functions to clean the raw exchange rates


func ScrapeExchageRates() []RawExchangeRate {
	url := "https://www.federalreserve.gov/releases/h10/current/"

	// TODO: make the countries slice an input to the function
	countries := []string{
		"AUSTRALIA",
		"CANADA",
		"EMU MEMBERS",
		"UNITED KINGDOM",
	}

	c := colly.NewCollector()

	var ed string  // exchange date

	rates :=  []RawExchangeRate{}

	// get exchange date
	c.OnHTML("thead tr th:last-of-type", func(e *colly.HTMLElement) {
		ed = e.Text
	})

	// get rows of data
	c.OnHTML("tbody tr", func(tr *colly.HTMLElement) {
		tr.ForEach("th", func(_ int, th *colly.HTMLElement) {
			if containsOneOf(th.Text, countries) {
				rate := RawExchangeRate{}
				
				rate.Country = th.Text

				// get raw currency
				tr.ForEach("td:nth-of-type(1)", func(_ int, td *colly.HTMLElement) {
					rate.Currency = td.Text
				}) 
				
				// get raw exchange rate
				tr.ForEach("td:last-of-type", func(_ int, td *colly.HTMLElement) {
					rate.Rate = td.Text
				}) 
				rates = append(rates, rate)
			}
		})
	})

	// Before making a request print the url
	c.OnRequest(func(r *colly.Request) {
		// fmt.Println("Visiting", r.URL.String())
	})

	sd := time.Now().UTC()  // scrape date

	// add exchange date and scrape date to all raw rates
	for _, rate := range rates {
		rate.ScrapeDate = sd
		rate.ExchangeDate = ed
	}

	c.Visit(url)

	return rates
}


func containsOneOf(s string, substrs []string) bool {
	// Determines if a string contains any of a list of substrings
	for _, ss := range substrs {
		if strings.Contains(s, ss) {
			return true
		}
	}	
	return false
}
