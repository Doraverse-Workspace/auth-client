package connector

import (
	"encoding/json"
	"fmt"

	client "github.com/Doraverse-Workspace/auth-client/v1"
	"github.com/Doraverse-Workspace/auth-client/v1/model"
)

const pathConnectorGoogleDrive = "api/v1/auth/connectors/google-drive"

// GoogleDriveAuthURL gets the google drive auth url
// Required headers:
// - Authorization: Bearer <access_token>
// - Content-Type: application/json
// Query:
// - callbackUrl: the callback URL
// Response:
//
//	{
//		"authUrl": "https://accounts.google.com/o/oauth2/auth?client_id=1234567890&redirect_uri=https://example.com/callback&response_type=code&scope=https://www.googleapis.com/auth/drive.readonly"
//	}
func (cr *connector) GoogleDriveAuthURL(callbackURL string) (*model.GoogleDriveAuthURLResponse, error) {
	var (
		data model.Response
		res  model.GoogleDriveAuthURLResponse
		c    = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}
	resp, err := request.R().
		SetHeaders(cr.Headers.ConstructHeaders()).
		Get(fmt.Sprintf("%s/%s/auth-url?callbackUrl=%s", c.BaseURL, pathConnectorGoogleDrive, callbackURL))
	if err != nil {
		client.Errorf(err, "failed to get google drive auth url", c.IsDebug)
		return nil, err
	}
	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to get google drive auth url"), "failed to get google drive auth url", c.IsDebug)
		return nil, fmt.Errorf("failed to get google drive auth url")
	}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		client.Errorf(err, "failed to unmarshal google drive auth url", c.IsDebug)
		return nil, err
	}
	b, err := json.Marshal(data.Data)
	if err != nil {
		client.Errorf(err, "failed to marshal google drive auth url", c.IsDebug)
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		client.Errorf(err, "failed to unmarshal google drive auth url", c.IsDebug)
		return nil, err
	}
	return &res, nil
}

// GoogleDriveExchangeToken exchanges the google drive token
// Required headers:
// - Authorization: Bearer <access_token>
// - Content-Type: application/json
// Body:
//
//	{
//		"code": "1234567890",
//		"state": "..."
//	}
//
// Response:
//
//	{
//		"accessToken": "1234567890",
//		"expiresIn": 3600
//	}
func (cr *connector) GoogleDriveExchangeToken(code, state string) (*model.GoogleDriveExchangeTokenResponse, error) {
	var (
		data model.Response
		res  *model.GoogleDriveExchangeTokenResponse
		c    = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}
	resp, err := request.R().
		SetHeaders(cr.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(model.GoogleDriveExchangeTokenRequest{
			Code:  code,
			State: state,
		}).
		Post(fmt.Sprintf("%s/%s/exchange-token", c.BaseURL, pathConnectorGoogleDrive))
	if err != nil {
		client.Errorf(err, "failed to exchange google drive token", c.IsDebug)
		return nil, err
	}
	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to exchange google drive token"), "failed to exchange google drive token", c.IsDebug)
		return nil, fmt.Errorf("failed to exchange google drive token")
	}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		client.Errorf(err, "failed to unmarshal google drive exchange token", c.IsDebug)
		return nil, err
	}
	b, err := json.Marshal(data.Data)
	if err != nil {
		client.Errorf(err, "failed to marshal google drive exchange token", c.IsDebug)
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		client.Errorf(err, "failed to unmarshal google drive exchange token", c.IsDebug)
		return nil, err
	}
	return res, nil
}

// GoogleDriveGetTokenByAccessToken gets the google drive token by access token
// Required headers:
// - Authorization: Bearer <access_token>
// - Content-Type: application/json
// Response:
//
//	{
//		"accessToken": "1234567890",
//		"expiresIn": 3600
//	}
func (cr *connector) GoogleDriveGetTokenByAccessToken() (*model.GoogleDriveExchangeTokenResponse, error) {
	var (
		data model.Response
		res  *model.GoogleDriveExchangeTokenResponse
		c    = client.GetClient()
	)

	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}
	resp, err := request.R().
		SetHeaders(cr.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		Get(fmt.Sprintf("%s/%s/token", c.BaseURL, pathConnectorGoogleDrive))
	if err != nil {
		client.Errorf(err, "failed to get google drive token by access token", c.IsDebug)
		return nil, err
	}
	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to get google drive token by access token"), "failed to get google drive token by access token", c.IsDebug)
		return nil, fmt.Errorf("failed to get google drive token by access token")
	}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		client.Errorf(err, "failed to unmarshal google drive token by access token", c.IsDebug)
		return nil, err
	}
	b, err := json.Marshal(data.Data)
	if err != nil {
		client.Errorf(err, "failed to marshal google drive token by access token", c.IsDebug)
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		client.Errorf(err, "failed to unmarshal google drive token by access token", c.IsDebug)
		return nil, err
	}
	return res, nil
}
