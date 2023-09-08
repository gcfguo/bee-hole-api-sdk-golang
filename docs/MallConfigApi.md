# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminMallconfigListPost**](MallConfigApi.md#AdminMallconfigListPost) | **Post** /admin/mallconfig/list | 商城配置列表
[**AdminMallconfigSavePost**](MallConfigApi.md#AdminMallconfigSavePost) | **Post** /admin/mallconfig/save | 保存商城配置

# **AdminMallconfigListPost**
> InlineResponse20060 AdminMallconfigListPost(ctx, body)
商城配置列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormListReq**](XbuyAppFormAdminFormListReq.md)|  | 

### Return type

[**InlineResponse20060**](inline_response_200_60.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMallconfigSavePost**
> InlineResponse2007 AdminMallconfigSavePost(ctx, body)
保存商城配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSaveReq**](XbuyAppFormAdminFormSaveReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

