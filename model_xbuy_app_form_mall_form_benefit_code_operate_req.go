/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormMallFormBenefitCodeOperateReq struct {
	// 折扣码
	PromotionCode string `json:"promotion_code"`
	// 预支付ID
	AdvPayId string `json:"adv_pay_id"`
	// 预览流水号
	PreviewSerialId string `json:"preview_serial_id"`
	// 赠品信息
	GiftItemsSummary []XbuyAppModelEntityMallOrderItems `json:"gift_items_summary,omitempty"`
	// 折扣码操作 使用/删除
	Operate string `json:"operate"`
}