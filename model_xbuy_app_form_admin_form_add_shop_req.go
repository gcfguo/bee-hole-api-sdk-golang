/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormAdminFormAddShopReq struct {
	// 店铺id
	ShopId int64 `json:"shop_id,omitempty"`
	// 店铺名称
	Title string `json:"title"`
	// 店铺地址
	Address string `json:"address,omitempty"`
}
