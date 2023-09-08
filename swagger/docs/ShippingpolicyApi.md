# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminShippingpolicyPriorityCreatePost**](ShippingpolicyApi.md#AdminShippingpolicyPriorityCreatePost) | **Post** /admin/shippingpolicy/priority/create | 创建优先发货策略
[**AdminShippingpolicyPriorityUpdatePost**](ShippingpolicyApi.md#AdminShippingpolicyPriorityUpdatePost) | **Post** /admin/shippingpolicy/priority/update | 修改优先发货策略

# **AdminShippingpolicyPriorityCreatePost**
> InlineResponse200149 AdminShippingpolicyPriorityCreatePost(ctx, body)
创建优先发货策略

创建优先发货策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormPriorityPolicyCreateReq**](XbuyAppFormAdminFormPriorityPolicyCreateReq.md)|  | 

### Return type

[**InlineResponse200149**](inline_response_200_149.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShippingpolicyPriorityUpdatePost**
> InlineResponse2007 AdminShippingpolicyPriorityUpdatePost(ctx, body)
修改优先发货策略

修改优先发货策略

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormPriorityPolicyUpdateReq**](XbuyAppFormAdminFormPriorityPolicyUpdateReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

