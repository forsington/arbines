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
type GetCharactersCharacterIdFleetOk struct {
	// The character's current fleet ID
	FleetId int64 `json:"fleet_id"`
	// Member’s role in fleet
	Role string `json:"role"`
	// ID of the squad the member is in. If not applicable, will be set to -1
	SquadId int64 `json:"squad_id"`
	// ID of the wing the member is in. If not applicable, will be set to -1
	WingId int64 `json:"wing_id"`
}
