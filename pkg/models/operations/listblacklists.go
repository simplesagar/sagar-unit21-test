// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

// ListBlacklists200ApplicationJSON - OK
type ListBlacklists200ApplicationJSON struct {
	// Number of objects in the response
	ResponseCount *int64 `json:"response_count,omitempty"`
	// Total number of objects in the platform.
	TotalCount *int64 `json:"total_count,omitempty"`
}

type ListBlacklistsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ListBlacklists200ApplicationJSONObject *ListBlacklists200ApplicationJSON
}
