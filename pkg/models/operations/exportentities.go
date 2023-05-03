// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unit/pkg/models/shared"
)

type ExportEntitiesRequestBodyFiltersStatusesEnum string

const (
	ExportEntitiesRequestBodyFiltersStatusesEnumActive   ExportEntitiesRequestBodyFiltersStatusesEnum = "active"
	ExportEntitiesRequestBodyFiltersStatusesEnumInactive ExportEntitiesRequestBodyFiltersStatusesEnum = "inactive"
)

func (e ExportEntitiesRequestBodyFiltersStatusesEnum) ToPointer() *ExportEntitiesRequestBodyFiltersStatusesEnum {
	return &e
}

func (e *ExportEntitiesRequestBodyFiltersStatusesEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "active":
		fallthrough
	case "inactive":
		*e = ExportEntitiesRequestBodyFiltersStatusesEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ExportEntitiesRequestBodyFiltersStatusesEnum: %v", v)
	}
}

// ExportEntitiesRequestBodyFilters - Filter to narrow down entities in export
type ExportEntitiesRequestBodyFilters struct {
	// Entities that belong to this Child Org (Numerical ID).
	ChildOrgIds []int64 `json:"child_org_ids,omitempty"`
	// Entity creation date end.
	EndDate *string `json:"end_date,omitempty"`
	// Numerical IDs of the entities.
	EntityIds []int64 `json:"entity_ids,omitempty"`
	// Describes a user such as `employee` or `business`
	EntityType *string `json:"entity_type,omitempty"`
	// Type of entity.
	InternalEntityType []string `json:"internal_entity_type,omitempty"`
	// Numerical IDs of the rules.
	RuleIds []int64 `json:"rule_ids,omitempty"`
	// Entity creation date start.
	StartDate *string `json:"start_date,omitempty"`
	// Status of the entities.
	Status *string `json:"status,omitempty"`
	// Statuses for the entities.
	Statuses []ExportEntitiesRequestBodyFiltersStatusesEnum `json:"statuses,omitempty"`
	// Numerical IDs of the tags.
	TagIds []int64 `json:"tag_ids,omitempty"`
}

type ExportEntitiesRequestBody struct {
	// Array of the unique identifiers of the entity IDs.
	EntityIds []int64 `json:"entity_ids,omitempty"`
	// Filter to narrow down entities in export
	Filters *ExportEntitiesRequestBodyFilters `json:"filters,omitempty"`
}

// ExportEntitiesMessageGeneralResponse - OK
type ExportEntitiesMessageGeneralResponse struct {
	// Error message
	ErrorCode *string `json:"error_code,omitempty"`
	// Detailed message
	Message *string `json:"message,omitempty"`
}

type ExportEntitiesResponse struct {
	ContentType string
	// OK
	MessageGeneralResponse *ExportEntitiesMessageGeneralResponse
	StatusCode             int
	RawResponse            *http.Response
	// OK
	MessageResponse *shared.MessageResponse
}