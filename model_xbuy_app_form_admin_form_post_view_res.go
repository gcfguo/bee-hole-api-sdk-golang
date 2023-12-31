/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type XbuyAppFormAdminFormPostViewRes struct {
	// 岗位ID
	Id int64 `json:"id,omitempty"`
	// 岗位编码
	Code string `json:"code,omitempty"`
	// 岗位名称
	Name string `json:"name,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 显示顺序
	Sort int32 `json:"sort,omitempty"`
	// 状态
	Status string `json:"status,omitempty"`
	MemberId int64 `json:"member_id,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 更新时间
	UpdatedAt string `json:"updated_at,omitempty"`
}
