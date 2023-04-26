// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExportAlertsRequestBodyFiltersAlertTypesEnum string

const (
	ExportAlertsRequestBodyFiltersAlertTypesEnumChainalysis ExportAlertsRequestBodyFiltersAlertTypesEnum = "CHAINALYSIS"
	ExportAlertsRequestBodyFiltersAlertTypesEnumTm          ExportAlertsRequestBodyFiltersAlertTypesEnum = "TM"
	ExportAlertsRequestBodyFiltersAlertTypesEnumKyc         ExportAlertsRequestBodyFiltersAlertTypesEnum = "KYC"
	ExportAlertsRequestBodyFiltersAlertTypesEnumManual      ExportAlertsRequestBodyFiltersAlertTypesEnum = "MANUAL"
	ExportAlertsRequestBodyFiltersAlertTypesEnumCar         ExportAlertsRequestBodyFiltersAlertTypesEnum = "CAR"
)

func (e ExportAlertsRequestBodyFiltersAlertTypesEnum) ToPointer() *ExportAlertsRequestBodyFiltersAlertTypesEnum {
	return &e
}

func (e *ExportAlertsRequestBodyFiltersAlertTypesEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "CHAINALYSIS":
		fallthrough
	case "TM":
		fallthrough
	case "KYC":
		fallthrough
	case "MANUAL":
		fallthrough
	case "CAR":
		*e = ExportAlertsRequestBodyFiltersAlertTypesEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ExportAlertsRequestBodyFiltersAlertTypesEnum: %s", s)
	}
}

type ExportAlertsRequestBodyFiltersSourcesEnum string

const (
	ExportAlertsRequestBodyFiltersSourcesEnumInternal ExportAlertsRequestBodyFiltersSourcesEnum = "INTERNAL"
	ExportAlertsRequestBodyFiltersSourcesEnumExternal ExportAlertsRequestBodyFiltersSourcesEnum = "EXTERNAL"
)

func (e ExportAlertsRequestBodyFiltersSourcesEnum) ToPointer() *ExportAlertsRequestBodyFiltersSourcesEnum {
	return &e
}

func (e *ExportAlertsRequestBodyFiltersSourcesEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "INTERNAL":
		fallthrough
	case "EXTERNAL":
		*e = ExportAlertsRequestBodyFiltersSourcesEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ExportAlertsRequestBodyFiltersSourcesEnum: %s", s)
	}
}

type ExportAlertsRequestBodyFiltersStatusesEnum string

const (
	ExportAlertsRequestBodyFiltersStatusesEnumOpen   ExportAlertsRequestBodyFiltersStatusesEnum = "OPEN"
	ExportAlertsRequestBodyFiltersStatusesEnumClosed ExportAlertsRequestBodyFiltersStatusesEnum = "CLOSED"
)

func (e ExportAlertsRequestBodyFiltersStatusesEnum) ToPointer() *ExportAlertsRequestBodyFiltersStatusesEnum {
	return &e
}

func (e *ExportAlertsRequestBodyFiltersStatusesEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "OPEN":
		fallthrough
	case "CLOSED":
		*e = ExportAlertsRequestBodyFiltersStatusesEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for ExportAlertsRequestBodyFiltersStatusesEnum: %s", s)
	}
}

// ExportAlertsRequestBodyFilters - Filter to narrow down alerts in export
type ExportAlertsRequestBodyFilters struct {
	// Numerical IDs of the agents.
	AgentIds []int64 `json:"agent_ids,omitempty"`
	// Numerical ID of the alerts.
	AlertIds []int64 `json:"alert_ids,omitempty"`
	// Numerical IDs of the alert queues.
	AlertQueueIds []int64 `json:"alert_queue_ids,omitempty"`
	// Sources for the alerts.
	AlertTypes []ExportAlertsRequestBodyFiltersAlertTypesEnum `json:"alert_types,omitempty"`
	// Numerical ID of the alerting org.
	AlertingOrgIds []int64 `json:"alerting_org_ids,omitempty"`
	// Disposition name.
	Disposition *string `json:"disposition,omitempty"`
	// Disposition end date.
	DispositionEndDate *string `json:"disposition_end_date,omitempty"`
	// Disposition start date.
	DispositionStartDate *string `json:"disposition_start_date,omitempty"`
	// Alert creation end date.
	EndDate *string `json:"end_date,omitempty"`
	// Numerical IDs of the entities.
	EntityIds []int64 `json:"entity_ids,omitempty"`
	// Maximum amount in the alert.
	MaximumAmount *int64 `json:"maximum_amount,omitempty"`
	// Minimum amount in the alert.
	MinimumAmount *int64 `json:"minimum_amount,omitempty"`
	// ID or title of the alert.
	Phrase *string `json:"phrase,omitempty"`
	// Numerical IDs of the rules.
	RuleIds []int64 `json:"rule_ids,omitempty"`
	// Sources for the alerts.
	Sources []ExportAlertsRequestBodyFiltersSourcesEnum `json:"sources,omitempty"`
	// Alert creation start date.
	StartDate *string `json:"start_date,omitempty"`
	// Status for the alert.
	Status *string `json:"status,omitempty"`
	// Statuses for the alerts.
	Statuses []ExportAlertsRequestBodyFiltersStatusesEnum `json:"statuses,omitempty"`
	// Numerical IDs of the agents.
	SubdispositionIds []int64 `json:"subdisposition_ids,omitempty"`
	// Numerical IDs of the subdisposition.
	SubdispositionOptionIds []int64 `json:"subdisposition_option_ids,omitempty"`
	// Numerical IDs of the tags.
	TagIds []int64 `json:"tag_ids,omitempty"`
	// Numerical IDs of the teams.
	TeamIds []int64 `json:"team_ids,omitempty"`
}

// ExportAlertsRequestBody - To filter your response.
type ExportAlertsRequestBody struct {
	// Array of the unique identifiers of the alert IDs.
	AlertIds []int64 `json:"alert_ids,omitempty"`
	// Filter to narrow down alerts in export
	Filters *ExportAlertsRequestBodyFilters `json:"filters,omitempty"`
}

type ExportAlertsResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
