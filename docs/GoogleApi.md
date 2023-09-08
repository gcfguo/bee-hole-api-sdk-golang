# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminGoogleAddPost**](GoogleApi.md#AdminGoogleAddPost) | **Post** /admin/google/add | 添加域名
[**AdminGoogleListPost**](GoogleApi.md#AdminGoogleListPost) | **Post** /admin/google/list | 获取谷歌站点列表
[**AdminGoogleOauthCallbackGet**](GoogleApi.md#AdminGoogleOauthCallbackGet) | **Get** /admin/google/oauth/callback | 谷歌授权回调
[**AdminGoogleSearchPost**](GoogleApi.md#AdminGoogleSearchPost) | **Post** /admin/google/search | 获取网站搜索分析数据
[**AdminGoogleSubmitPost**](GoogleApi.md#AdminGoogleSubmitPost) | **Post** /admin/google/submit | 提交网站地图
[**AdminGoogleUrlindexPost**](GoogleApi.md#AdminGoogleUrlindexPost) | **Post** /admin/google/urlindex | 获取网址索引

# **AdminGoogleAddPost**
> InlineResponse20050 AdminGoogleAddPost(ctx, body)
添加域名

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoogleAddReq**](XbuyAppFormAdminFormGoogleAddReq.md)|  | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoogleListPost**
> InlineResponse20051 AdminGoogleListPost(ctx, body)
获取谷歌站点列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoogleListReq**](XbuyAppFormAdminFormGoogleListReq.md)|  | 

### Return type

[**InlineResponse20051**](inline_response_200_51.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoogleOauthCallbackGet**
> InlineResponse2007 AdminGoogleOauthCallbackGet(ctx, code, optional)
谷歌授权回调

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**|  | 
 **optional** | ***GoogleApiAdminGoogleOauthCallbackGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GoogleApiAdminGoogleOauthCallbackGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **state** | **optional.String**|  | 
 **scope** | **optional.String**|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoogleSearchPost**
> InlineResponse20052 AdminGoogleSearchPost(ctx, body)
获取网站搜索分析数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoogleSearchAnalyticsReq**](XbuyAppFormAdminFormGoogleSearchAnalyticsReq.md)|  | 

### Return type

[**InlineResponse20052**](inline_response_200_52.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoogleSubmitPost**
> InlineResponse20050 AdminGoogleSubmitPost(ctx, body)
提交网站地图

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoogleSubmitReq**](XbuyAppFormAdminFormGoogleSubmitReq.md)|  | 

### Return type

[**InlineResponse20050**](inline_response_200_50.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoogleUrlindexPost**
> InlineResponse20053 AdminGoogleUrlindexPost(ctx, body)
获取网址索引

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoogleUrlIndexReq**](XbuyAppFormAdminFormGoogleUrlIndexReq.md)|  | 

### Return type

[**InlineResponse20053**](inline_response_200_53.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

