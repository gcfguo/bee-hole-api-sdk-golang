/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type XbuyAppServiceMsgserviceMailserviceMessage struct {
	From string `json:"from,omitempty"`
	To []string `json:"to,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body string `json:"body,omitempty"`
	Attachment []string `json:"attachment,omitempty"`
	Date string `json:"date,omitempty"`
}