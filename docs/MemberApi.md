# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminMemberBindMemberToShopPost**](MemberApi.md#AdminMemberBindMemberToShopPost) | **Post** /admin/member/bind_member_to_shop | 绑定会员到商店
[**AdminMemberGetMemberByGroupIdAndShopIdPost**](MemberApi.md#AdminMemberGetMemberByGroupIdAndShopIdPost) | **Post** /admin/member/get_member_by_group_id_and_shop_id | 根据分组ID和商店ID获取会员
[**AdminMemberGetMemberByLevelIdAndShopIdPost**](MemberApi.md#AdminMemberGetMemberByLevelIdAndShopIdPost) | **Post** /admin/member/get_member_by_level_id_and_shop_id | 根据等级ID和商店ID获取会员
[**AdminMemberGetMemberByTagIdAndShopIdPost**](MemberApi.md#AdminMemberGetMemberByTagIdAndShopIdPost) | **Post** /admin/member/get_member_by_tag_id_and_shop_id | 根据标签ID和商店ID获取会员
[**AdminMemberGetMemberListPost**](MemberApi.md#AdminMemberGetMemberListPost) | **Post** /admin/member/get_member_list | 获取会员列表
[**AdminMemberIsShopBlacklistPost**](MemberApi.md#AdminMemberIsShopBlacklistPost) | **Post** /admin/member/is_shop_blacklist | 是否在商店黑名单
[**AdminMemberShopBlacklistAddPost**](MemberApi.md#AdminMemberShopBlacklistAddPost) | **Post** /admin/member/shop_blacklist/add | 将用户加入商店黑名单
[**AdminMemberShopBlacklistPost**](MemberApi.md#AdminMemberShopBlacklistPost) | **Post** /admin/member/shop_blacklist | 指定商店黑名单
[**AdminMemberShopBlacklistRemovePost**](MemberApi.md#AdminMemberShopBlacklistRemovePost) | **Post** /admin/member/shop_blacklist/remove | 将用户移出商店黑名单
[**AdminMemberShopMemberListPost**](MemberApi.md#AdminMemberShopMemberListPost) | **Post** /admin/member/shop_member_list | 商城会员列表
[**AdminMemberUnbindMemberToShopPost**](MemberApi.md#AdminMemberUnbindMemberToShopPost) | **Post** /admin/member/unbind_member_to_shop | 解绑会员到商店
[**MallMemberCaptchaGet**](MemberApi.md#MallMemberCaptchaGet) | **Get** /mall/member/captcha | 验证码
[**MallMemberChangePwdPost**](MemberApi.md#MallMemberChangePwdPost) | **Post** /mall/member/change_pwd | 修改密码
[**MallMemberEmailVerifyPost**](MemberApi.md#MallMemberEmailVerifyPost) | **Post** /mall/member/email/verify | 验证注册邮箱
[**MallMemberLoginPost**](MemberApi.md#MallMemberLoginPost) | **Post** /mall/member/login | 登录
[**MallMemberLogoutPost**](MemberApi.md#MallMemberLogoutPost) | **Post** /mall/member/logout | 注销
[**MallMemberPhoneVerifyPost**](MemberApi.md#MallMemberPhoneVerifyPost) | **Post** /mall/member/phone/verify | 验证注册手机号
[**MallMemberProfilePost**](MemberApi.md#MallMemberProfilePost) | **Post** /mall/member/profile | 个人信息
[**MallMemberRefreshTokenPost**](MemberApi.md#MallMemberRefreshTokenPost) | **Post** /mall/member/refresh_token | 刷新token
[**MallMemberRegisterPost**](MemberApi.md#MallMemberRegisterPost) | **Post** /mall/member/register | 注册
[**MallMemberResetPasswordApplyPost**](MemberApi.md#MallMemberResetPasswordApplyPost) | **Post** /mall/member/reset_password/apply | 重置密码
[**MallMemberResetPasswordConfirmPost**](MemberApi.md#MallMemberResetPasswordConfirmPost) | **Post** /mall/member/reset_password/confirm | 重置密码确认
[**MallMemberUpdatePost**](MemberApi.md#MallMemberUpdatePost) | **Post** /mall/member/update | 更新
[**MallMemberVerifyEmailApplyPost**](MemberApi.md#MallMemberVerifyEmailApplyPost) | **Post** /mall/member/verify_email/apply | 验证邮箱
[**MallMemberVerifyEmailConfirmPost**](MemberApi.md#MallMemberVerifyEmailConfirmPost) | **Post** /mall/member/verify_email/confirm | 验证邮箱确认

# **AdminMemberBindMemberToShopPost**
> InlineResponse20061 AdminMemberBindMemberToShopPost(ctx, body)
绑定会员到商店

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormBindMemberToShopReq**](XbuyAppFormAdminFormBindMemberToShopReq.md)|  | 

### Return type

[**InlineResponse20061**](inline_response_200_61.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGetMemberByGroupIdAndShopIdPost**
> InlineResponse20063 AdminMemberGetMemberByGroupIdAndShopIdPost(ctx, body)
根据分组ID和商店ID获取会员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberByGroupIdAndShopIdReq**](XbuyAppFormAdminFormGetMemberByGroupIdAndShopIdReq.md)|  | 

### Return type

[**InlineResponse20063**](inline_response_200_63.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGetMemberByLevelIdAndShopIdPost**
> InlineResponse20064 AdminMemberGetMemberByLevelIdAndShopIdPost(ctx, body)
根据等级ID和商店ID获取会员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberByLevelIdAndShopIdReq**](XbuyAppFormAdminFormGetMemberByLevelIdAndShopIdReq.md)|  | 

### Return type

[**InlineResponse20064**](inline_response_200_64.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGetMemberByTagIdAndShopIdPost**
> InlineResponse20065 AdminMemberGetMemberByTagIdAndShopIdPost(ctx, body)
根据标签ID和商店ID获取会员

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberByTagIdAndShopIdReq**](XbuyAppFormAdminFormGetMemberByTagIdAndShopIdReq.md)|  | 

### Return type

[**InlineResponse20065**](inline_response_200_65.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGetMemberListPost**
> InlineResponse20066 AdminMemberGetMemberListPost(ctx, body)
获取会员列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberListReq**](XbuyAppFormAdminFormGetMemberListReq.md)|  | 

### Return type

[**InlineResponse20066**](inline_response_200_66.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberIsShopBlacklistPost**
> InlineResponse20068 AdminMemberIsShopBlacklistPost(ctx, body)
是否在商店黑名单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormIsShopBlackListReq**](XbuyAppFormAdminFormIsShopBlackListReq.md)|  | 

### Return type

[**InlineResponse20068**](inline_response_200_68.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberShopBlacklistAddPost**
> InlineResponse2007 AdminMemberShopBlacklistAddPost(ctx, body)
将用户加入商店黑名单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormShopBlackListAddReq**](XbuyAppFormAdminFormShopBlackListAddReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberShopBlacklistPost**
> InlineResponse20074 AdminMemberShopBlacklistPost(ctx, body)
指定商店黑名单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormShopBlackListReq**](XbuyAppFormAdminFormShopBlackListReq.md)|  | 

### Return type

[**InlineResponse20074**](inline_response_200_74.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberShopBlacklistRemovePost**
> InlineResponse2007 AdminMemberShopBlacklistRemovePost(ctx, body)
将用户移出商店黑名单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormShopBlackListRemoveReq**](XbuyAppFormAdminFormShopBlackListRemoveReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberShopMemberListPost**
> InlineResponse20075 AdminMemberShopMemberListPost(ctx, body)
商城会员列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormShopMemberListReq**](XbuyAppFormAdminFormShopMemberListReq.md)|  | 

### Return type

[**InlineResponse20075**](inline_response_200_75.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberUnbindMemberToShopPost**
> InlineResponse20061 AdminMemberUnbindMemberToShopPost(ctx, body)
解绑会员到商店

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUnBindMemberToShopReq**](XbuyAppFormAdminFormUnBindMemberToShopReq.md)|  | 

### Return type

[**InlineResponse20061**](inline_response_200_61.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberCaptchaGet**
> InlineResponse20057 MallMemberCaptchaGet(ctx, )
验证码

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20057**](inline_response_200_57.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberChangePwdPost**
> InlineResponse200201 MallMemberChangePwdPost(ctx, body)
修改密码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormMallMemberChangePassWordReq**](XbuyAppFormMallFormMallMemberChangePassWordReq.md)|  | 

### Return type

[**InlineResponse200201**](inline_response_200_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberEmailVerifyPost**
> InlineResponse2007 MallMemberEmailVerifyPost(ctx, body)
验证注册邮箱

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormVerifyEmailReq**](XbuyAppFormMallFormVerifyEmailReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberLoginPost**
> InlineResponse200202 MallMemberLoginPost(ctx, body)
登录

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormMallMemberLoginReq**](XbuyAppFormMallFormMallMemberLoginReq.md)|  | 

### Return type

[**InlineResponse200202**](inline_response_200_202.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberLogoutPost**
> InlineResponse2007 MallMemberLogoutPost(ctx, body)
注销

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormMallMemberLogoutReq**](XbuyAppFormMallFormMallMemberLogoutReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberPhoneVerifyPost**
> InlineResponse2007 MallMemberPhoneVerifyPost(ctx, body)
验证注册手机号

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormVerifyPhoneApplyReq**](XbuyAppFormMallFormVerifyPhoneApplyReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberProfilePost**
> InlineResponse200203 MallMemberProfilePost(ctx, body)
个人信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormMallMemberProfileReq**](XbuyAppFormMallFormMallMemberProfileReq.md)|  | 

### Return type

[**InlineResponse200203**](inline_response_200_203.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberRefreshTokenPost**
> InlineResponse200202 MallMemberRefreshTokenPost(ctx, body)
刷新token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormMallReFreshTokenReq**](XbuyAppFormMallFormMallReFreshTokenReq.md)|  | 

### Return type

[**InlineResponse200202**](inline_response_200_202.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberRegisterPost**
> InlineResponse200204 MallMemberRegisterPost(ctx, body)
注册

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormMallMemberRegisterReq**](XbuyAppFormMallFormMallMemberRegisterReq.md)|  | 

### Return type

[**InlineResponse200204**](inline_response_200_204.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberResetPasswordApplyPost**
> InlineResponse2007 MallMemberResetPasswordApplyPost(ctx, body)
重置密码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormResetPasswordApplyReq**](XbuyAppFormMallFormResetPasswordApplyReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberResetPasswordConfirmPost**
> InlineResponse2007 MallMemberResetPasswordConfirmPost(ctx, body)
重置密码确认

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormResetPasswordConfirmReq**](XbuyAppFormMallFormResetPasswordConfirmReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberUpdatePost**
> InlineResponse200204 MallMemberUpdatePost(ctx, body)
更新

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormMallMemberUpdateReq**](XbuyAppFormMallFormMallMemberUpdateReq.md)|  | 

### Return type

[**InlineResponse200204**](inline_response_200_204.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberVerifyEmailApplyPost**
> InlineResponse2007 MallMemberVerifyEmailApplyPost(ctx, body)
验证邮箱

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormVerifyEmailApplyReq**](XbuyAppFormMallFormVerifyEmailApplyReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallMemberVerifyEmailConfirmPost**
> InlineResponse2007 MallMemberVerifyEmailConfirmPost(ctx, body)
验证邮箱确认

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormVerifyEmailConfirmReq**](XbuyAppFormMallFormVerifyEmailConfirmReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

