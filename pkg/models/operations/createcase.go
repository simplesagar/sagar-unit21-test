// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unit/pkg/models/shared"
)

// CreateCaseCaseData - Top-level case data
type CreateCaseCaseData struct {
	// Array of the unique identifiers of the alert IDs.
	AlertIds []int64 `json:"alert_ids,omitempty"`
	// Unique identifier of the case on the customer's platform
	CaseID string `json:"case_id"`
	// Description of the case
	Description *string `json:"description,omitempty"`
	// Labels that describe the outcome of an alert or case investigation. More info in [this knowledge base article about dispositions](https://docs.unit21.ai/docs/concept-dispositions).
	Disposition *string `json:"disposition,omitempty"`
	// Free form text documenting reasoning and investigation notes
	DispositionNotes *string `json:"disposition_notes,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	EndDate *int64 `json:"end_date,omitempty"`
	// Array of the unique identifiers of the entity IDs.
	EntityIds []int64 `json:"entity_ids,omitempty"`
	// Array of the unique identifiers of the event IDs.
	EventIds []int64 `json:"event_ids,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	StartDate int64 `json:"start_date"`
	// Investigation status, either `OPEN` or `ClOSED`
	Status *shared.InvestigationStatusEnum `json:"status,omitempty"`
	// List of string tags, in the format `keyString:valueString` (note that the Key strings are NOT enclosed in `"`)
	Tags []string `json:"tags,omitempty"`
	// Title of the case
	Title string `json:"title"`
}

type CreateCaseResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	CreateCaseResponse *shared.CreateCaseResponse
}