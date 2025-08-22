package model

import "fmt"

// RequestHeaders is the headers for the API request
// UserAgent is the user agent of the request
// BearerToken is the bearer token of the request
// ClientIP is the client IP of the request
// Example:
//
//	{
//		"UserAgent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
//		"BearerToken": "1234567890",
//		"ClientIP": "127.0.0.1"
//	}
type RequestHeaders struct {
	UserAgent    string
	BearerToken  string
	ClientIP     string
	ContentType  string
	ClientSource string
}

// ConstructHeaders constructs the headers for the API request
// returns a map of headers
func (s RequestHeaders) ConstructHeaders() map[string]string {
	headers := make(map[string]string)
	// if client ip is set, use the client ip
	if s.ClientIP != "" {
		headers["Client-User-IP"] = s.ClientIP
	}
	// if bearer token is set, use the bearer token
	if s.BearerToken != "" {
		headers["Authorization"] = fmt.Sprintf("Bearer %s", s.BearerToken)
	}
	// if user agent is set, use the user agent
	if s.UserAgent != "" {
		headers["Client-User-Agent"] = s.UserAgent
	}
	// if client source is set, use the client source
	if s.ClientSource != "" {
		headers["X-Client-Scope"] = s.ClientSource
	}
	// default content type is application/json
	headers["Content-Type"] = "application/json"
	// if content type is set, use the content type
	if s.ContentType != "" {
		headers["Content-Type"] = s.ContentType
	}
	return headers
}

// Response is the response body for the API
// Data is the data of the response
// Message is the message of the response
// Code is the code of the response
// Example:
//
//	{
//		"data": {},
//		"message": "success",
//		"code": 200
//	}
type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Code    string      `json:"code"`
}
