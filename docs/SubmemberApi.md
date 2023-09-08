# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminSubmemberDeletePost**](SubmemberApi.md#AdminSubmemberDeletePost) | **Post** /admin/submember/delete | 删除子账户
[**AdminSubmemberEditPost**](SubmemberApi.md#AdminSubmemberEditPost) | **Post** /admin/submember/edit | 修改/新增子账户
[**AdminSubmemberEmailUniqueGet**](SubmemberApi.md#AdminSubmemberEmailUniqueGet) | **Get** /admin/submember/email_unique | 邮箱是否唯一
[**AdminSubmemberListGet**](SubmemberApi.md#AdminSubmemberListGet) | **Get** /admin/submember/list | 获取会员列表
[**AdminSubmemberMobileUniqueGet**](SubmemberApi.md#AdminSubmemberMobileUniqueGet) | **Get** /admin/submember/mobile_unique | 手机号是否唯一
[**AdminSubmemberNameUniqueGet**](SubmemberApi.md#AdminSubmemberNameUniqueGet) | **Get** /admin/submember/name_unique | 会员名称是否唯一
[**AdminSubmemberResetPwdPost**](SubmemberApi.md#AdminSubmemberResetPwdPost) | **Post** /admin/submember/reset_pwd | 重置密码
[**AdminSubmemberViewGet**](SubmemberApi.md#AdminSubmemberViewGet) | **Get** /admin/submember/view | 获取指定信息

# **AdminSubmemberDeletePost**
> InlineResponse2007 AdminSubmemberDeletePost(ctx, body)
删除子账户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSubmemberDeleteReq**](XbuyAppFormAdminFormSubmemberDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSubmemberEditPost**
> InlineResponse2007 AdminSubmemberEditPost(ctx, body)
修改/新增子账户

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSubmemberEditReq**](XbuyAppFormAdminFormSubmemberEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSubmemberEmailUniqueGet**
> InlineResponse20022 AdminSubmemberEmailUniqueGet(ctx, email, optional)
邮箱是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| 邮箱 | 
 **optional** | ***SubmemberApiAdminSubmemberEmailUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubmemberApiAdminSubmemberEmailUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 会员ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSubmemberListGet**
> InlineResponse200166 AdminSubmemberListGet(ctx, optional)
获取会员列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubmemberApiAdminSubmemberListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubmemberApiAdminSubmemberListGetOpts struct
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

[**InlineResponse200166**](inline_response_200_166.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSubmemberMobileUniqueGet**
> InlineResponse20022 AdminSubmemberMobileUniqueGet(ctx, mobile, optional)
手机号是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mobile** | **string**| 手机号 | 
 **optional** | ***SubmemberApiAdminSubmemberMobileUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubmemberApiAdminSubmemberMobileUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 会员ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSubmemberNameUniqueGet**
> InlineResponse20022 AdminSubmemberNameUniqueGet(ctx, username, optional)
会员名称是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 会员名称 | 
 **optional** | ***SubmemberApiAdminSubmemberNameUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubmemberApiAdminSubmemberNameUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 会员ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSubmemberResetPwdPost**
> InlineResponse2007 AdminSubmemberResetPwdPost(ctx, body)
重置密码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSubmemberResetPwdReq**](XbuyAppFormAdminFormSubmemberResetPwdReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSubmemberViewGet**
> InlineResponse200167 AdminSubmemberViewGet(ctx, optional)
获取指定信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SubmemberApiAdminSubmemberViewGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SubmemberApiAdminSubmemberViewGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| 会员ID | 

### Return type

[**InlineResponse200167**](inline_response_200_167.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

