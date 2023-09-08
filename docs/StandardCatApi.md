# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminGoodInfoAddPost**](StandardCatApi.md#AdminGoodInfoAddPost) | **Post** /admin/good/info/add | 标准类目参数新增/编辑
[**AdminGoodInfoDeleteOnePost**](StandardCatApi.md#AdminGoodInfoDeleteOnePost) | **Post** /admin/good/info/delete/one | 标准类目参数删除
[**AdminScatSavePost**](StandardCatApi.md#AdminScatSavePost) | **Post** /admin/scat/save | 标准类目列表
[**AdminStandardCatPost**](StandardCatApi.md#AdminStandardCatPost) | **Post** /admin/standard/cat | 标准类目列表

# **AdminGoodInfoAddPost**
> InlineResponse20046 AdminGoodInfoAddPost(ctx, body)
标准类目参数新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodInfoAddReq**](XbuyAppFormAdminFormGoodInfoAddReq.md)|  | 

### Return type

[**InlineResponse20046**](inline_response_200_46.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminGoodInfoDeleteOnePost**
> InlineResponse2007 AdminGoodInfoDeleteOnePost(ctx, body)
标准类目参数删除

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodInfoDeleteOneReq**](XbuyAppFormAdminFormGoodInfoDeleteOneReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminScatSavePost**
> InlineResponse20080 AdminScatSavePost(ctx, body)
标准类目列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSaveCatReq**](XbuyAppFormAdminFormSaveCatReq.md)|  | 

### Return type

[**InlineResponse20080**](inline_response_200_80.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminStandardCatPost**
> InlineResponse200165 AdminStandardCatPost(ctx, body)
标准类目列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormStandardCatReq**](XbuyAppFormAdminFormStandardCatReq.md)|  | 

### Return type

[**InlineResponse200165**](inline_response_200_165.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

