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

// devicesAPI - Devices representing any computer or physical device used to execute an event. Devices are most suitable when events can be traced back to a specific device fingerprint. The `/devices` endpoint can create, list, and update instruments.
type devicesAPI struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newDevicesAPI(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *devicesAPI {
	return &devicesAPI{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateDevice - Create a device
// Creates a new device, sending device data in the request body.
//
// Recommended values for `device_type` include: laptop, home_computer, work_laptop, phone, smartphone, scanner.
//
// If the `/devices/create` API is called for an device that already exists in our system (i.e.  has an existing `device_id`, it is treated it as an  [upsert](https://docs.unit21.ai/u21/reference/should-i-update-or-upsert) and an update on the existing device is performed. The response to the request will then contain the entry `previously_existed: true`.
//
// Unit21 selectively ignores upserts if the request is identical to a previous request.  The response to any isgnored upsert will  contain the field `ignored: true`.
//
// Updates to an device's `device_id` are not allowed.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/u21/reference/endpoint-options)
//   - [Custom data](https://docs.unit21.ai/u21/reference/best-practices-for-custom-data)
//   - [Batch uploads](https://docs.unit21.ai/u21/reference/batch-request-examples)
//   - [Modifying tags](https://docs.unit21.ai/u21/reference/modifying-tags)
//
//
// The response will consist of the following fields:
//
//   | Field                    | Type     | Description                                            |
//   |--------------------------|----------|--------------------------------------------------------|
//   | `device_id`	             | String   | 	Unique identifier of the device on your platform     |
//   | `unit21_id`	             | String   | 	Internal ID of the device within Unit21's system     |
//   | `previously_existed`	   | Boolean  | 	If entity (with the same `device_id`) already exists |
//

func (s *devicesAPI) CreateDevice(ctx context.Context, request operations.CreateDeviceDeviceData) (*operations.CreateDeviceResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/devices/create"

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

	res := &operations.CreateDeviceResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *shared.CreateDeviceResponse
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateDeviceResponse = out
		}
	}

	return res, nil
}

// ExportDevices - Bulk export devices
// Initiates an email and dashboard export of devices. The export will be as a CSV file.
//
// Either the agent `ID` or `email` is required to begin the export.
//
// Either the `filters` or the list of `device IDs` are required for the export.
//

func (s *devicesAPI) ExportDevices(ctx context.Context, request operations.ExportDevicesRequestBody) (*operations.ExportDevicesResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/devices/bulk-export"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "Request", "json")
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

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

	res := &operations.ExportDevicesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// GetDeviceByExternal - Get a device using external ID
// Returns all data objects belonging to a single device.
//
// This endpoint requires the `device_id` which is a unique ID created by your organization to identify the device. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.

func (s *devicesAPI) GetDeviceByExternal(ctx context.Context, request operations.GetDeviceByExternalRequest) (*operations.GetDeviceByExternalResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/{org_name}/devices/{device_id}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

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

	res := &operations.GetDeviceByExternalResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// ListDevices - List devices
// Returns paginated list of of top-level information about devices.
// Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
// * `limit`  indicates how many objects the request returns (the page maximum is 50)
// * `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record. The offset is relative to the number of pages (not the total count of objects).
//
// The `total_count` field contains the total number of devices where the  `response_count` field contains the number of devices included in the response.
//

func (s *devicesAPI) ListDevices(ctx context.Context, request shared.ListRequest) (*operations.ListDevicesResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/devices/list"

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

	res := &operations.ListDevicesResponse{
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

// UpdateDeviceByExternal - Update device using external ID
// Updates an device's information using the `device_id` from your platform.
//
// Updating a device has no required fields. You MAY send any  subset of the fields that the `devices/create` endpoint accepts.
//
// This endpoint requires the `device_id` which is a unique ID created by your organization to identify the device. The `org_name` is your Unit21 appointed organization name such as `google` or `acme`.
//
// Note that you can also update a device using an upsert through `/devices/create`.
//
// Follow the links for more information:
//   - [Endpoint options](https://docs.unit21.ai/u21/reference/endpoint-options)
//   - [Custom data](https://docs.unit21.ai/u21/reference/best-practices-for-custom-data)
//   - [Batch uploads](https://docs.unit21.ai/u21/reference/batch-request-examples)
//   - [Modifying tags](https://docs.unit21.ai/u21/reference/modifying-tags)
//
//
// The response will consist of the following fields:
//
//   | Field                    | Type     | Description                                            |
//   |--------------------------|----------|--------------------------------------------------------|
//   | `device_id`	             | String   | 	Unique identifier of the device on your platform     |
//   | `unit21_id`	             | String   | 	Internal ID of the device within Unit21's system     |
//   | `previously_existed`	   | Boolean  | 	If entity (with the same `device_id`) already exists |

func (s *devicesAPI) UpdateDeviceByExternal(ctx context.Context, request operations.UpdateDeviceByExternalRequest) (*operations.UpdateDeviceByExternalResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/{org_name}/devices/{device_id}/update", request, nil)
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

	res := &operations.UpdateDeviceByExternalResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}