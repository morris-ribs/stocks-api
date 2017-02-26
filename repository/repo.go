package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"stocks/models"
)

var stocksList models.Stocks

//sliceIndex is a generic function to find an item in a slice
func sliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

// loadStocks loads the list of stocks
func loadStocks() models.Stocks {
	absPath, _ := filepath.Abs("./data.json")
	file, e := ioutil.ReadFile(absPath)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	results := models.Stocks{}
	if err := json.Unmarshal(file, &results); err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}

// GetStocks returns the list of stocks
func GetStocks() models.Stocks {
	if stocksList == nil {
		stocksList = loadStocks()
	}
	fmt.Println(stocksList)
	return stocksList
}

// AddStock inserts a stock in the DB
func AddStock(stock models.Stock) bool {
	stocksList = append(stocksList, stock)

	return true
}

// UpdateStock updates a stock in the DB
func UpdateStock(stock models.Stock) bool {
	index := sliceIndex(len(stocksList), func(i int) bool { return stocksList[i].Name == stock.Name })
	fmt.Println(index)
	fmt.Println(stocksList)
	stocksList[index] = stock

	return true
}
