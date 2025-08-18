package model

// ExchangeTokenRequest is the request body for the exchange token endpoint.
// Code is required.
// Example:
//
//	{
//		"code": "1234567890"
//	}
type ExchangeTokenRequest struct {
	Code string `json:"code"`
}

// ExchangeTokenResponse is the response body for the exchange token endpoint.
// RefreshToken is the refresh token.
// AccessToken is the access token.
// ExpiresIn is the expiration time of the access token in seconds.
// Example:
//
//	{
//		"refreshToken": "ey...aw...",
//		"accessToken": "ey...aw...",
//		"expiresIn": 1800 // 30 minutes
//	}
type ExchangeTokenResponse struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
	ExpiresIn    int    `json:"expiresIn"`
}

// GetAuthCodeRequest is the request body for the get auth code endpoint.
// UserID is the user ID.
// WorkspaceID is the workspace ID.
// Example:
//
//	{
//		"userId": "1234567890",
//		"workspaceId": "1234567890"
//	}
type GetAuthCodeRequest struct {
	UserID      string `json:"userId"`
	WorkspaceID string `json:"workspaceId"`
}

// GetAuthCodeResponse is the response body for the get auth code endpoint.
// Code is the auth code.
// Example:
//
//	{
//		"code": "1234567890"
//	}
type GetAuthCodeResponse struct {
	Code string `json:"code"`
}

// GetAccessTokenByRefreshTokenRequest is the request body for the get access token by refresh token endpoint.
// RefreshToken is the refresh token.
// Example:
//
//	{
//		"refreshToken": "1234567890"
//	}
type GetAccessTokenByRefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken"`
}

// GetAccessTokenByRefreshTokenResponse is the response body for the get access token by refresh token endpoint.
// AccessToken is the access token.
// ExpiresIn is the expiration time of the access token in seconds.
// Example:
//
//	{
//		"accessToken": "ey...aw...",
//		"refreshToken": "ey...aw...",
//		"expiresIn": 1800 // 30 minutes
//	}
type GetAccessTokenByRefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    int    `json:"expiresIn"`
}
