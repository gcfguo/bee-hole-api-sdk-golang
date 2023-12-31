/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppModelEntityMallOrderWaybillExpress struct {
	// 主键
	Id int32 `json:"id,omitempty"`
	// 订单号
	OrderId string `json:"order_id,omitempty"`
	// 店铺ID
	ShopId int64 `json:"shop_id,omitempty"`
	// 快递单号
	ShippingCode string `json:"shipping_code,omitempty"`
	// 快递公司
	ShippingCorp string `json:"shipping_corp,omitempty"`
	// 是否主单，1=>主单，0=>不是主单
	IsMain int32 `json:"is_main,omitempty"`
	// 排序
	Sort int32 `json:"sort,omitempty"`
	// 原数据
	Data string `json:"data,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty"`
}
