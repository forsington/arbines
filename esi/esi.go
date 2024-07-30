package esi

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/antihax/optional"
	"github.com/forsington/arbines/esi/swagger"
)

func NewClient(server string) *Client {
	client := &http.Client{
		Timeout: time.Minute,
	}
	config := swagger.Configuration{
		BasePath: "https://esi.evetech.net/latest",
		// UserAgent:  "Testing ESI for a trading script, contact Repping in game",
		HTTPClient: client,
	}
	esi := swagger.NewAPIClient(&config)
	return &Client{
		esi:    esi,
		calls:  0,
		server: server,
	}
}

type Client struct {
	esi    *swagger.APIClient
	calls  int
	server string
}

func (c *Client) Calls() int {
	return c.calls
}

func (c *Client) GetMarketOrders(regionId int, typeId int) ([]swagger.GetMarketsRegionIdOrders200Ok, error) {
	opts := &swagger.MarketApiGetMarketsRegionIdOrdersOpts{
		Datasource: optional.NewString(c.server),
		TypeId:     optional.NewInt32(int32(typeId)),
		Page:       optional.NewInt32(1),
	}

	orders, resp, err := c.esi.MarketApi.GetMarketsRegionIdOrders(context.Background(), "all", int32(regionId), opts)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("error fetching market orders for item %d from ESI: %v", typeId, err)
	}
	c.calls = c.calls + 1

	return orders, err
}
