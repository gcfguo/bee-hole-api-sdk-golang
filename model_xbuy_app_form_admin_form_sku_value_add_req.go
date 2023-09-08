/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormAdminFormSkuValueAddReq struct {
	// 商品编号
	GoodId int64 `json:"good_id"`
	// 视图Id
	ViewId int64 `json:"view_id,omitempty"`
	// 商品规格编号
	PropId int64 `json:"prop_id"`
	// 商品规格值Id
	ValueId int64 `json:"value_id,omitempty"`
	// 商品规格值
	PropValue string `json:"prop_value,omitempty"`
	// 是否启用ai翻译
	Ai int32 `json:"ai,omitempty"`
}