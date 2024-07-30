/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.25
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// ally object
type GetWarsWarIdAlly struct {
	// Alliance ID if and only if this ally is an alliance
	AllianceId int32 `json:"alliance_id,omitempty"`
	// Corporation ID if and only if this ally is a corporation
	CorporationId int32 `json:"corporation_id,omitempty"`
}
