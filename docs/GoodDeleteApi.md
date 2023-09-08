# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminGoodDeletePost**](GoodDeleteApi.md#AdminGoodDeletePost) | **Post** /admin/good/delete | 删除商品
[**AdminGoodInfoValueDeleteOnePost**](GoodDeleteApi.md#AdminGoodInfoValueDeleteOnePost) | **Post** /admin/good/info/value/delete/one | 删除商品参数值
[**AdminSkuDeletePost**](GoodDeleteApi.md#AdminSkuDeletePost) | **Post** /admin/sku/delete | 删除sku
[**AdminSkuPropDeletePost**](GoodDeleteApi.md#AdminSkuPropDeletePost) | **Post** /admin/sku/prop/delete | 删除规格
[**AdminSkuValueDeletePost**](GoodDeleteApi.md#AdminSkuValueDeletePost) | **Post** /admin/sku/value/delete | 删除规格值
[**AdminTaxRateDeletePost**](GoodDeleteApi.md#AdminTaxRateDeletePost) | **Post** /admin/tax_rate/delete | 删除sku税率

# **AdminGoodDeletePost**
> InlineResponse2007 AdminGoodDeletePost(ctx, body)
删除商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodDeleteReq**](XbuyAppFormAdminFormGoodDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoodInfoValueDeleteOnePost**
> InlineResponse2007 AdminGoodInfoValueDeleteOnePost(ctx, body)
删除商品参数值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodInfoValueDeleteOneReq**](XbuyAppFormAdminFormGoodInfoValueDeleteOneReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuDeletePost**
> InlineResponse200159 AdminSkuDeletePost(ctx, body)
删除sku

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuDeleteReq**](XbuyAppFormAdminFormSkuDeleteReq.md)|  | 

### Return type

[**InlineResponse200159**](inline_response_200_159.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuPropDeletePost**
> InlineResponse200159 AdminSkuPropDeletePost(ctx, body)
删除规格

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuPropsDeleteReq**](XbuyAppFormAdminFormSkuPropsDeleteReq.md)|  | 

### Return type

[**InlineResponse200159**](inline_response_200_159.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuValueDeletePost**
> InlineResponse200159 AdminSkuValueDeletePost(ctx, body)
删除规格值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuValueDeleteReq**](XbuyAppFormAdminFormSkuValueDeleteReq.md)|  | 

### Return type

[**InlineResponse200159**](inline_response_200_159.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTaxRateDeletePost**
> InlineResponse200169 AdminTaxRateDeletePost(ctx, body)
删除sku税率

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDeleteSkuTaxReq**](XbuyAppFormAdminFormDeleteSkuTaxReq.md)|  | 

### Return type

[**InlineResponse200169**](inline_response_200_169.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

