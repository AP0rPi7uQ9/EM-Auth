package application


const (
	Version_Service = "1.1.0"
)

/*
	User Service
*/
const (
	Auth_NoVerify = 0
	Auth_BasicVerify = 1
	Auth_AdvancedVerify = 2
)

// Client Service name
// 客户端服务名
const (
	Service_Attachment = "em_AttachmentRpcService"
)

// Relationship name
// 关系名
const (
	Relationship_User_Avatar = "user-avatar"
)

// Cache variable
// 缓存变量
var (
	Cache_MenuGetAll = "MenuGetAll"
	Cache_UserGetCurrent = "UserGetCurrent"
	Cache_PermissionGetAll = "PermissionGetAll"
	Cache_UserGetAll = "UserGetAll"
	Cache_RoleGetAll = "RoleGetAll"
)

// Route
var NoAuthRoute = []string{
	"/api/auth/*/user/login",
	"/api/auth/*/user/getCurrent",
}
