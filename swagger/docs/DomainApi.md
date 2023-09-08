# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminDomainAddPost**](DomainApi.md#AdminDomainAddPost) | **Post** /admin/domain/add | 域名新增/编辑
[**AdminDomainDeletePost**](DomainApi.md#AdminDomainDeletePost) | **Post** /admin/domain/delete | 删除域名
[**AdminDomainGetMainPost**](DomainApi.md#AdminDomainGetMainPost) | **Post** /admin/domain/get_main | 查询主域名
[**AdminDomainGetPost**](DomainApi.md#AdminDomainGetPost) | **Post** /admin/domain/get | 查询域名
[**AdminDomainListPost**](DomainApi.md#AdminDomainListPost) | **Post** /admin/domain/list | 域名列表

# **AdminDomainAddPost**
> InlineResponse20033 AdminDomainAddPost(ctx, body)
域名新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDomainAddReq**](XbuyAppFormAdminFormDomainAddReq.md)|  | 

### Return type

[**InlineResponse20033**](inline_response_200_33.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDomainDeletePost**
> InlineResponse2007 AdminDomainDeletePost(ctx, body)
删除域名

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDomainDeleteReq**](XbuyAppFormAdminFormDomainDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDomainGetMainPost**
> InlineResponse20034 AdminDomainGetMainPost(ctx, body)
查询主域名

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDomainGetMainReq**](XbuyAppFormAdminFormDomainGetMainReq.md)|  | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDomainGetPost**
> InlineResponse20034 AdminDomainGetPost(ctx, body)
查询域名

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDomainGetReq**](XbuyAppFormAdminFormDomainGetReq.md)|  | 

### Return type

[**InlineResponse20034**](inline_response_200_34.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDomainListPost**
> InlineResponse20035 AdminDomainListPost(ctx, body)
域名列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDomainListReq**](XbuyAppFormAdminFormDomainListReq.md)|  | 

### Return type

[**InlineResponse20035**](inline_response_200_35.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

