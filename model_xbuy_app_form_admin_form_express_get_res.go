/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type XbuyAppFormAdminFormExpressGetRes struct {
	// 快递公司编号
	ExpressId int64 `json:"express_id,omitempty"`
	// 快递公司编码
	LogisticsCode string `json:"logistics_code,omitempty"`
	// 快递公司名称
	LogisticsName string `json:"logistics_name,omitempty"`
	// 是否为外国快递公司 1=否2=是
	ForeignExpress int32 `json:"foreign_express,omitempty"`
	// 顺序
	Sort int32 `json:"sort,omitempty"`
}