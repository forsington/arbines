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
type GetUniverseRaces200Ok struct {
	// The alliance generally associated with this race
	AllianceId int32 `json:"alliance_id"`
	// description string
	Description string `json:"description"`
	// name string
	Name string `json:"name"`
	// race_id integer
	RaceId int32 `json:"race_id"`
}
