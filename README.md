# 🔐 Auth Client

Go client library for interacting with authentication and MFA APIs of the Doraverse system.

## ✨ Features

- **🔑 Authentication**: Exchange authorization code for access token and refresh token
- **🛡️ Multi-Factor Authentication (MFA)**: Support for requesting and verifying OTP codes
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
    client := v1.NewClient(config)
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
authClient := auth.NewAuthClient(headers)

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
mfaClient := mfa.NewMFA(headers)

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
│   └── model/
│       ├── auth.go        # Auth request/response models
│       ├── common.go      # Common models and headers
│       └── mfa.go         # MFA request/response models
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
    Token string `json:"token"`
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

- `POST /api/v1/auth/exchange-token` - 🔑 Exchange authorization code
- `POST /api/v1/mfa/otp` - 📱 Request OTP code
- `POST /api/v1/mfa/verify-otp` - ✅ Verify OTP code

## 💬 Support

If you have any issues or questions, please create an issue on the GitHub repository.
