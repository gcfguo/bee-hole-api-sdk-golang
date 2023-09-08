/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type XbuyAppFormAdminFormMemberInfoRes struct {
	// 默认用户配置
	DefaultPortalConfig []XbuyAppFormAdminFormPortalConfig `json:"defaultPortalConfig,omitempty"`
	// 应用版本号
	LincenseInfo string `json:"lincenseInfo,omitempty"`
	// 权限
	Permissions []string `json:"permissions,omitempty"`
	// 角色
	Roles []string `json:"roles,omitempty"`
	// 系统公告
	SysNoticeList []XbuyAppModelEntityAdminNotice `json:"sysNoticeList,omitempty"`
	// 用户配置
	UserPortalConfig []XbuyAppFormAdminFormPortalConfig `json:"userPortalConfig,omitempty"`
	User *XbuyAppModelIdentity `json:"user,omitempty"`
}
