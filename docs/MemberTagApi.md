# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminMemberTagAddPost**](MemberTagApi.md#AdminMemberTagAddPost) | **Post** /admin/member_tag/add | 添加会员标签
[**AdminMemberTagBindPost**](MemberTagApi.md#AdminMemberTagBindPost) | **Post** /admin/member_tag/bind | 绑定会员标签
[**AdminMemberTagBoundListPost**](MemberTagApi.md#AdminMemberTagBoundListPost) | **Post** /admin/member_tag/bound/list | 获取标签绑定会员列表
[**AdminMemberTagBoundListTagMemberPost**](MemberTagApi.md#AdminMemberTagBoundListTagMemberPost) | **Post** /admin/member_tag/bound/list/tag/member | 根据标签ID查询对应的用户信息列表，包含用户信息和是否存在绑定关系
[**AdminMemberTagBoundListTagPost**](MemberTagApi.md#AdminMemberTagBoundListTagPost) | **Post** /admin/member_tag/bound/list/tag | 标签ID查询已经绑定的会员列表
[**AdminMemberTagBoundMemberListPost**](MemberTagApi.md#AdminMemberTagBoundMemberListPost) | **Post** /admin/member_tag/bound/member/list | 获取会员绑定标签列表
[**AdminMemberTagDeletePost**](MemberTagApi.md#AdminMemberTagDeletePost) | **Post** /admin/member_tag/delete | 删除会员标签
[**AdminMemberTagListPost**](MemberTagApi.md#AdminMemberTagListPost) | **Post** /admin/member_tag/list | 获取会员标签列表
[**AdminMemberTagListShopPost**](MemberTagApi.md#AdminMemberTagListShopPost) | **Post** /admin/member_tag/list/shop | 获取店铺下的会员标签列表
[**AdminMemberTagUnboundListTagPost**](MemberTagApi.md#AdminMemberTagUnboundListTagPost) | **Post** /admin/member_tag/unbound/list/tag | 标签ID查询未绑定的会员列表
[**AdminMemberTagUnboundPost**](MemberTagApi.md#AdminMemberTagUnboundPost) | **Post** /admin/member_tag/unbound | 解绑会员标签
[**AdminMemberTagUpdateNamePost**](MemberTagApi.md#AdminMemberTagUpdateNamePost) | **Post** /admin/member_tag/update/name | 更新标签名称

# **AdminMemberTagAddPost**
> InlineResponse200103 AdminMemberTagAddPost(ctx, body)
添加会员标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAddMemberTagReq**](XbuyAppFormAdminFormAddMemberTagReq.md)|  | 

### Return type

[**InlineResponse200103**](inline_response_200_103.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagBindPost**
> InlineResponse200104 AdminMemberTagBindPost(ctx, body)
绑定会员标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormBindMemberTagReq**](XbuyAppFormAdminFormBindMemberTagReq.md)|  | 

### Return type

[**InlineResponse200104**](inline_response_200_104.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagBoundListPost**
> InlineResponse200105 AdminMemberTagBoundListPost(ctx, body)
获取标签绑定会员列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberTagBoundListReq**](XbuyAppFormAdminFormGetMemberTagBoundListReq.md)|  | 

### Return type

[**InlineResponse200105**](inline_response_200_105.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagBoundListTagMemberPost**
> InlineResponse200107 AdminMemberTagBoundListTagMemberPost(ctx, body)
根据标签ID查询对应的用户信息列表，包含用户信息和是否存在绑定关系

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberTagBoundListByTagIdWithMemberInfoReq**](XbuyAppFormAdminFormGetMemberTagBoundListByTagIdWithMemberInfoReq.md)|  | 

### Return type

[**InlineResponse200107**](inline_response_200_107.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagBoundListTagPost**
> InlineResponse200106 AdminMemberTagBoundListTagPost(ctx, body)
标签ID查询已经绑定的会员列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberTagBoundListByTagIdReq**](XbuyAppFormAdminFormGetMemberTagBoundListByTagIdReq.md)|  | 

### Return type

[**InlineResponse200106**](inline_response_200_106.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagBoundMemberListPost**
> InlineResponse200108 AdminMemberTagBoundMemberListPost(ctx, body)
获取会员绑定标签列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberBoundTagListReq**](XbuyAppFormAdminFormGetMemberBoundTagListReq.md)|  | 

### Return type

[**InlineResponse200108**](inline_response_200_108.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagDeletePost**
> InlineResponse200109 AdminMemberTagDeletePost(ctx, body)
删除会员标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDeleteMemberTagReq**](XbuyAppFormAdminFormDeleteMemberTagReq.md)|  | 

### Return type

[**InlineResponse200109**](inline_response_200_109.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagListPost**
> InlineResponse200110 AdminMemberTagListPost(ctx, body)
获取会员标签列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberTagListReq**](XbuyAppFormAdminFormGetMemberTagListReq.md)|  | 

### Return type

[**InlineResponse200110**](inline_response_200_110.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagListShopPost**
> InlineResponse200111 AdminMemberTagListShopPost(ctx, body)
获取店铺下的会员标签列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberTagListByShopIdReq**](XbuyAppFormAdminFormGetMemberTagListByShopIdReq.md)|  | 

### Return type

[**InlineResponse200111**](inline_response_200_111.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagUnboundListTagPost**
> InlineResponse200113 AdminMemberTagUnboundListTagPost(ctx, body)
标签ID查询未绑定的会员列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberTagUnBoundListByTagIdReq**](XbuyAppFormAdminFormGetMemberTagUnBoundListByTagIdReq.md)|  | 

### Return type

[**InlineResponse200113**](inline_response_200_113.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagUnboundPost**
> InlineResponse200112 AdminMemberTagUnboundPost(ctx, body)
解绑会员标签

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUnBoundMemberTagReq**](XbuyAppFormAdminFormUnBoundMemberTagReq.md)|  | 

### Return type

[**InlineResponse200112**](inline_response_200_112.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberTagUpdateNamePost**
> InlineResponse200109 AdminMemberTagUpdateNamePost(ctx, body)
更新标签名称

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUpdateMemberTagNameReq**](XbuyAppFormAdminFormUpdateMemberTagNameReq.md)|  | 

### Return type

[**InlineResponse200109**](inline_response_200_109.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

