package main

import "time"

type Prices struct {
	ID          int  `json:"id"`
	Whitelisted bool `json:"whitelisted"`
	Buys struct {
		Quantity  int `json:"quantity"`
		UnitPrice int `json:"unit_price"`
	} `json:"buys"`
	Sells struct {
		Quantity  int `json:"quantity"`
		UnitPrice int `json:"unit_price"`
	} `json:"sells"`
	timestamp time.Time
}
