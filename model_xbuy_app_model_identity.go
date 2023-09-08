/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppModelIdentity struct {
	// 会员ID
	Id int64 `json:"id,omitempty"`
	// 店铺ID
	ShopId int64 `json:"shop_id,omitempty"`
	// 子账户ID
	SubMemberId int64 `json:"sub_member_id,omitempty"`
	// 主账号ID
	PmemberId int64 `json:"pmember_id,omitempty"`
	// 用户名
	Username string `json:"username,omitempty"`
	// 昵称
	Realname string `json:"realname,omitempty"`
	// 头像
	Avatar string `json:"avatar,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 手机号码
	Mobile string `json:"mobile,omitempty"`
	// 访问次数
	VisitCount int32 `json:"visit_count,omitempty"`
	// 最后一次登录时间
	LastTime int32 `json:"last_time,omitempty"`
	// 最后一次登录ip
	LastIp string `json:"last_ip,omitempty"`
	// 权限
	Role int64 `json:"role,omitempty"`
	// 登录有效期截止时间戳
	Exp int64 `json:"exp,omitempty"`
	// 登录有效期
	Expires int64 `json:"expires,omitempty"`
	// 登录应用
	App string `json:"app,omitempty"`
	// 语言
	Lang string `json:"lang,omitempty"`
	// 时区
	Zone string `json:"zone,omitempty"`
}
