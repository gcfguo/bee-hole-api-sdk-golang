/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormMallFormMallMemberUpdateReq struct {
	// 邮箱
	Email string `json:"email,omitempty"`
	// 名
	FirstName string `json:"first_name"`
	// 姓
	LastName string `json:"last_name"`
	// 昵称
	Nick string `json:"nick,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
	// 头像地址
	Avatar string `json:"avatar,omitempty"`
}