# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminPostfeeAddPost**](PostfeeApi.md#AdminPostfeeAddPost) | **Post** /admin/postfee/add | 运费模板新增/编辑
[**AdminPostfeeListPost**](PostfeeApi.md#AdminPostfeeListPost) | **Post** /admin/postfee/list | 运费列表
[**AdminPostfeeModeAddPost**](PostfeeApi.md#AdminPostfeeModeAddPost) | **Post** /admin/postfee/mode-add | 运费方式新增/编辑
[**AdminPostfeeModeListPost**](PostfeeApi.md#AdminPostfeeModeListPost) | **Post** /admin/postfee/mode-list | 运费方式列表

# **AdminPostfeeAddPost**
> InlineResponse200133 AdminPostfeeAddPost(ctx, body)
运费模板新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormPostFeeAddReq**](XbuyAppFormAdminFormPostFeeAddReq.md)|  | 

### Return type

[**InlineResponse200133**](inline_response_200_133.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPostfeeListPost**
> InlineResponse200134 AdminPostfeeListPost(ctx, body)
运费列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormPostFeeListReq**](XbuyAppFormAdminFormPostFeeListReq.md)|  | 

### Return type

[**InlineResponse200134**](inline_response_200_134.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPostfeeModeAddPost**
> InlineResponse200135 AdminPostfeeModeAddPost(ctx, body)
运费方式新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormPostFeeModeAddReq**](XbuyAppFormAdminFormPostFeeModeAddReq.md)|  | 

### Return type

[**InlineResponse200135**](inline_response_200_135.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPostfeeModeListPost**
> InlineResponse200136 AdminPostfeeModeListPost(ctx, body)
运费方式列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormModeListReq**](XbuyAppFormAdminFormModeListReq.md)|  | 

### Return type

[**InlineResponse200136**](inline_response_200_136.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

