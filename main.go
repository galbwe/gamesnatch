package main

import (
	"encoding/json"
	"fmt"

	"github.com/galbwe/gamesnatch/currency"
)

func main() {
	// items := bgg.SearchForGame("clinic")
	// out, err := json.Marshal(items[:3])
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(out))
	rates := currency.ScrapeExchageRates()

	bs, e := json.Marshal(rates)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(bs))
}
