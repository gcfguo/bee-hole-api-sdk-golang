/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppFormAdminFormUpdateMemberGroupReq struct {
	Id int64 `json:"Id"`
	GroupName string `json:"GroupName"`
	GroupType string `json:"GroupType"`
	GroupCondition string `json:"GroupCondition"`
	ShopId int64 `json:"ShopId"`
}
