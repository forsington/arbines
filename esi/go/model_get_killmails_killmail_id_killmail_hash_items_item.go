/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.25
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// item object
type GetKillmailsKillmailIdKillmailHashItemsItem struct {
	// flag integer
	Flag int32 `json:"flag"`
	// item_type_id integer
	ItemTypeId int32 `json:"item_type_id"`
	// quantity_destroyed integer
	QuantityDestroyed int64 `json:"quantity_destroyed,omitempty"`
	// quantity_dropped integer
	QuantityDropped int64 `json:"quantity_dropped,omitempty"`
	// singleton integer
	Singleton int32 `json:"singleton"`
}
