// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// DigitalDataArrayProperties - Associated digital properties - IP, device, browser, client info etc.
type DigitalDataArrayProperties struct {
	// List of unique addresses.
	ClientFingerprints []string `json:"client_fingerprints,omitempty"`
	// List of IP addresses. MUST be in either IPv4 or IPv6 format.
	IPAddresses []string `json:"ip_addresses,omitempty"`
}
