/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppDaoSiteEntry struct {
	PermissionLevel string `json:"permissionLevel,omitempty"`
	SiteUrl string `json:"siteUrl,omitempty"`
	Url string `json:"url,omitempty"`
	SitemapResponse []XbuyAppDaoGetSitemapResponse `json:"sitemapResponse,omitempty"`
}