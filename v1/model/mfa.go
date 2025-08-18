package model

// GetMFATokenResponse is the response body for the GetMFAToken API
// MFAToken is the MFA token
// Example:
//
//	{
//		"token": "1234567890"
//	}
type GetMFATokenResponse struct {
	MFAToken string `json:"token"`
}

// VerifyMFATokenRequest is the request body for the VerifyMFAToken API
// Code is the MFA code
// Example:
//
//	{
//		"code": "1234567890"
//	}
type VerifyMFATokenRequest struct {
	Code string `json:"code"`
}

// VerifyMFATokenResponse is the response body for the VerifyMFAToken API
// Token is the MFA token
// Example:
//
//	{
//		"token": "1234567890"
//	}
type VerifyMFATokenResponse struct {
	Token string `json:"token"`
}

// ValidateMFATokenRequest is the request body for the ValidateMFAToken API
// Token is the MFA token
// Example:
//
//	{
//		"token": "1234567890"
//	}
type ValidateMFATokenRequest struct {
	Token string `json:"token"`
}
