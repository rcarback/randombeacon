// main.go
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
// currently a driver to grab stock data
package main

import (
	"os"													
	"fmt"
	"strings"
	"io/ioutil"
	"time"
	"./finance"
	"./yahoo"
)

type Stock finance.Stock

// Default filename containing the list of stocks you want to check.
const DefFileName = "SPX.txt"

// Get the stock names from this file and skip any that look wrong.
func getStockNames(FileName string) []string {
	// FIXME: 5 is probably too small, but should work for SPX and DJIA 
	//        averages.
	const MAXSTOCKNAMELEN = 5
	
	data, err := ioutil.ReadFile(FileName)
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "Error reading file %s: %s\n", FileName, err)
	}

	CandidateStockNames := strings.Split(string(data), "\n")

	// Perform some dumb error checking
	var StockNames []string
	for i, name := range CandidateStockNames {
		nameLen := len(name)
		if nameLen == 0 || nameLen > MAXSTOCKNAMELEN {
			fmt.Fprintf(os.Stderr, "Line %d: Stock Name '%s' invalid!\n", i, name) 
		} else {
			// FIXME: We need HTML characters for things like ^, but probably
			//        not an issue for now.
			StockNames = append(StockNames, name)
		}
	}
	return StockNames
}


func main() {
	var SPX string
	var stockData []yahoo.Stock

	if len(os.Args) > 1 {
		SPX = os.Args[1]
	} else {
		SPX = DefFileName
	}

	tickers := getStockNames(SPX)
	stockData = yahoo.GetStockData(tickers)
	for _,stock := range stockData {
		fmt.Printf("%s, %f, %s\n", stock.Name, stock.Price, 
			stock.LastUpdate.Format(time.ANSIC))
	}
}
