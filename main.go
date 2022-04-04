package main

import (
	"time"

	"github.com/cbrissonCoveo/CoveStonks/api"
	"github.com/cbrissonCoveo/CoveStonks/msg"
	"github.com/cbrissonCoveo/CoveStonks/utils"
	"github.com/pterm/pterm"
)

func main() {

	utils.Clear()
	utils.TerminalTitle(10)
	pterm.Print("\n\n")

	area, _ := pterm.DefaultArea.Start()
	for {
		stock := api.Quote()
		price := msg.StockPrice(stock.Price(), stock.PriceState())
		utils.TerminalTitle(stock.Price())
		statistics := msg.Statistics(stock, 43830000)
		area.Update(price + statistics)
		time.Sleep(time.Second * 5)
	}
}
