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
type GetCharactersCharacterIdLocationOk struct {
	// solar_system_id integer
	SolarSystemId int32 `json:"solar_system_id"`
	// station_id integer
	StationId int32 `json:"station_id,omitempty"`
	// structure_id integer
	StructureId int64 `json:"structure_id,omitempty"`
}
