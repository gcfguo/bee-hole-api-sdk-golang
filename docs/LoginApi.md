# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminLoginCaptchaGet**](LoginApi.md#AdminLoginCaptchaGet) | **Get** /admin/login/captcha | 获取登录验证码
[**AdminLoginLogoutPost**](LoginApi.md#AdminLoginLogoutPost) | **Post** /admin/login/logout | 注销登录
[**AdminLoginSignPost**](LoginApi.md#AdminLoginSignPost) | **Post** /admin/login/sign | 商户登录
[**AdminLoginSubAuthPost**](LoginApi.md#AdminLoginSubAuthPost) | **Post** /admin/login/sub/auth | 子账户授权
[**AdminLoginSubSignPost**](LoginApi.md#AdminLoginSubSignPost) | **Post** /admin/login/sub/sign | 商户子账号登录

# **AdminLoginCaptchaGet**
> InlineResponse20057 AdminLoginCaptchaGet(ctx, )
获取登录验证码

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

# **AdminLoginLogoutPost**
> InlineResponse2007 AdminLoginLogoutPost(ctx, body)
注销登录

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormLoginLogoutReq**](XbuyAppFormAdminFormLoginLogoutReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminLoginSignPost**
> InlineResponse20058 AdminLoginSignPost(ctx, body)
商户登录

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormLoginReq**](XbuyAppFormAdminFormLoginReq.md)|  | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminLoginSubAuthPost**
> InlineResponse20058 AdminLoginSubAuthPost(ctx, body)
子账户授权

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormLoginTokenReq**](XbuyAppFormAdminFormLoginTokenReq.md)|  | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminLoginSubSignPost**
> InlineResponse20058 AdminLoginSubSignPost(ctx, body)
商户子账号登录

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSubLoginReq**](XbuyAppFormAdminFormSubLoginReq.md)|  | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

