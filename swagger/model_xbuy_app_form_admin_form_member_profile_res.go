/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type XbuyAppFormAdminFormMemberProfileRes struct {
	// 岗位名称
	PostGroup string `json:"postGroup,omitempty"`
	// 角色名称
	RoleGroup string `json:"roleGroup,omitempty"`
	User *XbuyAppFormInputAdminMemberViewModel `json:"user,omitempty"`
	SubUser *XbuyAppFormInputAdminSubmemberViewModel `json:"subUser,omitempty"`
	SysDept *XbuyAppFormInputAdminDeptViewModel `json:"sysDept,omitempty"`
	// 角色列表
	SysRoles []XbuyAppFormInputAdminRoleListModel `json:"sysRoles,omitempty"`
	// 当前岗位
	PostIds int64 `json:"postIds,omitempty"`
	// 当前角色
	RoleIds int64 `json:"roleIds,omitempty"`
}
