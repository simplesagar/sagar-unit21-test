// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// LinkVerificationResultVerificationTypeEnum - Type of verification, in one of the categories that the Unit21 system recognizes
type LinkVerificationResultVerificationTypeEnum string

const (
	LinkVerificationResultVerificationTypeEnumIDVerification        LinkVerificationResultVerificationTypeEnum = "ID_VERIFICATION"
	LinkVerificationResultVerificationTypeEnumDocVerification       LinkVerificationResultVerificationTypeEnum = "DOC_VERIFICATION"
	LinkVerificationResultVerificationTypeEnumBusinessVerification  LinkVerificationResultVerificationTypeEnum = "BUSINESS_VERIFICATION"
	LinkVerificationResultVerificationTypeEnumWatchlistScreening    LinkVerificationResultVerificationTypeEnum = "WATCHLIST_SCREENING"
	LinkVerificationResultVerificationTypeEnumAdverseMediaScreening LinkVerificationResultVerificationTypeEnum = "ADVERSE_MEDIA_SCREENING"
	LinkVerificationResultVerificationTypeEnumCryptoForensics       LinkVerificationResultVerificationTypeEnum = "CRYPTO_FORENSICS"
)

func (e LinkVerificationResultVerificationTypeEnum) ToPointer() *LinkVerificationResultVerificationTypeEnum {
	return &e
}

func (e *LinkVerificationResultVerificationTypeEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ID_VERIFICATION":
		fallthrough
	case "DOC_VERIFICATION":
		fallthrough
	case "BUSINESS_VERIFICATION":
		fallthrough
	case "WATCHLIST_SCREENING":
		fallthrough
	case "ADVERSE_MEDIA_SCREENING":
		fallthrough
	case "CRYPTO_FORENSICS":
		*e = LinkVerificationResultVerificationTypeEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LinkVerificationResultVerificationTypeEnum: %v", v)
	}
}

// LinkVerificationResult - Fields to link external verification results to entity in the Unit21 system
type LinkVerificationResult struct {
	// JSON-formatted response from verification provider
	Content map[string]interface{} `json:"content"`
	// name of KYC provider
	ProviderName string `json:"provider_name"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	VerificationTimestamp *int64 `json:"verification_timestamp,omitempty"`
	// Type of verification, in one of the categories that the Unit21 system recognizes
	VerificationType LinkVerificationResultVerificationTypeEnum `json:"verification_type"`
}
