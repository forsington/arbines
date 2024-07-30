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
type GetCorporationsCorporationIdStandings200Ok struct {
	// from_id integer
	FromId int32 `json:"from_id"`
	// from_type string
	FromType string `json:"from_type"`
	// standing number
	Standing float32 `json:"standing"`
}
