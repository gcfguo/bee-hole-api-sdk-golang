/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppModelEntityMallScats struct {
	// 标准类目ID
	Cid int64 `json:"cid,omitempty"`
	// 标准类目名称
	Name string `json:"name,omitempty"`
	// 父级类目ID
	ParentId int64 `json:"parent_id,omitempty"`
}
