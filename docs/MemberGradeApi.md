# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminMemberGradeAddPost**](MemberGradeApi.md#AdminMemberGradeAddPost) | **Post** /admin/member_grade/add | 创建会员等级
[**AdminMemberGradeBindPost**](MemberGradeApi.md#AdminMemberGradeBindPost) | **Post** /admin/member_grade/bind | 绑定会员等级
[**AdminMemberGradeCheckBindExistPost**](MemberGradeApi.md#AdminMemberGradeCheckBindExistPost) | **Post** /admin/member_grade/check_bind_exist | 检查会员等级绑定是否存在
[**AdminMemberGradeDeletePost**](MemberGradeApi.md#AdminMemberGradeDeletePost) | **Post** /admin/member_grade/delete | 删除会员等级
[**AdminMemberGradeGetBindListByGradePost**](MemberGradeApi.md#AdminMemberGradeGetBindListByGradePost) | **Post** /admin/member_grade/get_bind_list_by_grade | 根据等级id获取会员等级绑定用户列表
[**AdminMemberGradeGetBindListByMemberPost**](MemberGradeApi.md#AdminMemberGradeGetBindListByMemberPost) | **Post** /admin/member_grade/get_bind_list_by_member | 根据会员id获取用户会员等级绑定列表
[**AdminMemberGradeGetListByShopIdPost**](MemberGradeApi.md#AdminMemberGradeGetListByShopIdPost) | **Post** /admin/member_grade/get_list_by_shop_id | 根据店铺id获取会员等级列表
[**AdminMemberGradeGetMemberListByGradeIdPost**](MemberGradeApi.md#AdminMemberGradeGetMemberListByGradeIdPost) | **Post** /admin/member_grade/get_member_list_by_grade_id | 根据等级ID获取会员详细信息
[**AdminMemberGradeGetMemberListByGradeIdWithBindPost**](MemberGradeApi.md#AdminMemberGradeGetMemberListByGradeIdWithBindPost) | **Post** /admin/member_grade/get_member_list_by_grade_id_with_bind | 根据会员等级ID查询对应的用户信息列表，包含用户信息和是否存在绑定关系
[**AdminMemberGradeGetUnbindMemberListByGradeIdPost**](MemberGradeApi.md#AdminMemberGradeGetUnbindMemberListByGradeIdPost) | **Post** /admin/member_grade/get_unbind_member_list_by_grade_id | 根据等级ID获取未绑定的会员详细信息
[**AdminMemberGradeListPost**](MemberGradeApi.md#AdminMemberGradeListPost) | **Post** /admin/member_grade/list | 获取会员等级列表
[**AdminMemberGradeUnbindPost**](MemberGradeApi.md#AdminMemberGradeUnbindPost) | **Post** /admin/member_grade/unbind | 解绑会员等级
[**AdminMemberGradeUpdateBindPost**](MemberGradeApi.md#AdminMemberGradeUpdateBindPost) | **Post** /admin/member_grade/update_bind | 更新会员等级绑定
[**AdminMemberGradeUpdatePost**](MemberGradeApi.md#AdminMemberGradeUpdatePost) | **Post** /admin/member_grade/update | 更新会员等级

# **AdminMemberGradeAddPost**
> InlineResponse20077 AdminMemberGradeAddPost(ctx, body)
创建会员等级

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAddMemberGradeReq**](XbuyAppFormAdminFormAddMemberGradeReq.md)|  | 

### Return type

[**InlineResponse20077**](inline_response_200_77.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeBindPost**
> InlineResponse20078 AdminMemberGradeBindPost(ctx, body)
绑定会员等级

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormBindMemberGradeReq**](XbuyAppFormAdminFormBindMemberGradeReq.md)|  | 

### Return type

[**InlineResponse20078**](inline_response_200_78.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeCheckBindExistPost**
> InlineResponse20079 AdminMemberGradeCheckBindExistPost(ctx, body)
检查会员等级绑定是否存在

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormCheckBindMemberGradeExistReq**](XbuyAppFormAdminFormCheckBindMemberGradeExistReq.md)|  | 

### Return type

[**InlineResponse20079**](inline_response_200_79.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeDeletePost**
> InlineResponse20080 AdminMemberGradeDeletePost(ctx, body)
删除会员等级

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormDeleteMemberGradeReq**](XbuyAppFormAdminFormDeleteMemberGradeReq.md)|  | 

### Return type

[**InlineResponse20080**](inline_response_200_80.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeGetBindListByGradePost**
> InlineResponse20081 AdminMemberGradeGetBindListByGradePost(ctx, body)
根据等级id获取会员等级绑定用户列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetBindMemberGradeListByGradeReq**](XbuyAppFormAdminFormGetBindMemberGradeListByGradeReq.md)|  | 

### Return type

[**InlineResponse20081**](inline_response_200_81.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeGetBindListByMemberPost**
> InlineResponse20082 AdminMemberGradeGetBindListByMemberPost(ctx, body)
根据会员id获取用户会员等级绑定列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetBindMemberGradeListByMemberReq**](XbuyAppFormAdminFormGetBindMemberGradeListByMemberReq.md)|  | 

### Return type

[**InlineResponse20082**](inline_response_200_82.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeGetListByShopIdPost**
> InlineResponse20083 AdminMemberGradeGetListByShopIdPost(ctx, body)
根据店铺id获取会员等级列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberGradeListByShopIdReq**](XbuyAppFormAdminFormGetMemberGradeListByShopIdReq.md)|  | 

### Return type

[**InlineResponse20083**](inline_response_200_83.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeGetMemberListByGradeIdPost**
> InlineResponse20084 AdminMemberGradeGetMemberListByGradeIdPost(ctx, body)
根据等级ID获取会员详细信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberListByGradeIdReq**](XbuyAppFormAdminFormGetMemberListByGradeIdReq.md)|  | 

### Return type

[**InlineResponse20084**](inline_response_200_84.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeGetMemberListByGradeIdWithBindPost**
> InlineResponse20085 AdminMemberGradeGetMemberListByGradeIdWithBindPost(ctx, body)
根据会员等级ID查询对应的用户信息列表，包含用户信息和是否存在绑定关系

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberListByGradeIdWithBindReq**](XbuyAppFormAdminFormGetMemberListByGradeIdWithBindReq.md)|  | 

### Return type

[**InlineResponse20085**](inline_response_200_85.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeGetUnbindMemberListByGradeIdPost**
> InlineResponse20086 AdminMemberGradeGetUnbindMemberListByGradeIdPost(ctx, body)
根据等级ID获取未绑定的会员详细信息

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetUnBindMemberListByGradeIdReq**](XbuyAppFormAdminFormGetUnBindMemberListByGradeIdReq.md)|  | 

### Return type

[**InlineResponse20086**](inline_response_200_86.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeListPost**
> InlineResponse20087 AdminMemberGradeListPost(ctx, body)
获取会员等级列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberGradeListReq**](XbuyAppFormAdminFormGetMemberGradeListReq.md)|  | 

### Return type

[**InlineResponse20087**](inline_response_200_87.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeUnbindPost**
> InlineResponse20088 AdminMemberGradeUnbindPost(ctx, body)
解绑会员等级

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUnBindMemberGradeReq**](XbuyAppFormAdminFormUnBindMemberGradeReq.md)|  | 

### Return type

[**InlineResponse20088**](inline_response_200_88.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeUpdateBindPost**
> InlineResponse20090 AdminMemberGradeUpdateBindPost(ctx, body)
更新会员等级绑定

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUpdateBindMemberGradeReq**](XbuyAppFormAdminFormUpdateBindMemberGradeReq.md)|  | 

### Return type

[**InlineResponse20090**](inline_response_200_90.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberGradeUpdatePost**
> InlineResponse20089 AdminMemberGradeUpdatePost(ctx, body)
更新会员等级

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormUpdateMemberGradeReq**](XbuyAppFormAdminFormUpdateMemberGradeReq.md)|  | 

### Return type

[**InlineResponse20089**](inline_response_200_89.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

