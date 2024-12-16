package main

import (
	"fmt"
	"github.com/putongyong/go-stock-scraper/scraper"
	"github.com/putongyong/go-stock-scraper/utils"
	"log"
	"sync"
	"time"
)

const interval int = 10

func main() {
	schedule := time.NewTicker(time.Duration(interval) * time.Second)
	defer schedule.Stop()
	fmt.Println("-------------START---------------")
	executeTask()
	for range schedule.C {
		fmt.Printf("-------------%d seconds---------------\n", interval)
		executeTask()
	}
}

func executeTask() {
	tickers, err := utils.ReadTickersFromFile("tickers.txt")
	if err != nil {
		log.Fatalf("Error reading tickers: %v", err)
	}
	var wg sync.WaitGroup

	for _, tickerSymbol := range tickers {
		wg.Add(1)
		go func(ticker string) {
			defer wg.Done()
			Printtickers(ticker)
		}(tickerSymbol)
	}
	wg.Wait()
}

func Printtickers(tickerSymbol string) {
	currentTime := utils.GetCurrentTimestamp()
	stockData := scraper.ScrapeStockData(tickerSymbol)
	fmt.Printf("Stock Data for %s:\n%s\n", tickerSymbol, currentTime)
	for key, value := range stockData {
		fmt.Printf("%s: %s\n", key, value)
	}
	fmt.Println("----------------------------")
}
