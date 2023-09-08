# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminPmethodCurrencyGetGet**](PmethodApi.md#AdminPmethodCurrencyGetGet) | **Get** /admin/pmethod/currency/get | 获取币种类型
[**AdminPmethodCurrencyratesSetPost**](PmethodApi.md#AdminPmethodCurrencyratesSetPost) | **Post** /admin/pmethod/currencyrates/set | 获取币种汇率关系

# **AdminPmethodCurrencyGetGet**
> InlineResponse200129 AdminPmethodCurrencyGetGet(ctx, )
获取币种类型

获取币种类型

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse200129**](inline_response_200_129.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPmethodCurrencyratesSetPost**
> InlineResponse200130 AdminPmethodCurrencyratesSetPost(ctx, body)
获取币种汇率关系

获取币种汇率关系

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormSetCurrencyRatesReq**](XbuyAppFormAdminFormSetCurrencyRatesReq.md)|  | 

### Return type

[**InlineResponse200130**](inline_response_200_130.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

