/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type XbuyAppFormAdminFormOrderDataRes struct {
	// 30天内付款订单数据
	OrderDataList map[string]int32 `json:"order_data_list,omitempty"`
	// 30天内未付款订单数据
	OrderDataOtherList map[string]int32 `json:"order_data_other_list,omitempty"`
}