/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type InlineResponse20035Data struct {
	// 域名列表
	List []XbuyAppFormAdminFormDomains `json:"list,omitempty"`
	// 当前页码
	Page int32 `json:"page,omitempty"`
	// 每页数量
	Limit int32 `json:"limit,omitempty"`
	// 全部数据量
	TotalCount int64 `json:"total_count,omitempty"`
}
