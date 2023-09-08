# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminGoodInfoListPost**](GoodViewApi.md#AdminGoodInfoListPost) | **Post** /admin/good/info/list | 标准类目列表
[**AdminGoodInfoPost**](GoodViewApi.md#AdminGoodInfoPost) | **Post** /admin/good/info | 商品详情
[**AdminGoodListPost**](GoodViewApi.md#AdminGoodListPost) | **Post** /admin/good/list | 商品列表
[**AdminSkuListPost**](GoodViewApi.md#AdminSkuListPost) | **Post** /admin/sku/list | sku列表
[**AdminSkuPropsListPost**](GoodViewApi.md#AdminSkuPropsListPost) | **Post** /admin/sku/props/list | 规格列表
[**AdminTaxRateGetRatePost**](GoodViewApi.md#AdminTaxRateGetRatePost) | **Post** /admin/tax_rate/get_rate | 获取sku税率
[**AdminTaxRateRateListPost**](GoodViewApi.md#AdminTaxRateRateListPost) | **Post** /admin/tax_rate/rate_list | sku税率列表
[**AdminTaxRateReadPost**](GoodViewApi.md#AdminTaxRateReadPost) | **Post** /admin/tax_rate/read | 查询sku税率

# **AdminGoodInfoListPost**
> InlineResponse20047 AdminGoodInfoListPost(ctx, body)
标准类目列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodInfoListReq**](XbuyAppFormAdminFormGoodInfoListReq.md)|  | 

### Return type

[**InlineResponse20047**](inline_response_200_47.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoodInfoPost**
> InlineResponse20045 AdminGoodInfoPost(ctx, body)
商品详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodInfoReq**](XbuyAppFormAdminFormGoodInfoReq.md)|  | 

### Return type

[**InlineResponse20045**](inline_response_200_45.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoodListPost**
> InlineResponse20048 AdminGoodListPost(ctx, body)
商品列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodListReq**](XbuyAppFormAdminFormGoodListReq.md)|  | 

### Return type

[**InlineResponse20048**](inline_response_200_48.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuListPost**
> InlineResponse200160 AdminSkuListPost(ctx, body)
sku列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuListReq**](XbuyAppFormAdminFormSkuListReq.md)|  | 

### Return type

[**InlineResponse200160**](inline_response_200_160.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuPropsListPost**
> InlineResponse200162 AdminSkuPropsListPost(ctx, body)
规格列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuPropsListReq**](XbuyAppFormAdminFormSkuPropsListReq.md)|  | 

### Return type

[**InlineResponse200162**](inline_response_200_162.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTaxRateGetRatePost**
> InlineResponse200170 AdminTaxRateGetRatePost(ctx, body)
获取sku税率

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetTaxRateByIdReq**](XbuyAppFormAdminFormGetTaxRateByIdReq.md)|  | 

### Return type

[**InlineResponse200170**](inline_response_200_170.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTaxRateRateListPost**
> InlineResponse200171 AdminTaxRateRateListPost(ctx, body)
sku税率列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuTaxRateListReq**](XbuyAppFormAdminFormSkuTaxRateListReq.md)|  | 

### Return type

[**InlineResponse200171**](inline_response_200_171.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTaxRateReadPost**
> InlineResponse200172 AdminTaxRateReadPost(ctx, body)
查询sku税率

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormReadSkuTaxReq**](XbuyAppFormAdminFormReadSkuTaxReq.md)|  | 

### Return type

[**InlineResponse200172**](inline_response_200_172.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

