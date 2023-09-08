# XbuyAppModelEntityMallBenefitCondition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 主键 | [optional] [default to null]
**CodeId** | **int32** | 折扣码ID | [optional] [default to null]
**BenefitAmount** | **int32** | 优惠金额  以分为单位存储 | [optional] [default to null]
**FixedPrice** | **int32** | 固定减免金额   以分为单位存储 | [optional] [default to null]
**Discount** | **int32** | 折扣 | [optional] [default to null]
**OffPercent** | **int32** | 减免折扣  discount为90，那此处应为10     此字段+discount 必须等于100 | [optional] [default to null]
**BenefitCount** | **int32** | 满送件数 | [optional] [default to null]
**BenefitSeq** | **string** | 优惠码流水号 | [optional] [default to null]
**BenefitScopeType** | **int32** | 优惠类型范围  1.指定商品；2.指定商品分类；3.指定商品款式(Sku) | [optional] [default to null]
**EffectScopeType** | **int32** | 固定折扣适用商品类型   0所有商品；1.指定商品；2.指定商品分类；3.指定商品款式(Sku) | [optional] [default to null]
**EffectSkuList** | **string** | 折扣sku列表  [{\&quot;good_id\&quot;:\&quot;\&quot;,\&quot;sku_id\&quot;:\&quot;\&quot;}] | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

