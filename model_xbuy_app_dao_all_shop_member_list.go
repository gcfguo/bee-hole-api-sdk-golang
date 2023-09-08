/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppDaoAllShopMemberList struct {
	// 用户ID
	MemberId int32 `json:"member_id,omitempty"`
	// 用户昵称
	Nick string `json:"nick,omitempty"`
	// 姓
	LastName string `json:"last_name,omitempty"`
	// 名
	FirstName string `json:"first_name,omitempty"`
	// 用户邮箱
	Email string `json:"email,omitempty"`
	// 密码
	Password string `json:"password,omitempty"`
	// 手机号国家码
	PhoneCountryCode string `json:"phone_country_code,omitempty"`
	// 用户手机号
	Phone string `json:"phone,omitempty"`
	// 用户头像
	Avatar string `json:"avatar,omitempty"`
	// 密码盐
	Salt string `json:"salt,omitempty"`
	// 注册时间
	CreateTime string `json:"create_time,omitempty"`
	// 注册IP
	LastIp string `json:"last_ip,omitempty"`
	// 全局ban，0没ban，1ban了
	IsBanned int32 `json:"is_banned,omitempty"`
	// 邮箱验证状态0(未验证)1(已验证)
	EmailVerifyStatus int32 `json:"email_verify_status,omitempty"`
	ShopId string `json:"shop_id,omitempty"`
}
