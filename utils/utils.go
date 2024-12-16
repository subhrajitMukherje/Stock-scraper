package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func GetCurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func ReadTickersFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open ticker file: %v", err)
	}
	defer file.Close()

	var tickers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			tickers = append(tickers, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading ticker file: %v", err)
	}

	return tickers, nil
}
