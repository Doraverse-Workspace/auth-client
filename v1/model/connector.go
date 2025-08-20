package model

// GoogleDriveAuthURLRequest is the request body for the google drive auth url endpoint.
// CallbackURL is the callback URL.
// Example:
//
//	{
//		"callbackUrl": "https://example.com/callback"
//	}
type GoogleDriveAuthURLRequest struct {
	CallbackURL string `json:"callbackUrl"`
}

// GoogleDriveAuthURLResponse is the response body for the google drive auth url endpoint.
// AuthURL is the auth URL.
// Example:
//
//	{
//		"authUrl": "https://accounts.google.com/o/oauth2/auth?client_id=1234567890&redirect_uri=https://example.com/callback&response_type=code&scope=https://www.googleapis.com/auth/drive.readonly"
//	}
type GoogleDriveAuthURLResponse struct {
	AuthURL string `json:"authUrl"`
}

// GoogleDriveExchangeTokenRequest is the request body for the google drive exchange token endpoint.
// Code is the code.
// Example:
//
//	{
//		"code": "1234567890"
//		"state": "..."
//	}
type GoogleDriveExchangeTokenRequest struct {
	Code  string `json:"code"`
	State string `json:"state"`
}

// GoogleDriveExchangeTokenResponse is the response body for the google drive exchange token endpoint.
// AccessToken is the access token.
// ExpiresIn is the expires in.
// Example:
//
//	{
//		"accessToken": "1234567890",
//		"expiresIn": 3600
//	}
type GoogleDriveExchangeTokenResponse struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   int    `json:"expiresIn"`
}
