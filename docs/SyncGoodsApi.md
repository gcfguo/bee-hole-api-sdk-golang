# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminInitDataPost**](SyncGoodsApi.md#AdminInitDataPost) | **Post** /admin/init/data | 初始化数据
[**AdminSynconfigDelPost**](SyncGoodsApi.md#AdminSynconfigDelPost) | **Post** /admin/synconfig/del | 删除配置
[**AdminSynconfigListPost**](SyncGoodsApi.md#AdminSynconfigListPost) | **Post** /admin/synconfig/list | 配置列表
[**AdminSynconfigSavePost**](SyncGoodsApi.md#AdminSynconfigSavePost) | **Post** /admin/synconfig/save | 保存配置
[**AdminSyngoodsTriggerPost**](SyncGoodsApi.md#AdminSyngoodsTriggerPost) | **Post** /admin/syngoods/trigger | 同步商品

# **AdminInitDataPost**
> InlineResponse20054 AdminInitDataPost(ctx, body)
初始化数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMakeTestDataReq**](XbuyAppFormAdminFormMakeTestDataReq.md)|  | 

### Return type

[**InlineResponse20054**](inline_response_200_54.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSynconfigDelPost**
> InlineResponse2007 AdminSynconfigDelPost(ctx, body)
删除配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSynDeleteReq**](XbuyAppFormAdminFormSynDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSynconfigListPost**
> InlineResponse200168 AdminSynconfigListPost(ctx, body)
配置列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSynListReq**](XbuyAppFormAdminFormSynListReq.md)|  | 

### Return type

[**InlineResponse200168**](inline_response_200_168.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSynconfigSavePost**
> InlineResponse20080 AdminSynconfigSavePost(ctx, body)
保存配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSynAddReq**](XbuyAppFormAdminFormSynAddReq.md)|  | 

### Return type

[**InlineResponse20080**](inline_response_200_80.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminSyngoodsTriggerPost**
> InlineResponse2007 AdminSyngoodsTriggerPost(ctx, body)
同步商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSynGoodsReq**](XbuyAppFormAdminFormSynGoodsReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

