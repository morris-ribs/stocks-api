package models

//StockHistory represents a StockHistory
type StockHistory struct {
	Date  string  `json:"date"`
	Price float32 `json:"price"`
}

//StockHistoryList represents a list of StockHistory
type StockHistoryList []StockHistory

//Stock represents a stock
type Stock struct {
	Name         string           `json:"name"`
	CurrentPrice float32          `json:"cur_price"`
	CompanyName  string           `json:"company_name"`
	HistoryList  StockHistoryList `json:"history_prices"`
}

//Stocks is an array of Stock
type Stocks []Stock
