// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"unit/pkg/models/shared"
)

// ListCasesRequestBodyOptions - Options for the data included in the returned case. Removing unneeded options can improve response speed.
type ListCasesRequestBodyOptions struct {
	// If `true`, the response includes list of all actions taken on the case, including disposition changes, status changes, reassignments, etcetera.
	IncludeActions *bool `json:"include_actions,omitempty"`
	// If `true`, the response includes all rule IDs and SAR filing IDs associated with the case under `rules` and `descendant_sar_filing`.
	IncludeAssociations *bool `json:"include_associations,omitempty"`
}

// ListCasesRequestBody - To filter your response to a subset of alerts, use these fields.
type ListCasesRequestBody struct {
	// Only objects associated with the listed entities' `unit21_id` values.
	AssociatedEntities []int64 `json:"associated_entities,omitempty"`
	// Only objects associated with the listed events' `unit21_id` values.
	AssociatedEvents []int64 `json:"associated_events,omitempty"`
	// Only objects associated with the listed instruments' `unit21_id` values.
	AssociatedInstruments []int64 `json:"associated_instruments,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	CreatedAfter *int64 `json:"created_after,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	CreatedBefore *int64 `json:"created_before,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	DispositionedAfter *int64 `json:"dispositioned_after,omitempty"`
	// Date in seconds since 1 Jan 1970 00:00:00 UTC (i.e. in [Unix time](https://en.wikipedia.org/wiki/Unix_time)).
	DispositionedBefore *int64 `json:"dispositioned_before,omitempty"`
	// List of agent emails. Returns only objects whose dispositions were most recently changed by listed agents.
	DispositionedBy []string `json:"dispositioned_by,omitempty"`
	// Set of `disposition` values to include.
	Dispositions []string `json:"dispositions,omitempty"`
	// Number of objects per page.
	Limit *int64 `json:"limit,omitempty"`
	// Pagination offset. A value of 1 returns a response beginning with the first record. The offset is relative to the number of pages (not the total count of objects)
	Offset *int64 `json:"offset,omitempty"`
	// Options for the data included in the returned case. Removing unneeded options can improve response speed.
	Options *ListCasesRequestBodyOptions `json:"options,omitempty"`
	// Only objects associated with the listed rules' `unit21_id` values.
	Rules []int64 `json:"rules,omitempty"`
	// Only `INTERNAL`, only `EXTERNAL`, or both.
	Sources []shared.SourceArrayEnum `json:"sources,omitempty"`
	// Only objects from the listed set of `status` values. E.g. only `OPEN`
	Statuses []shared.InvestigationStatusEnum `json:"statuses,omitempty"`
}

type ListCasesResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	ListResponse *shared.ListResponse
}
