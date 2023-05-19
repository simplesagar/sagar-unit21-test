// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"unit/pkg/models/operations"
	"unit/pkg/models/shared"
	"unit/pkg/utils"
)

// alertsAPI - Alerts have two origins. Typically, alerts are generated whenever a Unit21 detection tool (like a rule) flags an object, like an entity. However, your organization can also send alerts generated from your in-house detection systems. The `/alerts` endpoint can create, list, and update alerts.
type alertsAPI struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newAlertsAPI(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *alertsAPI {
	return &alertsAPI{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateAlert - Create alerts
// Creates a new alert, sending alert data in the request body.
// To create an Alert, you MUST include the following fields: `alert_id`, `alert_type`, `created_at`, `title`, and `status`. The other top-level fields are optional.
//
// If we receive a request to create an alert for an `alert_id` that already exists in our system,  we will respond with a **409 error code** indicating that this alert cannot be created/updated. You must use the `/alert/update` endpoint to update an alert.
//
// You can add the following objects to an alert:
//
//	| Field                    | Type     | Description                                                                                                           |
//	|--------------------------|----------|-----------------------------------------------------------------------------------------------------------------------|
//	| `rules`	                 | String[] | Unique identifier of the rules/triggers/scenarios that triggered this alert                                           |
//	| `events`	               | Object[] | Transactions affiliated with the alert. Each object must consist of a `event_id` and `event_type` field               |
//	| `entities`	             | Object[] | Users or businesses affiliated with the alert. Each object must consist of an `entity_id` and `entity_type` field     |
//	| `instruments`	           | String[] | Unique identifiers of any instruments affiliated with the alert                                                       |
//
// Please note that if `verification_result_id` is included, it will link the entity that is associated  with the verification result with the alert.
//
// Updates to an alert's `alert_id` are not allowed.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
//   - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
//   - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)
//
// The response will consist of the following fields:
//
//	| Field                    | Type     | Description                                             |
//	|--------------------------|----------|---------------------------------------------------------|
//	| `alert_id`	             | String   | Unique identifier of the alert on your platform         |
//	| `unit21_id`	             | String   | Internal ID of the alert within Unit21's system         |
//	| `previously_existed`	   | Boolean  | If alert (with the same `alert_id`) already exists      |
func (s *alertsAPI) CreateAlert(ctx context.Context, request operations.CreateAlertRules) (*operations.CreateAlertResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/alerts/create"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}
	if bodyReader == nil {
		return nil, fmt.Errorf("request body is required")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.CreateAlertResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.CreateAlertResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateAlertResponse = out
		}
	}

	return res, nil
}

// ExportAlerts - Bulk export alerts
// Initiates an email and dashboard export of alerts. The export will be as a CSV file.
//
// Either the agent `ID` or `email` is required to begin the export.
//
// Either the `filters` or the list of `alert IDs` are required for the export.
//
// Custom data filters are not supported for bulk exports at this time.
func (s *alertsAPI) ExportAlerts(ctx context.Context, request operations.ExportAlertsRequestBody) (*operations.ExportAlertsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/alerts/bulk-export"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ExportAlertsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetAlertByUnit21ID - Get an alert
// Returns all data objects belonging to a single alert.
//
// This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the entity is first created.
func (s *alertsAPI) GetAlertByUnit21ID(ctx context.Context, request operations.GetAlertByUnit21IDRequest) (*operations.GetAlertByUnit21IDResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/alerts/{unit21_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetAlertByUnit21IDResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.GetAlertByUnit21ID200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// LinkMediaToAlert - Add media to an alert
// Adds rich media objects (images, videos, etc.) to an existing alert.
//
// This endpoint is useful for sending in rich media such as profile pictures, ID card scans, official documents etc.  that you want available for investigative and verification purposes.
//
// Supported file types are: txt, pdf, video (mp4, mov, wmv, avi, mkv), images (png, jpg, tiff, gif, raw, eps).
//
// The payload to this endpoint can either be a **form-data** or a **base64** encoded media file via the requests JSON body.
//
// **Form-data** sent to this endpoint must use the key `media_key` and the `value` as the media file.  If you wish to provide optional information, use the `media_key` and provide stringified JSON data as the value.  There are no required fields in each media file's supplementary form data. However, if a recognized `media_type` value is provided,  the Unit21 system will be able to use the media object for purposes such as document verification.
//
// ```
//
//	--form 'document_front=@/src/103031/images/document_front.jpg' \
//	--form 'document_front={"media_type": "IMAGE_ID_CARD_FRONT", "source": "passport_app", "timestamp": 1572673229}'
//
// ```
//
// **Base64** encoded media objects must follow the format:
//
// ```json
//
//	{
//	  "media": "iVBORw0KGgoAAAANSUhEUgAAAQMAAADCCAYAAABNEqduAAAgAElEQVR4Aey9CbgmV1Xv...",
//	  "name": "Drivers_License.png",
//	  "media_type": "IMAGE_DRIVERS_LICENSE_FRONT"
//	}
//
// ```
//
// `media` and `name` are the only required fields for each media object. The `name“ must include the file extension such a `File.pdf`. Supplementary form data is sent through the optional `custom_data` object.
//
// Recognized values of `media_type` are:
//
//	| media_type                  |
//	|-----------------------------|
//	| IMAGE_PROFILE_PICTURE       |
//	| IMAGE_DRIVERS_LICENSE_FRONT |
//	| IMAGE_DRIVERS_LICENSE_BACK  |
//	| IMAGE_PASSPORT_FRONT        |
//	| IMAGE_ID_CARD_FRONT         |
//	| IMAGE_ID_CARD_BACK          |
//	| IMAGE_FACE_IMAGE            |
func (s *alertsAPI) LinkMediaToAlert(ctx context.Context, request operations.LinkMediaToAlertRequest) (*operations.LinkMediaToAlertResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/alerts/{unit21_id}/link-media", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.LinkMediaToAlertResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// ListAlerts - List alerts
// Returns an array of top-level information about alerts in your environment.
//
// Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
// * `limit`  indicates how many objects the request returns (the page maximum is 50)
// * `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record.
//
// To narrow down your alert search, we provide filter parameters to this endpoint. Note that all list inputs function as an "or" filter, as in any one of the values must match the selected alert(s):
//
//	| Field                   | Type        | Description                                                                                                       |
//	| ----------------------- | ----------- | ----------------------------------------------------------------------------------------------------------------- |
//	| `case_id`               | Numeric     | Only alerts with the associated case ID will be shown.                                                            |
//	| `types`                 | String[]    | Must be list of alert types: `tm`, `kyc`, `chainalysis`                                                           |
//	| `created_after`         | Numeric     | Alerts created on or after this unix timestamp                                                                    |
//	| `created_before`        | Numeric     | Alerts created before this unix timestamp                                                                         |
//	| `dispositions`          | String[]    | List of alert disposition states (defined on an integration basis)                                                |
//	| `dispositioned_after`   | Numeric     | Alerts with a disposition most recently updated after this unix timestamp                                         |
//	| `dispositioned_before`  | Numeric     | Alerts with a disposition most recently updated before this unix timestamp                                        |
//	| `dispositioned_by`      | String[]    | List of agent emails. Returns alerts with a disposition most recently changed by agents in the list               |
//	| `rules`                 | Numeric[]   | List of Unit21 rule ids that are associated with the alert                                                        |
//	| `associated_entities`   | Numeric[]   | List of Unit21 entity ids associated with this alert                                                              |
//	| `associated_events`     | Numeric[]   | List of Unit21 event ids associated with this alert                                                               |
//	| `associated_instruments`| Numeric[]   | List of Unit21 instrument ids associated with this alert                                                          |
//	| `sources`               | String[]    | Must be list of alert sources: `INTERNAL`, `EXTERNAL`                                                             |
//	| `statuses`              | String[]    | Must be list of alert statuses: `OPEN`, `CLOSED`                                                                  |
//	| `tag_filters`           | String[]    | List of string tags (`key:value`) or keys to associate this alert with (e.g. `alert_type:high_velocity` or `alert_type`). If only the key is provided, we will match against all tags with that key        |
//	| `limit`                 | Numeric     | A limit on the number of objects to be returned. Limit can range between 1 and 50, and the default is 10          |
//	| `offset`                | Numeric     | The offset for pagination. Default is 1                                                                           |
//	| `options`               | Object      | Options for the data included in the returned alerts. Removing unneeded options can improve response speed        |
//
// The `total_count` field contains the total number of alerts where the  `response_count` field contains the number of alerts included in the response.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
func (s *alertsAPI) ListAlerts(ctx context.Context, request operations.ListAlertsRequestBody) (*operations.ListAlertsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/alerts/list"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListAlertsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.ListResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListResponse = out
		}
	}

	return res, nil
}

// UpdateAlert - Update alert
// Updates an alert's information using the `unit21_id`. ONLY EXTERNAL ALERTS CAN BE UPDATED!
//
// Updating an alert has no required fields. You MAY send any subset of the fields that the `alerts/create` endpoint accepts.
//
// This endpoint requires the `unit21_id` which is a unique ID created by Unit21 when the entity is first created.
//
// Please note that if `verification_result_id` is included, it will link the entity that is associated with the  verification result with the alert.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/reference/endpoint-options)
//   - [Custom data](https://docs.unit21.ai/reference/best-practices-for-custom-data)
//   - [Batch uploads](https://docs.unit21.ai/reference/batch-request-examples)
//   - [Modifying tags](https://docs.unit21.ai/reference/modifying-tags)
//
// The response will consist of the following fields:
//
//	| Field                    | Type     | Description                                             |
//	|--------------------------|----------|---------------------------------------------------------|
//	| `alert_id`	             | String   | Unique identifier of the alert on your platform         |
//	| `unit21_id`	             | String   | Internal ID of the alert within Unit21's system         |
//	| `previously_existed`	   | Boolean  | If alert (with the same `alert_id`) already exists      |
func (s *alertsAPI) UpdateAlert(ctx context.Context, request operations.UpdateAlertRequest) (*operations.UpdateAlertResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/alerts/{unit21_id}/update", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", fmt.Sprintf("speakeasy-sdk/%s %s %s", s.language, s.sdkVersion, s.genVersion))

	req.Header.Set("Content-Type", reqContentType)

	client := s.securityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}
	defer httpRes.Body.Close()

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.UpdateAlertResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out interface{}
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.UpdateAlert200ApplicationJSONAny = out
		}
	}

	return res, nil
}
