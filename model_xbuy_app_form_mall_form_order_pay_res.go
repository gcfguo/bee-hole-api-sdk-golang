/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormMallFormOrderPayRes struct {
	H5 *XbuyAppDomainOrderdomainPaymentH5Result `json:"h5,omitempty"`
	Native *XbuyAppDomainOrderdomainPaymentNativeResult `json:"native,omitempty"`
	Jsapi *XbuyAppDomainOrderdomainPaymentJsapiResult `json:"jsapi,omitempty"`
	LockTtl int64 `json:"lock_ttl,omitempty"`
}