package main

import (
	"encoding/json"
	"os"
)

type Item struct {
	TypeId int
	Name   string
	ISK    float64
}

type Pack struct {
	Name  string
	Plex  int
	Items []Item
}

func main() {
	jsonFile := "packs.json"

	packsFromJson, err := loadPacksFromJson(jsonFile)
	if err != nil {
		println("Error loading packs from json file")
		os.Exit(1)
	}

	items := []Item{}
	items = append(items, Item{Name: "Plex", TypeId: 44992})
	for _, pack := range packsFromJson {
		for _, item := range pack.Items {
			items = append(items, item)
		}
	}

	// Load the prices from the market
	pricedItems := loadPrices(items)

	// Look for arbitrage
}

func loadPacksFromJson(jsonFile string) ([]Pack, error) {
	file, _ := os.Open(jsonFile)
	decoder := json.NewDecoder(file)
	packs := []Pack{}
	err := decoder.Decode(&packs)

	if err != nil {
		return nil, err
	}

	return packs, nil
}

func loadPrices(items []Item) ([]Item, error) {
	// ESI client

}
