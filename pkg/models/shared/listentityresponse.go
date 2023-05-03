// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// ListEntityResponse - OK
type ListEntityResponse struct {
	// List of entity objects.
	Entities []EntityList `json:"entities,omitempty"`
	// Number of objects in the paginated response
	ResponseCount *int64 `json:"response_count,omitempty"`
	// Total number of objects in the platform.
	TotalCount *int64 `json:"total_count,omitempty"`
}