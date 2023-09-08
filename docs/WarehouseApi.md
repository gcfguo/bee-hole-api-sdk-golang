# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminWarehouseCreatePost**](WarehouseApi.md#AdminWarehouseCreatePost) | **Post** /admin/warehouse/create | 创建仓库
[**AdminWarehouseDeletePost**](WarehouseApi.md#AdminWarehouseDeletePost) | **Post** /admin/warehouse/delete | 删除仓库
[**AdminWarehouseReadGet**](WarehouseApi.md#AdminWarehouseReadGet) | **Get** /admin/warehouse/read | 读取仓库列表
[**AdminWarehouseStockCreatePost**](WarehouseApi.md#AdminWarehouseStockCreatePost) | **Post** /admin/warehouse/stock/create | 创建仓库商品库存
[**AdminWarehouseStockPagingPost**](WarehouseApi.md#AdminWarehouseStockPagingPost) | **Post** /admin/warehouse/stock/paging | 读取仓库商品库存
[**AdminWarehouseStockReadPost**](WarehouseApi.md#AdminWarehouseStockReadPost) | **Post** /admin/warehouse/stock/read | 读取仓库商品库存
[**AdminWarehouseStockUpdatePost**](WarehouseApi.md#AdminWarehouseStockUpdatePost) | **Post** /admin/warehouse/stock/update | 更新仓库商品库存
[**AdminWarehouseUpdatePost**](WarehouseApi.md#AdminWarehouseUpdatePost) | **Post** /admin/warehouse/update | 更新仓库

# **AdminWarehouseCreatePost**
> InlineResponse200181 AdminWarehouseCreatePost(ctx, body)
创建仓库

创建仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormWareHouseCreateReq**](XbuyAppFormAdminFormWareHouseCreateReq.md)|  | 

### Return type

[**InlineResponse200181**](inline_response_200_181.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminWarehouseDeletePost**
> InlineResponse200182 AdminWarehouseDeletePost(ctx, body)
删除仓库

删除仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormWareHouseDeleteReq**](XbuyAppFormAdminFormWareHouseDeleteReq.md)|  | 

### Return type

[**InlineResponse200182**](inline_response_200_182.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminWarehouseReadGet**
> InlineResponse200183 AdminWarehouseReadGet(ctx, shopId)
读取仓库列表

读取仓库列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **shopId** | **int64**|  | 

### Return type

[**InlineResponse200183**](inline_response_200_183.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminWarehouseStockCreatePost**
> InlineResponse200184 AdminWarehouseStockCreatePost(ctx, body)
创建仓库商品库存

创建仓库商品库存

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormWareHouseStockCreateReq**](XbuyAppFormAdminFormWareHouseStockCreateReq.md)|  | 

### Return type

[**InlineResponse200184**](inline_response_200_184.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminWarehouseStockPagingPost**
> InlineResponse200185 AdminWarehouseStockPagingPost(ctx, body)
读取仓库商品库存

读取仓库商品库存

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormWareHouseStockPagingReq**](XbuyAppFormAdminFormWareHouseStockPagingReq.md)|  | 

### Return type

[**InlineResponse200185**](inline_response_200_185.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminWarehouseStockReadPost**
> InlineResponse200186 AdminWarehouseStockReadPost(ctx, body)
读取仓库商品库存

读取仓库商品库存

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormWareHouseStockReadReq**](XbuyAppFormAdminFormWareHouseStockReadReq.md)|  | 

### Return type

[**InlineResponse200186**](inline_response_200_186.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminWarehouseStockUpdatePost**
> InlineResponse200184 AdminWarehouseStockUpdatePost(ctx, body)
更新仓库商品库存

更新仓库商品库存

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormWareHouseStockUpdateReq**](XbuyAppFormAdminFormWareHouseStockUpdateReq.md)|  | 

### Return type

[**InlineResponse200184**](inline_response_200_184.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminWarehouseUpdatePost**
> InlineResponse2007 AdminWarehouseUpdatePost(ctx, body)
更新仓库

更新仓库

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormWareHouseUpdateReq**](XbuyAppFormAdminFormWareHouseUpdateReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

