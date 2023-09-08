# XbuyAppFormAdminFormBenefitCodeAddReq

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityName** | **string** | 活动标题 | [default to null]
**ActivityDesc** | **string** | 活动描述 | [optional] [default to null]
**PromotionCode** | **string** | 折扣码 | [default to null]
**BenefitType** | **int32** | 折扣码类型 | [default to null]
**BenefitEvent** | **string** | 优惠条件 | [default to null]
**BenefitCondition** | **string** | 优惠内容 | [default to null]
**EffectScopeType** | **int32** | 折扣适用商品类型 | [optional] [default to null]
**EffectScopeGoodsList** | **string** | 折扣商品范围列表 | [optional] [default to null]
**TargetUserType** | **int32** | 是否指定客户 | [optional] [default to null]
**TargetUserInfoList** | **string** | 参与优惠客户信息 | [optional] [default to null]
**ActivityStatus** | **int32** | 活动状态 | [default to null]
**OverlayAutoDiscounts** | **int32** | 允许与自动折扣叠加使用 | [optional] [default to null]
**OverlayDiscountCode** | **int32** | 允许与其他折扣码叠加使用 | [optional] [default to null]
**PrerequisiteShippingPriceRange** | **int32** | 免运费金额上限 | [optional] [default to null]
**EntitledCountryTargetSelection** | **string** | 国家免运费 | [default to null]
**EntitledCountryIdList** | **string** | 免运费国家ID列表 | [optional] [default to null]
**OrderUseLimit** | **int32** | 订单限制 | [optional] [default to null]
**PublishNum** | **int32** | 折扣码共能使用次数 | [default to null]
**AcquirePerUserLimit** | **int32** | 每人可用次数 | [default to null]
**DiscountCodeDisplay** | **int32** | 商品详情页折扣码显示 | [default to null]
**MaxBenefitAmount** | **int32** | 是否设置每笔订单最高优惠金额 | [default to null]
**MaxAccumulateNum** | **int32** | 满送类型设置每个订单的最大使用次数 | [default to null]
**BenefitScopeType** | **int32** | 优惠类型范围 | [optional] [default to null]
**BenefitScopeGoodsList** | **string** | 优惠商品信息 | [optional] [default to null]
**EnvCustomInfo** | **string** | 自定义海报/折扣券样式 | [optional] [default to null]
**EffectSkuList** | [**[]XbuyAppDaoEffectSkuList**](xbuy.app.dao.EffectSkuList.md) | 折扣sku列表 | [optional] [default to null]
**StartTime** | **string** | 开始时间 | [default to null]
**EndTime** | **string** | 结束时间 | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

