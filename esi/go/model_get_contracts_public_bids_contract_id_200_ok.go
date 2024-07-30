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
type GetContractsPublicBidsContractId200Ok struct {
	// The amount bid, in ISK
	Amount float32 `json:"amount"`
	// Unique ID for the bid
	BidId int32 `json:"bid_id"`
	// Datetime when the bid was placed
	DateBid time.Time `json:"date_bid"`
}
