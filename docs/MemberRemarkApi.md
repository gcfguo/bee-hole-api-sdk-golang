# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminMemberRemarkAddPost**](MemberRemarkApi.md#AdminMemberRemarkAddPost) | **Post** /admin/member/remark/add | 添加会员备注
[**AdminMemberRemarkListPost**](MemberRemarkApi.md#AdminMemberRemarkListPost) | **Post** /admin/member/remark/list | 获取会员备注列表
[**AdminMemberRemarkListPost_0**](MemberRemarkApi.md#AdminMemberRemarkListPost_0) | **Post** /admin/member/remark_list | 会员备注列表

# **AdminMemberRemarkAddPost**
> InlineResponse20071 AdminMemberRemarkAddPost(ctx, body)
添加会员备注

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormAddMemberRemarkReq**](XbuyAppFormAdminFormAddMemberRemarkReq.md)|  | 

### Return type

[**InlineResponse20071**](inline_response_200_71.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberRemarkListPost**
> InlineResponse20072 AdminMemberRemarkListPost(ctx, body)
获取会员备注列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberRemarkListReq**](XbuyAppFormAdminFormGetMemberRemarkListReq.md)|  | 

### Return type

[**InlineResponse20072**](inline_response_200_72.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AdminMemberRemarkListPost_0**
> InlineResponse20073 AdminMemberRemarkListPost_0(ctx, body)
会员备注列表

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**XbuyAppFormAdminFormGetMemberRemarkListByMemberIdReq**](XbuyAppFormAdminFormGetMemberRemarkListByMemberIdReq.md)|  | 

### Return type

[**InlineResponse20073**](inline_response_200_73.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

