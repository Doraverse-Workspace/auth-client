# 🔐 Auth Client

Go client library for interacting with authentication and MFA APIs of the Doraverse system.

## ✨ Features

- **🔑 Authentication**: Exchange authorization code for access token and refresh token
- **🛡️ Multi-Factor Authentication (MFA)**: Support for requesting, verifying, and validating OTP codes  
- **👤 User Management**: Get user profile information, logout, and session management
- **🏢 Workspace Management**: Create workspaces, invite members, update member roles, and manage member status
- **🔗 Third-party Connectors**: Integration with external services like Google Drive
- **🔒 TLS Support**: Support for mutual TLS authentication with certificates
- **🐛 Debug Mode**: Debug mode for tracking requests/responses
- **📋 Custom Headers**: Support for custom headers for tracking and authorization

## 📦 Installation

```bash
go get github.com/Doraverse-Workspace/auth-client
```

## 🚀 Usage

### 1. 🔧 Initialize Client

```go
package main

import (
    "github.com/Doraverse-Workspace/auth-client/v1"
    "github.com/Doraverse-Workspace/auth-client/v1/model"
)

func main() {
    // Configure client
    config := v1.ClientConfig{
        BaseURL: "https://api.doraverse.com",
        TLS: v1.TLSConfig{
            CertFile: "path/to/cert.pem",
            KeyFile:  "path/to/key.pem",
        },
        IsDebug: true, // Enable debug mode
    }
    
    // Create client
    v1.NewClient(config)
}
```

### 2. 🔑 Authentication - Exchange Code for Token

```go
import (
    "github.com/Doraverse-Workspace/auth-client/v1/auth"
    "github.com/Doraverse-Workspace/auth-client/v1/model"
)

// Create headers for request
headers := model.RequestHeaders{
    UserAgent:   "MyApp/1.0.0",
    ClientIP:    "192.168.1.1",
    ContentType: "application/json",
}

// Create auth client
authClient := auth.New(headers)

// Exchange authorization code for tokens
response, err := authClient.ExchangeToken("your_authorization_code")
if err != nil {
    panic(err)
}

fmt.Printf("Access Token: %s\n", response.AccessToken)
fmt.Printf("Refresh Token: %s\n", response.RefreshToken)
fmt.Printf("Expires In: %d seconds\n", response.ExpiresIn)
```

### 3. 🛡️ Multi-Factor Authentication (MFA)

```go
import (
    "github.com/Doraverse-Workspace/auth-client/v1/mfa"
    "github.com/Doraverse-Workspace/auth-client/v1/model"
)

// Create headers with access token
headers := model.RequestHeaders{
    UserAgent:   "MyApp/1.0.0",
    BearerToken: "your_access_token", // Token from authentication step
    ClientIP:    "192.168.1.1",
}

// Create MFA client
mfaClient := mfa.New(headers)

// Request OTP code
err := mfaClient.RequestMFAOtp()
if err != nil {
    panic(err)
}

// Verify OTP code
mfaResponse, err := mfaClient.VerifyMFA("123456") // OTP code from user
if err != nil {
    panic(err)
}

fmt.Printf("MFA Token: %s\n", mfaResponse.Token)
fmt.Printf("Expires In: %d seconds\n", mfaResponse.ExpiresIn)

// Validate MFA token
validateResponse, err := mfaClient.ValidateMFAToken(mfaResponse.Token)
if err != nil {
    panic(err)
}

fmt.Printf("User ID: %s\n", validateResponse.UserID)
fmt.Printf("Username: %s\n", validateResponse.Username)
```

### 4. 👤 User Management

