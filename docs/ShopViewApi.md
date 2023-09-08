# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminShopViewAddPost**](ShopViewApi.md#AdminShopViewAddPost) | **Post** /admin/shop/view/add | 修改/新增
[**AdminShopViewCloseAutoTransPost**](ShopViewApi.md#AdminShopViewCloseAutoTransPost) | **Post** /admin/shop/view/close_auto_trans | 关闭自动翻译
[**AdminShopViewInfoPost**](ShopViewApi.md#AdminShopViewInfoPost) | **Post** /admin/shop/view/info | 详情
[**AdminShopViewListPost**](ShopViewApi.md#AdminShopViewListPost) | **Post** /admin/shop/view/list | 列表
[**AdminShopViewOpenAutoTransPost**](ShopViewApi.md#AdminShopViewOpenAutoTransPost) | **Post** /admin/shop/view/open_auto_trans | 开启自动翻译

# **AdminShopViewAddPost**
> InlineResponse200156 AdminShopViewAddPost(ctx, body)
修改/新增

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormShopViewAddReq**](XbuyAppFormAdminFormShopViewAddReq.md)|  | 

### Return type

[**InlineResponse200156**](inline_response_200_156.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopViewCloseAutoTransPost**
> InlineResponse2007 AdminShopViewCloseAutoTransPost(ctx, body)
关闭自动翻译

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCloseShopViewAutoTransReq**](XbuyAppFormAdminFormCloseShopViewAutoTransReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopViewInfoPost**
> InlineResponse200157 AdminShopViewInfoPost(ctx, body)
详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormShopViewInfoReq**](XbuyAppFormAdminFormShopViewInfoReq.md)|  | 

### Return type

[**InlineResponse200157**](inline_response_200_157.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopViewListPost**
> InlineResponse200158 AdminShopViewListPost(ctx, body)
列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormShopViewListReq**](XbuyAppFormAdminFormShopViewListReq.md)|  | 

### Return type

[**InlineResponse200158**](inline_response_200_158.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopViewOpenAutoTransPost**
> InlineResponse2007 AdminShopViewOpenAutoTransPost(ctx, body)
开启自动翻译

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOpenShopViewAutoTransReq**](XbuyAppFormAdminFormOpenShopViewAutoTransReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

