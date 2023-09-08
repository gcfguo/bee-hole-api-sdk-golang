/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppDaoModeListContent struct {
	CarryId int64 `json:"carry_id,omitempty"`
	// 店铺编号
	ShopId int64 `json:"shop_id,omitempty"`
	// 运费模板id
	FareId int64 `json:"fare_id,omitempty"`
	// 运费方式名称
	CarryName string `json:"carry_name,omitempty"`
	// 补充说明
	Remark string `json:"remark,omitempty"`
	// 计费方式 1=固定计费 2=首续重 3=首续件
	BillingMethod int32 `json:"billing_method,omitempty"`
	// 固定运费，专属billing_methods=1时使用
	FixedFreight float64 `json:"fixed_freight,omitempty"`
	// 首重、首件
	FirstPiece string `json:"first_piece,omitempty"`
	// 首重(首件)费用
	FirstAmount float64 `json:"first_amount,omitempty"`
	// 续重、续件
	SecondPiece string `json:"second_piece,omitempty"`
	// 续重(续件)费用
	SecondAmount float64 `json:"second_amount,omitempty"`
	// 是否启用额外条件 条件区间 0=关闭 1=开启
	ExtraCondition int32 `json:"extra_condition,omitempty"`
	// 条件区间 1=按商品总价 2=按包裹重量 3=按商品件数
	ConditionInterval int32 `json:"condition_interval,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty"`
	SubMode []XbuyAppModelEntityMallCarrySubMode `json:"sub_mode,omitempty"`
}