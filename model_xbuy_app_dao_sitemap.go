/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppDaoSitemap struct {
	Contents []XbuyAppDaoSitemapContent `json:"contents,omitempty"`
	LastDownloaded string `json:"lastDownloaded,omitempty"`
	LastSubmitted string `json:"lastSubmitted,omitempty"`
	Path string `json:"path,omitempty"`
	Type_ string `json:"type,omitempty"`
}
