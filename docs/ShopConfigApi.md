# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminMallConfigGetPost**](ShopConfigApi.md#AdminMallConfigGetPost) | **Post** /admin/mall_config/get | 获取店铺配置
[**AdminMallConfigSavePost**](ShopConfigApi.md#AdminMallConfigSavePost) | **Post** /admin/mall_config/save | 保存店铺配置

# **AdminMallConfigGetPost**
> InlineResponse20059 AdminMallConfigGetPost(ctx, body)
获取店铺配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMallShopConfigsReq**](XbuyAppFormAdminFormGetMallShopConfigsReq.md)|  | 

### Return type

[**InlineResponse20059**](inline_response_200_59.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMallConfigSavePost**
> InlineResponse2007 AdminMallConfigSavePost(ctx, body)
保存店铺配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSaveMallShopConfigsReq**](XbuyAppFormAdminFormSaveMallShopConfigsReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

