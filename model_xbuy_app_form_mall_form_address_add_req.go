/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 添加和更新
type XbuyAppFormMallFormAddressAddReq struct {
	AddressId int32 `json:"address_id,omitempty"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Street string `json:"street"`
	Apartment string `json:"apartment"`
	City string `json:"city"`
	State string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
	Phone string `json:"phone"`
	IsDefault int32 `json:"is_default,omitempty"`
	Sort int32 `json:"sort,omitempty"`
}
