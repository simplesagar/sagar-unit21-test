// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unit/pkg/models/shared"
)

// CreateAlertRulesAlertTypeEnum - Either transaction monitoring, `tm`, or know-your-customer `kyc`. Default is `tm`
type CreateAlertRulesAlertTypeEnum string

const (
	CreateAlertRulesAlertTypeEnumTm          CreateAlertRulesAlertTypeEnum = "tm"
	CreateAlertRulesAlertTypeEnumKyc         CreateAlertRulesAlertTypeEnum = "kyc"
	CreateAlertRulesAlertTypeEnumChainalysis CreateAlertRulesAlertTypeEnum = "chainalysis"
)

func (e CreateAlertRulesAlertTypeEnum) ToPointer() *CreateAlertRulesAlertTypeEnum {
	return &e
}

func (e *CreateAlertRulesAlertTypeEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "tm":
		fallthrough
	case "kyc":
		fallthrough
	case "chainalysis":
		*e = CreateAlertRulesAlertTypeEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAlertRulesAlertTypeEnum: %s", s)
	}
}

type CreateAlertRulesEntities struct {
	// Unique identifier of the entity. Entity IDs must be unique and only comprise of the characters -_:.@a-zA-Z0-9!#$%&*+/=?^`{'
	EntityID string `json:"entity_id"`
}

type CreateAlertRulesEvents struct {
	EventID string `json:"event_id"`
	// `transaction` for monetary flows, `action` for other state changes, like new logins.
	//
	EventType shared.EventTypeEnum `json:"event_type"`
}

type CreateAlertRules struct {
	// Unique identifier of the alert on the customer's platform.
	AlertID string `json:"alert_id"`
	// Either transaction monitoring, `tm`, or know-your-customer `kyc`. Default is `tm`
	AlertType CreateAlertRulesAlertTypeEnum `json:"alert_type"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	CreatedAt int64 `json:"created_at"`
	// Description of the alert
	Description *string `json:"description,omitempty"`
	// Labels that describe the outcome of an alert or case investigation. More info in [this knowledge base article about dispositions](https://docs.unit21.ai/docs/concept-dispositions).
	Disposition *string `json:"disposition,omitempty"`
	// Free form text documenting reasoning and investigation notes
	DispositionNotes *string `json:"disposition_notes,omitempty"`
	// List of the unique identifiers of the entity IDs.
	Entities []CreateAlertRulesEntities `json:"entities,omitempty"`
	// Array of event objects, consisting of `event_id` and `event_type`
	Events []CreateAlertRulesEvents `json:"events,omitempty"`
	// Array of `instrument_id` strings
	Instruments []string `json:"instruments,omitempty"`
	// Array of "rule_id" strings
	Rules []string `json:"rules,omitempty"`
	// Investigation status, either `OPEN` or `ClOSED`
	Status shared.InvestigationStatusEnum `json:"status"`
	// List of string tags, in the format `keyString:valueString` (note that the Key strings are NOT enclosed in `"`)
	Tags []string `json:"tags,omitempty"`
	// Title of the alert
	Title string `json:"title"`
	// Integer value greater than or equal to 1. Used when `alert_type` is `kyc`.
	VerificationResultID *int64 `json:"verification_result_id,omitempty"`
}

type CreateAlertResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	CreateAlertResponse *shared.CreateAlertResponse
}
