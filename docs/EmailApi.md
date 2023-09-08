# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminEmailBindPost**](EmailApi.md#AdminEmailBindPost) | **Post** /admin/email/bind | 绑定邮箱
[**AdminEmailListPost**](EmailApi.md#AdminEmailListPost) | **Post** /admin/email/list | 获取邮件列表
[**AdminEmailPullPost**](EmailApi.md#AdminEmailPullPost) | **Post** /admin/email/pull | 拉取邮件
[**AdminEmailSendPost**](EmailApi.md#AdminEmailSendPost) | **Post** /admin/email/send | 发送邮件
[**AdminEmailSettingPost**](EmailApi.md#AdminEmailSettingPost) | **Post** /admin/email/setting | 邮件服务设置

# **AdminEmailBindPost**
> InlineResponse2007 AdminEmailBindPost(ctx, body)
绑定邮箱

绑定邮箱

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormEmailBindReq**](XbuyAppFormAdminFormEmailBindReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminEmailListPost**
> InlineResponse20036 AdminEmailListPost(ctx, body)
获取邮件列表

获取邮件列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormEmailListReq**](XbuyAppFormAdminFormEmailListReq.md)|  | 

### Return type

[**InlineResponse20036**](inline_response_200_36.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminEmailPullPost**
> InlineResponse20037 AdminEmailPullPost(ctx, body)
拉取邮件

拉取邮件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormEmailPullReq**](XbuyAppFormAdminFormEmailPullReq.md)|  | 

### Return type

[**InlineResponse20037**](inline_response_200_37.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminEmailSendPost**
> InlineResponse2007 AdminEmailSendPost(ctx, body)
发送邮件

发送邮件

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormEmailSendReq**](XbuyAppFormAdminFormEmailSendReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminEmailSettingPost**
> InlineResponse2007 AdminEmailSettingPost(ctx, body)
邮件服务设置

邮件服务设置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormEmailSettingReq**](XbuyAppFormAdminFormEmailSettingReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

