/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 生成售后单
type XbuyAppFormMallFormSalesOrderCreateReq struct {
	OrderId string `json:"order_id"`
	SkuInfo []XbuyAppFormMallFormAfterSaleSkuInfo `json:"sku_info,omitempty"`
	Reason string `json:"reason"`
}
