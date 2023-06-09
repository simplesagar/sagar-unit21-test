// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdk

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"unit/pkg/models/operations"
	"unit/pkg/utils"
)

// verificationFormsAPI - With verification forms, you can automate ID verification and user collection. To gather user input, the `verification forms` endpoint creates a URL. This URL is only valid for a specified period of time.
type verificationFormsAPI struct {
	defaultClient  HTTPClient
	securityClient HTTPClient
	serverURL      string
	language       string
	sdkVersion     string
	genVersion     string
}

func newVerificationFormsAPI(defaultClient, securityClient HTTPClient, serverURL, language, sdkVersion, genVersion string) *verificationFormsAPI {
	return &verificationFormsAPI{
		defaultClient:  defaultClient,
		securityClient: securityClient,
		serverURL:      serverURL,
		language:       language,
		sdkVersion:     sdkVersion,
		genVersion:     genVersion,
	}
}

// CreateVerificationForm - Verification Forms API
// If you are verifying IDs and collecting user data, this endpoint creates a temporary URL to which you can redirect users.

func (s *verificationFormsAPI) CreateVerificationForm(ctx context.Context, request operations.CreateVerificationFormRequestBody) (*operations.CreateVerificationFormResponse, error) {
	baseURL := s.serverURL
	url := strings.TrimSuffix(baseURL, "/") + "/verification-forms/create"

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

	res := &operations.CreateVerificationFormResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out *operations.CreateVerificationForm200ApplicationJSON
			if err := utils.UnmarshalJsonFromResponseBody(httpRes.Body, &out); err != nil {
				return nil, err
			}

			res.CreateVerificationForm200ApplicationJSONObject = out
		}
	}

	return res, nil
}
