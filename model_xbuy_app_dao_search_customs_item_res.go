/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppDaoSearchCustomsItemRes struct {
	Flag string `json:"flag,omitempty"`
	Code string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	TotalCount int64 `json:"total_count,omitempty"`
	Data []ModelInterface `json:"data,omitempty"`
}
