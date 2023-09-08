# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminMenuCodeUniqueGet**](MenuApi.md#AdminMenuCodeUniqueGet) | **Get** /admin/menu/code_unique | 菜单编码是否唯一
[**AdminMenuDeletePost**](MenuApi.md#AdminMenuDeletePost) | **Post** /admin/menu/delete | 删除菜单
[**AdminMenuEditPost**](MenuApi.md#AdminMenuEditPost) | **Post** /admin/menu/edit | 修改/新增菜单
[**AdminMenuListGet**](MenuApi.md#AdminMenuListGet) | **Get** /admin/menu/list | 获取菜单列表
[**AdminMenuMaxSortGet**](MenuApi.md#AdminMenuMaxSortGet) | **Get** /admin/menu/max_sort | 菜单最大排序
[**AdminMenuNameUniqueGet**](MenuApi.md#AdminMenuNameUniqueGet) | **Get** /admin/menu/name_unique | 菜单名称是否唯一
[**AdminMenuRoleListGet**](MenuApi.md#AdminMenuRoleListGet) | **Get** /admin/menu/role_list | 查询角色菜单列表
[**AdminMenuSearchListGet**](MenuApi.md#AdminMenuSearchListGet) | **Get** /admin/menu/search_list | 获取菜单列表
[**AdminMenuSubRoleListGet**](MenuApi.md#AdminMenuSubRoleListGet) | **Get** /admin/menu/sub_role_list | 查询角色菜单列表
[**AdminMenuViewGet**](MenuApi.md#AdminMenuViewGet) | **Get** /admin/menu/view | 获取指定菜单信息

# **AdminMenuCodeUniqueGet**
> InlineResponse20022 AdminMenuCodeUniqueGet(ctx, code, optional)
菜单编码是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **code** | **string**| 菜单编码 | 
 **optional** | ***MenuApiAdminMenuCodeUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MenuApiAdminMenuCodeUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 菜单ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMenuDeletePost**
> InlineResponse2007 AdminMenuDeletePost(ctx, body)
删除菜单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMenuDeleteReq**](XbuyAppFormAdminFormMenuDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMenuEditPost**
> InlineResponse2007 AdminMenuEditPost(ctx, body)
修改/新增菜单

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMenuEditReq**](XbuyAppFormAdminFormMenuEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMenuListGet**
> InlineResponse200114 AdminMenuListGet(ctx, optional)
获取菜单列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MenuApiAdminMenuListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MenuApiAdminMenuListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **pid** | **optional.Int64**| 父ID | 

### Return type

[**InlineResponse200114**](inline_response_200_114.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMenuMaxSortGet**
> InlineResponse20021 AdminMenuMaxSortGet(ctx, optional)
菜单最大排序

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MenuApiAdminMenuMaxSortGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MenuApiAdminMenuMaxSortGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| 菜单ID | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMenuNameUniqueGet**
> InlineResponse20022 AdminMenuNameUniqueGet(ctx, name, optional)
菜单名称是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 菜单名称 | 
 **optional** | ***MenuApiAdminMenuNameUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MenuApiAdminMenuNameUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 菜单ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMenuRoleListGet**
> InlineResponse200115 AdminMenuRoleListGet(ctx, optional)
查询角色菜单列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MenuApiAdminMenuRoleListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MenuApiAdminMenuRoleListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleId** | **optional.String**| 角色ID | 

### Return type

[**InlineResponse200115**](inline_response_200_115.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMenuSearchListGet**
> InlineResponse200116 AdminMenuSearchListGet(ctx, optional)
获取菜单列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MenuApiAdminMenuSearchListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MenuApiAdminMenuSearchListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| 菜单名称 | 
 **status** | **optional.Int32**| 状态 | 

### Return type

[**InlineResponse200116**](inline_response_200_116.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMenuSubRoleListGet**
> InlineResponse200115 AdminMenuSubRoleListGet(ctx, optional)
查询角色菜单列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***MenuApiAdminMenuSubRoleListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MenuApiAdminMenuSubRoleListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **memberId** | **optional.Int64**| 用户ID | 
 **postId** | **optional.Int64**| 岗位ID | 

### Return type

[**InlineResponse200115**](inline_response_200_115.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMenuViewGet**
> InlineResponse200117 AdminMenuViewGet(ctx, id)
获取指定菜单信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| 菜单ID | 

### Return type

[**InlineResponse200117**](inline_response_200_117.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

