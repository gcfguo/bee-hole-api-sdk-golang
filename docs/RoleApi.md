# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRoleDynamicGet**](RoleApi.md#AdminRoleDynamicGet) | **Get** /admin/role/dynamic | 获取动态路由
[**AdminRoleEditPost**](RoleApi.md#AdminRoleEditPost) | **Post** /admin/role/edit | 编辑角色菜单权限
[**AdminRoleListGet**](RoleApi.md#AdminRoleListGet) | **Get** /admin/role/list | 获取角色列表
[**AdminRoleMemberListGet**](RoleApi.md#AdminRoleMemberListGet) | **Get** /admin/role/member_list | 获取角色下的会员列表
[**AdminRoleShopAddPost**](RoleApi.md#AdminRoleShopAddPost) | **Post** /admin/role/shop/add | 增加店铺授权
[**AdminRoleShopDelPost**](RoleApi.md#AdminRoleShopDelPost) | **Post** /admin/role/shop/del | 移除店铺授权
[**AdminRoleShopListPost**](RoleApi.md#AdminRoleShopListPost) | **Post** /admin/role/shop/list | 列表店铺授权

# **AdminRoleDynamicGet**
> InlineResponse200143 AdminRoleDynamicGet(ctx, )
获取动态路由

获取登录用户动态路由

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200143**](inline_response_200_143.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminRoleEditPost**
> InlineResponse2007 AdminRoleEditPost(ctx, body)
编辑角色菜单权限

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormRoleMenuEditReq**](XbuyAppFormAdminFormRoleMenuEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminRoleListGet**
> InlineResponse200144 AdminRoleListGet(ctx, optional)
获取角色列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoleApiAdminRoleListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleApiAdminRoleListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **startTime** | **optional.String**| 开始时间 | 
 **endTime** | **optional.String**| 结束时间 | 
 **status** | **optional.Int32**| 状态 | 
 **deptId** | **optional.Int32**| 部门ID | 
 **mobile** | **optional.Int32**| 手机号 | 
 **username** | **optional.String**| 用户名 | 
 **realname** | **optional.String**| 真实姓名 | 
 **name** | **optional.String**| 岗位名称 | 
 **code** | **optional.String**| 岗位编码 | 

### Return type

[**InlineResponse200144**](inline_response_200_144.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminRoleMemberListGet**
> InlineResponse20069 AdminRoleMemberListGet(ctx, optional)
获取角色下的会员列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RoleApiAdminRoleMemberListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RoleApiAdminRoleMemberListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **startTime** | **optional.String**| 开始时间 | 
 **endTime** | **optional.String**| 结束时间 | 
 **status** | **optional.Int32**| 状态 | 
 **role** | **optional.Int32**| 角色ID | 
 **deptId** | **optional.Int32**| 部门ID | 
 **mobile** | **optional.Int32**| 手机号 | 
 **username** | **optional.String**| 用户名 | 
 **realname** | **optional.String**| 真实姓名 | 
 **name** | **optional.String**| 岗位名称 | 
 **code** | **optional.String**| 岗位编码 | 

### Return type

[**InlineResponse20069**](inline_response_200_69.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminRoleShopAddPost**
> InlineResponse2007 AdminRoleShopAddPost(ctx, body)
增加店铺授权

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormRoleShopAddReq**](XbuyAppFormAdminFormRoleShopAddReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminRoleShopDelPost**
> InlineResponse2007 AdminRoleShopDelPost(ctx, body)
移除店铺授权

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormRoleShopDelReq**](XbuyAppFormAdminFormRoleShopDelReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminRoleShopListPost**
> InlineResponse200145 AdminRoleShopListPost(ctx, body)
列表店铺授权

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormRoleShopListReq**](XbuyAppFormAdminFormRoleShopListReq.md)|  | 

### Return type

[**InlineResponse200145**](inline_response_200_145.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

