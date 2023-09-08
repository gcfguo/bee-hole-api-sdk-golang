/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type InlineResponse20062 struct {
	// 状态码
	Code int32 `json:"code,omitempty"`
	// 提示消息
	Message string `json:"message,omitempty"`
	Data *InlineResponse20062Data `json:"data,omitempty"`
	// 服务器时间戳
	Timestamp int64 `json:"timestamp,omitempty"`
	// 唯一请求ID
	ReqId string `json:"req_id,omitempty"`
}
