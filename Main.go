package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

var gw2api string = "https://api.guildwars2.com/v2/"
var sellIdsEndpoint string = "commerce/prices/"

var sellIdsSlice *[]uint32

func main() {
	getAllSellableIds()
	goAllIds()
}

func getLastPrice(itemId uint32) {
	fmt.Println("getLastPrice: " + fmt.Sprint(itemId))
	response, err := http.Get(gw2api + sellIdsEndpoint + fmt.Sprint(itemId))
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {

		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}
		var item *Prices
		serializeError := json.Unmarshal([]byte(contents), &item)
		if serializeError != nil {
			fmt.Printf("Error: %s", serializeError)
		}
		go goInsertItemDb(item)
	}
}

func goInsertItemDb(item *Prices){
	//fmt.Println("testing")
	//fmt.Println("sell price: " + fmt.Sprint(item.Sells.UnitPrice))
}

func goAllIds() {
	for _, element := range *sellIdsSlice {
		go getLastPrice(element)
	}
}

func getAllSellableIds() *[]uint32 {
	response, err := http.Get(gw2api + sellIdsEndpoint)
	if err != nil {
		fmt.Printf("Error: %s", err)
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err.Error())
		}

		serializeError := json.Unmarshal([]byte(contents), &sellIdsSlice)
		if serializeError != nil {
			fmt.Printf("Error: %s", serializeError)
		}
	}

	return sellIdsSlice
}
