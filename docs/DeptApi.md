# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminDeptDeletePost**](DeptApi.md#AdminDeptDeletePost) | **Post** /admin/dept/delete | 删除部门
[**AdminDeptEditPost**](DeptApi.md#AdminDeptEditPost) | **Post** /admin/dept/edit | 修改/新增部门
[**AdminDeptListGet**](DeptApi.md#AdminDeptListGet) | **Get** /admin/dept/list | 获取部门列表
[**AdminDeptListTreeGet**](DeptApi.md#AdminDeptListTreeGet) | **Get** /admin/dept/list_tree | 获取部门列表树
[**AdminDeptMaxSortGet**](DeptApi.md#AdminDeptMaxSortGet) | **Get** /admin/dept/max_sort | 部门最大排序
[**AdminDeptNameUniqueGet**](DeptApi.md#AdminDeptNameUniqueGet) | **Get** /admin/dept/name_unique | 部门名称是否唯一
[**AdminDeptViewGet**](DeptApi.md#AdminDeptViewGet) | **Get** /admin/dept/view | 获取指定信息

# **AdminDeptDeletePost**
> InlineResponse2007 AdminDeptDeletePost(ctx, body)
删除部门

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDeptDeleteReq**](XbuyAppFormAdminFormDeptDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDeptEditPost**
> InlineResponse2007 AdminDeptEditPost(ctx, body)
修改/新增部门

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDeptEditReq**](XbuyAppFormAdminFormDeptEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDeptListGet**
> InlineResponse20025 AdminDeptListGet(ctx, optional)
获取部门列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeptApiAdminDeptListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeptApiAdminDeptListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pid** | **optional.Int64**| 父级部门ID | 

### Return type

[**InlineResponse20025**](inline_response_200_25.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDeptListTreeGet**
> InlineResponse20026 AdminDeptListTreeGet(ctx, optional)
获取部门列表树

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeptApiAdminDeptListTreeGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeptApiAdminDeptListTreeGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| 部门ID | 

### Return type

[**InlineResponse20026**](inline_response_200_26.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDeptMaxSortGet**
> InlineResponse20021 AdminDeptMaxSortGet(ctx, optional)
部门最大排序

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DeptApiAdminDeptMaxSortGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeptApiAdminDeptMaxSortGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| 部门ID | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDeptNameUniqueGet**
> InlineResponse20022 AdminDeptNameUniqueGet(ctx, name, optional)
部门名称是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 部门名称 | 
 **optional** | ***DeptApiAdminDeptNameUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DeptApiAdminDeptNameUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 部门ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDeptViewGet**
> InlineResponse20027 AdminDeptViewGet(ctx, id)
获取指定信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| 部门ID | 

### Return type

[**InlineResponse20027**](inline_response_200_27.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

