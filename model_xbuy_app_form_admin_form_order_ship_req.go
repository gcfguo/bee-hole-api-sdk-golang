/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 订单发货
type XbuyAppFormAdminFormOrderShipReq struct {
	OrderId string `json:"order_id"`
	LogisticsName string `json:"logistics_name"`
	ShipmentId string `json:"shipment_id"`
	IsUpdate bool `json:"is_update,omitempty"`
}
