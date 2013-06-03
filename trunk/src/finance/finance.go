// finance.go
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
// This encapsulates a data type for passing stock data
package finance

import (
	"time"
)

// This stores the basic stock information we'll be downloading from Yahoo.
type Stock struct {
	Name string
	Price float32
	LastUpdate time.Time
}
