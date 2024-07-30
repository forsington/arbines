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
type GetMarketsRegionIdHistory200Ok struct {
	// average number
	Average float64 `json:"average"`
	// The date of this historical statistic entry
	Date string `json:"date"`
	// highest number
	Highest float64 `json:"highest"`
	// lowest number
	Lowest float64 `json:"lowest"`
	// Total number of orders happened that day
	OrderCount int64 `json:"order_count"`
	// Total
	Volume int64 `json:"volume"`
}
