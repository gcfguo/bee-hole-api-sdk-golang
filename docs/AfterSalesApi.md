# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminAfterSalesListPost**](AfterSalesApi.md#AdminAfterSalesListPost) | **Post** /admin/after_sales/list | 售后单列表
[**AdminAfterSalesRefundPost**](AfterSalesApi.md#AdminAfterSalesRefundPost) | **Post** /admin/after_sales/refund | 发起退款
[**AdminAfterSalesSwitchPost**](AfterSalesApi.md#AdminAfterSalesSwitchPost) | **Post** /admin/after_sales/switch | 启用或关闭售后

# **AdminAfterSalesListPost**
> InlineResponse200 AdminAfterSalesListPost(ctx, body)
售后单列表

售后单列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAfterSalesListReq**](XbuyAppFormAdminFormAfterSalesListReq.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminAfterSalesRefundPost**
> XbuyAppFormAdminFormRefundRes AdminAfterSalesRefundPost(ctx, body)
发起退款

发起退款

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormRefundReq**](XbuyAppFormAdminFormRefundReq.md)|  | 

### Return type

[**XbuyAppFormAdminFormRefundRes**](xbuy.app.form.adminForm.RefundRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminAfterSalesSwitchPost**
> InlineResponse2001 AdminAfterSalesSwitchPost(ctx, body)
启用或关闭售后

启用或关闭售后

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAfterSalesSwitchReq**](XbuyAppFormAdminFormAfterSalesSwitchReq.md)|  | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

