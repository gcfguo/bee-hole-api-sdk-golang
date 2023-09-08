# XbuyAppDaoBenefitCode

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | 折扣码主键 | [optional] [default to null]
**ShopId** | **int64** | 店铺ID | [optional] [default to null]
**ActivityName** | **string** | 活动名称 | [optional] [default to null]
**ActivityDesc** | **string** | 活动描述 | [optional] [default to null]
**PromotionCode** | **string** | 折扣码 | [optional] [default to null]
**BenefitType** | **int32** | 折扣码类型  1百分比折扣（满100打9折）  2固定金额（满100减10元）  3买X送Y（买2件送1件）  4免运费（满3件免运费） | [optional] [default to null]
**BenefitEvent** | **string** | 优惠条件 {\&quot;type\&quot;:\&quot;\&quot;,\&quot;minThreshold\&quot;:\&quot;\&quot;}   满金额类型以分作为单位存储 | [optional] [default to null]
**TargetUserType** | **int32** | 是否指定客户  0否  1是 | [optional] [default to null]
**ActivityStatus** | **int32** | 活动状态  0关闭  1开启 | [optional] [default to null]
**OverlayAutoDiscounts** | **int32** | 允许与自动折扣叠加使用  0 &#x3D; false,  1 &#x3D; true | [optional] [default to null]
**OverlayDiscountCode** | **int32** | 允许与其他折扣码叠加使用  0 &#x3D; false,  1 &#x3D; true | [optional] [default to null]
**PrerequisiteShippingPriceRange** | **int32** | 免运费金额上限  -1为不限制，正数为限制 | [optional] [default to null]
**MaxBenefitAmount** | **int32** | 设置每笔订单最高优惠金额  -1为不限制，正数为限制 | [optional] [default to null]
**EntitledCountryTargetSelection** | **string** | 国家免运费  all &#x3D; 所有国家都免，entitled &#x3D; 部分国家免运费 | [optional] [default to null]
**EntitledCountryIdList** | **string** | 免运费国家ID列表 | [optional] [default to null]
**OrderUseLimit** | **int32** | 订单限制，0 &#x3D; false, 1 &#x3D; true | [optional] [default to null]
**PublishNum** | **int32** | 折扣码共能使用次数 | [optional] [default to null]
**AcquirePerUserLimit** | **int32** | 每人可用次数 | [optional] [default to null]
**DiscountCodeDisplay** | **int32** | 商品详情页折扣码显示, 0 &#x3D; false, 1 &#x3D; true | [optional] [default to null]
**MaxAccumulateNum** | **int32** | 满送类型设置每个订单的最大使用次数   -1 &#x3D; 不限制，正数 &#x3D; 限制 | [optional] [default to null]
**BenefitScopeType** | **int32** | 优惠类型范围  1.指定商品；2.指定商品分类；3.指定商品款式(Sku) | [optional] [default to null]
**EffectScopeType** | **int32** | 固定折扣适用商品类型   0所有商品；1.指定商品；2.指定商品分类；3.指定商品款式(Sku) | [optional] [default to null]
**StartTime** | **string** | 开始时间 | [optional] [default to null]
**EndTime** | **string** | 结束时间 | [optional] [default to null]
**BenefitCondition** | [***XbuyAppModelEntityMallBenefitCondition**](xbuy.app.model.entity.MallBenefitCondition.md) |  | [optional] [default to null]
**EnvCustomInfo** | [***XbuyAppModelEntityMallBenefitEnvCustomInfo**](xbuy.app.model.entity.MallBenefitEnvCustomInfo.md) |  | [optional] [default to null]
**BenefitScopeGoodsList** | [**[]XbuyAppDaoBenefitWith**](xbuy.app.dao.BenefitWith.md) |  | [optional] [default to null]
**EffectScopeGoodsList** | [**[]XbuyAppDaoEffectWith**](xbuy.app.dao.EffectWith.md) |  | [optional] [default to null]
**TargetUserInfoList** | [**[]XbuyAppDaoTargetWith**](xbuy.app.dao.TargetWith.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

