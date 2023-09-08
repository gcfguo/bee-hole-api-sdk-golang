# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminCatAddPost**](CatApi.md#AdminCatAddPost) | **Post** /admin/cat/add | 新增/编辑类目
[**AdminCatDeletePost**](CatApi.md#AdminCatDeletePost) | **Post** /admin/cat/delete | 删除类目
[**AdminCatInfoPost**](CatApi.md#AdminCatInfoPost) | **Post** /admin/cat/info | 类目详情
[**AdminCatListPost**](CatApi.md#AdminCatListPost) | **Post** /admin/cat/list | 类目列表
[**AdminCatParentListPost**](CatApi.md#AdminCatParentListPost) | **Post** /admin/cat/parent/list | 父级类目列表
[**AdminCatSubAddPost**](CatApi.md#AdminCatSubAddPost) | **Post** /admin/cat/sub/add | 类目副本新增/编辑

# **AdminCatAddPost**
> InlineResponse2004 AdminCatAddPost(ctx, body)
新增/编辑类目

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCatAddReq**](XbuyAppFormAdminFormCatAddReq.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminCatDeletePost**
> InlineResponse2007 AdminCatDeletePost(ctx, body)
删除类目

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCatDeleteReq**](XbuyAppFormAdminFormCatDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminCatInfoPost**
> InlineResponse20015 AdminCatInfoPost(ctx, body)
类目详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCatInfoReq**](XbuyAppFormAdminFormCatInfoReq.md)|  | 

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminCatListPost**
> InlineResponse20016 AdminCatListPost(ctx, body)
类目列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCatListReq**](XbuyAppFormAdminFormCatListReq.md)|  | 

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminCatParentListPost**
> InlineResponse20017 AdminCatParentListPost(ctx, body)
父级类目列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCatParentListReq**](XbuyAppFormAdminFormCatParentListReq.md)|  | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminCatSubAddPost**
> InlineResponse2004 AdminCatSubAddPost(ctx, body)
类目副本新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCatSubAddReq**](XbuyAppFormAdminFormCatSubAddReq.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

