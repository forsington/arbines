package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/forsington/arbines/esi"
	"github.com/forsington/arbines/esi/swagger"
)

var (
	salesTax = 0.0202
)

type Item struct {
	TypeId    int    `json:"type_id"`
	Name      string `json:"name"`
	Quantity  int    `json:"quantity"`
	BuyPrice  float64
	SellPrice float64
}

type Pack struct {
	Name  string `json:"name"`
	Plex  int    `json:"plex"`
	Items []Item `json:"items"`
}

func main() {
	logger := log.New(os.Stdout, "arbines: ", log.LstdFlags)

	logger.Printf("Starting up")

	jsonFile := "packs.json"
	packs, err := loadPacksFromJson(jsonFile)
	if err != nil {
		logger.Fatal("Error loading packs from json file")
	}

	plex := Item{Name: "Plex", TypeId: 44992}
	arr, err := loadPrices([]Item{plex})
	if err != nil {
		logger.Fatal("Error loading plex prices from the market")
	}
	plex = arr[0]

	// Look for arbitrage
	for _, pack := range packs {
		pack.Items, err = loadPrices(pack.Items)
		if err != nil {
			logger.Fatalf("Error loading prices for pack %s", pack.Name)
		}
		profit := pack.Arbitrage(plex.BuyPrice)

		logger.Printf("Arbitraging %s will yield %s ISK\n", pack.Name, humanize.Comma(int64(profit)))
	}

	logger.Printf("Scanned %d packs, all done\n", len(packs))
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
	loadedItems := []Item{}
	// ESI client
	esi := esi.NewClient("tranquility")

	for _, item := range items {
		orders, err := esi.GetMarketOrders(10000002, item.TypeId)
		if err != nil {
			return nil, err
		}

		item.BuyPrice = findHighestBuyOrder(orders)
		item.SellPrice = findLowestSellOrder(orders)

		loadedItems = append(loadedItems, item)
	}

	return loadedItems, nil
}

func findHighestBuyOrder(orders []swagger.GetMarketsRegionIdOrders200Ok) float64 {
	highest := 0.0
	for _, order := range orders {
		if order.IsBuyOrder && order.Price > highest {
			highest = order.Price
		}
	}
	return highest
}

func findLowestSellOrder(orders []swagger.GetMarketsRegionIdOrders200Ok) float64 {
	lowest := 1000000000.0
	for _, order := range orders {
		if !order.IsBuyOrder && order.Price < lowest {
			lowest = order.Price
		}
	}
	return lowest
}

func (i *Item) SansSalesTax() float64 {
	return i.BuyPrice * (1 - salesTax)
}

func (p *Pack) Arbitrage(plexPrice float64) float64 {
	// Calculate the cost of the pack
	cost := float64(p.Plex) * plexPrice

	// Calculate the value of the pack
	value := 0.0
	for _, item := range p.Items {
		value += item.SansSalesTax() * float64(item.Quantity)
	}

	return value - cost
}
