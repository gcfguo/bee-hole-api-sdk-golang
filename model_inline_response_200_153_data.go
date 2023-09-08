/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type InlineResponse200153Data struct {
	// 门店编号
	ShopId int64 `json:"shop_id,omitempty"`
	// 所属商家
	UserId int64 `json:"user_id,omitempty"`
	// 门店名称
	Title string `json:"title,omitempty"`
	// 门店图标
	Logo string `json:"logo,omitempty"`
	// 地址
	Address string `json:"address,omitempty"`
	// 发货策略,priority(优先发货),default(默认发货)
	ShippingPolicy string `json:"shipping_policy,omitempty"`
	CurrencySetter *XbuyAppDaoCurrencySetter `json:"currency_setter,omitempty"`
	PaymentSetter *XbuyAppDaoPaymentSetter `json:"payment_setter,omitempty"`
	EmailSetter *XbuyAppDaoEmailSetter `json:"email_setter,omitempty"`
	// 管理员邮箱,最多可绑定五个
	AdminEmails string `json:"admin_emails,omitempty"`
	// 店铺授权序列号
	AuthSerialNo int32 `json:"auth_serial_no,omitempty"`
	// 店铺授权秘钥
	AuthSecret string `json:"auth_secret,omitempty"`
	// 店铺资金，单位分
	Money float64 `json:"money,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty"`
}
