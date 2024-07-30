/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.25
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 200 ok object
type GetLoyaltyStoresCorporationIdOffers200Ok struct {
	// Analysis kredit cost
	AkCost int32 `json:"ak_cost,omitempty"`
	// isk_cost integer
	IskCost int64 `json:"isk_cost"`
	// lp_cost integer
	LpCost int32 `json:"lp_cost"`
	// offer_id integer
	OfferId int32 `json:"offer_id"`
	// quantity integer
	Quantity int32 `json:"quantity"`
	// required_items array
	RequiredItems []GetLoyaltyStoresCorporationIdOffersRequiredItem `json:"required_items"`
	// type_id integer
	TypeId int32 `json:"type_id"`
}
