/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormAdminFormReviewListReq struct {
	// 当前页码
	Page int32 `json:"page,omitempty"`
	// 每页数量
	Limit int32 `json:"limit,omitempty"`
	OrderId string `json:"order_id,omitempty"`
	GoodsId int64 `json:"goods_id,omitempty"`
	SkuId int64 `json:"sku_id,omitempty"`
	ShopId int64 `json:"shop_id,omitempty"`
}