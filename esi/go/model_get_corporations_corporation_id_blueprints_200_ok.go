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
type GetCorporationsCorporationIdBlueprints200Ok struct {
	// Unique ID for this item.
	ItemId int64 `json:"item_id"`
	// Type of the location_id
	LocationFlag string `json:"location_flag"`
	// References a station, a ship or an item_id if this blueprint is located within a container.
	LocationId int64 `json:"location_id"`
	// Material Efficiency Level of the blueprint.
	MaterialEfficiency int32 `json:"material_efficiency"`
	// A range of numbers with a minimum of -2 and no maximum value where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (e.g. no activities performed on them yet).
	Quantity int32 `json:"quantity"`
	// Number of runs remaining if the blueprint is a copy, -1 if it is an original.
	Runs int32 `json:"runs"`
	// Time Efficiency Level of the blueprint.
	TimeEfficiency int32 `json:"time_efficiency"`
	// type_id integer
	TypeId int32 `json:"type_id"`
}