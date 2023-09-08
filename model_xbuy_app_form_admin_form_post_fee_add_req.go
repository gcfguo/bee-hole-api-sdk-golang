/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormAdminFormPostFeeAddReq struct {
	// 运费模板编号
	FareId int64 `json:"fare_id,omitempty"`
	// 店铺编号
	ShopId int64 `json:"shop_id,omitempty"`
	// 模板名称
	FareName string `json:"fare_name,omitempty"`
	// 配送区域
	Region string `json:"region,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty"`
}