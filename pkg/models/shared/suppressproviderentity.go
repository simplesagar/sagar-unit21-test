// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SuppressProviderEntity struct {
	// Array of entity ID strings
	//
	ProviderEntityIds []string `json:"provider_entity_ids"`
	// Enable or disable the silencing of the entity. Use `true` to turn on and `false` to turn off.
	//
	Suppress bool `json:"suppress"`
	// Default: `false`. If `true`,  the response is returned in the request but may take slightly longer to complete.
	SynchronousResponse *bool `json:"synchronous_response,omitempty"`
}
