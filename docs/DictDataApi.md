# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminDictAttributeGet**](DictDataApi.md#AdminDictAttributeGet) | **Get** /admin/dict/attribute | 获取指定字典类型的属性数据
[**AdminDictDataDeletePost**](DictDataApi.md#AdminDictDataDeletePost) | **Post** /admin/dict_data/delete | 删除字典数据
[**AdminDictDataEditPost**](DictDataApi.md#AdminDictDataEditPost) | **Post** /admin/dict_data/edit | 修改/新增字典数据
[**AdminDictDataListGet**](DictDataApi.md#AdminDictDataListGet) | **Get** /admin/dict_data/list | 获取字典数据列表
[**AdminDictDataMaxSortGet**](DictDataApi.md#AdminDictDataMaxSortGet) | **Get** /admin/dict_data/max_sort | 查询字典数据最大排序
[**AdminDictDataUniqueGet**](DictDataApi.md#AdminDictDataUniqueGet) | **Get** /admin/dict_data/unique | 数据键值是否唯一
[**AdminDictDataViewGet**](DictDataApi.md#AdminDictDataViewGet) | **Get** /admin/dict_data/view | 获取指定字典数据信息
[**AdminDictTypeDeletePost**](DictDataApi.md#AdminDictTypeDeletePost) | **Post** /admin/dict_type/delete | 删除字典类型
[**AdminDictTypeEditPost**](DictDataApi.md#AdminDictTypeEditPost) | **Post** /admin/dict_type/edit | 修改/新增字典类型
[**AdminDictTypeExportGet**](DictDataApi.md#AdminDictTypeExportGet) | **Get** /admin/dict_type/export | 导出字典类型
[**AdminDictTypeListGet**](DictDataApi.md#AdminDictTypeListGet) | **Get** /admin/dict_type/list | 获取字典类型列表
[**AdminDictTypeRefreshCacheGet**](DictDataApi.md#AdminDictTypeRefreshCacheGet) | **Get** /admin/dict_type/refresh_cache | 刷新字典缓存
[**AdminDictTypeUniqueGet**](DictDataApi.md#AdminDictTypeUniqueGet) | **Get** /admin/dict_type/unique | 类型是否唯一
[**AdminDictTypeViewGet**](DictDataApi.md#AdminDictTypeViewGet) | **Get** /admin/dict_type/view | 获取指定字典类型信息
[**ApiDictAttributeGet**](DictDataApi.md#ApiDictAttributeGet) | **Get** /api/dict/attribute | 获取指定字典类型的属性数据

# **AdminDictAttributeGet**
> InlineResponse20028 AdminDictAttributeGet(ctx, type_)
获取指定字典类型的属性数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| 字典类型 | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictDataDeletePost**
> InlineResponse2007 AdminDictDataDeletePost(ctx, body)
删除字典数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDictDataDeleteReq**](XbuyAppFormAdminFormDictDataDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictDataEditPost**
> InlineResponse2007 AdminDictDataEditPost(ctx, body)
修改/新增字典数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDictDataEditReq**](XbuyAppFormAdminFormDictDataEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictDataListGet**
> InlineResponse20029 AdminDictDataListGet(ctx, type_, optional)
获取字典数据列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| 字典类型 | 
 **optional** | ***DictDataApiAdminDictDataListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DictDataApiAdminDictDataListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]

### Return type

[**InlineResponse20029**](inline_response_200_29.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictDataMaxSortGet**
> InlineResponse20021 AdminDictDataMaxSortGet(ctx, type_)
查询字典数据最大排序

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| 字典类型 | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictDataUniqueGet**
> InlineResponse20022 AdminDictDataUniqueGet(ctx, value, type_, optional)
数据键值是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **value** | **string**| 数据键值 | 
  **type_** | **string**| 字典类型 | 
 **optional** | ***DictDataApiAdminDictDataUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DictDataApiAdminDictDataUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **id** | **optional.Int64**| 字典数据ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictDataViewGet**
> InlineResponse20030 AdminDictDataViewGet(ctx, id)
获取指定字典数据信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 字典数据ID | 

### Return type

[**InlineResponse20030**](inline_response_200_30.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictTypeDeletePost**
> InlineResponse2007 AdminDictTypeDeletePost(ctx, body)
删除字典类型

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDictTypeDeleteReq**](XbuyAppFormAdminFormDictTypeDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictTypeEditPost**
> InlineResponse2007 AdminDictTypeEditPost(ctx, body)
修改/新增字典类型

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDictTypeEditReq**](XbuyAppFormAdminFormDictTypeEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictTypeExportGet**
> InlineResponse2007 AdminDictTypeExportGet(ctx, optional)
导出字典类型

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DictDataApiAdminDictTypeExportGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DictDataApiAdminDictTypeExportGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **startTime** | **optional.String**| 开始日期 | 
 **endTime** | **optional.String**| 结束日期 | 
 **status** | **optional.Int32**| 状态 | 
 **name** | **optional.String**| 字典名称 | 
 **type_** | **optional.String**| 字典类型 | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictTypeListGet**
> InlineResponse20031 AdminDictTypeListGet(ctx, optional)
获取字典类型列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DictDataApiAdminDictTypeListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DictDataApiAdminDictTypeListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **startTime** | **optional.String**| 开始日期 | 
 **endTime** | **optional.String**| 结束日期 | 
 **status** | **optional.Int32**| 状态 | 
 **name** | **optional.String**| 字典名称 | 
 **type_** | **optional.String**| 字典类型 | 

### Return type

[**InlineResponse20031**](inline_response_200_31.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictTypeRefreshCacheGet**
> InlineResponse2007 AdminDictTypeRefreshCacheGet(ctx, )
刷新字典缓存

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictTypeUniqueGet**
> InlineResponse20022 AdminDictTypeUniqueGet(ctx, type_, optional)
类型是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| 字典类型 | 
 **optional** | ***DictDataApiAdminDictTypeUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DictDataApiAdminDictTypeUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 字典类型ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminDictTypeViewGet**
> InlineResponse20032 AdminDictTypeViewGet(ctx, id)
获取指定字典类型信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **string**| 字典类型ID | 

### Return type

[**InlineResponse20032**](inline_response_200_32.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ApiDictAttributeGet**
> InlineResponse20028 ApiDictAttributeGet(ctx, type_)
获取指定字典类型的属性数据

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| 字典类型 | 

### Return type

[**InlineResponse20028**](inline_response_200_28.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

