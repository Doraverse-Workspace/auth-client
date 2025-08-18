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
