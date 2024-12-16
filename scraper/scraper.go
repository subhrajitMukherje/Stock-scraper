package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeStockData(tickerSymbol string) map[string]string {
	c := colly.NewCollector()
	stock := make(map[string]string)
	regularMarketSelectors := map[string]string{
		"regular_market_price":          fmt.Sprintf(`[data-symbol="%s"][data-field="regularMarketPrice"]`, tickerSymbol),
		"regular_market_change":         fmt.Sprintf(`[data-symbol="%s"][data-field="regularMarketChange"]`, tickerSymbol),
		"regular_market_change_percent": fmt.Sprintf(`[data-symbol="%s"][data-field="regularMarketChangePercent"]`, tickerSymbol),
	}
	for key, selector := range regularMarketSelectors {
		c.OnHTML(selector, func(e *colly.HTMLElement) {
			value := strings.TrimSpace(e.Text)
			if key == "regular_market_change_percent" {
				value = strings.ReplaceAll(value, "(", "")
				value = strings.ReplaceAll(value, ")", "")
			}
			if _, exists := stock[key]; exists {
				return
			}
			stock[key] = value
		})
	}
	err := c.Visit("https://finance.yahoo.com/quote/" + tickerSymbol)
	if err != nil {
		fmt.Printf("Error visiting page for %s: %v\n", tickerSymbol, err)
	}

	return stock
}
