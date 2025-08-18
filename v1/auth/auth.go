package auth

import (
	"encoding/json"
	"errors"
	"fmt"

	client "github.com/Doraverse-Workspace/auth-client/v1"
	"github.com/Doraverse-Workspace/auth-client/v1/model"
)

type auth struct {
	Headers model.RequestHeaders
}

func New(headers model.RequestHeaders) *auth {
	return &auth{
		Headers: headers,
	}
}

// GetAccessTokenByRefreshToken gets the access token by refresh token
// returns a GetAccessTokenByRefreshTokenResponse and an error if the get access token by refresh token fails
// refreshToken is the refresh token to get the access token by
// Headers:
// - Authorization: Bearer <access_token>
// - Content-Type: application/json
// Payload:
//
//	{
//		"refreshToken": "1234567890"
//	}
//
// Response:
//
//	{
//		"accessToken": "1234567890",
//		"refreshToken": "1234567890",
//		"expiresIn": 1800
//	}
func (a *auth) GetAccessTokenByRefreshToken(refreshToken string) (*model.GetAccessTokenByRefreshTokenResponse, error) {
	var (
		data model.Response
		res  model.GetAccessTokenByRefreshTokenResponse
		c    = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}
	resp, err := request.R().
		SetHeaders(a.Headers.ConstructHeaders()).
		SetBody(model.GetAccessTokenByRefreshTokenRequest{
			RefreshToken: refreshToken,
		}).
		Post(fmt.Sprintf("%s/api/v1/auth/refresh", c.BaseURL))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get access token by refresh token")
	}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		client.Errorf(err, "failed to unmarshal response", c.IsDebug)
		return nil, err
	}
	b, err := json.Marshal(data.Data)
	if err != nil {
		client.Errorf(err, "failed to marshal response", c.IsDebug)
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		client.Errorf(err, "failed to unmarshal response", c.IsDebug)
		return nil, err
	}
	return &res, nil
}

// ExchangeToken exchanges a code for a token
// returns a ExchangeTokenResponse and an error if the exchange fails
// code is the code to exchange
// Example:
// Payload:
//
//	{
//		"code": "1234567890"
//	}
//
// Response:
//
//	{
//		"refreshToken": "1234567890",
//		"accessToken": "1234567890",
//		"expiresIn": 1800
//	}
func (a *auth) ExchangeToken(code string) (*model.ExchangeTokenResponse, error) {
	var (
		data model.Response
		res  model.ExchangeTokenResponse
		c    = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}
	resp, err := request.R().
		SetHeaders(a.Headers.ConstructHeaders()).
		SetBody(model.ExchangeTokenRequest{
			Code: code,
		}).
		Post(fmt.Sprintf("%s/api/v1/auth/exchange-token", c.BaseURL))

	if err != nil {
		fmt.Println("❌ failed to exchange token", err)
		return nil, err
	}
	if resp.StatusCode() != 200 {
		fmt.Println("❌ failed to exchange token", resp.StatusCode())
		return nil, errors.New("failed to exchange token")
	}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		fmt.Println("❌ failed to unmarshal response", err)
		return nil, err
	}
	b, err := json.Marshal(data.Data)
	if err != nil {
		fmt.Println("❌ failed to marshal response", err)
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		fmt.Println("❌ failed to unmarshal response", err)
		return nil, err
	}
	return &res, nil
}

// GetAuthCodeRequest gets the auth code request
// Body:
//
//	{
//		"userId": "1234567890",
//		"workspaceId": "1234567890"
//	}
//
// Response:
//
//	{
//		"code": "1234567890"
//	}
func (a *auth) GetAuthCodeRequest(userId string, workspaceId string) (*model.GetAuthCodeResponse, error) {
	var (
		data model.Response
		res  model.GetAuthCodeResponse
		c    = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}
	resp, err := request.R().
		SetHeaders(a.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(model.GetAuthCodeRequest{
			UserID:      userId,
			WorkspaceID: workspaceId,
		}).
		Post(fmt.Sprintf("%s/api/v1/auth/auth-code", c.BaseURL))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get auth code")
	}
	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		client.Errorf(err, "failed to unmarshal response", c.IsDebug)
		return nil, err
	}
	b, err := json.Marshal(data.Data)
	if err != nil {
		client.Errorf(err, "failed to marshal response", c.IsDebug)
		return nil, err
	}
	err = json.Unmarshal(b, &res)
	if err != nil {
		client.Errorf(err, "failed to unmarshal response", c.IsDebug)
		return nil, err
	}
	return &res, nil
}
