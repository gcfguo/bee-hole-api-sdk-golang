/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type XbuyAppFormAdminFormAddMemberGroupRes struct {
	Id int64 `json:"id,omitempty"`
	GroupName string `json:"group_name,omitempty"`
	GroupType string `json:"group_type,omitempty"`
	GroupCondition string `json:"group_condition,omitempty"`
	CreateTime *ModelInterface `json:"create_time,omitempty"`
	ShopId int64 `json:"shop_id,omitempty"`
}
