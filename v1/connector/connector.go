package connector

import (
	"github.com/Doraverse-Workspace/auth-client/v1/model"
)

type connector struct {
	Headers model.RequestHeaders
}

func New(headers model.RequestHeaders) *connector {
	return &connector{
		Headers: headers,
	}
}
