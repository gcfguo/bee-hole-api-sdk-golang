/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormAdminFormAddSkuTaxReq struct {
	Id int64 `json:"id,omitempty"`
	SkuId int64 `json:"SkuId"`
	CountryCode string `json:"CountryCode"`
	StateCode string `json:"StateCode,omitempty"`
	Tax float64 `json:"Tax"`
	IsWorker string `json:"is_worker,omitempty"`
}