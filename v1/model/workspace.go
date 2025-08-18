package model

// CreateWorkspaceRequest is the request body for creating a workspace
// Example:
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
type CreateWorkspaceRequest struct {
	Members   []WorkspaceMember `json:"members"`
	Workspace Workspace         `json:"workspace"`
}

// Workspace is the workspace information
// Example:
//
//	{
//		"id": "1234567890",
//		"name": "Workspace Name",
//		"hostname": "workspace-name.doradora.vn"
//	}
type Workspace struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Hostname string `json:"hostname"`
}

// WorkspaceMember is the member information
// Example:
//
//	{
//		"departmentCode": "1234567890",
//		"email": "john.doe@example.com",
//		"isCreator": true,
//		"name": "John Doe",
//		"roleCode": "ADMIN"
//	}
type WorkspaceMember struct {
	DepartmentCode string `json:"departmentCode"`
	Email          string `json:"email"`
	IsCreator      bool   `json:"isCreator"`
	Name           string `json:"name"`
	RoleCode       string `json:"roleCode"`
}

// WorkspaceInfoResponse is the response body for getting workspace information
// Example:
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
type CreateWorkspaceResponse struct {
	WorkspaceID string          `json:"workspaceId"`
	Members     []MemberReponse `json:"members"`
}

type MemberReponse struct {
	DepartmentCode string `json:"departmentCode"`
	Email          string `json:"email"`
	IsCreator      bool   `json:"isCreator"`
	Name           string `json:"name"`
	RoleCode       string `json:"roleCode"`
	ID             string `json:"id"`
	UserID         string `json:"userId"`
}

// InviteMembersRequest is the request body for inviting members to a workspace
//
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
type InviteMembersRequest struct {
	Members []MemberInvite `json:"members"`
}

type MemberInvite struct {
	Email          string `json:"email"`
	RoleCode       string `json:"roleCode"`
	DepartmentCode string `json:"departmentCode"`
}

// InviteMembersResponse is the response body for inviting members to a workspace
// Example:
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
type InviteMembersResponse struct {
	Members     []MemberReponse `json:"members"`
	WorkspaceID string          `json:"workspaceId"`
}

// UpdateMemberRequest is the request body for updating a member
// Example:
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
type UpdateMemberRequest struct {
	RoleCode       string `json:"roleCode"`
	DepartmentCode string `json:"departmentCode"`
}

type UpdateMemberResponse struct {
	ID string `json:"id"`
}

// ChangeStatusMemberRequest is the request body for changing the status of a member
// Example:
//
//	{
//		"status": "active | inactive"
//	}
type ChangeStatusMemberRequest struct {
	Status string `json:"status"`
}
