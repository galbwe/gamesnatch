package main

import "github.com/galbwe/gamesnatch/currency"

func main() {
	// items := bgg.SearchForGame("clinic")
	// out, err := json.Marshal(items[:3])
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(out))
	currency.ScrapeExchageRates()
}
