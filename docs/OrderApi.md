# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminOrderAfterSalesApplyPost**](OrderApi.md#AdminOrderAfterSalesApplyPost) | **Post** /admin/order/after_sales/apply | 申请售后
[**AdminOrderCancelAfterSalesPost**](OrderApi.md#AdminOrderCancelAfterSalesPost) | **Post** /admin/order/cancel_after_sales | 取消售后
[**AdminOrderExportPost**](OrderApi.md#AdminOrderExportPost) | **Post** /admin/order/export | 订单导出
[**AdminOrderImportPost**](OrderApi.md#AdminOrderImportPost) | **Post** /admin/order/import | 订单导入
[**AdminOrderInsurePost**](OrderApi.md#AdminOrderInsurePost) | **Post** /admin/order/insure | 保单推送日志
[**AdminOrderItemsPost**](OrderApi.md#AdminOrderItemsPost) | **Post** /admin/order/items | 订单商品
[**AdminOrderListPost**](OrderApi.md#AdminOrderListPost) | **Post** /admin/order/list | 订单列表
[**AdminOrderPushLogPost**](OrderApi.md#AdminOrderPushLogPost) | **Post** /admin/order/push/log | 订单推送日志
[**AdminOrderPushPost**](OrderApi.md#AdminOrderPushPost) | **Post** /admin/order/push | 推送订单
[**AdminOrderReceiverUpdatePost**](OrderApi.md#AdminOrderReceiverUpdatePost) | **Post** /admin/order/receiver/update | 更新收货人信息
[**AdminOrderShipPost**](OrderApi.md#AdminOrderShipPost) | **Post** /admin/order/ship | 订单发货
[**AdminOrderWaybillListPost**](OrderApi.md#AdminOrderWaybillListPost) | **Post** /admin/order/waybill/list | 订单物流
[**MallOrderAfterSalesApplyPost**](OrderApi.md#MallOrderAfterSalesApplyPost) | **Post** /mall/order/after_sales/apply | 申请售后

# **AdminOrderAfterSalesApplyPost**
> XbuyAppFormAdminFormAfterSalesApplyRes AdminOrderAfterSalesApplyPost(ctx, body)
申请售后

申请售后

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAfterSalesApplyReq**](XbuyAppFormAdminFormAfterSalesApplyReq.md)|  | 

### Return type

[**XbuyAppFormAdminFormAfterSalesApplyRes**](xbuy.app.form.adminForm.AfterSalesApplyRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderCancelAfterSalesPost**
> InlineResponse2007 AdminOrderCancelAfterSalesPost(ctx, body)
取消售后

取消售后

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCancelAfterSalesReq**](XbuyAppFormAdminFormCancelAfterSalesReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderExportPost**
> InlineResponse200120 AdminOrderExportPost(ctx, body)
订单导出

订单导出

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormExportOrderReq**](XbuyAppFormAdminFormExportOrderReq.md)|  | 

### Return type

[**InlineResponse200120**](inline_response_200_120.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderImportPost**
> InlineResponse200121 AdminOrderImportPost(ctx, body)
订单导入

订单导入

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormImportOrderReq**](XbuyAppFormAdminFormImportOrderReq.md)|  | 

### Return type

[**InlineResponse200121**](inline_response_200_121.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderInsurePost**
> InlineResponse200122 AdminOrderInsurePost(ctx, body)
保单推送日志

保单推送日志

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOrderInsuredReq**](XbuyAppFormAdminFormOrderInsuredReq.md)|  | 

### Return type

[**InlineResponse200122**](inline_response_200_122.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderItemsPost**
> InlineResponse200123 AdminOrderItemsPost(ctx, body)
订单商品

订单商品

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOrderItemsReq**](XbuyAppFormAdminFormOrderItemsReq.md)|  | 

### Return type

[**InlineResponse200123**](inline_response_200_123.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderListPost**
> InlineResponse200124 AdminOrderListPost(ctx, body)
订单列表

订单列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOrderListReq**](XbuyAppFormAdminFormOrderListReq.md)|  | 

### Return type

[**InlineResponse200124**](inline_response_200_124.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderPushLogPost**
> InlineResponse200126 AdminOrderPushLogPost(ctx, body)
订单推送日志

订单推送日志

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOrderPushLogReq**](XbuyAppFormAdminFormOrderPushLogReq.md)|  | 

### Return type

[**InlineResponse200126**](inline_response_200_126.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderPushPost**
> InlineResponse200125 AdminOrderPushPost(ctx, body)
推送订单

推送订单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOrderPushReq**](XbuyAppFormAdminFormOrderPushReq.md)|  | 

### Return type

[**InlineResponse200125**](inline_response_200_125.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderReceiverUpdatePost**
> InlineResponse2007 AdminOrderReceiverUpdatePost(ctx, body)
更新收货人信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUpdateReceiverInfoReq**](XbuyAppFormAdminFormUpdateReceiverInfoReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderShipPost**
> InlineResponse200127 AdminOrderShipPost(ctx, body)
订单发货

订单发货

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOrderShipReq**](XbuyAppFormAdminFormOrderShipReq.md)|  | 

### Return type

[**InlineResponse200127**](inline_response_200_127.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminOrderWaybillListPost**
> InlineResponse200128 AdminOrderWaybillListPost(ctx, body)
订单物流

订单物流列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOrderWaybillExpressReq**](XbuyAppFormAdminFormOrderWaybillExpressReq.md)|  | 

### Return type

[**InlineResponse200128**](inline_response_200_128.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MallOrderAfterSalesApplyPost**
> XbuyAppFormMallFormAfterSalesApplyRes MallOrderAfterSalesApplyPost(ctx, body)
申请售后

申请售后

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormMallFormAfterSalesApplyReq**](XbuyAppFormMallFormAfterSalesApplyReq.md)|  | 

### Return type

[**XbuyAppFormMallFormAfterSalesApplyRes**](xbuy.app.form.mallForm.AfterSalesApplyRes.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

