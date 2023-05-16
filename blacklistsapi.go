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

// blacklistsAPI - Blacklists comprise one of the following categories:
//   - of entities (users or business)
//   - IPs (single or ranges)
//   - strings
type blacklistsAPI struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newBlacklistsAPI(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *blacklistsAPI {
	return &blacklistsAPI{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// AddBlacklistValues - Add items to a blacklist
// Add items to a blacklist, according to the blacklist's `type`.
//
// Each request must specify at least **1** object to blacklist. You may add up to **100**  values to an existing blacklist at once.
//
// The `/blacklists/<blacklist-id>/add-values` API will ignore entries provided that already exist  in the blacklist. No error will be thrown when this occurs.
//
// The response will consist of the following fields:
//
//	| Type       | Description                                                              | Example                           |
//	|------------|--------------------------------------------------------------------------|-----------------------------------|
//	| `STRING`	 | Plain strings to match against any text-type field.                      | 		"blacklist_value": "abcde"    |
//	| `IP_INET`	 | IPv4 or IPv6 IP addresses to blacklist.                                  | 	"ip_address": "255.255.255.255" |
//	| `IP_CIDR`	 | Classless Inter-Domain Routing (CIDR) [notation](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) IP address ranges to blacklist.  | 	"cidr": "255.255.255.255/32" |
//	| `USER`	   | 	Series of fields that a Unit21 user entity will be matched against.     | 	user_data object                |
//	| `BUSINESS` | Series of fields that a Unit21 business entity will be matched against.  | 	business_data object            |
func (s *blacklistsAPI) AddBlacklistValues(ctx context.Context, request operations.AddBlacklistValuesRequest) (*operations.AddBlacklistValuesResponse, error) {
	baseURL := s.serverURL
	url, err := utils.GenerateURL(ctx, baseURL, "/blacklists/{unit21_id}/add-values", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, "RequestBody", "json")
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

	res := &operations.AddBlacklistValuesResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
	}

	return res, nil
}

// CreateBlacklist - Create a blacklist
// Create a new blacklist sending blacklist data in the request body.
//
// Unit21 currently supports 5 types of blacklists:
//
//   - `STRING`: Plain strings to match against any text-type field.
//   - `IP_INET`: IPv4 or IPv6 IP addresses to blacklist.
//   - `IP_CIDR`: [Classless Inter-Domain Routing (CIDR) notation IP address ranges](https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing#CIDR_notation) to blacklist,
//   - `USER`: Series of fields that a Unit21 user entity will be matched against.
//   - `BUSINESS`: Series of fields that a Unit21 business entity will be matched against.
//
// If the `/blacklists/create` API is called multiple times, it will create a new blacklist each time.  This endpoint does not support updates/upserts.
//
// This endpoint does not support batch uploads.
//
// The response will consist of the following fields:
//
//	| Field           | Type     | Description                                           |
//	|-----------------|----------|-------------------------------------------------------|
//	| `blacklist_id`  | String   | 	Unique identifier of the entity on your platform     |
func (s *blacklistsAPI) CreateBlacklist(ctx context.Context, request shared.CreateBlacklist) (*operations.CreateBlacklistResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/blacklists/create"

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

	res := &operations.CreateBlacklistResponse{
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

			res.CreateBlacklist200ApplicationJSONAny = out
		}
	}

	return res, nil
}

// ListBlacklists - List blacklists
// Returns an array of blacklist in your environment.
//
// Because the response is paginated, the request body has a `limit` and `offset` field. At least one must be filled.
// * `limit`  indicates how many objects the request returns (the page maximum is 50)
// * `offset` indicates the offset for pagination. An `offset` value of 1 starts with the environment's first record.
//
// The `total_count` field contains the total number of blacklists where the  `response_count` field contains the number of blacklists included in the response.
func (s *blacklistsAPI) ListBlacklists(ctx context.Context, request shared.ListRequest) (*operations.ListBlacklistsResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/blacklists/list"

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

	res := &operations.ListBlacklistsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.ListBlacklists200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.ListBlacklists200ApplicationJSONObject = out
		}
	}

	return res, nil
}
