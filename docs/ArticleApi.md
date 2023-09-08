# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminArticleAddPost**](ArticleApi.md#AdminArticleAddPost) | **Post** /admin/article/add | 文章新增/编辑
[**AdminArticleCatAddPost**](ArticleApi.md#AdminArticleCatAddPost) | **Post** /admin/article/cat/add | 类目新增/编辑
[**AdminArticleCatInfoPost**](ArticleApi.md#AdminArticleCatInfoPost) | **Post** /admin/article/cat/info | 类目详情
[**AdminArticleCatListPost**](ArticleApi.md#AdminArticleCatListPost) | **Post** /admin/article/cat/list | 类目列表
[**AdminArticleCatSubAddPost**](ArticleApi.md#AdminArticleCatSubAddPost) | **Post** /admin/article/cat/sub/add | 类目副本新增/编辑
[**AdminArticleDeletePost**](ArticleApi.md#AdminArticleDeletePost) | **Post** /admin/article/delete | 删除文章
[**AdminArticleGetPost**](ArticleApi.md#AdminArticleGetPost) | **Post** /admin/article/get | 查询文章
[**AdminArticleListPost**](ArticleApi.md#AdminArticleListPost) | **Post** /admin/article/list | 文章列表
[**AdminArticleSubAddPost**](ArticleApi.md#AdminArticleSubAddPost) | **Post** /admin/article/sub/add | 文章副本新增/编辑

# **AdminArticleAddPost**
> InlineResponse2003 AdminArticleAddPost(ctx, body)
文章新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormArticleAddReq**](XbuyAppFormAdminFormArticleAddReq.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminArticleCatAddPost**
> InlineResponse2004 AdminArticleCatAddPost(ctx, body)
类目新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormArticleCatAddReq**](XbuyAppFormAdminFormArticleCatAddReq.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminArticleCatInfoPost**
> InlineResponse2005 AdminArticleCatInfoPost(ctx, body)
类目详情

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormArticleCatInfoReq**](XbuyAppFormAdminFormArticleCatInfoReq.md)|  | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminArticleCatListPost**
> InlineResponse2006 AdminArticleCatListPost(ctx, body)
类目列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormArticleCatListReq**](XbuyAppFormAdminFormArticleCatListReq.md)|  | 

### Return type

[**InlineResponse2006**](inline_response_200_6.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminArticleCatSubAddPost**
> InlineResponse2004 AdminArticleCatSubAddPost(ctx, body)
类目副本新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormArticlesCatSubAddReq**](XbuyAppFormAdminFormArticlesCatSubAddReq.md)|  | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminArticleDeletePost**
> InlineResponse2007 AdminArticleDeletePost(ctx, body)
删除文章

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormArticleDeleteReq**](XbuyAppFormAdminFormArticleDeleteReq.md)|  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminArticleGetPost**
> InlineResponse2008 AdminArticleGetPost(ctx, body)
查询文章

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormArticleGetReq**](XbuyAppFormAdminFormArticleGetReq.md)|  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminArticleListPost**
> InlineResponse2009 AdminArticleListPost(ctx, body)
文章列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormArticleListReq**](XbuyAppFormAdminFormArticleListReq.md)|  | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminArticleSubAddPost**
> InlineResponse2003 AdminArticleSubAddPost(ctx, body)
文章副本新增/编辑

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormArticleSubAddReq**](XbuyAppFormAdminFormArticleSubAddReq.md)|  | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

