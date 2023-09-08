/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type InlineResponse200167Data struct {
	Id int64 `json:"id,omitempty"`
	// 父级账号编号
	Pmemberid int64 `json:"pmemberid,omitempty"`
	// 部门ID
	DeptId int64 `json:"dept_id,omitempty"`
	// 帐号
	Username string `json:"username,omitempty"`
	// 密码
	PasswordHash string `json:"password_hash,omitempty"`
	// 密码盐
	Salt string `json:"salt,omitempty"`
	// 授权令牌
	AuthKey string `json:"auth_key,omitempty"`
	// 密码重置令牌
	PasswordResetToken string `json:"password_reset_token,omitempty"`
	// 1:普通管理员;10超级管理员
	Type_ string `json:"type,omitempty"`
	// 真实姓名
	Realname string `json:"realname,omitempty"`
	// 头像
	Avatar string `json:"avatar,omitempty"`
	// 性别[0:未知;1:男;2:女]
	Sex string `json:"sex,omitempty"`
	// qq
	Qq string `json:"qq,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 生日
	Birthday string `json:"birthday,omitempty"`
	// 省
	ProvinceId int32 `json:"province_id,omitempty"`
	// 城市
	CityId int32 `json:"city_id,omitempty"`
	// 地区
	AreaId int32 `json:"area_id,omitempty"`
	// 默认地址
	Address string `json:"address,omitempty"`
	// 手机号码
	Mobile string `json:"mobile,omitempty"`
	// 家庭号码
	HomePhone string `json:"home_phone,omitempty"`
	// 钉钉机器人token
	DingtalkRobotToken string `json:"dingtalk_robot_token,omitempty"`
	// 访问次数
	VisitCount int32 `json:"visit_count,omitempty"`
	// 最后一次登录时间
	LastTime int32 `json:"last_time,omitempty"`
	// 最后一次登录ip
	LastIp string `json:"last_ip,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 状态
	Status string `json:"status,omitempty"`
	// 创建时间
	CreatedAt string `json:"created_at,omitempty"`
	// 修改时间
	UpdatedAt string `json:"updated_at,omitempty"`
	// 可选岗位
	Posts []XbuyAppFormInputAdminPostListModel `json:"posts,omitempty"`
	// 当前岗位
	PostIds []int64 `json:"postIds,omitempty"`
	// 可选角色
	Roles []XbuyAppFormInputAdminRoleListModel `json:"roles,omitempty"`
	// 当前角色
	RoleIds []int64 `json:"roleIds,omitempty"`
	// 部门名称
	DeptName string `json:"dept_name,omitempty"`
}
