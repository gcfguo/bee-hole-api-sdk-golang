# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminTransTaskAddGoodJobPost**](TransTaskApi.md#AdminTransTaskAddGoodJobPost) | **Post** /admin/trans_task/add_good_job | 投递商品级翻译任务
[**AdminTransTaskAddShopJobPost**](TransTaskApi.md#AdminTransTaskAddShopJobPost) | **Post** /admin/trans_task/add_shop_job | 投递商店级翻译任务

# **AdminTransTaskAddGoodJobPost**
> InlineResponse2007 AdminTransTaskAddGoodJobPost(ctx, body)
投递商品级翻译任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAddGoodsTransJobReq**](XbuyAppFormAdminFormAddGoodsTransJobReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTransTaskAddShopJobPost**
> InlineResponse2007 AdminTransTaskAddShopJobPost(ctx, body)
投递商店级翻译任务

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAddShopTransJobReq**](XbuyAppFormAdminFormAddShopTransJobReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

