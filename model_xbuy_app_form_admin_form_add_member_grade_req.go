/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormAdminFormAddMemberGradeReq struct {
	GradeName string `json:"GradeName"`
	Condition string `json:"Condition,omitempty"`
	ShopId int64 `json:"ShopId"`
	ConsAmt int64 `json:"ConsAmt"`
	TotOrd int64 `json:"TotOrd"`
}
