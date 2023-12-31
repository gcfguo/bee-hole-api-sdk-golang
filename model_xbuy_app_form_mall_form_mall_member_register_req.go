/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormMallFormMallMemberRegisterReq struct {
	// 手机号
	Password string `json:"password"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 名
	FirstName string `json:"first_name"`
	// 姓
	LastName string `json:"last_name"`
	// 昵称
	Nick string `json:"nick"`
	// 手机号国家码
	PhoneCountryCode string `json:"phone_country_code,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
	// 验证码
	Code string `json:"code,omitempty"`
	// 头像地址
	Avatar string `json:"avatar,omitempty"`
}
