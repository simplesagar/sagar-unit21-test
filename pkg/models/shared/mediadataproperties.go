// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// MediaDataProperties - Media data. See `/link-media` endpoint
type MediaDataProperties struct {
	// Epoch
	CreatedAt *string `json:"created_at,omitempty"`
	// Base64 encoded media file
	Media *string `json:"media,omitempty"`
	// Type of media. The dropdown contains Unit21-recognized types.
	//
	// Though you can provide any string, Unit21 processes and display recognized `media_type` values with custom formatting. Unit21 can also use recognized media object types for dedicated requests, like document verification.
	//
	MediaType *MediaType `json:"media_type,omitempty"`
	// Media file name
	Name *string `json:"name,omitempty"`
}
