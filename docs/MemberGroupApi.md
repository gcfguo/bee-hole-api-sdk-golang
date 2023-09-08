# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminMemberGroupAddPost**](MemberGroupApi.md#AdminMemberGroupAddPost) | **Post** /admin/member_group/add | 添加会员组
[**AdminMemberGroupDeletePost**](MemberGroupApi.md#AdminMemberGroupDeletePost) | **Post** /admin/member_group/delete | 删除会员组
[**AdminMemberGroupGetListByIdsPost**](MemberGroupApi.md#AdminMemberGroupGetListByIdsPost) | **Post** /admin/member_group/get_list_by_ids | 根据id获取会员组列表
[**AdminMemberGroupGetListByShopIdPost**](MemberGroupApi.md#AdminMemberGroupGetListByShopIdPost) | **Post** /admin/member_group/get_list_by_shop_id | 根据店铺id获取会员组列表
[**AdminMemberGroupListPost**](MemberGroupApi.md#AdminMemberGroupListPost) | **Post** /admin/member_group/list | 会员组列表
[**AdminMemberGroupRelationAddPost**](MemberGroupApi.md#AdminMemberGroupRelationAddPost) | **Post** /admin/member_group_relation/add | 添加会员组关系
[**AdminMemberGroupRelationDeletePost**](MemberGroupApi.md#AdminMemberGroupRelationDeletePost) | **Post** /admin/member_group_relation/delete | 删除会员组关系
[**AdminMemberGroupRelationGetByIdsPost**](MemberGroupApi.md#AdminMemberGroupRelationGetByIdsPost) | **Post** /admin/member_group_relation/get_by_ids | 根据ids获取会员组关系
[**AdminMemberGroupRelationGetListByGroupIdNotBindPost**](MemberGroupApi.md#AdminMemberGroupRelationGetListByGroupIdNotBindPost) | **Post** /admin/member_group_relation/get_list_by_group_id_not_bind | 根据群组id获取未绑定的会员详细信息列表
[**AdminMemberGroupRelationGetListByGroupIdPost**](MemberGroupApi.md#AdminMemberGroupRelationGetListByGroupIdPost) | **Post** /admin/member_group_relation/get_list_by_group_id | 根据群组id获取已经绑定的会员详细信息列表
[**AdminMemberGroupRelationGetListPost**](MemberGroupApi.md#AdminMemberGroupRelationGetListPost) | **Post** /admin/member_group_relation/get_list | 获取会员组关系列表
[**AdminMemberGroupRelationUpdatePost**](MemberGroupApi.md#AdminMemberGroupRelationUpdatePost) | **Post** /admin/member_group_relation/update | 更新会员组关系
[**AdminMemberGroupUpdatePost**](MemberGroupApi.md#AdminMemberGroupUpdatePost) | **Post** /admin/member_group/update | 更新会员组

# **AdminMemberGroupAddPost**
> InlineResponse20091 AdminMemberGroupAddPost(ctx, body)
添加会员组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAddMemberGroupReq**](XbuyAppFormAdminFormAddMemberGroupReq.md)|  | 

### Return type

[**InlineResponse20091**](inline_response_200_91.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupDeletePost**
> InlineResponse20080 AdminMemberGroupDeletePost(ctx, body)
删除会员组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDeleteMemberGroupReq**](XbuyAppFormAdminFormDeleteMemberGroupReq.md)|  | 

### Return type

[**InlineResponse20080**](inline_response_200_80.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupGetListByIdsPost**
> InlineResponse20092 AdminMemberGroupGetListByIdsPost(ctx, body)
根据id获取会员组列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberGroupListByIdsReq**](XbuyAppFormAdminFormGetMemberGroupListByIdsReq.md)|  | 

### Return type

[**InlineResponse20092**](inline_response_200_92.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupGetListByShopIdPost**
> InlineResponse20093 AdminMemberGroupGetListByShopIdPost(ctx, body)
根据店铺id获取会员组列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberGroupListByShopIdReq**](XbuyAppFormAdminFormGetMemberGroupListByShopIdReq.md)|  | 

### Return type

[**InlineResponse20093**](inline_response_200_93.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupListPost**
> InlineResponse20094 AdminMemberGroupListPost(ctx, body)
会员组列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormMemberGroupListReq**](XbuyAppFormAdminFormMemberGroupListReq.md)|  | 

### Return type

[**InlineResponse20094**](inline_response_200_94.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupRelationAddPost**
> InlineResponse20096 AdminMemberGroupRelationAddPost(ctx, body)
添加会员组关系

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAddMemberGroupRelationReq**](XbuyAppFormAdminFormAddMemberGroupRelationReq.md)|  | 

### Return type

[**InlineResponse20096**](inline_response_200_96.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupRelationDeletePost**
> InlineResponse20097 AdminMemberGroupRelationDeletePost(ctx, body)
删除会员组关系

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDeleteMemberGroupRelationReq**](XbuyAppFormAdminFormDeleteMemberGroupRelationReq.md)|  | 

### Return type

[**InlineResponse20097**](inline_response_200_97.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupRelationGetByIdsPost**
> InlineResponse20098 AdminMemberGroupRelationGetByIdsPost(ctx, body)
根据ids获取会员组关系

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberGroupRelationByIdReq**](XbuyAppFormAdminFormGetMemberGroupRelationByIdReq.md)|  | 

### Return type

[**InlineResponse20098**](inline_response_200_98.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupRelationGetListByGroupIdNotBindPost**
> InlineResponse200101 AdminMemberGroupRelationGetListByGroupIdNotBindPost(ctx, body)
根据群组id获取未绑定的会员详细信息列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdNotBindReq**](XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdNotBindReq.md)|  | 

### Return type

[**InlineResponse200101**](inline_response_200_101.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupRelationGetListByGroupIdPost**
> InlineResponse200100 AdminMemberGroupRelationGetListByGroupIdPost(ctx, body)
根据群组id获取已经绑定的会员详细信息列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdReq**](XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdReq.md)|  | 

### Return type

[**InlineResponse200100**](inline_response_200_100.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupRelationGetListPost**
> InlineResponse20099 AdminMemberGroupRelationGetListPost(ctx, body)
获取会员组关系列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberGroupRelationListReq**](XbuyAppFormAdminFormGetMemberGroupRelationListReq.md)|  | 

### Return type

[**InlineResponse20099**](inline_response_200_99.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupRelationUpdatePost**
> InlineResponse200102 AdminMemberGroupRelationUpdatePost(ctx, body)
更新会员组关系

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUpdateMemberGroupRelationReq**](XbuyAppFormAdminFormUpdateMemberGroupRelationReq.md)|  | 

### Return type

[**InlineResponse200102**](inline_response_200_102.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGroupUpdatePost**
> InlineResponse20095 AdminMemberGroupUpdatePost(ctx, body)
更新会员组

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUpdateMemberGroupReq**](XbuyAppFormAdminFormUpdateMemberGroupReq.md)|  | 

### Return type

[**InlineResponse20095**](inline_response_200_95.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

