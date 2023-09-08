# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminSkuStockGetPost**](GoodStockApi.md#AdminSkuStockGetPost) | **Post** /admin/sku/stock/get | 获取sku库存
[**AdminSkuStockUpdatePost**](GoodStockApi.md#AdminSkuStockUpdatePost) | **Post** /admin/sku/stock/update | 更新sku库存

# **AdminSkuStockGetPost**
> InlineResponse200163 AdminSkuStockGetPost(ctx, body)
获取sku库存

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetSkuStockReq**](XbuyAppFormAdminFormGetSkuStockReq.md)|  | 

### Return type

[**InlineResponse200163**](inline_response_200_163.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuStockUpdatePost**
> InlineResponse2007 AdminSkuStockUpdatePost(ctx, body)
更新sku库存

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUpdateSkuStockReq**](XbuyAppFormAdminFormUpdateSkuStockReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

