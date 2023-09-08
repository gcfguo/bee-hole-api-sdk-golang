/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type InlineResponse20053Data struct {
	OauthUrl string `json:"oauth_url,omitempty"`
	InspectUrlIndex []XbuyAppModelEntityMallSeoGoogle `json:"inspect_url_index,omitempty"`
	TotalCount int32 `json:"total_count,omitempty"`
	// 当前页码
	Page int32 `json:"page,omitempty"`
	// 每页数量
	Limit int32 `json:"limit,omitempty"`
}