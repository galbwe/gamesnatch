package currency

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)


func TestCleanExchangeRate(t *testing.T) {
	raw := RawExchangeRate{
		Country: "*AUSTRALIA",
		Currency: "DOLLAR",
		ExchangeDate: "      Oct. 6  ",
		ScrapeDate: time.Date(2000, 2, 1, 12, 30, 0, 0, time.UTC),
		Rate: "0.6394",
	}
	clean := CleanExchangeRate(raw)
	assert.Equal(t, "au", clean.Country)
}
