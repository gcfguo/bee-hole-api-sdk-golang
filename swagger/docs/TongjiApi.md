# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminTongjiGoodsNumPost**](TongjiApi.md#AdminTongjiGoodsNumPost) | **Post** /admin/tongji/goods-num | 商品总数
[**AdminTongjiGoodsSalesPost**](TongjiApi.md#AdminTongjiGoodsSalesPost) | **Post** /admin/tongji/goods-sales | 商品销售额
[**AdminTongjiOrderDataPost**](TongjiApi.md#AdminTongjiOrderDataPost) | **Post** /admin/tongji/order-data | 订单数据
[**AdminTongjiOrderMoneyPost**](TongjiApi.md#AdminTongjiOrderMoneyPost) | **Post** /admin/tongji/order-money | 订单金额
[**AdminTongjiRealTimeDataPost**](TongjiApi.md#AdminTongjiRealTimeDataPost) | **Post** /admin/tongji/real-time-data | 实时数据
[**AdminTongjiUserBuyPost**](TongjiApi.md#AdminTongjiUserBuyPost) | **Post** /admin/tongji/user-buy | 用户购买统计

# **AdminTongjiGoodsNumPost**
> InlineResponse200173 AdminTongjiGoodsNumPost(ctx, body)
商品总数

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodNumReq**](XbuyAppFormAdminFormGoodNumReq.md)|  | 

### Return type

[**InlineResponse200173**](inline_response_200_173.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTongjiGoodsSalesPost**
> InlineResponse200174 AdminTongjiGoodsSalesPost(ctx, body)
商品销售额

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGoodsSalesReq**](XbuyAppFormAdminFormGoodsSalesReq.md)|  | 

### Return type

[**InlineResponse200174**](inline_response_200_174.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTongjiOrderDataPost**
> InlineResponse200175 AdminTongjiOrderDataPost(ctx, body)
订单数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOrderDataReq**](XbuyAppFormAdminFormOrderDataReq.md)|  | 

### Return type

[**InlineResponse200175**](inline_response_200_175.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTongjiOrderMoneyPost**
> InlineResponse200176 AdminTongjiOrderMoneyPost(ctx, body)
订单金额

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormOrderMoneyReq**](XbuyAppFormAdminFormOrderMoneyReq.md)|  | 

### Return type

[**InlineResponse200176**](inline_response_200_176.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTongjiRealTimeDataPost**
> InlineResponse200177 AdminTongjiRealTimeDataPost(ctx, body)
实时数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormRealTimeDataReq**](XbuyAppFormAdminFormRealTimeDataReq.md)|  | 

### Return type

[**InlineResponse200177**](inline_response_200_177.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminTongjiUserBuyPost**
> InlineResponse200178 AdminTongjiUserBuyPost(ctx, body)
用户购买统计

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUserBuyReq**](XbuyAppFormAdminFormUserBuyReq.md)|  | 

### Return type

[**InlineResponse200178**](inline_response_200_178.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

