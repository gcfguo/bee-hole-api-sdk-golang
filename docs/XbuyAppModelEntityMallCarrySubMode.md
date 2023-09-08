# XbuyAppModelEntityMallCarrySubMode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | [optional] [default to null]
**CarryId** | **int64** | 运费方式编号 | [optional] [default to null]
**BillingMethod** | **int32** | 计费方式 1&#x3D;固定计费 2&#x3D;首续重 3&#x3D;首续件 | [optional] [default to null]
**ConditionInterval** | **int32** | 条件区间 1&#x3D;按商品总价 2&#x3D;按包裹重量 3&#x3D;按商品件数 | [optional] [default to null]
**StartInterval** | **string** | 开始 | [optional] [default to null]
**EndInterval** | **string** | 结束 | [optional] [default to null]
**FixedFreight** | **float64** | 固定运费，专属billing_methods&#x3D;1时使用 | [optional] [default to null]
**FirstPiece** | **string** | 首重、首件 | [optional] [default to null]
**FirstAmount** | **float64** | 首重(首件)费用 | [optional] [default to null]
**SecondPiece** | **string** | 续重、续件 | [optional] [default to null]
**SecondAmount** | **float64** | 续重(续件)费用 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

