// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ExportRulesRequestBodyFiltersStatusesEnum string

const (
	ExportRulesRequestBodyFiltersStatusesEnumActive     ExportRulesRequestBodyFiltersStatusesEnum = "ACTIVE"
	ExportRulesRequestBodyFiltersStatusesEnumInactive   ExportRulesRequestBodyFiltersStatusesEnum = "INACTIVE"
	ExportRulesRequestBodyFiltersStatusesEnumValidation ExportRulesRequestBodyFiltersStatusesEnum = "VALIDATION"
)

func (e ExportRulesRequestBodyFiltersStatusesEnum) ToPointer() *ExportRulesRequestBodyFiltersStatusesEnum {
	return &e
}

func (e *ExportRulesRequestBodyFiltersStatusesEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE":
		fallthrough
	case "INACTIVE":
		fallthrough
	case "VALIDATION":
		*e = ExportRulesRequestBodyFiltersStatusesEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExportRulesRequestBodyFiltersStatusesEnum: %v", v)
	}
}

// ExportRulesRequestBodyFilters - Filter to narrow down rules in export.
type ExportRulesRequestBodyFilters struct {
	// Numerical IDs of the agents.
	AgentIds []int64 `json:"agent_ids,omitempty"`
	// Rule Creation end date.
	EndDate *string `json:"end_date,omitempty"`
	// Numerical IDs of the rules.
	RuleIds []int64 `json:"rule_ids,omitempty"`
	// Rule creation start date.
	StartDate *string `json:"start_date,omitempty"`
	// Status for the rule.
	Status *string `json:"status,omitempty"`
	// Status for the rule.
	Statuses []ExportRulesRequestBodyFiltersStatusesEnum `json:"statuses,omitempty"`
	// Numerical IDs of the tags.
	TagIds []int64 `json:"tag_ids,omitempty"`
}

type ExportRulesRequestBody struct {
	// Filter to narrow down rules in export.
	Filters *ExportRulesRequestBodyFilters `json:"filters,omitempty"`
	// Array of the unique identifiers of the rule IDs.
	RuleIds []int64 `json:"rule_ids,omitempty"`
}

type ExportRulesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
