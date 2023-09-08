/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppModelEntityMallArticlesSub struct {
	// 文章编号
	ArticleId int64 `json:"article_id,omitempty"`
	// 店铺编号
	ShopId int64 `json:"shop_id,omitempty"`
	// 所属语言
	ViewId int64 `json:"view_id,omitempty"`
	// 类型
	Type_ int32 `json:"type,omitempty"`
	// 类目编号
	CatId int64 `json:"cat_id,omitempty"`
	// 文章标题
	Title string `json:"title,omitempty"`
	// 文章内容
	Content string `json:"content,omitempty"`
	// 文章顺序
	Sort int32 `json:"sort,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// 启用ai：0不启用，1启用，默认1
	Ai int32 `json:"ai,omitempty"`
	// 文章状态
	Status int32 `json:"status,omitempty"`
}