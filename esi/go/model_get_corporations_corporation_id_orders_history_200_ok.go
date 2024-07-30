/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.25
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// 200 ok object
type GetCorporationsCorporationIdOrdersHistory200Ok struct {
	// Number of days the order was valid for (starting from the issued date). An order expires at time issued + duration
	Duration int32 `json:"duration"`
	// For buy orders, the amount of ISK in escrow
	Escrow float64 `json:"escrow,omitempty"`
	// True if the order is a bid (buy) order
	IsBuyOrder bool `json:"is_buy_order,omitempty"`
	// Date and time when this order was issued
	Issued time.Time `json:"issued"`
	// The character who issued this order
	IssuedBy int32 `json:"issued_by,omitempty"`
	// ID of the location where order was placed
	LocationId int64 `json:"location_id"`
	// For buy orders, the minimum quantity that will be accepted in a matching sell order
	MinVolume int32 `json:"min_volume,omitempty"`
	// Unique order ID
	OrderId int64 `json:"order_id"`
	// Cost per unit for this order
	Price float64 `json:"price"`
	// Valid order range, numbers are ranges in jumps
	Range_ string `json:"range"`
	// ID of the region where order was placed
	RegionId int32 `json:"region_id"`
	// Current order state
	State string `json:"state"`
	// The type ID of the item transacted in this order
	TypeId int32 `json:"type_id"`
	// Quantity of items still required or offered
	VolumeRemain int32 `json:"volume_remain"`
	// Quantity of items required or offered at time order was placed
	VolumeTotal int32 `json:"volume_total"`
	// The corporation wallet division used for this order
	WalletDivision int32 `json:"wallet_division"`
}