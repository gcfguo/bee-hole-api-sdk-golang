/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppDaoPreorderResult struct {
	PreviewId string `json:"preview_id,omitempty"`
	PrepayId string `json:"prepay_id,omitempty"`
	TotalCount float64 `json:"total_count,omitempty"`
	Payment float64 `json:"payment,omitempty"`
	Orders map[string]XbuyAppDaoShopAndOrderCont `json:"orders,omitempty"`
	Expresses map[string][]XbuyAppDaoExpress `json:"expresses,omitempty"`
	ExchangeRates map[string]XbuyAppDaoCurrencyConversion `json:"exchange_rates,omitempty"`
	TransCurrency string `json:"trans_currency,omitempty"`
	MemberId int64 `json:"member_id,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	Errs []XbuyXerrorError `json:"errs,omitempty"`
	PromoCodes []string `json:"promo_codes,omitempty"`
	BillingAddress *XbuyAppModelEntityMallBillingAddress `json:"billing_address,omitempty"`
	HasIdentityCheck bool `json:"has_identity_check,omitempty"`
	Discount *XbuyAppDaoDiscount `json:"discount,omitempty"`
}
