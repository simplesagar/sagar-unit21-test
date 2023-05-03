// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// UpdateEventEventOptionsOptionsLinkedEntityEnum - Possible values are `sender`, `receiver`, and `both`. Defaults to `both`. If `link_digital_data_to_entity` is flagged on transaction events, this specifies which entities to associate the `digital_data` to. If there is no `digital_data` or entities, no exception is thrown.
type UpdateEventEventOptionsOptionsLinkedEntityEnum string

const (
	UpdateEventEventOptionsOptionsLinkedEntityEnumSender   UpdateEventEventOptionsOptionsLinkedEntityEnum = "sender"
	UpdateEventEventOptionsOptionsLinkedEntityEnumReceiver UpdateEventEventOptionsOptionsLinkedEntityEnum = "receiver"
	UpdateEventEventOptionsOptionsLinkedEntityEnumBoth     UpdateEventEventOptionsOptionsLinkedEntityEnum = "both"
)

func (e UpdateEventEventOptionsOptionsLinkedEntityEnum) ToPointer() *UpdateEventEventOptionsOptionsLinkedEntityEnum {
	return &e
}

func (e *UpdateEventEventOptionsOptionsLinkedEntityEnum) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "sender":
		fallthrough
	case "receiver":
		fallthrough
	case "both":
		*e = UpdateEventEventOptionsOptionsLinkedEntityEnum(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateEventEventOptionsOptionsLinkedEntityEnum: %v", v)
	}
}

type UpdateEventEventOptionsOptions struct {
	// Whether or not to link the included `digital_data` with the related entities. Includes geoip information if resolve_geoip is enabled as well. On action events, defaults to `true`
	LinkDigitalDataToEntity *bool `json:"link_digital_data_to_entity,omitempty"`
	// Possible values are `sender`, `receiver`, and `both`. Defaults to `both`. If `link_digital_data_to_entity` is flagged on transaction events, this specifies which entities to associate the `digital_data` to. If there is no `digital_data` or entities, no exception is thrown.
	LinkedEntity *UpdateEventEventOptionsOptionsLinkedEntityEnum `json:"linked_entity,omitempty"`
	// Only relevant for updates/upserts, ignored otherwise. See [custom data merge strategy](doc:how-data-merges-on-updates#custom-data-merge-strategy) for more details. **Default**: `false`
	MergeCustomData *bool `json:"merge_custom_data,omitempty"`
	// Whether or not to monitor this event (defaults to `true`). Typically used to signal Unit21 to not flag such events or include them in calculations i.e. to prevent double counting, or to ignore applying monitoring to unimportant events that you still want to associate with users
	Monitor *bool `json:"monitor,omitempty"`
	// If `false`, does not resolve the geographic location from the provided IP. If `true` and `digital_data.ip_addresses` is empty, Unit21 ignores the field. **Default**: `true`
	ResolveGeoip *bool `json:"resolve_geoip,omitempty"`
	// If POST request includes an object that already exists when  `upsert_on_conflict` is `false`, API returns a 409 error code and the object is not overwritten. **Default**: `true`
	UpsertOnConflict *bool `json:"upsert_on_conflict,omitempty"`
}

type UpdateEventEventOptions struct {
	Options *UpdateEventEventOptionsOptions `json:"options,omitempty"`
}

type UpdateEventRequest struct {
	RequestBody *UpdateEventEventOptions `request:"mediaType=application/json"`
	// Unique identifier of the event on your platform
	EventID string `pathParam:"style=simple,explode=false,name=event_id"`
	// Name of organization in your environment
	OrgName string `pathParam:"style=simple,explode=false,name=org_name"`
}

type UpdateEventResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
