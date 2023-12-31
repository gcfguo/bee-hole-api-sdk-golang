/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormAdminFormSkuListReq struct {
	// 视图Id
	ViewId int64 `json:"view_id,omitempty"`
	// 店铺编号
	ShopId int64 `json:"shop_id,omitempty"`
	// 页码
	Page int32 `json:"page"`
	// 每页条数
	Size int32 `json:"size"`
	Q string `json:"q,omitempty"`
}