```go
import (
    "github.com/Doraverse-Workspace/auth-client/v1/user"
    "github.com/Doraverse-Workspace/auth-client/v1/model"
)

// Create headers with access token
headers := model.RequestHeaders{
    UserAgent:   "MyApp/1.0.0",
    BearerToken: "your_access_token", // Token from authentication step
    ClientIP:    "192.168.1.1",
}

// Create user client
userClient := user.New(headers)

// Get user profile information
userInfo, err := userClient.GetUserInfo()
if err != nil {
    panic(err)
}

fmt.Printf("User ID: %s\n", userInfo.ID)
fmt.Printf("Email: %s\n", userInfo.Email)
fmt.Printf("Name: %s\n", userInfo.Name)

// Logout user
err = userClient.Logout()
if err != nil {
    panic(err)
}

// Remove user session (admin operation)
err = userClient.RemoveSessionUser("user123", "workspace456")
if err != nil {
    panic(err)
}
```

### 5. 🏢 Workspace Management

```go
import (
    "github.com/Doraverse-Workspace/auth-client/v1/workspace"
    "github.com/Doraverse-Workspace/auth-client/v1/model"
)

// Create headers with access token
headers := model.RequestHeaders{
    UserAgent:   "MyApp/1.0.0",
    BearerToken: "your_access_token", // Token from authentication step
    ClientIP:    "192.168.1.1",
}

// Create workspace client
workspaceClient := workspace.New(headers)

// Create a new workspace
createWorkspaceRequest := model.CreateWorkspaceRequest{
    Members: []model.WorkspaceMember{
        {
            DepartmentCode: "DEPT001",
            Email:          "admin@example.com",
            IsCreator:      true,
            Name:           "Admin User",
            RoleCode:       "ADMIN",
        },
    },
    Workspace: model.Workspace{
        ID:       "workspace123",
        Name:     "My Workspace",
        Hostname: "my-workspace.doradora.vn",
    },
}

workspaceResponse, err := workspaceClient.CreateWorkspace(createWorkspaceRequest)
if err != nil {
    panic(err)
}

fmt.Printf("Workspace ID: %s\n", workspaceResponse.WorkspaceID)

// Invite members to workspace
inviteRequest := model.InviteMembersRequest{
    Members: []model.MemberInvite{
        {
            Email:          "user@example.com",
            RoleCode:       "USER",
            DepartmentCode: "DEPT002",
        },
    },
}

inviteResponse, err := workspaceClient.InviteMembers("workspace123", inviteRequest)
if err != nil {
    panic(err)
}

fmt.Printf("Invited members count: %d\n", len(inviteResponse.Members))

// Update member role
updateRequest := model.UpdateMemberRequest{
    RoleCode:       "ADMIN",
    DepartmentCode: "DEPT001",
}

updateResponse, err := workspaceClient.UpdateMember("workspace123", "member456", updateRequest)
if err != nil {
    panic(err)
}

fmt.Printf("Updated member ID: %s\n", updateResponse.ID)

// Change member status (activate/deactivate)
err = workspaceClient.ChangeStatusMember("workspace123", "member456", "active")
if err != nil {
    panic(err)
}

fmt.Printf("Member status updated successfully\n")
```

### 6. 🔗 Google Drive Connector

```go
import (
    "github.com/Doraverse-Workspace/auth-client/v1/connector"
    "github.com/Doraverse-Workspace/auth-client/v1/model"
)

// Create headers with access token
headers := model.RequestHeaders{
    UserAgent:   "MyApp/1.0.0",
    BearerToken: "your_access_token", // Token from authentication step
    ClientIP:    "192.168.1.1",
}

// Create connector client
connectorClient := connector.New(headers)

// Get Google Drive authorization URL
authResponse, err := connectorClient.GoogleDriveAuthURL("https://yourapp.com/callback")
if err != nil {
    panic(err)
}

fmt.Printf("Auth URL: %s\n", authResponse.AuthURL)

// Exchange authorization code for Google Drive token
tokenResponse, err := connectorClient.GoogleDriveExchangeToken("auth_code", "state")
if err != nil {
    panic(err)
}

fmt.Printf("Google Drive Access Token: %s\n", tokenResponse.AccessToken)
fmt.Printf("Expires In: %d seconds\n", tokenResponse.ExpiresIn)

// Get existing Google Drive token
existingToken, err := connectorClient.GoogleDriveGetTokenByAccessToken()
if err != nil {
    panic(err)
}

fmt.Printf("Existing Token: %s\n", existingToken.AccessToken)
```

