# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminBenefitCodeAddPost**](BenefitCodeApi.md#AdminBenefitCodeAddPost) | **Post** /admin/benefit_code/add | 添加折扣码
[**AdminBenefitCodeCountryListPost**](BenefitCodeApi.md#AdminBenefitCodeCountryListPost) | **Post** /admin/benefit_code/country_list | 国家列表
[**AdminBenefitCodeListPost**](BenefitCodeApi.md#AdminBenefitCodeListPost) | **Post** /admin/benefit_code/list | 折扣码列表
[**AdminBenefitCodeStatisticsPost**](BenefitCodeApi.md#AdminBenefitCodeStatisticsPost) | **Post** /admin/benefit_code/statistics | 统计信息
[**AdminBenefitCodeZendCountryListPost**](BenefitCodeApi.md#AdminBenefitCodeZendCountryListPost) | **Post** /admin/benefit_code/zend_country_list | 树状国家列表

# **AdminBenefitCodeAddPost**
> InlineResponse20010 AdminBenefitCodeAddPost(ctx, body)
添加折扣码

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormBenefitCodeAddReq**](XbuyAppFormAdminFormBenefitCodeAddReq.md)|  | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminBenefitCodeCountryListPost**
> InlineResponse20011 AdminBenefitCodeCountryListPost(ctx, body)
国家列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCountryListReq**](XbuyAppFormAdminFormCountryListReq.md)|  | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminBenefitCodeListPost**
> InlineResponse20012 AdminBenefitCodeListPost(ctx, body)
折扣码列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormBenefitCodeListReq**](XbuyAppFormAdminFormBenefitCodeListReq.md)|  | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminBenefitCodeStatisticsPost**
> InlineResponse20013 AdminBenefitCodeStatisticsPost(ctx, body)
统计信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormBenefitCodeStatisticsReq**](XbuyAppFormAdminFormBenefitCodeStatisticsReq.md)|  | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminBenefitCodeZendCountryListPost**
> InlineResponse20014 AdminBenefitCodeZendCountryListPost(ctx, body)
树状国家列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormZendCountryListReq**](XbuyAppFormAdminFormZendCountryListReq.md)|  | 

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

