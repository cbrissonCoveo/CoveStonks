package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
)

// Quote structure based on Yahoo's API
type QuoteStructure struct {
	Symbol                     string
	ShortName                  string
	RegularMarketPrice         float64 // WE EXPECT HIGH NUMBERS, RIGHT
	RegularMarketChange        float64
	RegularMarketChangePercent float64
	RegularMarketVolume        int
	PreMarketPrice             float64
	PreMarketChange            float64
	PreMarketChangePercent     float64
	PostMarketPrice            float64
	PostMarketChange           float64
	PostMarketChangePercent    float64
	MarketState                string
	Currency                   string
	Exchange                   string
}

// Quote structure based on Yahoo's API
type Result struct {
	Result []QuoteStructure `json:"result"`
}

// Quote structure based on Yahoo's API
type QuoteResponse struct {
	QuoteResponse Result `json:"quoteResponse"`
}

func Quote() QuoteStructure {
	url := "https://query1.finance.yahoo.com/v7/finance/quote?symbols=CVO.TO"
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	api_data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var quote QuoteResponse
	jsonErr := json.Unmarshal(api_data, &quote)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return quote.QuoteResponse.Result[0]
}

// My doc
func (stock *QuoteStructure) Price() float64 {
	return stock.RegularMarketPrice
}

func (stock *QuoteStructure) PriceState() bool {
	return !math.Signbit(stock.RegularMarketChange)
}

func (stock *QuoteStructure) State() string {
	switch stock.MarketState {
	case "PRE":
		return "Pre"
	case "POST":
		return "After hours"
	case "CLOSED", "POSTPOST":
		return "Closed"
	default:
		return "Open"

	}
}

func (stock *QuoteStructure) MarketChangePercent() string {
	return fmt.Sprintf("%.2f", stock.RegularMarketChangePercent) + "%"
}

func (stock *QuoteStructure) MarketChange() string {
	return fmt.Sprintf("%.2f", stock.RegularMarketChange) + stock.Currency
}

func (stock *QuoteStructure) MarketVolume() int {
	if stock.MarketState == "REGULAR" {
		return stock.RegularMarketVolume
	}

	return 0
}

func (stock *QuoteStructure) Name() string {
	return stock.ShortName
}

func (stock *QuoteStructure) StockCurrency() string {
	return stock.Currency
}
