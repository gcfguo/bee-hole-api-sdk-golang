/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppModelEntitySysResources struct {
	// api编号
	ResourceId int64 `json:"resource_id,omitempty"`
	// 访问路径
	Path string `json:"path,omitempty"`
	// 访问方法
	Method string `json:"method,omitempty"`
	// 类别
	Cat string `json:"cat,omitempty"`
	// 说明
	Title string `json:"title,omitempty"`
	// 标签
	Tags string `json:"tags,omitempty"`
	// 添加时间
	CreatedAt string `json:"created_at,omitempty"`
}
