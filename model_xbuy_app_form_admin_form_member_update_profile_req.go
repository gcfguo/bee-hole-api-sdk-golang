/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormAdminFormMemberUpdateProfileReq struct {
	// 手机号
	Mobile int32 `json:"mobile,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 真实姓名
	Realname string `json:"realname,omitempty"`
	// 性别
	Sex string `json:"sex,omitempty"`
}
