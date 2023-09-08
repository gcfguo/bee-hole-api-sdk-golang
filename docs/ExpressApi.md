# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminExpressAddPost**](ExpressApi.md#AdminExpressAddPost) | **Post** /admin/express/add | 快递公司新增/编辑
[**AdminExpressDeletePost**](ExpressApi.md#AdminExpressDeletePost) | **Post** /admin/express/delete | 快递公司删除
[**AdminExpressGetPost**](ExpressApi.md#AdminExpressGetPost) | **Post** /admin/express/get | 查询快递公司
[**AdminExpressListPost**](ExpressApi.md#AdminExpressListPost) | **Post** /admin/express/list | 快递公司列表

# **AdminExpressAddPost**
> InlineResponse20038 AdminExpressAddPost(ctx, body)
快递公司新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormExpressAddReq**](XbuyAppFormAdminFormExpressAddReq.md)|  | 

### Return type

[**InlineResponse20038**](inline_response_200_38.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminExpressDeletePost**
> InlineResponse2007 AdminExpressDeletePost(ctx, body)
快递公司删除

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormExpressDeleteReq**](XbuyAppFormAdminFormExpressDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminExpressGetPost**
> InlineResponse20039 AdminExpressGetPost(ctx, body)
查询快递公司

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormExpressGetReq**](XbuyAppFormAdminFormExpressGetReq.md)|  | 

### Return type

[**InlineResponse20039**](inline_response_200_39.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminExpressListPost**
> InlineResponse20040 AdminExpressListPost(ctx, body)
快递公司列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormExpressListReq**](XbuyAppFormAdminFormExpressListReq.md)|  | 

### Return type

[**InlineResponse20040**](inline_response_200_40.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

