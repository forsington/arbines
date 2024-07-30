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
type GetFwSystems200Ok struct {
	// contested string
	Contested string `json:"contested"`
	// occupier_faction_id integer
	OccupierFactionId int32 `json:"occupier_faction_id"`
	// owner_faction_id integer
	OwnerFactionId int32 `json:"owner_faction_id"`
	// solar_system_id integer
	SolarSystemId int32 `json:"solar_system_id"`
	// victory_points integer
	VictoryPoints int32 `json:"victory_points"`
	// victory_points_threshold integer
	VictoryPointsThreshold int32 `json:"victory_points_threshold"`
}