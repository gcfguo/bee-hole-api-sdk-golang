/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 创建优先发货策略
type XbuyAppFormAdminFormPriorityPolicyCreateReq struct {
	WarehouseId int32 `json:"warehouse_id"`
	Weight int32 `json:"weight,omitempty"`
}
