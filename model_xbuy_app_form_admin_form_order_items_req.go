/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 订单商品
type XbuyAppFormAdminFormOrderItemsReq struct {
	OrderId string `json:"order_id,omitempty"`
	GoodId int64 `json:"good_id,omitempty"`
	SkuId int64 `json:"sku_id,omitempty"`
}
