/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 保单推送日志
type XbuyAppFormAdminFormOrderInsuredReq struct {
	OrderId string `json:"order_id"`
	// 当前页码
	Page int32 `json:"page,omitempty"`
	// 每页数量
	Limit int32 `json:"limit,omitempty"`
}
