/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppDaoCurrencyConversion struct {
	BaseCurrency string `json:"base_currency,omitempty"`
	TargetCurrency *XbuyAppDaoExchangeRate `json:"target_currency,omitempty"`
}
