/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.25
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Summary of victory points gained for the given faction
type GetFwStatsVictoryPoints struct {
	// Last week's victory points gained
	LastWeek int32 `json:"last_week"`
	// Total victory points gained since faction warfare began
	Total int32 `json:"total"`
	// Yesterday's victory points gained
	Yesterday int32 `json:"yesterday"`
}
