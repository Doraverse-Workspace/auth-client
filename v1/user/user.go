package user

import (
	"encoding/json"
	"fmt"

	client "github.com/Doraverse-Workspace/auth-client/v1"
	"github.com/Doraverse-Workspace/auth-client/v1/model"
)

type user struct {
	Headers model.RequestHeaders
}

func New(headers model.RequestHeaders) *user {
	return &user{
		Headers: headers,
	}
}

// GetUserInfo gets the user info
// Required headers:
// - Authorization: Bearer <access_token>
// - Content-Type: application/json
// Example:
//
//	{
//		"id": "1234567890",
//		"email": "john.doe@example.com",
//		"firstName": "John",
//		"lastName": "Doe"
//	}
func (u *user) GetUserInfo() (*model.UserInfoResponse, error) {
	var (
		data model.Response
		res  model.UserInfoResponse
		c    = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}
	resp, err := request.R().
		SetHeaders(u.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		Get(fmt.Sprintf("%s/api/v1/user/profile", c.BaseURL))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("failed to get user info")
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

// Logout logs out the user
// Required headers:
// - Authorization: Bearer <access_token>
// - Content-Type: application/json
func (u *user) Logout(refreshToken string) error {
	var (
		c = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return err
	}
	resp, err := request.R().
		SetHeaders(u.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(model.LogoutRequest{
			RefreshToken: refreshToken,
		}).
		Post(fmt.Sprintf("%s/api/v1/auth/logout", c.BaseURL))
	if err != nil {
		client.Errorf(err, "failed to logout", c.IsDebug)
		return err
	}
	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to logout"), "failed to logout", c.IsDebug)
		return fmt.Errorf("failed to logout")
	}
	return nil
}

// RemoveSessionUser deletes the user session
// Required headers:
// - Authorization: Bearer <access_token>
// - Content-Type: application/json
func (u *user) RemoveSessionUser(userId string, workspaceId string) error {
	var (
		c = client.GetClient()
	)
	request, err := c.NewRequest()
	if err != nil {
		return err
	}
	resp, err := request.R().
		SetHeaders(u.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(model.RemoveSessionUserRequest{
			UserID:      userId,
			WorkspaceID: workspaceId,
		}).
		Delete(fmt.Sprintf("%s/api/v1/auth/user-session", c.BaseURL))
	if err != nil {
		client.Errorf(err, "failed to remove session user", c.IsDebug)
		return err
	}
	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to remove session user"), "failed to remove session user", c.IsDebug)
		return fmt.Errorf("failed to remove session user")
	}
	return nil
}