## 📂 Project Structure

```
auth-client/
├── v1/
│   ├── client.go          # Main client and TLS configuration
│   ├── utils.go           # Utility functions (error handling)
│   ├── auth/
│   │   └── auth.go        # Authentication functions
│   ├── mfa/
│   │   └── mfa.go         # Multi-Factor Authentication functions
│   ├── user/
│   │   └── user.go        # User management functions
│   ├── workspace/
│   │   └── workspace.go   # Workspace management functions
│   ├── connector/
│   │   ├── connector.go   # Base connector functionality
│   │   └── google_drive.go # Google Drive integration
│   └── model/
│       ├── auth.go        # Auth request/response models
│       ├── common.go      # Common models and headers
│       ├── mfa.go         # MFA request/response models
│       ├── user.go        # User models
│       ├── workspace.go   # Workspace models
│       └── connector.go   # Connector models
├── go.mod
└── README.md
```

## 📋 API Models

### 🔑 Authentication Models

```go
// Request to exchange authorization code
type ExchangeTokenRequest struct {
    Code string `json:"code"`
}

// Response containing tokens
type ExchangeTokenResponse struct {
    RefreshToken string `json:"refreshToken"`
    AccessToken  string `json:"accessToken"`
    ExpiresIn    int    `json:"expiresIn"` // Expiration time (seconds)
}
```

### 🛡️ MFA Models

```go
// Request to verify MFA OTP
type VerifyMFATokenRequest struct {
    Code string `json:"code"`
}

// Response containing MFA token
type VerifyMFATokenResponse struct {
    Token     string `json:"token"`
    ExpiresIn int    `json:"expiresIn"` // Expiration time (seconds)
}

// Request to validate MFA token
type ValidateMFATokenRequest struct {
    Token string `json:"token"`
}

// Response containing user information after MFA validation
type ValidateMFATokenResponse struct {
    UserID   string `json:"userId"`
    Username string `json:"username"`
}
```

### 👤 User Management Models

```go
// Response containing user profile information
type UserInfoResponse struct {
    ID        string                 `json:"id"`
    Email     string                 `json:"email"`
    Name      string                 `json:"name"`
    Metadata  map[string]interface{} `json:"metadata"`
    CreatedAt string                 `json:"createdAt"`
    UpdatedAt string                 `json:"updatedAt"`
}

// Request to remove user session
type RemoveSessionUserRequest struct {
    UserID      string `json:"userId"`
    WorkspaceID string `json:"workspaceId"` // optional
}
```

### 🏢 Workspace Management Models

```go
// Request to create a workspace
type CreateWorkspaceRequest struct {
    Members   []WorkspaceMember `json:"members"`
    Workspace Workspace         `json:"workspace"`
}

// Workspace information
type Workspace struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Hostname string `json:"hostname"`
}

// Workspace member information
type WorkspaceMember struct {
    DepartmentCode string `json:"departmentCode"`
    Email          string `json:"email"`
    IsCreator      bool   `json:"isCreator"`
    Name           string `json:"name"`
    RoleCode       string `json:"roleCode"`
}

// Response after creating workspace
type CreateWorkspaceResponse struct {
    WorkspaceID string          `json:"workspaceId"`
    Members     []MemberReponse `json:"members"`
}

// Request to invite members to workspace
type InviteMembersRequest struct {
    Members []MemberInvite `json:"members"`
}

// Member invitation information
type MemberInvite struct {
    Email          string `json:"email"`
    RoleCode       string `json:"roleCode"`
    DepartmentCode string `json:"departmentCode"`
}

// Request to update member information
type UpdateMemberRequest struct {
    RoleCode       string `json:"roleCode"`
    DepartmentCode string `json:"departmentCode"`
}

// Response after updating member
type UpdateMemberResponse struct {
    ID string `json:"id"`
}

// Request to change member status
type ChangeStatusMemberRequest struct {
    Status string `json:"status"` // "active" or "inactive"
}
```

