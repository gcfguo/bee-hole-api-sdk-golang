# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminConfigCopyrightPost**](ConfigApi.md#AdminConfigCopyrightPost) | **Post** /admin/config/copyright | 蜂洞授权信息
[**AdminConfigDeletePost**](ConfigApi.md#AdminConfigDeletePost) | **Post** /admin/config/delete | 删除配置
[**AdminConfigEditPost**](ConfigApi.md#AdminConfigEditPost) | **Post** /admin/config/edit | 修改/新增配置
[**AdminConfigGetValueGet**](ConfigApi.md#AdminConfigGetValueGet) | **Get** /admin/config/get_value | 获取指定配置键的值
[**AdminConfigListGet**](ConfigApi.md#AdminConfigListGet) | **Get** /admin/config/list | 获取配置列表
[**AdminConfigMaxSortGet**](ConfigApi.md#AdminConfigMaxSortGet) | **Get** /admin/config/max_sort | 配置最大排序
[**AdminConfigNameUniqueGet**](ConfigApi.md#AdminConfigNameUniqueGet) | **Get** /admin/config/name_unique | 配置名称是否唯一
[**AdminConfigSysGetPost**](ConfigApi.md#AdminConfigSysGetPost) | **Post** /admin/config/sys/get | 获取指定系统级别配置的值
[**AdminConfigSysSetPost**](ConfigApi.md#AdminConfigSysSetPost) | **Post** /admin/config/sys/set | 获取指定系统级别配置的值
[**AdminConfigViewGet**](ConfigApi.md#AdminConfigViewGet) | **Get** /admin/config/view | 获取指定信息

# **AdminConfigCopyrightPost**
> InlineResponse20018 AdminConfigCopyrightPost(ctx, body)
蜂洞授权信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCopyrightReq**](XbuyAppFormAdminFormCopyrightReq.md)|  | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminConfigDeletePost**
> InlineResponse2007 AdminConfigDeletePost(ctx, body)
删除配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormConfigDeleteReq**](XbuyAppFormAdminFormConfigDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminConfigEditPost**
> InlineResponse2007 AdminConfigEditPost(ctx, body)
修改/新增配置

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormConfigEditReq**](XbuyAppFormAdminFormConfigEditReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminConfigGetValueGet**
> InlineResponse20019 AdminConfigGetValueGet(ctx, key)
获取指定配置键的值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **key** | **string**| 配置键 | 

### Return type

[**InlineResponse20019**](inline_response_200_19.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminConfigListGet**
> InlineResponse20020 AdminConfigListGet(ctx, optional)
获取配置列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiAdminConfigListGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiAdminConfigListGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **optional.Int32**| 当前页码 | [default to 1]
 **limit** | **optional.Int32**| 每页数量 | [default to 10]
 **startTime** | **optional.String**| 开始日期 | 
 **endTime** | **optional.String**| 结束日期 | 
 **status** | **optional.Int32**| 状态 | 
 **name** | **optional.String**| 配置名称 | 

### Return type

[**InlineResponse20020**](inline_response_200_20.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminConfigMaxSortGet**
> InlineResponse20021 AdminConfigMaxSortGet(ctx, optional)
配置最大排序

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ConfigApiAdminConfigMaxSortGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiAdminConfigMaxSortGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **optional.Int64**| 配置ID | 

### Return type

[**InlineResponse20021**](inline_response_200_21.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminConfigNameUniqueGet**
> InlineResponse20022 AdminConfigNameUniqueGet(ctx, name, optional)
配置名称是否唯一

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| 配置名称 | 
 **optional** | ***ConfigApiAdminConfigNameUniqueGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ConfigApiAdminConfigNameUniqueGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **optional.Int64**| 配置ID | 

### Return type

[**InlineResponse20022**](inline_response_200_22.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminConfigSysGetPost**
> InlineResponse20023 AdminConfigSysGetPost(ctx, body)
获取指定系统级别配置的值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormConfigSysGetReq**](XbuyAppFormAdminFormConfigSysGetReq.md)|  | 

### Return type

[**InlineResponse20023**](inline_response_200_23.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminConfigSysSetPost**
> InlineResponse2007 AdminConfigSysSetPost(ctx, body)
获取指定系统级别配置的值

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormConfigSysSetReq**](XbuyAppFormAdminFormConfigSysSetReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminConfigViewGet**
> InlineResponse20024 AdminConfigViewGet(ctx, id)
获取指定信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| 配置ID | 

### Return type

[**InlineResponse20024**](inline_response_200_24.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

