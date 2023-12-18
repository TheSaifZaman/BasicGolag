package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Response Define data structure
type Response struct {
	TradeID int
	Price   string
	Size    string
	Bid     string
	Ask     string
	Volume  string
	Time    string
}

func getContent() {
	// json data
	url := "https://api.gdax.com/products/BTC-EUR/ticker"

	res, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}
	var v Response
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	err = json.NewDecoder(res.Body).Decode(&v)
	if err != nil {
		return
	}

	// print values of the object
	fmt.Printf("Price: $ %s\n", v.Price)
	fmt.Printf("Price: $ %s\n", v.Bid)
	fmt.Printf("Price: $ %s\n", v.Ask)

	os.Exit(0)
}

func main() {
	getContent()
}
