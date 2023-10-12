package main

import (
	"encoding/json"
	"fmt"

	"github.com/galbwe/gamesnatch/bgg"
)

func main() {
	items := bgg.SearchForGame("clinic")
	out, err := json.Marshal(items[:3])
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
}
