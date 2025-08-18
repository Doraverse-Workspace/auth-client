package mfa

import (
	"encoding/json"
	"fmt"

	client "github.com/Doraverse-Workspace/auth-client/v1"
	"github.com/Doraverse-Workspace/auth-client/v1/model"
)

type mfa struct {
	Headers model.RequestHeaders
}

func NewMFA(headers model.RequestHeaders) *mfa {
	return &mfa{
		Headers: headers,
	}
}

// RequestMFAOtp requests an MFA OTP
// Required headers:
// - Authorization: Bearer <access_token>
func (m *mfa) RequestMFAOtp() error {
	var (
		c = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return err
	}
	resp, err := request.R().
		SetHeaders(m.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		Post(fmt.Sprintf("%s/api/v1/mfa/otp", c.BaseURL))
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return fmt.Errorf("failed to get MFA OTP")
	}
	return nil
}

// VerifyMFA verifies an MFA OTP
// Required headers:
// - Authorization: Bearer <access_token>
// - Content-Type: application/json
// Body:
//
//	{
//		"code": "123456"
//	}
//
// Response:
//
//	{
//		"token": "token_mfa"
//	}
func (m *mfa) VerifyMFA(code string) (*model.VerifyMFATokenResponse, error) {
	var (
		data model.Response
		res  model.VerifyMFATokenResponse
		c    = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}
	resp, err := request.R().
		SetHeaders(m.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(model.VerifyMFATokenRequest{
			Code: code,
		}).
		Post(fmt.Sprintf("%s/api/v1/mfa/verify-otp", c.BaseURL))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to verify MFA"), "failed to verify MFA", c.IsDebug)
		return nil, fmt.Errorf("failed to verify MFA")
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
