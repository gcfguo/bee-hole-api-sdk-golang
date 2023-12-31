/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppModelEntityMallSubscribes struct {
	// 主键ID
	Id int64 `json:"id,omitempty"`
	// 用户ID
	MemberId int64 `json:"member_id,omitempty"`
	// 店铺ID
	ShopId int64 `json:"shop_id,omitempty"`
	// 用户邮箱
	Email string `json:"email,omitempty"`
	// 添加时间
	CreateTime string `json:"create_time,omitempty"`
}
