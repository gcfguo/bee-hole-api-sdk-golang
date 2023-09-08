# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminPostCodeUniqueGet**](PostApi.md#AdminPostCodeUniqueGet) | **Get** /admin/post/code_unique | 岗位编码是否唯一
[**AdminPostDeletePost**](PostApi.md#AdminPostDeletePost) | **Post** /admin/post/delete | 删除岗位
[**AdminPostEditPost**](PostApi.md#AdminPostEditPost) | **Post** /admin/post/edit | 修改/新增岗位
[**AdminPostListGet**](PostApi.md#AdminPostListGet) | **Get** /admin/post/list | 获取岗位列表
[**AdminPostMaxSortGet**](PostApi.md#AdminPostMaxSortGet) | **Get** /admin/post/max_sort | 岗位最大排序
[**AdminPostNameUniqueGet**](PostApi.md#AdminPostNameUniqueGet) | **Get** /admin/post/name_unique | 岗位名称是否唯一
[**AdminPostViewGet**](PostApi.md#AdminPostViewGet) | **Get** /admin/post/view | 获取指定信息
[**AdminRolePostEditPost**](PostApi.md#AdminRolePostEditPost) | **Post** /admin/role/post_edit | 岗位权限菜单编辑

# **AdminPostCodeUniqueGet**
> InlineResponse20022 AdminPostCodeUniqueGet(ctx, code, optional)
岗位编码是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| 岗位编码 | 
 **optional** | ***PostApiAdminPostCodeUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostApiAdminPostCodeUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 岗位ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPostDeletePost**
> InlineResponse2007 AdminPostDeletePost(ctx, body)
删除岗位

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormPostDeleteReq**](XbuyAppFormAdminFormPostDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPostEditPost**
> InlineResponse2007 AdminPostEditPost(ctx, body)
修改/新增岗位

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormPostEditReq**](XbuyAppFormAdminFormPostEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPostListGet**
> InlineResponse200131 AdminPostListGet(ctx, optional)
获取岗位列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostApiAdminPostListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostApiAdminPostListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **startTime** | **optional.String**| 开始日期 | 
 **endTime** | **optional.String**| 结束日期 | 
 **status** | **optional.Int32**| 状态 | 
 **name** | **optional.String**| 岗位名称 | 
 **code** | **optional.String**| 岗位编码 | 

### Return type

[**InlineResponse200131**](inline_response_200_131.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPostMaxSortGet**
> InlineResponse20021 AdminPostMaxSortGet(ctx, optional)
岗位最大排序

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PostApiAdminPostMaxSortGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostApiAdminPostMaxSortGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| 岗位ID | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPostNameUniqueGet**
> InlineResponse20022 AdminPostNameUniqueGet(ctx, name, optional)
岗位名称是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 岗位名称 | 
 **optional** | ***PostApiAdminPostNameUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostApiAdminPostNameUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 岗位ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminPostViewGet**
> InlineResponse200132 AdminPostViewGet(ctx, id)
获取指定信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 岗位ID | 

### Return type

[**InlineResponse200132**](inline_response_200_132.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminRolePostEditPost**
> InlineResponse2007 AdminRolePostEditPost(ctx, body)
岗位权限菜单编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormPostMenuEditReq**](XbuyAppFormAdminFormPostMenuEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

