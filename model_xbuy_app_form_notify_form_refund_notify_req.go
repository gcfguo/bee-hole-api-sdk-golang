/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 退款通知
type XbuyAppFormNotifyFormRefundNotifyReq struct {
	OrderId string `json:"order_id,omitempty"`
	Attach string `json:"attach,omitempty"`
	RefundNo string `json:"refund_no,omitempty"`
	RefundedTime string `json:"refunded_time,omitempty"`
}
