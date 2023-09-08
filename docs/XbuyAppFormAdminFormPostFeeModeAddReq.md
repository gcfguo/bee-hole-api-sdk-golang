# XbuyAppFormAdminFormPostFeeModeAddReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarryId** | **int64** | 运费方式编号 | [optional] [default to null]
**FareId** | **int64** | 运费模板编号 | [optional] [default to null]
**ShopId** | **int64** | 店铺编号 | [optional] [default to null]
**CarryName** | **string** | 运费方式名称 | [optional] [default to null]
**Remark** | **string** | 说明 | [optional] [default to null]
**BillingMethod** | **int32** | 计费方式 1&#x3D;固定计费 2&#x3D;首续重 3&#x3D;首续件 | [optional] [default to null]
**FixedFreight** | **string** | 固定运费 专属billing_methods&#x3D;1时使用 | [optional] [default to null]
**FirstPiece** | **string** | 首重、首件 | [optional] [default to null]
**FirstAmount** | **string** | 首重(首件)费用 | [optional] [default to null]
**SecondPiece** | **string** | 续重、续件 | [optional] [default to null]
**SecondAmount** | **string** | 续重(续件)费用 | [optional] [default to null]
**ExtraCondition** | **int32** | 是否启用额外条件 条件区间 0&#x3D;关闭 1&#x3D;开启 | [optional] [default to null]
**ConditionInterval** | **int32** | 条件区间 1&#x3D;按商品总价 2&#x3D;按包裹重量 3&#x3D;按商品件数 | [optional] [default to null]
**CreateTime** | **string** | 创建时间 | [optional] [default to null]
**UpdateTime** | **string** | 更新时间 | [optional] [default to null]
**CarrySubMode** | **string** | 启用运费方式子条件 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

