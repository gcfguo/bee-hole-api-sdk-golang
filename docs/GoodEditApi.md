# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminGoodAddPost**](GoodEditApi.md#AdminGoodAddPost) | **Post** /admin/good/add | 商品新增/编辑
[**AdminGoodFilingAddPost**](GoodEditApi.md#AdminGoodFilingAddPost) | **Post** /admin/good/filing/add | 商品备案
[**AdminGoodOnShelvesPost**](GoodEditApi.md#AdminGoodOnShelvesPost) | **Post** /admin/good/on_shelves | 商品上下架
[**AdminGoodSubAddPost**](GoodEditApi.md#AdminGoodSubAddPost) | **Post** /admin/good/sub/add | 副本商品新增/编辑
[**AdminSkuAddPost**](GoodEditApi.md#AdminSkuAddPost) | **Post** /admin/sku/add | sku新增或编辑
[**AdminSkuPropsAddPost**](GoodEditApi.md#AdminSkuPropsAddPost) | **Post** /admin/sku/props/add | 规格新增/编辑
[**AdminSkuPropsSubAddPost**](GoodEditApi.md#AdminSkuPropsSubAddPost) | **Post** /admin/sku/props/sub/add | 规格副本新增/编辑
[**AdminSkuSubAddPost**](GoodEditApi.md#AdminSkuSubAddPost) | **Post** /admin/sku/sub/add | sku副本新增或编辑
[**AdminSkuValueAddPost**](GoodEditApi.md#AdminSkuValueAddPost) | **Post** /admin/sku/value/add | 规格值新增/编辑
[**AdminSkuValueSubAddPost**](GoodEditApi.md#AdminSkuValueSubAddPost) | **Post** /admin/sku/value/sub/add | 规格值副本新增/编辑
[**AdminTaxRateCreatePost**](GoodEditApi.md#AdminTaxRateCreatePost) | **Post** /admin/tax_rate/create | 创建sku税率
[**AdminTaxRateUpdatePost**](GoodEditApi.md#AdminTaxRateUpdatePost) | **Post** /admin/tax_rate/update | 更新sku税率

# **AdminGoodAddPost**
> InlineResponse20041 AdminGoodAddPost(ctx, body)
商品新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodAddReq**](XbuyAppFormAdminFormGoodAddReq.md)|  | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoodFilingAddPost**
> InlineResponse20043 AdminGoodFilingAddPost(ctx, body)
商品备案

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodFilingsReq**](XbuyAppFormAdminFormGoodFilingsReq.md)|  | 

### Return type

[**InlineResponse20043**](inline_response_200_43.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoodOnShelvesPost**
> InlineResponse2007 AdminGoodOnShelvesPost(ctx, body)
商品上下架

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOnShelvesReq**](XbuyAppFormAdminFormOnShelvesReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoodSubAddPost**
> InlineResponse20041 AdminGoodSubAddPost(ctx, body)
副本商品新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodSubAddReq**](XbuyAppFormAdminFormGoodSubAddReq.md)|  | 

### Return type

[**InlineResponse20041**](inline_response_200_41.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuAddPost**
> InlineResponse200159 AdminSkuAddPost(ctx, body)
sku新增或编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuAddReq**](XbuyAppFormAdminFormSkuAddReq.md)|  | 

### Return type

[**InlineResponse200159**](inline_response_200_159.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuPropsAddPost**
> InlineResponse200161 AdminSkuPropsAddPost(ctx, body)
规格新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuPropsAddReq**](XbuyAppFormAdminFormSkuPropsAddReq.md)|  | 

### Return type

[**InlineResponse200161**](inline_response_200_161.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuPropsSubAddPost**
> InlineResponse200161 AdminSkuPropsSubAddPost(ctx, body)
规格副本新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuPropsSubAddReq**](XbuyAppFormAdminFormSkuPropsSubAddReq.md)|  | 

### Return type

[**InlineResponse200161**](inline_response_200_161.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuSubAddPost**
> InlineResponse200159 AdminSkuSubAddPost(ctx, body)
sku副本新增或编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuSubAddReq**](XbuyAppFormAdminFormSkuSubAddReq.md)|  | 

### Return type

[**InlineResponse200159**](inline_response_200_159.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuValueAddPost**
> InlineResponse200164 AdminSkuValueAddPost(ctx, body)
规格值新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuValueAddReq**](XbuyAppFormAdminFormSkuValueAddReq.md)|  | 

### Return type

[**InlineResponse200164**](inline_response_200_164.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSkuValueSubAddPost**
> InlineResponse200164 AdminSkuValueSubAddPost(ctx, body)
规格值副本新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSkuValueSubAddReq**](XbuyAppFormAdminFormSkuValueSubAddReq.md)|  | 

### Return type

[**InlineResponse200164**](inline_response_200_164.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTaxRateCreatePost**
> InlineResponse200169 AdminTaxRateCreatePost(ctx, body)
创建sku税率

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAddSkuTaxReq**](XbuyAppFormAdminFormAddSkuTaxReq.md)|  | 

### Return type

[**InlineResponse200169**](inline_response_200_169.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTaxRateUpdatePost**
> InlineResponse200169 AdminTaxRateUpdatePost(ctx, body)
更新sku税率

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUpdateSkuTaxReq**](XbuyAppFormAdminFormUpdateSkuTaxReq.md)|  | 

### Return type

[**InlineResponse200169**](inline_response_200_169.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