### 🔗 Google Drive Connector Models

```go
// Request for Google Drive authorization URL
type GoogleDriveAuthURLRequest struct {
    CallbackURL string `json:"callbackUrl"`
}

// Response containing Google Drive authorization URL
type GoogleDriveAuthURLResponse struct {
    AuthURL string `json:"authUrl"`
}

// Request to exchange Google Drive authorization code
type GoogleDriveExchangeTokenRequest struct {
    Code  string `json:"code"`
    State string `json:"state"`
}

// Response containing Google Drive access token
type GoogleDriveExchangeTokenResponse struct {
    AccessToken string `json:"accessToken"`
    ExpiresIn   int    `json:"expiresIn"`
}
```

### 🔄 Common Models

```go
// Headers for API requests
type RequestHeaders struct {
    UserAgent   string // User agent of the application
    BearerToken string // Access token for authorization
    ClientIP    string // Client IP address
    ContentType string // Content type (default: application/json)
}

// Common response structure from API
type Response struct {
    Data    interface{} `json:"data"`
    Message string      `json:"message"`
    Code    int         `json:"code"`
}
```

## 🔒 TLS Configuration

The client supports mutual TLS authentication. You need to provide:

- **📄 CertFile**: Path to certificate file (PEM format)
- **🔑 KeyFile**: Path to private key file (PEM format)

```go
tlsConfig := v1.TLSConfig{
    CertFile: "/path/to/client.crt",
    KeyFile:  "/path/to/client.key",
}
```

## 🐛 Debug Mode

When debug mode is enabled (`IsDebug: true`), the client will print detailed information about requests/responses to assist with debugging.

## ⚠️ Error Handling

The library uses Go's built-in error handling. All functions return an error if something goes wrong.

```go
response, err := authClient.ExchangeToken(code)
if err != nil {
    // Handle error
    log.Printf("Exchange token failed: %v", err)
    return
}
```

## 📦 Dependencies

- **🌐 resty/v2**: HTTP client library for Go
- **🐹 Go 1.24+**: Minimum Go version required

## 🤝 Contributing

1. 🍴 Fork the repository
2. 🌿 Create a feature branch (`git checkout -b feature/amazing-feature`)
3. 💾 Commit your changes (`git commit -m 'Add some amazing feature'`)
4. 📤 Push to the branch (`git push origin feature/amazing-feature`)
5. 🔄 Create a Pull Request

## 📄 License

This project belongs to Doraverse Workspace.

## 🌐 API Endpoints

The client interacts with the following endpoints:

### 🔑 Authentication
- `POST /api/v1/auth/exchange-token` - Exchange authorization code for tokens
- `POST /api/v1/auth/logout` - Logout user session

### 🛡️ Multi-Factor Authentication
- `POST /api/v1/auth/mfa/otp` - Request OTP code
- `POST /api/v1/auth/mfa/verify-otp` - Verify OTP code
- `POST /api/v1/auth/mfa/validate` - Validate MFA token

### 👤 User Management
- `GET /api/v1/user/profile` - Get user profile information
- `DELETE /api/v1/auth/user-session` - Remove user session

### 🏢 Workspace Management
- `POST /api/v1/workspace` - Create a new workspace
- `POST /api/v1/workspace/{workspaceId}/members` - Invite members to workspace
- `PUT /api/v1/workspace/{workspaceId}/member/{memberId}` - Update member information
- `PATCH /api/v1/workspace/{workspaceId}/member/{memberId}/status` - Change member status

### 🔗 Google Drive Connector
- `GET /api/v1/auth/connectors/google-drive/auth-url` - Get Google Drive authorization URL
- `POST /api/v1/auth/connectors/google-drive/exchange-token` - Exchange Google Drive authorization code
- `GET /api/v1/auth/connectors/google-drive/token` - Get existing Google Drive token

## 💬 Support

If you have any issues or questions, please create an issue on the GitHub repository.
