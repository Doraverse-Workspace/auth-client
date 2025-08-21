package workspace

import (
	"encoding/json"
	"fmt"

	client "github.com/Doraverse-Workspace/auth-client/v1"
	"github.com/Doraverse-Workspace/auth-client/v1/model"
)

const path = "api/v1/workspace"

type Workspace struct {
	Headers model.RequestHeaders
}

func New(headers model.RequestHeaders) *Workspace {
	return &Workspace{Headers: headers}
}

// SyncData ...
func (w *Workspace) SyncData(payload model.SyncDataRequest) error {
	var (
		c = client.GetClient()
	)

	request, err := c.NewRequest()
	if err != nil {
		return err
	}

	resp, err := request.R().
		SetHeaders(w.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(payload).
		Post(fmt.Sprintf("%s/api/v1/migration/workspace", c.BaseURL))
	if err != nil {
		client.Errorf(err, "failed to sync data", c.IsDebug)
		return err
	}

	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to sync data"), "failed to sync data", c.IsDebug)
		return fmt.Errorf("failed to sync data")
	}
	return nil
}

// CreateWorkspace creates a new workspace
// Required headers:
// - Content-Type: application/json
// Body:
//
//	{
//		"members": [
//			{
//				"departmentCode": "1234567890",
//				"email": "john.doe@example.com",
//				"isCreator": true,
//				"name": "John Doe",
//				"roleCode": "ADMIN"
//			}
//		],
//		"workspace": {
//			"id": "1234567890",
//			"name": "Workspace Name",
//			"hostname": "workspace-name.doradora.vn"
//		}
//	}
//
// Response:
//
//	{
//		"workspaceId": "1234567890",
//		"members": [
//			{
//				"departmentCode": "1234567890",
//				"email": "john.doe@example.com",
//				"isCreator": true,
//				"name": "John Doe",
//				"roleCode": "ADMIN"
//				"id": "1234567890",
//				"userId": "1234567890"
//			}
//		]
//	}
func (w *Workspace) CreateWorkspace(payload model.CreateWorkspaceRequest) (*model.CreateWorkspaceResponse, error) {
	var (
		data model.Response
		res  model.CreateWorkspaceResponse
		c    = client.GetClient()
	)

	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}

	resp, err := request.R().
		SetHeaders(w.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(payload).
		Post(fmt.Sprintf("%s/%s", c.BaseURL, path))
	if err != nil {
		client.Errorf(err, "failed to create workspace", c.IsDebug)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to create workspace"), "failed to create workspace", c.IsDebug)
		return nil, fmt.Errorf("failed to create workspace")
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

// InviteMembers invites members to a workspace
// Required headers:
// - Content-Type: application/json
// Body:
//
//	{
//		"members": [
//			{
//				"email": "john.doe@example.com",
//				"departmentCode": "1234567890",
//				"roleCode": "ADMIN"
//			}
//		]
//	}
//
// Response:
//
//	{
//		"workspaceId": "1234567890",
//		"members": [
//			{
//				"departmentCode": "1234567890",
//				"email": "john.doe@example.com",
//				"isCreator": true,
//				"name": "John Doe",
//				"roleCode": "ADMIN"
//				"id": "1234567890",
//				"userId": "1234567890"
//			}
//		]
//	}
func (w *Workspace) InviteMembers(workspaceId string, payload model.InviteMembersRequest) (*model.InviteMembersResponse, error) {
	var (
		data model.Response
		res  model.InviteMembersResponse
		c    = client.GetClient()
	)

	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}

	resp, err := request.R().
		SetHeaders(w.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(payload).
		Post(fmt.Sprintf("%s/%s/%s/member", c.BaseURL, path, workspaceId))
	if err != nil {
		client.Errorf(err, "failed to invite members", c.IsDebug)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to invite members"), "failed to invite members", c.IsDebug)
		return nil, fmt.Errorf("failed to invite members")
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

// UpdateMember updates a member
// Request Params:
// - workspaceId: string
// - memberId: string
// Required headers:
// - Content-Type: application/json
// Body:
//
//	{
//		"roleCode": "ADMIN",
//		"departmentCode": "1234567890"
//	}
//
// Response:
//
//	{
//		"id": "1234567890"
//	}
func (w *Workspace) UpdateMember(workspaceId string, memberId string, payload model.UpdateMemberRequest) (*model.UpdateMemberResponse, error) {
	var (
		data model.Response
		res  model.UpdateMemberResponse
		c    = client.GetClient()
	)

	request, err := c.NewRequest()
	if err != nil {
		return nil, err
	}

	resp, err := request.R().
		SetHeaders(w.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(payload).
		Put(fmt.Sprintf("%s/%s/%s/member/%s", c.BaseURL, path, workspaceId, memberId))
	if err != nil {
		client.Errorf(err, "failed to update member", c.IsDebug)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to update member"), "failed to update member", c.IsDebug)
		return nil, fmt.Errorf("failed to update member")
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

// ChangeStatusMember changes the status of a member
// Request Params:
// - workspaceId: string
// - memberId: string
// Required headers:
// - Content-Type: application/json
// Body:
//
//	{
//		"status": "active | inactive"
//	}
//
// Response:
//
//	{
//		"userId": "1234567890"
//	}
func (w *Workspace) ChangeStatusMember(workspaceId string, memberId, status string) error {
	var (
		data model.Response
		c    = client.GetClient()
	)

	request, err := c.NewRequest()
	if err != nil {
		return err
	}

	resp, err := request.R().
		SetHeaders(w.Headers.ConstructHeaders()).
		SetDebug(c.IsDebug).
		SetBody(model.ChangeStatusMemberRequest{Status: status}).
		Patch(fmt.Sprintf("%s/%s/%s/member/%s/status", c.BaseURL, path, workspaceId, memberId))
	if err != nil {
		client.Errorf(err, "failed to change status member", c.IsDebug)
		return err
	}

	if resp.StatusCode() != 200 {
		client.Errorf(fmt.Errorf("failed to change status member"), "failed to change status member", c.IsDebug)
		return fmt.Errorf("failed to change status member")
	}

	err = json.Unmarshal(resp.Body(), &data)
	if err != nil {
		client.Errorf(err, "failed to unmarshal response", c.IsDebug)
		return err
	}
	return nil
}
