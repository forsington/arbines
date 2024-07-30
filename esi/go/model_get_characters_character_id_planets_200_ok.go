/*
 * EVE Swagger Interface
 *
 * An OpenAPI for EVE Online
 *
 * API version: 1.25
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger
import (
	"time"
)

// 200 ok object
type GetCharactersCharacterIdPlanets200Ok struct {
	// last_update string
	LastUpdate time.Time `json:"last_update"`
	// num_pins integer
	NumPins int32 `json:"num_pins"`
	// owner_id integer
	OwnerId int32 `json:"owner_id"`
	// planet_id integer
	PlanetId int32 `json:"planet_id"`
	// planet_type string
	PlanetType string `json:"planet_type"`
	// solar_system_id integer
	SolarSystemId int32 `json:"solar_system_id"`
	// upgrade_level integer
	UpgradeLevel int32 `json:"upgrade_level"`
}
