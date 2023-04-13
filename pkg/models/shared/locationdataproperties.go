// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// LocationDataProperties - Address/location data
type LocationDataProperties struct {
	// Building number of the primary residence (max 24 characters)
	BuildingNumber *string `json:"building_number,omitempty"`
	// Associated city (max 128 characters)
	City *string `json:"city,omitempty"`
	// Country of primary residence such as "USA", "United States", "US", "The United States of America"
	Country *string `json:"country,omitempty"`
	// Associated ZIP code or postal code. For US addresses, can be either 5-digit ZIP Code (99999) or ZIP+4 Code (99999-9999) formats
	PostalCode *string `json:"postal_code,omitempty"`
	// Any string representing a state such as "CA" or "California"
	State *string `json:"state,omitempty"`
	// Street name of primary residence (max 128 characters)
	StreetName *string `json:"street_name,omitempty"`
	// A string field indicating the type of address e.g. `SHIPPING`, `BILLING` (max 24 characters)
	Type *string `json:"type,omitempty"`
	// Flat/unit/apartment number of the location associated with the event (max 24 characters)
	UnitNumber *string `json:"unit_number,omitempty"`
	// Date when location was verified, in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time))".
	VerifiedOn *int64 `json:"verified_on,omitempty"`
}
