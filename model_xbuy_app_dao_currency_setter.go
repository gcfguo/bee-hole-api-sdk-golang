/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppDaoCurrencySetter struct {
	QuoteCurrency string `json:"quote_currency,omitempty"`
	ExchangeRates []XbuyAppDaoExchangeRate `json:"exchange_rates,omitempty"`
	IsDefault bool `json:"is_default,omitempty"`
}
