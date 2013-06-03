// yahoo.go
//
//-===============================LICENSE INFO===============================-// 
//
// Copyright (C) 2013 by Richard Carback
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 3 of the License, or
// (at your option) any later version.
// 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.
//
//-=============================END LICENSE INFO=============================-// 
//
// Last Updated: 6/2/2013
// By: Richard Carback <rick.carback@gmail.com>
//
// This is a simple tool to suck down and cache stock data using the Yahoo
// API.
package yahoo

import (
	"os"
	"time"
	"strings"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"../finance"
)

// retype "finance.Stock" to "Stock
type Stock finance.Stock

// The Yahoo API takes 2 parameters:
//   * s - The stock ticker symbol (usually 1-5 characters)
//   * f - The format, (symbol, last trading price, date, time) 
// The API is described in detail at:
//   https://code.google.com/p/yahoo-finance-managed/wiki/csvQuotesDownload
const YahooURL = "http://finance.yahoo.com/d/quotes.csv?s=%s&f=sl1d1t1"
// Intervals between URL requests to the server.
var Delay = 1000
// Maximum number of stocks you can request at once.
var MaxStockNum = 50
// Format string for time
var TimeFmt = "20060102"

// Check to see if we have already downloaded data in a specific filename
func getCachedStockData(filename string, stockData *[]Stock) bool {
	data, err := ioutil.ReadFile(filename)
	if (err == nil) {
		err := json.Unmarshal(data, &stockData)
		if (err == nil) {
			return true
		}
	}
	fmt.Println(err)
	return false
}

// Save retrieved stock data using the given date.
// func cacheStockData(stocks []Stock, date time.Time) {
// }

// Download and parse data using the given URL. This function expects
// to be called as a thread.
func downloadStockData(url string) []Stock {
	var res []Stock
	return res
}

// Use the given tickers to download or retrieve the stock data from cache.
func GetStockData(tickers []string) []Stock {
	var res []Stock = make([]Stock, len(tickers))

	if len(tickers) == 0 {
		return res
	}

	today := time.Now()
	cacheFName := fmt.Sprintf("%s-stocks.json", today.Format(TimeFmt))

	// Retrieve cached data if possible
	if !getCachedStockData(cacheFName, &res) {
		// Split tickers into groups of MaxStockNum
		// For each group, fire off a request to download the data
		for i := len(tickers); i > 0; i -= MaxStockNum {		
			urlStr := fmt.Sprintf(YahooURL,
				strings.Join(tickers[max(0,i-MaxStockNum):i], ","))
			
			fmt.Println(urlStr)
		}

		// Cache the retrieved data
		cacheData, err := json.Marshal(res)
		if (err == nil) {
			err := ioutil.WriteFile(cacheFName, cacheData, os.ModePerm)
			if (err != nil) {
				fmt.Fprintln(os.Stderr, "Unable to write cache data!")
			}
		} else {
			fmt.Fprintln(os.Stderr, "Unable to marshal cache data!")
		}
	}

	// TODO: Check that each stock requested ended up in the data and print any that 
	// are missing. NOTE: this can be caused if a ticker is bad, or the historical
	// data doesn't have it. 

	return res
}

// I died a little inside when I found out math only does it's magic in float64
func max (x, y int) int {
	if (y > x) {
		return y
	}
	return x
}
