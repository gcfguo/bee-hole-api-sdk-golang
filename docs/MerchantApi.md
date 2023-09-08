# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminMemberBlacklistAddPost**](MerchantApi.md#AdminMemberBlacklistAddPost) | **Post** /admin/member/blacklist/add | 将用户加入黑名单(全局）
[**AdminMemberBlacklistPost**](MerchantApi.md#AdminMemberBlacklistPost) | **Post** /admin/member/blacklist | 全局黑名单
[**AdminMemberBlacklistRemovePost**](MerchantApi.md#AdminMemberBlacklistRemovePost) | **Post** /admin/member/blacklist/remove | 将用户移出黑名单(全局)
[**AdminMemberDeletePost**](MerchantApi.md#AdminMemberDeletePost) | **Post** /admin/member/delete | 删除会员
[**AdminMemberEditPost**](MerchantApi.md#AdminMemberEditPost) | **Post** /admin/member/edit | 修改/新增会员
[**AdminMemberEmailUniqueGet**](MerchantApi.md#AdminMemberEmailUniqueGet) | **Get** /admin/member/email_unique | 邮箱是否唯一
[**AdminMemberEmailVerifyPost**](MerchantApi.md#AdminMemberEmailVerifyPost) | **Post** /admin/member/email/verify | 验证注册邮箱
[**AdminMemberInfoGet**](MerchantApi.md#AdminMemberInfoGet) | **Get** /admin/member/info | 获取登录用户信息
[**AdminMemberListGet**](MerchantApi.md#AdminMemberListGet) | **Get** /admin/member/list | 获取会员列表
[**AdminMemberMaxSortGet**](MerchantApi.md#AdminMemberMaxSortGet) | **Get** /admin/member/max_sort | 会员最大排序
[**AdminMemberMobileUniqueGet**](MerchantApi.md#AdminMemberMobileUniqueGet) | **Get** /admin/member/mobile_unique | 手机号是否唯一
[**AdminMemberNameUniqueGet**](MerchantApi.md#AdminMemberNameUniqueGet) | **Get** /admin/member/name_unique | 会员名称是否唯一
[**AdminMemberProfileGet**](MerchantApi.md#AdminMemberProfileGet) | **Get** /admin/member/profile | 获取登录用户的基本信息
[**AdminMemberRegisterPost**](MerchantApi.md#AdminMemberRegisterPost) | **Post** /admin/member/register | 用户注册
[**AdminMemberResetPwdPost**](MerchantApi.md#AdminMemberResetPwdPost) | **Post** /admin/member/reset_pwd | 重置密码
[**AdminMemberUpdateAvatarPost**](MerchantApi.md#AdminMemberUpdateAvatarPost) | **Post** /admin/member/update_avatar | 更新会员头像
[**AdminMemberUpdateProfilePost**](MerchantApi.md#AdminMemberUpdateProfilePost) | **Post** /admin/member/update_profile | 更新会员资料
[**AdminMemberUpdatePwdPost**](MerchantApi.md#AdminMemberUpdatePwdPost) | **Post** /admin/member/update_pwd | 重置密码
[**AdminMemberViewGet**](MerchantApi.md#AdminMemberViewGet) | **Get** /admin/member/view | 获取指定信息

# **AdminMemberBlacklistAddPost**
> InlineResponse2007 AdminMemberBlacklistAddPost(ctx, body)
将用户加入黑名单(全局）

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormBlackListAddReq**](XbuyAppFormAdminFormBlackListAddReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberBlacklistPost**
> InlineResponse20062 AdminMemberBlacklistPost(ctx, body)
全局黑名单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormBlackListReq**](XbuyAppFormAdminFormBlackListReq.md)|  | 

### Return type

[**InlineResponse20062**](inline_response_200_62.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberBlacklistRemovePost**
> InlineResponse2007 AdminMemberBlacklistRemovePost(ctx, body)
将用户移出黑名单(全局)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormBlackListRemoveReq**](XbuyAppFormAdminFormBlackListRemoveReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberDeletePost**
> InlineResponse2007 AdminMemberDeletePost(ctx, body)
删除会员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMemberDeleteReq**](XbuyAppFormAdminFormMemberDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberEditPost**
> InlineResponse2007 AdminMemberEditPost(ctx, body)
修改/新增会员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMemberEditReq**](XbuyAppFormAdminFormMemberEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberEmailUniqueGet**
> InlineResponse20022 AdminMemberEmailUniqueGet(ctx, email, optional)
邮箱是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **email** | **string**| 邮箱 | 
 **optional** | ***MerchantApiAdminMemberEmailUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiAdminMemberEmailUniqueGetOpts struct
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

# **AdminMemberEmailVerifyPost**
> InlineResponse2007 AdminMemberEmailVerifyPost(ctx, body)
验证注册邮箱

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormVerifyEmailApplyReq**](XbuyAppFormAdminFormVerifyEmailApplyReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberInfoGet**
> InlineResponse20067 AdminMemberInfoGet(ctx, optional)
获取登录用户信息

获取管理后台的登录用户信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MerchantApiAdminMemberInfoGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiAdminMemberInfoGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shopId** | **optional.Int64**| 店铺ID | 

### Return type

[**InlineResponse20067**](inline_response_200_67.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberListGet**
> InlineResponse20069 AdminMemberListGet(ctx, optional)
获取会员列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MerchantApiAdminMemberListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiAdminMemberListGetOpts struct
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

[**InlineResponse20069**](inline_response_200_69.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberMaxSortGet**
> InlineResponse20021 AdminMemberMaxSortGet(ctx, optional)
会员最大排序

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MerchantApiAdminMemberMaxSortGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiAdminMemberMaxSortGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| 会员ID | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberMobileUniqueGet**
> InlineResponse20022 AdminMemberMobileUniqueGet(ctx, mobile, optional)
手机号是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mobile** | **string**| 手机号 | 
 **optional** | ***MerchantApiAdminMemberMobileUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiAdminMemberMobileUniqueGetOpts struct
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

# **AdminMemberNameUniqueGet**
> InlineResponse20022 AdminMemberNameUniqueGet(ctx, username, optional)
会员名称是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **username** | **string**| 会员名称 | 
 **optional** | ***MerchantApiAdminMemberNameUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiAdminMemberNameUniqueGetOpts struct
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

# **AdminMemberProfileGet**
> InlineResponse20070 AdminMemberProfileGet(ctx, )
获取登录用户的基本信息

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20070**](inline_response_200_70.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberRegisterPost**
> InlineResponse2007 AdminMemberRegisterPost(ctx, body)
用户注册

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMemberNewRegisterReq**](XbuyAppFormAdminFormMemberNewRegisterReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberResetPwdPost**
> InlineResponse2007 AdminMemberResetPwdPost(ctx, body)
重置密码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMemberResetPwdReq**](XbuyAppFormAdminFormMemberResetPwdReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberUpdateAvatarPost**
> InlineResponse2007 AdminMemberUpdateAvatarPost(ctx, body)
更新会员头像

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMemberUpdateAvatarReq**](XbuyAppFormAdminFormMemberUpdateAvatarReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberUpdateProfilePost**
> InlineResponse2007 AdminMemberUpdateProfilePost(ctx, body)
更新会员资料

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMemberUpdateProfileReq**](XbuyAppFormAdminFormMemberUpdateProfileReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberUpdatePwdPost**
> InlineResponse2007 AdminMemberUpdatePwdPost(ctx, body)
重置密码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMemberUpdatePwdReq**](XbuyAppFormAdminFormMemberUpdatePwdReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberViewGet**
> InlineResponse20076 AdminMemberViewGet(ctx, optional)
获取指定信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MerchantApiAdminMemberViewGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MerchantApiAdminMemberViewGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| 会员ID | 

### Return type

[**InlineResponse20076**](inline_response_200_76.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

