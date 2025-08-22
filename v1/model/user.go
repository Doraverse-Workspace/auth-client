package model

// UserInfoResponse is the response body for the user info endpoint.
// ID is the user ID.
// Email is the user email.
// Name is the user name.
// Metadata is the user metadata.
// CreatedAt is the user created at.
// UpdatedAt is the user updated at.
// Example:
//
//	{
//		"id": "1234567890",
//		"email": "john.doe@example.com",
//		"name": "John Doe",
//		"metadata": {},
//		"createdAt": "2025-08-08T02:44:24.541Z",
//		"updatedAt": "2025-08-18T03:46:34.175Z"
//	}
type UserInfoResponse struct {
	ID        string                 `json:"id"`
	Email     string                 `json:"email"`
	Name      string                 `json:"name"`
	Metadata  map[string]interface{} `json:"metadata"`
	CreatedAt string                 `json:"createdAt"`
	UpdatedAt string                 `json:"updatedAt"`
}

// RemoveSessionUserRequest is the request body for the remove session user endpoint.
// UserID is the user ID.
// Example:
//
//	{
//		"userId": "1234567890",
//		"workspaceId": "1234567890" // optional
//	}
type RemoveSessionUserRequest struct {
	UserID      string `json:"userId"`
	WorkspaceID string `json:"workspaceId"` // optional
}

type RemoveSessionUserResponse struct {
}

// LogoutRequest is the request body for the logout endpoint.
// RefreshToken is the refresh token.
// Example:
//
//	{
//		"refreshToken": "1234567890"
//	}
type LogoutRequest struct {
	RefreshToken string `json:"refreshToken"`
}
