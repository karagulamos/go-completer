package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/*QuoteResponse is a container for the last stock price*/
type QuoteResponse struct {
	Name      string
	LastPrice float32
}

func main() {
	start := time.Now()

	stockURL := "http://dev.markitondemand.com/MODApis/Api/v2/quote?symbol="

	stockSymbols := []string{
		"googl", "msft", "aapl", "bbry", "hpq", "vz", "t", "tmus", "s",
	}

	completer := NewCompleter()

	for idx, symbol := range stockSymbols {
		completer.Register()

		go func(symbol string, completer *Completer, idx int) {
			index := idx + 1

			resp, _ := http.Get(fmt.Sprintf("%s%s", stockURL, symbol))

			body, _ := ioutil.ReadAll(resp.Body)

			defer resp.Body.Close()

			if msg := string(body); strings.HasPrefix(msg, "Request blocked") {
				fmt.Printf("%d. %s - %s\n", index, symbol, msg)
			} else {
				quote := &QuoteResponse{}
				xml.Unmarshal(body, quote)

				fmt.Printf("%d. %s: %.2f\n", index, quote.Name, quote.LastPrice)
			}

			completer.Signal()
		}(symbol, completer, idx)
	}

	completer.WaitAll()

	elapsed := time.Since(start)

	fmt.Printf("Execution time: %s", elapsed)
}
