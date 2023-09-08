# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminRechargeGet**](ShopApi.md#AdminRechargeGet) | **Get** /admin/recharge | 充值
[**AdminShopAddPost**](ShopApi.md#AdminShopAddPost) | **Post** /admin/shop/add | 新增店铺
[**AdminShopChangeShopPost**](ShopApi.md#AdminShopChangeShopPost) | **Post** /admin/shop/change_shop | 切换店铺
[**AdminShopFundsGet**](ShopApi.md#AdminShopFundsGet) | **Get** /admin/shop/funds | 查询店铺资金
[**AdminShopGetBindShopPost**](ShopApi.md#AdminShopGetBindShopPost) | **Post** /admin/shop/get_bind_shop | 已绑定店铺列表
[**AdminShopGetinfoPost**](ShopApi.md#AdminShopGetinfoPost) | **Post** /admin/shop/getinfo | 店铺详细信息
[**AdminShopInfoPost**](ShopApi.md#AdminShopInfoPost) | **Post** /admin/shop/info | 店铺信息
[**AdminShopListPost**](ShopApi.md#AdminShopListPost) | **Post** /admin/shop/list | 店铺列表
[**AdminShopPaywaysGet**](ShopApi.md#AdminShopPaywaysGet) | **Get** /admin/shop/payways | 查询支付方式
[**AdminShopUpdateBindPost**](ShopApi.md#AdminShopUpdateBindPost) | **Post** /admin/shop/update_bind | 商户子账户绑定店铺

# **AdminRechargeGet**
> InlineResponse200137 AdminRechargeGet(ctx, amount)
充值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amount** | **float64**|  | 

### Return type

[**InlineResponse200137**](inline_response_200_137.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopAddPost**
> InlineResponse200150 AdminShopAddPost(ctx, body)
新增店铺

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAddShopReq**](XbuyAppFormAdminFormAddShopReq.md)|  | 

### Return type

[**InlineResponse200150**](inline_response_200_150.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopChangeShopPost**
> InlineResponse20058 AdminShopChangeShopPost(ctx, body)
切换店铺

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormChangeShopReq**](XbuyAppFormAdminFormChangeShopReq.md)|  | 

### Return type

[**InlineResponse20058**](inline_response_200_58.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopFundsGet**
> InlineResponse200151 AdminShopFundsGet(ctx, )
查询店铺资金

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200151**](inline_response_200_151.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopGetBindShopPost**
> InlineResponse200152 AdminShopGetBindShopPost(ctx, body)
已绑定店铺列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetBindShopReq**](XbuyAppFormAdminFormGetBindShopReq.md)|  | 

### Return type

[**InlineResponse200152**](inline_response_200_152.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopGetinfoPost**
> InlineResponse200153 AdminShopGetinfoPost(ctx, body)
店铺详细信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetShopInfoReq**](XbuyAppFormAdminFormGetShopInfoReq.md)|  | 

### Return type

[**InlineResponse200153**](inline_response_200_153.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopInfoPost**
> InlineResponse200153 AdminShopInfoPost(ctx, body)
店铺信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetShopReq**](XbuyAppFormAdminFormGetShopReq.md)|  | 

### Return type

[**InlineResponse200153**](inline_response_200_153.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopListPost**
> InlineResponse200154 AdminShopListPost(ctx, body)
店铺列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormShopListReq**](XbuyAppFormAdminFormShopListReq.md)|  | 

### Return type

[**InlineResponse200154**](inline_response_200_154.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopPaywaysGet**
> InlineResponse200155 AdminShopPaywaysGet(ctx, )
查询支付方式

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200155**](inline_response_200_155.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminShopUpdateBindPost**
> InlineResponse2007 AdminShopUpdateBindPost(ctx, body)
商户子账户绑定店铺

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUpdateBindShopReq**](XbuyAppFormAdminFormUpdateBindShopReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

