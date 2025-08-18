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

func NewAuthClient(headers model.RequestHeaders) *auth {
	return &auth{
		Headers: headers,
	}
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
