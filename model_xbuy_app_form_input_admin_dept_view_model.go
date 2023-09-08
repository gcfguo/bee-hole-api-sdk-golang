/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormInputAdminDeptViewModel struct {
	// 部门id
	Id int64 `json:"id,omitempty"`
	// 父部门id
	Pid int64 `json:"pid,omitempty"`
	// 祖级列表
	Ancestors string `json:"ancestors,omitempty"`
	// 部门名称
	Name string `json:"name,omitempty"`
	// 部门编码
	Code string `json:"code,omitempty"`
	// 部门类型
	Type_ string `json:"type,omitempty"`
	// 负责人
	Leader string `json:"leader,omitempty"`
	// 联系电话
	Phone string `json:"phone,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 排序
	Sort int32 `json:"sort,omitempty"`
	// 部门状态
	Status string `json:"status,omitempty"`
	// 会员ID
	MemberId int64 `json:"member_id,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
}
