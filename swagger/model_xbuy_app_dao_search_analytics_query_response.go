/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppDaoSearchAnalyticsQueryResponse struct {
	Url string `json:"url,omitempty"`
	ResponseAggregationType string `json:"responseAggregationType,omitempty"`
	Rows []XbuyAppDaoRows `json:"rows,omitempty"`
}
