# Go API client

xbuy大黄蜂买手项目

- API version: 
- Package version: 1.0.0
- Build package: io.swagger.codegen.v3.generators.go.GoClientCodegen
For more information, please visit [https://www.dhfapp.com](https://www.dhfapp.com)

## Installation
如何使用请参考example下example_test文件的示例代码


## Documentation for API Endpoints

All URIs are relative to */*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**ApiBaseExportGet**](docs/DefaultApi.md#apibaseexportget) | **Get** /api/base/export | 导出字典类型
*DefaultApi* | [**ApiBaseIpLocationGet**](docs/DefaultApi.md#apibaseiplocationget) | **Get** /api/base/ip_location | 获取IP归属地信息
*DefaultApi* | [**ApiBaseLangGet**](docs/DefaultApi.md#apibaselangget) | **Get** /api/base/lang | 获取lang信息
*DefaultApi* | [**ApiLogClearPost**](docs/DefaultApi.md#apilogclearpost) | **Post** /api/log/clear | 清空日志
*DefaultApi* | [**ApiLogExamineReceivePost**](docs/DefaultApi.md#apilogexaminereceivepost) | **Post** /api/log/examine/receive | 接收申报回执
*DefaultApi* | [**ApiLogExportGet**](docs/DefaultApi.md#apilogexportget) | **Get** /api/log/export | 导出日志
*DefaultApi* | [**ApiLogListGet**](docs/DefaultApi.md#apiloglistget) | **Get** /api/log/list | 获取日志列表
*DefaultApi* | [**ApiLoginCheckGet**](docs/DefaultApi.md#apilogincheckget) | **Get** /api/login/check | 登录效验
*DefaultApi* | [**ApiLoginLogoutGet**](docs/DefaultApi.md#apiloginlogoutget) | **Get** /api/login/logout | 注销登录
*DefaultApi* | [**ApiLoginSignPost**](docs/DefaultApi.md#apiloginsignpost) | **Post** /api/login/sign | 提交登录
*DefaultApi* | [**ApiMemberProfileGet**](docs/DefaultApi.md#apimemberprofileget) | **Get** /api/member/profile | 获取登录用户的基本信息
*DefaultApi* | [**ApiOrderReceiveshipPost**](docs/DefaultApi.md#apiorderreceiveshippost) | **Post** /api/order/receiveship | 接收运单回执
*DefaultApi* | [**ApiSkuReceiveratePost**](docs/DefaultApi.md#apiskureceiveratepost) | **Post** /api/sku/receiverate | 接收商品税率
*DefaultApi* | [**MallAddressAddPost**](docs/DefaultApi.md#malladdressaddpost) | **Post** /mall/address/add | 收货地址添加和更新
*DefaultApi* | [**MallAddressDeletePost**](docs/DefaultApi.md#malladdressdeletepost) | **Post** /mall/address/delete | 删除收货地址
*DefaultApi* | [**MallAddressListPost**](docs/DefaultApi.md#malladdresslistpost) | **Post** /mall/address/list | 收货地址列表
*DefaultApi* | [**MallAddressReceivingCountryPost**](docs/DefaultApi.md#malladdressreceivingcountrypost) | **Post** /mall/address/receiving_country | 收货国家列表
*DefaultApi* | [**MallArticleCatListPost**](docs/DefaultApi.md#mallarticlecatlistpost) | **Post** /mall/article/cat/list | 类目列表
*DefaultApi* | [**MallArticleGetPost**](docs/DefaultApi.md#mallarticlegetpost) | **Post** /mall/article/get | 查询文章
*DefaultApi* | [**MallArticleListPost**](docs/DefaultApi.md#mallarticlelistpost) | **Post** /mall/article/list | 商城底部菜单列表
*DefaultApi* | [**MallArticleNewListPost**](docs/DefaultApi.md#mallarticlenewlistpost) | **Post** /mall/article/new_list | 文章列表
*DefaultApi* | [**MallBenefitCodeOperatePost**](docs/DefaultApi.md#mallbenefitcodeoperatepost) | **Post** /mall/benefit_code/operate | 使用折扣码
*DefaultApi* | [**MallCartAddPost**](docs/DefaultApi.md#mallcartaddpost) | **Post** /mall/cart/add | 添加购物车
*DefaultApi* | [**MallCartClearPost**](docs/DefaultApi.md#mallcartclearpost) | **Post** /mall/cart/clear | 清空购物车
*DefaultApi* | [**MallCartGetPost**](docs/DefaultApi.md#mallcartgetpost) | **Post** /mall/cart/get | 获取购物车商品信息
*DefaultApi* | [**MallCartNumPost**](docs/DefaultApi.md#mallcartnumpost) | **Post** /mall/cart/num | 获取商品数量
*DefaultApi* | [**MallCartRemovePost**](docs/DefaultApi.md#mallcartremovepost) | **Post** /mall/cart/remove | 购物车商品移出
*DefaultApi* | [**MallCartUpdatePost**](docs/DefaultApi.md#mallcartupdatepost) | **Post** /mall/cart/update | 购物车商品数量更新
*DefaultApi* | [**MallCatInfoPost**](docs/DefaultApi.md#mallcatinfopost) | **Post** /mall/cat/info | 类目详情
*DefaultApi* | [**MallCatListPost**](docs/DefaultApi.md#mallcatlistpost) | **Post** /mall/cat/list | 类目列表
*DefaultApi* | [**MallConfigAllPost**](docs/DefaultApi.md#mallconfigallpost) | **Post** /mall/config/all | 配置列表
*DefaultApi* | [**MallConfigListPost**](docs/DefaultApi.md#mallconfiglistpost) | **Post** /mall/config/list | 获取店铺单个配置
*DefaultApi* | [**MallFavAddPost**](docs/DefaultApi.md#mallfavaddpost) | **Post** /mall/fav/add | 添加收藏
*DefaultApi* | [**MallFavDeletePost**](docs/DefaultApi.md#mallfavdeletepost) | **Post** /mall/fav/delete | 取消收藏
*DefaultApi* | [**MallFavListPost**](docs/DefaultApi.md#mallfavlistpost) | **Post** /mall/fav/list | 收藏列表
*DefaultApi* | [**MallFavNumPost**](docs/DefaultApi.md#mallfavnumpost) | **Post** /mall/fav/num | 收藏总数量
*DefaultApi* | [**MallGoodInfoPost**](docs/DefaultApi.md#mallgoodinfopost) | **Post** /mall/good/info | 商品详情
*DefaultApi* | [**MallGoodListPost**](docs/DefaultApi.md#mallgoodlistpost) | **Post** /mall/good/list | 商品列表
*DefaultApi* | [**MallGoodSameCatPost**](docs/DefaultApi.md#mallgoodsamecatpost) | **Post** /mall/good/same/cat | 同类目商品
*DefaultApi* | [**MallOrderCreatePost**](docs/DefaultApi.md#mallordercreatepost) | **Post** /mall/order/create | 订单创建
*DefaultApi* | [**MallOrderInsuredPost**](docs/DefaultApi.md#mallorderinsuredpost) | **Post** /mall/order/insured | 生成保险单
*DefaultApi* | [**MallOrderListPost**](docs/DefaultApi.md#mallorderlistpost) | **Post** /mall/order/list | 订单列表
*DefaultApi* | [**MallOrderPayerUpdatePost**](docs/DefaultApi.md#mallorderpayerupdatepost) | **Post** /mall/order/payer/update | 更新支付人信息
*DefaultApi* | [**MallOrderPaymentCredentialPost**](docs/DefaultApi.md#mallorderpaymentcredentialpost) | **Post** /mall/order/payment/credential | 生成支付凭据
*DefaultApi* | [**MallOrderPaymentPost**](docs/DefaultApi.md#mallorderpaymentpost) | **Post** /mall/order/payment | 订单支付
*DefaultApi* | [**MallOrderPreorderPost**](docs/DefaultApi.md#mallorderpreorderpost) | **Post** /mall/order/preorder | 商品预下单
*DefaultApi* | [**MallOrderPrepayListPost**](docs/DefaultApi.md#mallorderprepaylistpost) | **Post** /mall/order/prepay/list | 获取预支付订单
*DefaultApi* | [**MallOrderQueryGet**](docs/DefaultApi.md#mallorderqueryget) | **Get** /mall/order/query | 交易查询
*DefaultApi* | [**MallOrderSalesOrderCreatePost**](docs/DefaultApi.md#mallordersalesordercreatepost) | **Post** /mall/order/sales_order/create | 生成售后单
*DefaultApi* | [**MallOrderSalesOrderListPost**](docs/DefaultApi.md#mallordersalesorderlistpost) | **Post** /mall/order/sales_order/list | 售后单列表
*DefaultApi* | [**MallOrderWaybillListPost**](docs/DefaultApi.md#mallorderwaybilllistpost) | **Post** /mall/order/waybill/list | 订单物流
*DefaultApi* | [**MallPmethodGetPost**](docs/DefaultApi.md#mallpmethodgetpost) | **Post** /mall/pmethod/get | 获取支付方式
*DefaultApi* | [**MallReviewAddPost**](docs/DefaultApi.md#mallreviewaddpost) | **Post** /mall/review/add | 订单评论增加
*DefaultApi* | [**MallReviewListPost**](docs/DefaultApi.md#mallreviewlistpost) | **Post** /mall/review/list | 订单评论列表
*DefaultApi* | [**MallShopConfigAllPost**](docs/DefaultApi.md#mallshopconfigallpost) | **Post** /mall/shop-config/all | 店铺配置列表
*DefaultApi* | [**MallShopConfigGetPost**](docs/DefaultApi.md#mallshopconfiggetpost) | **Post** /mall/shop-config/get | 店铺配置单个查询
*DefaultApi* | [**MallShopListGet**](docs/DefaultApi.md#mallshoplistget) | **Get** /mall/shop/list | 店铺列表
*DefaultApi* | [**MallShopViewPost**](docs/DefaultApi.md#mallshopviewpost) | **Post** /mall/shop/view | 列表
*DefaultApi* | [**MallStandardCatPost**](docs/DefaultApi.md#mallstandardcatpost) | **Post** /mall/standard/cat | 商城标准类目列表
*DefaultApi* | [**MallSubscribesAddPost**](docs/DefaultApi.md#mallsubscribesaddpost) | **Post** /mall/subscribes/add | 添加订阅
*DefaultApi* | [**MallSubscribesDeletePost**](docs/DefaultApi.md#mallsubscribesdeletepost) | **Post** /mall/subscribes/delete | 删除订阅
*DefaultApi* | [**MallSubscribesListPost**](docs/DefaultApi.md#mallsubscribeslistpost) | **Post** /mall/subscribes/list | 订阅列表
*DefaultApi* | [**MallTransAddPost**](docs/DefaultApi.md#malltransaddpost) | **Post** /mall/trans/add | 前端翻译添加
*DefaultApi* | [**NotifyPaidPost**](docs/DefaultApi.md#notifypaidpost) | **Post** /notify/paid | 
*DefaultApi* | [**NotifyRefundedPost**](docs/DefaultApi.md#notifyrefundedpost) | **Post** /notify/refunded | 
*DefaultApi* | [**OauthFacebookCallbackGet**](docs/DefaultApi.md#oauthfacebookcallbackget) | **Get** /oauth/facebook/callback | Facebook授权回调
*DefaultApi* | [**OauthFacebookGet**](docs/DefaultApi.md#oauthfacebookget) | **Get** /oauth/facebook | Facebook授权
*DefaultApi* | [**OauthGoogleCallbackGet**](docs/DefaultApi.md#oauthgooglecallbackget) | **Get** /oauth/google/callback | Google授权回调
*DefaultApi* | [**OauthGoogleGet**](docs/DefaultApi.md#oauthgoogleget) | **Get** /oauth/google | Google授权
*DefaultApi* | [**OauthMicrosoftCallbackGet**](docs/DefaultApi.md#oauthmicrosoftcallbackget) | **Get** /oauth/microsoft/callback | 微软授权回调
*DefaultApi* | [**OauthMicrosoftGet**](docs/DefaultApi.md#oauthmicrosoftget) | **Get** /oauth/microsoft | 微软授权
*DefaultApi* | [**OauthTwitterCallbackGet**](docs/DefaultApi.md#oauthtwittercallbackget) | **Get** /oauth/twitter/callback | Twitter授权回调
*DefaultApi* | [**OauthTwitterGet**](docs/DefaultApi.md#oauthtwitterget) | **Get** /oauth/twitter | Twitter授权
*DefaultApi* | [**OauthWechatCallbackGet**](docs/DefaultApi.md#oauthwechatcallbackget) | **Get** /oauth/wechat/callback | 微信授权回调
*DefaultApi* | [**OauthWechatCheckGet**](docs/DefaultApi.md#oauthwechatcheckget) | **Get** /oauth/wechat/check | 微信授权检测
*DefaultApi* | [**OauthWechatGet**](docs/DefaultApi.md#oauthwechatget) | **Get** /oauth/wechat | 微信授权
*AfterSalesApi* | [**AdminAfterSalesListPost**](docs/AfterSalesApi.md#adminaftersaleslistpost) | **Post** /admin/after_sales/list | 售后单列表
*AfterSalesApi* | [**AdminAfterSalesRefundPost**](docs/AfterSalesApi.md#adminaftersalesrefundpost) | **Post** /admin/after_sales/refund | 发起退款
*AfterSalesApi* | [**AdminAfterSalesSwitchPost**](docs/AfterSalesApi.md#adminaftersalesswitchpost) | **Post** /admin/after_sales/switch | 启用或关闭售后
*AiApi* | [**AdminAiTransPost**](docs/AiApi.md#adminaitranspost) | **Post** /admin/ai/trans | 翻译
*ArticleApi* | [**AdminArticleAddPost**](docs/ArticleApi.md#adminarticleaddpost) | **Post** /admin/article/add | 文章新增/编辑
*ArticleApi* | [**AdminArticleCatAddPost**](docs/ArticleApi.md#adminarticlecataddpost) | **Post** /admin/article/cat/add | 类目新增/编辑
*ArticleApi* | [**AdminArticleCatInfoPost**](docs/ArticleApi.md#adminarticlecatinfopost) | **Post** /admin/article/cat/info | 类目详情
*ArticleApi* | [**AdminArticleCatListPost**](docs/ArticleApi.md#adminarticlecatlistpost) | **Post** /admin/article/cat/list | 类目列表
*ArticleApi* | [**AdminArticleCatSubAddPost**](docs/ArticleApi.md#adminarticlecatsubaddpost) | **Post** /admin/article/cat/sub/add | 类目副本新增/编辑
*ArticleApi* | [**AdminArticleDeletePost**](docs/ArticleApi.md#adminarticledeletepost) | **Post** /admin/article/delete | 删除文章
*ArticleApi* | [**AdminArticleGetPost**](docs/ArticleApi.md#adminarticlegetpost) | **Post** /admin/article/get | 查询文章
*ArticleApi* | [**AdminArticleListPost**](docs/ArticleApi.md#adminarticlelistpost) | **Post** /admin/article/list | 文章列表
*ArticleApi* | [**AdminArticleSubAddPost**](docs/ArticleApi.md#adminarticlesubaddpost) | **Post** /admin/article/sub/add | 文章副本新增/编辑
*BenefitCodeApi* | [**AdminBenefitCodeAddPost**](docs/BenefitCodeApi.md#adminbenefitcodeaddpost) | **Post** /admin/benefit_code/add | 添加折扣码
*BenefitCodeApi* | [**AdminBenefitCodeCountryListPost**](docs/BenefitCodeApi.md#adminbenefitcodecountrylistpost) | **Post** /admin/benefit_code/country_list | 国家列表
*BenefitCodeApi* | [**AdminBenefitCodeListPost**](docs/BenefitCodeApi.md#adminbenefitcodelistpost) | **Post** /admin/benefit_code/list | 折扣码列表
*BenefitCodeApi* | [**AdminBenefitCodeStatisticsPost**](docs/BenefitCodeApi.md#adminbenefitcodestatisticspost) | **Post** /admin/benefit_code/statistics | 统计信息
*BenefitCodeApi* | [**AdminBenefitCodeZendCountryListPost**](docs/BenefitCodeApi.md#adminbenefitcodezendcountrylistpost) | **Post** /admin/benefit_code/zend_country_list | 树状国家列表
*CatApi* | [**AdminCatAddPost**](docs/CatApi.md#admincataddpost) | **Post** /admin/cat/add | 新增/编辑类目
*CatApi* | [**AdminCatDeletePost**](docs/CatApi.md#admincatdeletepost) | **Post** /admin/cat/delete | 删除类目
*CatApi* | [**AdminCatInfoPost**](docs/CatApi.md#admincatinfopost) | **Post** /admin/cat/info | 类目详情
*CatApi* | [**AdminCatListPost**](docs/CatApi.md#admincatlistpost) | **Post** /admin/cat/list | 类目列表
*CatApi* | [**AdminCatParentListPost**](docs/CatApi.md#admincatparentlistpost) | **Post** /admin/cat/parent/list | 父级类目列表
*CatApi* | [**AdminCatSubAddPost**](docs/CatApi.md#admincatsubaddpost) | **Post** /admin/cat/sub/add | 类目副本新增/编辑
*ConfigApi* | [**AdminConfigCopyrightPost**](docs/ConfigApi.md#adminconfigcopyrightpost) | **Post** /admin/config/copyright | 蜂洞授权信息
*ConfigApi* | [**AdminConfigDeletePost**](docs/ConfigApi.md#adminconfigdeletepost) | **Post** /admin/config/delete | 删除配置
*ConfigApi* | [**AdminConfigEditPost**](docs/ConfigApi.md#adminconfigeditpost) | **Post** /admin/config/edit | 修改/新增配置
*ConfigApi* | [**AdminConfigGetValueGet**](docs/ConfigApi.md#adminconfiggetvalueget) | **Get** /admin/config/get_value | 获取指定配置键的值
*ConfigApi* | [**AdminConfigListGet**](docs/ConfigApi.md#adminconfiglistget) | **Get** /admin/config/list | 获取配置列表
*ConfigApi* | [**AdminConfigMaxSortGet**](docs/ConfigApi.md#adminconfigmaxsortget) | **Get** /admin/config/max_sort | 配置最大排序
*ConfigApi* | [**AdminConfigNameUniqueGet**](docs/ConfigApi.md#adminconfignameuniqueget) | **Get** /admin/config/name_unique | 配置名称是否唯一
*ConfigApi* | [**AdminConfigSysGetPost**](docs/ConfigApi.md#adminconfigsysgetpost) | **Post** /admin/config/sys/get | 获取指定系统级别配置的值
*ConfigApi* | [**AdminConfigSysSetPost**](docs/ConfigApi.md#adminconfigsyssetpost) | **Post** /admin/config/sys/set | 获取指定系统级别配置的值
*ConfigApi* | [**AdminConfigViewGet**](docs/ConfigApi.md#adminconfigviewget) | **Get** /admin/config/view | 获取指定信息
*DefaultApi* | [**DocsAdminGet**](docs/DefaultApi.md#docsadminget) | **Get** /docs/admin | 
*DefaultApi* | [**DocsMallGet**](docs/DefaultApi.md#docsmallget) | **Get** /docs/mall | 
*DefaultApi* | [**DocsOpenGet**](docs/DefaultApi.md#docsopenget) | **Get** /docs/open | 
*DeptApi* | [**AdminDeptDeletePost**](docs/DeptApi.md#admindeptdeletepost) | **Post** /admin/dept/delete | 删除部门
*DeptApi* | [**AdminDeptEditPost**](docs/DeptApi.md#admindepteditpost) | **Post** /admin/dept/edit | 修改/新增部门
*DeptApi* | [**AdminDeptListGet**](docs/DeptApi.md#admindeptlistget) | **Get** /admin/dept/list | 获取部门列表
*DeptApi* | [**AdminDeptListTreeGet**](docs/DeptApi.md#admindeptlisttreeget) | **Get** /admin/dept/list_tree | 获取部门列表树
*DeptApi* | [**AdminDeptMaxSortGet**](docs/DeptApi.md#admindeptmaxsortget) | **Get** /admin/dept/max_sort | 部门最大排序
*DeptApi* | [**AdminDeptNameUniqueGet**](docs/DeptApi.md#admindeptnameuniqueget) | **Get** /admin/dept/name_unique | 部门名称是否唯一
*DeptApi* | [**AdminDeptViewGet**](docs/DeptApi.md#admindeptviewget) | **Get** /admin/dept/view | 获取指定信息
*DictDataApi* | [**AdminDictAttributeGet**](docs/DictDataApi.md#admindictattributeget) | **Get** /admin/dict/attribute | 获取指定字典类型的属性数据
*DictDataApi* | [**AdminDictDataDeletePost**](docs/DictDataApi.md#admindictdatadeletepost) | **Post** /admin/dict_data/delete | 删除字典数据
*DictDataApi* | [**AdminDictDataEditPost**](docs/DictDataApi.md#admindictdataeditpost) | **Post** /admin/dict_data/edit | 修改/新增字典数据
*DictDataApi* | [**AdminDictDataListGet**](docs/DictDataApi.md#admindictdatalistget) | **Get** /admin/dict_data/list | 获取字典数据列表
*DictDataApi* | [**AdminDictDataMaxSortGet**](docs/DictDataApi.md#admindictdatamaxsortget) | **Get** /admin/dict_data/max_sort | 查询字典数据最大排序
*DictDataApi* | [**AdminDictDataUniqueGet**](docs/DictDataApi.md#admindictdatauniqueget) | **Get** /admin/dict_data/unique | 数据键值是否唯一
*DictDataApi* | [**AdminDictDataViewGet**](docs/DictDataApi.md#admindictdataviewget) | **Get** /admin/dict_data/view | 获取指定字典数据信息
*DictDataApi* | [**AdminDictTypeDeletePost**](docs/DictDataApi.md#admindicttypedeletepost) | **Post** /admin/dict_type/delete | 删除字典类型
*DictDataApi* | [**AdminDictTypeEditPost**](docs/DictDataApi.md#admindicttypeeditpost) | **Post** /admin/dict_type/edit | 修改/新增字典类型
*DictDataApi* | [**AdminDictTypeExportGet**](docs/DictDataApi.md#admindicttypeexportget) | **Get** /admin/dict_type/export | 导出字典类型
*DictDataApi* | [**AdminDictTypeListGet**](docs/DictDataApi.md#admindicttypelistget) | **Get** /admin/dict_type/list | 获取字典类型列表
*DictDataApi* | [**AdminDictTypeRefreshCacheGet**](docs/DictDataApi.md#admindicttyperefreshcacheget) | **Get** /admin/dict_type/refresh_cache | 刷新字典缓存
*DictDataApi* | [**AdminDictTypeUniqueGet**](docs/DictDataApi.md#admindicttypeuniqueget) | **Get** /admin/dict_type/unique | 类型是否唯一
*DictDataApi* | [**AdminDictTypeViewGet**](docs/DictDataApi.md#admindicttypeviewget) | **Get** /admin/dict_type/view | 获取指定字典类型信息
*DictDataApi* | [**ApiDictAttributeGet**](docs/DictDataApi.md#apidictattributeget) | **Get** /api/dict/attribute | 获取指定字典类型的属性数据
*DomainApi* | [**AdminDomainAddPost**](docs/DomainApi.md#admindomainaddpost) | **Post** /admin/domain/add | 域名新增/编辑
*DomainApi* | [**AdminDomainDeletePost**](docs/DomainApi.md#admindomaindeletepost) | **Post** /admin/domain/delete | 删除域名
*DomainApi* | [**AdminDomainGetMainPost**](docs/DomainApi.md#admindomaingetmainpost) | **Post** /admin/domain/get_main | 查询主域名
*DomainApi* | [**AdminDomainGetPost**](docs/DomainApi.md#admindomaingetpost) | **Post** /admin/domain/get | 查询域名
*DomainApi* | [**AdminDomainListPost**](docs/DomainApi.md#admindomainlistpost) | **Post** /admin/domain/list | 域名列表
*EmailApi* | [**AdminEmailBindPost**](docs/EmailApi.md#adminemailbindpost) | **Post** /admin/email/bind | 绑定邮箱
*EmailApi* | [**AdminEmailListPost**](docs/EmailApi.md#adminemaillistpost) | **Post** /admin/email/list | 获取邮件列表
*EmailApi* | [**AdminEmailPullPost**](docs/EmailApi.md#adminemailpullpost) | **Post** /admin/email/pull | 拉取邮件
*EmailApi* | [**AdminEmailSendPost**](docs/EmailApi.md#adminemailsendpost) | **Post** /admin/email/send | 发送邮件
*EmailApi* | [**AdminEmailSettingPost**](docs/EmailApi.md#adminemailsettingpost) | **Post** /admin/email/setting | 邮件服务设置
*ExpressApi* | [**AdminExpressAddPost**](docs/ExpressApi.md#adminexpressaddpost) | **Post** /admin/express/add | 快递公司新增/编辑
*ExpressApi* | [**AdminExpressDeletePost**](docs/ExpressApi.md#adminexpressdeletepost) | **Post** /admin/express/delete | 快递公司删除
*ExpressApi* | [**AdminExpressGetPost**](docs/ExpressApi.md#adminexpressgetpost) | **Post** /admin/express/get | 查询快递公司
*ExpressApi* | [**AdminExpressListPost**](docs/ExpressApi.md#adminexpresslistpost) | **Post** /admin/express/list | 快递公司列表
*GoodDeleteApi* | [**AdminGoodDeletePost**](docs/GoodDeleteApi.md#admingooddeletepost) | **Post** /admin/good/delete | 删除商品
*GoodDeleteApi* | [**AdminGoodInfoValueDeleteOnePost**](docs/GoodDeleteApi.md#admingoodinfovaluedeleteonepost) | **Post** /admin/good/info/value/delete/one | 删除商品参数值
*GoodDeleteApi* | [**AdminSkuDeletePost**](docs/GoodDeleteApi.md#adminskudeletepost) | **Post** /admin/sku/delete | 删除sku
*GoodDeleteApi* | [**AdminSkuPropDeletePost**](docs/GoodDeleteApi.md#adminskupropdeletepost) | **Post** /admin/sku/prop/delete | 删除规格
*GoodDeleteApi* | [**AdminSkuValueDeletePost**](docs/GoodDeleteApi.md#adminskuvaluedeletepost) | **Post** /admin/sku/value/delete | 删除规格值
*GoodDeleteApi* | [**AdminTaxRateDeletePost**](docs/GoodDeleteApi.md#admintaxratedeletepost) | **Post** /admin/tax_rate/delete | 删除sku税率
*GoodEditApi* | [**AdminGoodAddPost**](docs/GoodEditApi.md#admingoodaddpost) | **Post** /admin/good/add | 商品新增/编辑
*GoodEditApi* | [**AdminGoodFilingAddPost**](docs/GoodEditApi.md#admingoodfilingaddpost) | **Post** /admin/good/filing/add | 商品备案
*GoodEditApi* | [**AdminGoodOnShelvesPost**](docs/GoodEditApi.md#admingoodonshelvespost) | **Post** /admin/good/on_shelves | 商品上下架
*GoodEditApi* | [**AdminGoodSubAddPost**](docs/GoodEditApi.md#admingoodsubaddpost) | **Post** /admin/good/sub/add | 副本商品新增/编辑
*GoodEditApi* | [**AdminSkuAddPost**](docs/GoodEditApi.md#adminskuaddpost) | **Post** /admin/sku/add | sku新增或编辑
*GoodEditApi* | [**AdminSkuPropsAddPost**](docs/GoodEditApi.md#adminskupropsaddpost) | **Post** /admin/sku/props/add | 规格新增/编辑
*GoodEditApi* | [**AdminSkuPropsSubAddPost**](docs/GoodEditApi.md#adminskupropssubaddpost) | **Post** /admin/sku/props/sub/add | 规格副本新增/编辑
*GoodEditApi* | [**AdminSkuSubAddPost**](docs/GoodEditApi.md#adminskusubaddpost) | **Post** /admin/sku/sub/add | sku副本新增或编辑
*GoodEditApi* | [**AdminSkuValueAddPost**](docs/GoodEditApi.md#adminskuvalueaddpost) | **Post** /admin/sku/value/add | 规格值新增/编辑
*GoodEditApi* | [**AdminSkuValueSubAddPost**](docs/GoodEditApi.md#adminskuvaluesubaddpost) | **Post** /admin/sku/value/sub/add | 规格值副本新增/编辑
*GoodEditApi* | [**AdminTaxRateCreatePost**](docs/GoodEditApi.md#admintaxratecreatepost) | **Post** /admin/tax_rate/create | 创建sku税率
*GoodEditApi* | [**AdminTaxRateUpdatePost**](docs/GoodEditApi.md#admintaxrateupdatepost) | **Post** /admin/tax_rate/update | 更新sku税率
*GoodStockApi* | [**AdminSkuStockGetPost**](docs/GoodStockApi.md#adminskustockgetpost) | **Post** /admin/sku/stock/get | 获取sku库存
*GoodStockApi* | [**AdminSkuStockUpdatePost**](docs/GoodStockApi.md#adminskustockupdatepost) | **Post** /admin/sku/stock/update | 更新sku库存
*GoodViewApi* | [**AdminGoodInfoListPost**](docs/GoodViewApi.md#admingoodinfolistpost) | **Post** /admin/good/info/list | 标准类目列表
*GoodViewApi* | [**AdminGoodInfoPost**](docs/GoodViewApi.md#admingoodinfopost) | **Post** /admin/good/info | 商品详情
*GoodViewApi* | [**AdminGoodListPost**](docs/GoodViewApi.md#admingoodlistpost) | **Post** /admin/good/list | 商品列表
*GoodViewApi* | [**AdminSkuListPost**](docs/GoodViewApi.md#adminskulistpost) | **Post** /admin/sku/list | sku列表
*GoodViewApi* | [**AdminSkuPropsListPost**](docs/GoodViewApi.md#adminskupropslistpost) | **Post** /admin/sku/props/list | 规格列表
*GoodViewApi* | [**AdminTaxRateGetRatePost**](docs/GoodViewApi.md#admintaxrategetratepost) | **Post** /admin/tax_rate/get_rate | 获取sku税率
*GoodViewApi* | [**AdminTaxRateRateListPost**](docs/GoodViewApi.md#admintaxrateratelistpost) | **Post** /admin/tax_rate/rate_list | sku税率列表
*GoodViewApi* | [**AdminTaxRateReadPost**](docs/GoodViewApi.md#admintaxratereadpost) | **Post** /admin/tax_rate/read | 查询sku税率
*GoogleApi* | [**AdminGoogleAddPost**](docs/GoogleApi.md#admingoogleaddpost) | **Post** /admin/google/add | 添加域名
*GoogleApi* | [**AdminGoogleListPost**](docs/GoogleApi.md#admingooglelistpost) | **Post** /admin/google/list | 获取谷歌站点列表
*GoogleApi* | [**AdminGoogleOauthCallbackGet**](docs/GoogleApi.md#admingoogleoauthcallbackget) | **Get** /admin/google/oauth/callback | 谷歌授权回调
*GoogleApi* | [**AdminGoogleSearchPost**](docs/GoogleApi.md#admingooglesearchpost) | **Post** /admin/google/search | 获取网站搜索分析数据
*GoogleApi* | [**AdminGoogleSubmitPost**](docs/GoogleApi.md#admingooglesubmitpost) | **Post** /admin/google/submit | 提交网站地图
*GoogleApi* | [**AdminGoogleUrlindexPost**](docs/GoogleApi.md#admingoogleurlindexpost) | **Post** /admin/google/urlindex | 获取网址索引
*LangApi* | [**AdminLangGetPost**](docs/LangApi.md#adminlanggetpost) | **Post** /admin/lang/get | 获取语言
*LogApi* | [**AdminLogClearPost**](docs/LogApi.md#adminlogclearpost) | **Post** /admin/log/clear | 清空日志
*LogApi* | [**AdminLogExportGet**](docs/LogApi.md#adminlogexportget) | **Get** /admin/log/export | 导出日志
*LogApi* | [**AdminLogListGet**](docs/LogApi.md#adminloglistget) | **Get** /admin/log/list | 获取日志列表
*LoginApi* | [**AdminLoginCaptchaGet**](docs/LoginApi.md#adminlogincaptchaget) | **Get** /admin/login/captcha | 获取登录验证码
*LoginApi* | [**AdminLoginLogoutPost**](docs/LoginApi.md#adminloginlogoutpost) | **Post** /admin/login/logout | 注销登录
*LoginApi* | [**AdminLoginSignPost**](docs/LoginApi.md#adminloginsignpost) | **Post** /admin/login/sign | 商户登录
*LoginApi* | [**AdminLoginSubAuthPost**](docs/LoginApi.md#adminloginsubauthpost) | **Post** /admin/login/sub/auth | 子账户授权
*LoginApi* | [**AdminLoginSubSignPost**](docs/LoginApi.md#adminloginsubsignpost) | **Post** /admin/login/sub/sign | 商户子账号登录
*MallConfigApi* | [**AdminMallconfigListPost**](docs/MallConfigApi.md#adminmallconfiglistpost) | **Post** /admin/mallconfig/list | 商城配置列表
*MallConfigApi* | [**AdminMallconfigSavePost**](docs/MallConfigApi.md#adminmallconfigsavepost) | **Post** /admin/mallconfig/save | 保存商城配置
*MemberApi* | [**AdminMemberBindMemberToShopPost**](docs/MemberApi.md#adminmemberbindmembertoshoppost) | **Post** /admin/member/bind_member_to_shop | 绑定会员到商店
*MemberApi* | [**AdminMemberGetMemberByGroupIdAndShopIdPost**](docs/MemberApi.md#adminmembergetmemberbygroupidandshopidpost) | **Post** /admin/member/get_member_by_group_id_and_shop_id | 根据分组ID和商店ID获取会员
*MemberApi* | [**AdminMemberGetMemberByLevelIdAndShopIdPost**](docs/MemberApi.md#adminmembergetmemberbylevelidandshopidpost) | **Post** /admin/member/get_member_by_level_id_and_shop_id | 根据等级ID和商店ID获取会员
*MemberApi* | [**AdminMemberGetMemberByTagIdAndShopIdPost**](docs/MemberApi.md#adminmembergetmemberbytagidandshopidpost) | **Post** /admin/member/get_member_by_tag_id_and_shop_id | 根据标签ID和商店ID获取会员
*MemberApi* | [**AdminMemberGetMemberListPost**](docs/MemberApi.md#adminmembergetmemberlistpost) | **Post** /admin/member/get_member_list | 获取会员列表
*MemberApi* | [**AdminMemberIsShopBlacklistPost**](docs/MemberApi.md#adminmemberisshopblacklistpost) | **Post** /admin/member/is_shop_blacklist | 是否在商店黑名单
*MemberApi* | [**AdminMemberShopBlacklistAddPost**](docs/MemberApi.md#adminmembershopblacklistaddpost) | **Post** /admin/member/shop_blacklist/add | 将用户加入商店黑名单
*MemberApi* | [**AdminMemberShopBlacklistPost**](docs/MemberApi.md#adminmembershopblacklistpost) | **Post** /admin/member/shop_blacklist | 指定商店黑名单
*MemberApi* | [**AdminMemberShopBlacklistRemovePost**](docs/MemberApi.md#adminmembershopblacklistremovepost) | **Post** /admin/member/shop_blacklist/remove | 将用户移出商店黑名单
*MemberApi* | [**AdminMemberShopMemberListPost**](docs/MemberApi.md#adminmembershopmemberlistpost) | **Post** /admin/member/shop_member_list | 商城会员列表
*MemberApi* | [**AdminMemberUnbindMemberToShopPost**](docs/MemberApi.md#adminmemberunbindmembertoshoppost) | **Post** /admin/member/unbind_member_to_shop | 解绑会员到商店
*MemberApi* | [**MallMemberCaptchaGet**](docs/MemberApi.md#mallmembercaptchaget) | **Get** /mall/member/captcha | 验证码
*MemberApi* | [**MallMemberChangePwdPost**](docs/MemberApi.md#mallmemberchangepwdpost) | **Post** /mall/member/change_pwd | 修改密码
*MemberApi* | [**MallMemberEmailVerifyPost**](docs/MemberApi.md#mallmemberemailverifypost) | **Post** /mall/member/email/verify | 验证注册邮箱
*MemberApi* | [**MallMemberLoginPost**](docs/MemberApi.md#mallmemberloginpost) | **Post** /mall/member/login | 登录
*MemberApi* | [**MallMemberLogoutPost**](docs/MemberApi.md#mallmemberlogoutpost) | **Post** /mall/member/logout | 注销
*MemberApi* | [**MallMemberPhoneVerifyPost**](docs/MemberApi.md#mallmemberphoneverifypost) | **Post** /mall/member/phone/verify | 验证注册手机号
*MemberApi* | [**MallMemberProfilePost**](docs/MemberApi.md#mallmemberprofilepost) | **Post** /mall/member/profile | 个人信息
*MemberApi* | [**MallMemberRefreshTokenPost**](docs/MemberApi.md#mallmemberrefreshtokenpost) | **Post** /mall/member/refresh_token | 刷新token
*MemberApi* | [**MallMemberRegisterPost**](docs/MemberApi.md#mallmemberregisterpost) | **Post** /mall/member/register | 注册
*MemberApi* | [**MallMemberResetPasswordApplyPost**](docs/MemberApi.md#mallmemberresetpasswordapplypost) | **Post** /mall/member/reset_password/apply | 重置密码
*MemberApi* | [**MallMemberResetPasswordConfirmPost**](docs/MemberApi.md#mallmemberresetpasswordconfirmpost) | **Post** /mall/member/reset_password/confirm | 重置密码确认
*MemberApi* | [**MallMemberUpdatePost**](docs/MemberApi.md#mallmemberupdatepost) | **Post** /mall/member/update | 更新
*MemberApi* | [**MallMemberVerifyEmailApplyPost**](docs/MemberApi.md#mallmemberverifyemailapplypost) | **Post** /mall/member/verify_email/apply | 验证邮箱
*MemberApi* | [**MallMemberVerifyEmailConfirmPost**](docs/MemberApi.md#mallmemberverifyemailconfirmpost) | **Post** /mall/member/verify_email/confirm | 验证邮箱确认
*MemberGradeApi* | [**AdminMemberGradeAddPost**](docs/MemberGradeApi.md#adminmembergradeaddpost) | **Post** /admin/member_grade/add | 创建会员等级
*MemberGradeApi* | [**AdminMemberGradeBindPost**](docs/MemberGradeApi.md#adminmembergradebindpost) | **Post** /admin/member_grade/bind | 绑定会员等级
*MemberGradeApi* | [**AdminMemberGradeCheckBindExistPost**](docs/MemberGradeApi.md#adminmembergradecheckbindexistpost) | **Post** /admin/member_grade/check_bind_exist | 检查会员等级绑定是否存在
*MemberGradeApi* | [**AdminMemberGradeDeletePost**](docs/MemberGradeApi.md#adminmembergradedeletepost) | **Post** /admin/member_grade/delete | 删除会员等级
*MemberGradeApi* | [**AdminMemberGradeGetBindListByGradePost**](docs/MemberGradeApi.md#adminmembergradegetbindlistbygradepost) | **Post** /admin/member_grade/get_bind_list_by_grade | 根据等级id获取会员等级绑定用户列表
*MemberGradeApi* | [**AdminMemberGradeGetBindListByMemberPost**](docs/MemberGradeApi.md#adminmembergradegetbindlistbymemberpost) | **Post** /admin/member_grade/get_bind_list_by_member | 根据会员id获取用户会员等级绑定列表
*MemberGradeApi* | [**AdminMemberGradeGetListByShopIdPost**](docs/MemberGradeApi.md#adminmembergradegetlistbyshopidpost) | **Post** /admin/member_grade/get_list_by_shop_id | 根据店铺id获取会员等级列表
*MemberGradeApi* | [**AdminMemberGradeGetMemberListByGradeIdPost**](docs/MemberGradeApi.md#adminmembergradegetmemberlistbygradeidpost) | **Post** /admin/member_grade/get_member_list_by_grade_id | 根据等级ID获取会员详细信息
*MemberGradeApi* | [**AdminMemberGradeGetMemberListByGradeIdWithBindPost**](docs/MemberGradeApi.md#adminmembergradegetmemberlistbygradeidwithbindpost) | **Post** /admin/member_grade/get_member_list_by_grade_id_with_bind | 根据会员等级ID查询对应的用户信息列表，包含用户信息和是否存在绑定关系
*MemberGradeApi* | [**AdminMemberGradeGetUnbindMemberListByGradeIdPost**](docs/MemberGradeApi.md#adminmembergradegetunbindmemberlistbygradeidpost) | **Post** /admin/member_grade/get_unbind_member_list_by_grade_id | 根据等级ID获取未绑定的会员详细信息
*MemberGradeApi* | [**AdminMemberGradeListPost**](docs/MemberGradeApi.md#adminmembergradelistpost) | **Post** /admin/member_grade/list | 获取会员等级列表
*MemberGradeApi* | [**AdminMemberGradeUnbindPost**](docs/MemberGradeApi.md#adminmembergradeunbindpost) | **Post** /admin/member_grade/unbind | 解绑会员等级
*MemberGradeApi* | [**AdminMemberGradeUpdateBindPost**](docs/MemberGradeApi.md#adminmembergradeupdatebindpost) | **Post** /admin/member_grade/update_bind | 更新会员等级绑定
*MemberGradeApi* | [**AdminMemberGradeUpdatePost**](docs/MemberGradeApi.md#adminmembergradeupdatepost) | **Post** /admin/member_grade/update | 更新会员等级
*MemberGroupApi* | [**AdminMemberGroupAddPost**](docs/MemberGroupApi.md#adminmembergroupaddpost) | **Post** /admin/member_group/add | 添加会员组
*MemberGroupApi* | [**AdminMemberGroupDeletePost**](docs/MemberGroupApi.md#adminmembergroupdeletepost) | **Post** /admin/member_group/delete | 删除会员组
*MemberGroupApi* | [**AdminMemberGroupGetListByIdsPost**](docs/MemberGroupApi.md#adminmembergroupgetlistbyidspost) | **Post** /admin/member_group/get_list_by_ids | 根据id获取会员组列表
*MemberGroupApi* | [**AdminMemberGroupGetListByShopIdPost**](docs/MemberGroupApi.md#adminmembergroupgetlistbyshopidpost) | **Post** /admin/member_group/get_list_by_shop_id | 根据店铺id获取会员组列表
*MemberGroupApi* | [**AdminMemberGroupListPost**](docs/MemberGroupApi.md#adminmembergrouplistpost) | **Post** /admin/member_group/list | 会员组列表
*MemberGroupApi* | [**AdminMemberGroupRelationAddPost**](docs/MemberGroupApi.md#adminmembergrouprelationaddpost) | **Post** /admin/member_group_relation/add | 添加会员组关系
*MemberGroupApi* | [**AdminMemberGroupRelationDeletePost**](docs/MemberGroupApi.md#adminmembergrouprelationdeletepost) | **Post** /admin/member_group_relation/delete | 删除会员组关系
*MemberGroupApi* | [**AdminMemberGroupRelationGetByIdsPost**](docs/MemberGroupApi.md#adminmembergrouprelationgetbyidspost) | **Post** /admin/member_group_relation/get_by_ids | 根据ids获取会员组关系
*MemberGroupApi* | [**AdminMemberGroupRelationGetListByGroupIdNotBindPost**](docs/MemberGroupApi.md#adminmembergrouprelationgetlistbygroupidnotbindpost) | **Post** /admin/member_group_relation/get_list_by_group_id_not_bind | 根据群组id获取未绑定的会员详细信息列表
*MemberGroupApi* | [**AdminMemberGroupRelationGetListByGroupIdPost**](docs/MemberGroupApi.md#adminmembergrouprelationgetlistbygroupidpost) | **Post** /admin/member_group_relation/get_list_by_group_id | 根据群组id获取已经绑定的会员详细信息列表
*MemberGroupApi* | [**AdminMemberGroupRelationGetListPost**](docs/MemberGroupApi.md#adminmembergrouprelationgetlistpost) | **Post** /admin/member_group_relation/get_list | 获取会员组关系列表
*MemberGroupApi* | [**AdminMemberGroupRelationUpdatePost**](docs/MemberGroupApi.md#adminmembergrouprelationupdatepost) | **Post** /admin/member_group_relation/update | 更新会员组关系
*MemberGroupApi* | [**AdminMemberGroupUpdatePost**](docs/MemberGroupApi.md#adminmembergroupupdatepost) | **Post** /admin/member_group/update | 更新会员组
*MemberRemarkApi* | [**AdminMemberRemarkAddPost**](docs/MemberRemarkApi.md#adminmemberremarkaddpost) | **Post** /admin/member/remark/add | 添加会员备注
*MemberRemarkApi* | [**AdminMemberRemarkListPost**](docs/MemberRemarkApi.md#adminmemberremarklistpost) | **Post** /admin/member/remark/list | 获取会员备注列表
*MemberRemarkApi* | [**AdminMemberRemarkListPost_0**](docs/MemberRemarkApi.md#adminmemberremarklistpost_0) | **Post** /admin/member/remark_list | 会员备注列表
*MemberTagApi* | [**AdminMemberTagAddPost**](docs/MemberTagApi.md#adminmembertagaddpost) | **Post** /admin/member_tag/add | 添加会员标签
*MemberTagApi* | [**AdminMemberTagBindPost**](docs/MemberTagApi.md#adminmembertagbindpost) | **Post** /admin/member_tag/bind | 绑定会员标签
*MemberTagApi* | [**AdminMemberTagBoundListPost**](docs/MemberTagApi.md#adminmembertagboundlistpost) | **Post** /admin/member_tag/bound/list | 获取标签绑定会员列表
*MemberTagApi* | [**AdminMemberTagBoundListTagMemberPost**](docs/MemberTagApi.md#adminmembertagboundlisttagmemberpost) | **Post** /admin/member_tag/bound/list/tag/member | 根据标签ID查询对应的用户信息列表，包含用户信息和是否存在绑定关系
*MemberTagApi* | [**AdminMemberTagBoundListTagPost**](docs/MemberTagApi.md#adminmembertagboundlisttagpost) | **Post** /admin/member_tag/bound/list/tag | 标签ID查询已经绑定的会员列表
*MemberTagApi* | [**AdminMemberTagBoundMemberListPost**](docs/MemberTagApi.md#adminmembertagboundmemberlistpost) | **Post** /admin/member_tag/bound/member/list | 获取会员绑定标签列表
*MemberTagApi* | [**AdminMemberTagDeletePost**](docs/MemberTagApi.md#adminmembertagdeletepost) | **Post** /admin/member_tag/delete | 删除会员标签
*MemberTagApi* | [**AdminMemberTagListPost**](docs/MemberTagApi.md#adminmembertaglistpost) | **Post** /admin/member_tag/list | 获取会员标签列表
*MemberTagApi* | [**AdminMemberTagListShopPost**](docs/MemberTagApi.md#adminmembertaglistshoppost) | **Post** /admin/member_tag/list/shop | 获取店铺下的会员标签列表
*MemberTagApi* | [**AdminMemberTagUnboundListTagPost**](docs/MemberTagApi.md#adminmembertagunboundlisttagpost) | **Post** /admin/member_tag/unbound/list/tag | 标签ID查询未绑定的会员列表
*MemberTagApi* | [**AdminMemberTagUnboundPost**](docs/MemberTagApi.md#adminmembertagunboundpost) | **Post** /admin/member_tag/unbound | 解绑会员标签
*MemberTagApi* | [**AdminMemberTagUpdateNamePost**](docs/MemberTagApi.md#adminmembertagupdatenamepost) | **Post** /admin/member_tag/update/name | 更新标签名称
*MenuApi* | [**AdminMenuCodeUniqueGet**](docs/MenuApi.md#adminmenucodeuniqueget) | **Get** /admin/menu/code_unique | 菜单编码是否唯一
*MenuApi* | [**AdminMenuDeletePost**](docs/MenuApi.md#adminmenudeletepost) | **Post** /admin/menu/delete | 删除菜单
*MenuApi* | [**AdminMenuEditPost**](docs/MenuApi.md#adminmenueditpost) | **Post** /admin/menu/edit | 修改/新增菜单
*MenuApi* | [**AdminMenuListGet**](docs/MenuApi.md#adminmenulistget) | **Get** /admin/menu/list | 获取菜单列表
*MenuApi* | [**AdminMenuMaxSortGet**](docs/MenuApi.md#adminmenumaxsortget) | **Get** /admin/menu/max_sort | 菜单最大排序
*MenuApi* | [**AdminMenuNameUniqueGet**](docs/MenuApi.md#adminmenunameuniqueget) | **Get** /admin/menu/name_unique | 菜单名称是否唯一
*MenuApi* | [**AdminMenuRoleListGet**](docs/MenuApi.md#adminmenurolelistget) | **Get** /admin/menu/role_list | 查询角色菜单列表
*MenuApi* | [**AdminMenuSearchListGet**](docs/MenuApi.md#adminmenusearchlistget) | **Get** /admin/menu/search_list | 获取菜单列表
*MenuApi* | [**AdminMenuSubRoleListGet**](docs/MenuApi.md#adminmenusubrolelistget) | **Get** /admin/menu/sub_role_list | 查询角色菜单列表
*MenuApi* | [**AdminMenuViewGet**](docs/MenuApi.md#adminmenuviewget) | **Get** /admin/menu/view | 获取指定菜单信息
*MerchantApi* | [**AdminMemberBlacklistAddPost**](docs/MerchantApi.md#adminmemberblacklistaddpost) | **Post** /admin/member/blacklist/add | 将用户加入黑名单(全局）
*MerchantApi* | [**AdminMemberBlacklistPost**](docs/MerchantApi.md#adminmemberblacklistpost) | **Post** /admin/member/blacklist | 全局黑名单
*MerchantApi* | [**AdminMemberBlacklistRemovePost**](docs/MerchantApi.md#adminmemberblacklistremovepost) | **Post** /admin/member/blacklist/remove | 将用户移出黑名单(全局)
*MerchantApi* | [**AdminMemberDeletePost**](docs/MerchantApi.md#adminmemberdeletepost) | **Post** /admin/member/delete | 删除会员
*MerchantApi* | [**AdminMemberEditPost**](docs/MerchantApi.md#adminmembereditpost) | **Post** /admin/member/edit | 修改/新增会员
*MerchantApi* | [**AdminMemberEmailUniqueGet**](docs/MerchantApi.md#adminmemberemailuniqueget) | **Get** /admin/member/email_unique | 邮箱是否唯一
*MerchantApi* | [**AdminMemberEmailVerifyPost**](docs/MerchantApi.md#adminmemberemailverifypost) | **Post** /admin/member/email/verify | 验证注册邮箱
*MerchantApi* | [**AdminMemberInfoGet**](docs/MerchantApi.md#adminmemberinfoget) | **Get** /admin/member/info | 获取登录用户信息
*MerchantApi* | [**AdminMemberListGet**](docs/MerchantApi.md#adminmemberlistget) | **Get** /admin/member/list | 获取会员列表
*MerchantApi* | [**AdminMemberMaxSortGet**](docs/MerchantApi.md#adminmembermaxsortget) | **Get** /admin/member/max_sort | 会员最大排序
*MerchantApi* | [**AdminMemberMobileUniqueGet**](docs/MerchantApi.md#adminmembermobileuniqueget) | **Get** /admin/member/mobile_unique | 手机号是否唯一
*MerchantApi* | [**AdminMemberNameUniqueGet**](docs/MerchantApi.md#adminmembernameuniqueget) | **Get** /admin/member/name_unique | 会员名称是否唯一
*MerchantApi* | [**AdminMemberProfileGet**](docs/MerchantApi.md#adminmemberprofileget) | **Get** /admin/member/profile | 获取登录用户的基本信息
*MerchantApi* | [**AdminMemberRegisterPost**](docs/MerchantApi.md#adminmemberregisterpost) | **Post** /admin/member/register | 用户注册
*MerchantApi* | [**AdminMemberResetPwdPost**](docs/MerchantApi.md#adminmemberresetpwdpost) | **Post** /admin/member/reset_pwd | 重置密码
*MerchantApi* | [**AdminMemberUpdateAvatarPost**](docs/MerchantApi.md#adminmemberupdateavatarpost) | **Post** /admin/member/update_avatar | 更新会员头像
*MerchantApi* | [**AdminMemberUpdateProfilePost**](docs/MerchantApi.md#adminmemberupdateprofilepost) | **Post** /admin/member/update_profile | 更新会员资料
*MerchantApi* | [**AdminMemberUpdatePwdPost**](docs/MerchantApi.md#adminmemberupdatepwdpost) | **Post** /admin/member/update_pwd | 重置密码
*MerchantApi* | [**AdminMemberViewGet**](docs/MerchantApi.md#adminmemberviewget) | **Get** /admin/member/view | 获取指定信息
*NoticeApi* | [**AdminNoticeDeletePost**](docs/NoticeApi.md#adminnoticedeletepost) | **Post** /admin/notice/delete | 删除公告
*NoticeApi* | [**AdminNoticeEditPost**](docs/NoticeApi.md#adminnoticeeditpost) | **Post** /admin/notice/edit | 修改/新增公告
*NoticeApi* | [**AdminNoticeListGet**](docs/NoticeApi.md#adminnoticelistget) | **Get** /admin/notice/list | 获取公告列表
*NoticeApi* | [**AdminNoticeMaxSortGet**](docs/NoticeApi.md#adminnoticemaxsortget) | **Get** /admin/notice/max_sort | 公告最大排序
*NoticeApi* | [**AdminNoticeNameUniqueGet**](docs/NoticeApi.md#adminnoticenameuniqueget) | **Get** /admin/notice/name_unique | 公告名称是否唯一
*NoticeApi* | [**AdminNoticeViewGet**](docs/NoticeApi.md#adminnoticeviewget) | **Get** /admin/notice/view | 获取指定信息
*OrderApi* | [**AdminOrderAfterSalesApplyPost**](docs/OrderApi.md#adminorderaftersalesapplypost) | **Post** /admin/order/after_sales/apply | 申请售后
*OrderApi* | [**AdminOrderCancelAfterSalesPost**](docs/OrderApi.md#adminordercancelaftersalespost) | **Post** /admin/order/cancel_after_sales | 取消售后
*OrderApi* | [**AdminOrderExportPost**](docs/OrderApi.md#adminorderexportpost) | **Post** /admin/order/export | 订单导出
*OrderApi* | [**AdminOrderImportPost**](docs/OrderApi.md#adminorderimportpost) | **Post** /admin/order/import | 订单导入
*OrderApi* | [**AdminOrderInsurePost**](docs/OrderApi.md#adminorderinsurepost) | **Post** /admin/order/insure | 保单推送日志
*OrderApi* | [**AdminOrderItemsPost**](docs/OrderApi.md#adminorderitemspost) | **Post** /admin/order/items | 订单商品
*OrderApi* | [**AdminOrderListPost**](docs/OrderApi.md#adminorderlistpost) | **Post** /admin/order/list | 订单列表
*OrderApi* | [**AdminOrderPushLogPost**](docs/OrderApi.md#adminorderpushlogpost) | **Post** /admin/order/push/log | 订单推送日志
*OrderApi* | [**AdminOrderPushPost**](docs/OrderApi.md#adminorderpushpost) | **Post** /admin/order/push | 推送订单
*OrderApi* | [**AdminOrderReceiverUpdatePost**](docs/OrderApi.md#adminorderreceiverupdatepost) | **Post** /admin/order/receiver/update | 更新收货人信息
*OrderApi* | [**AdminOrderShipPost**](docs/OrderApi.md#adminordershippost) | **Post** /admin/order/ship | 订单发货
*OrderApi* | [**AdminOrderWaybillListPost**](docs/OrderApi.md#adminorderwaybilllistpost) | **Post** /admin/order/waybill/list | 订单物流
*OrderApi* | [**MallOrderAfterSalesApplyPost**](docs/OrderApi.md#mallorderaftersalesapplypost) | **Post** /mall/order/after_sales/apply | 申请售后
*PmethodApi* | [**AdminPmethodCurrencyGetGet**](docs/PmethodApi.md#adminpmethodcurrencygetget) | **Get** /admin/pmethod/currency/get | 获取币种类型
*PmethodApi* | [**AdminPmethodCurrencyratesSetPost**](docs/PmethodApi.md#adminpmethodcurrencyratessetpost) | **Post** /admin/pmethod/currencyrates/set | 获取币种汇率关系
*PostApi* | [**AdminPostCodeUniqueGet**](docs/PostApi.md#adminpostcodeuniqueget) | **Get** /admin/post/code_unique | 岗位编码是否唯一
*PostApi* | [**AdminPostDeletePost**](docs/PostApi.md#adminpostdeletepost) | **Post** /admin/post/delete | 删除岗位
*PostApi* | [**AdminPostEditPost**](docs/PostApi.md#adminposteditpost) | **Post** /admin/post/edit | 修改/新增岗位
*PostApi* | [**AdminPostListGet**](docs/PostApi.md#adminpostlistget) | **Get** /admin/post/list | 获取岗位列表
*PostApi* | [**AdminPostMaxSortGet**](docs/PostApi.md#adminpostmaxsortget) | **Get** /admin/post/max_sort | 岗位最大排序
*PostApi* | [**AdminPostNameUniqueGet**](docs/PostApi.md#adminpostnameuniqueget) | **Get** /admin/post/name_unique | 岗位名称是否唯一
*PostApi* | [**AdminPostViewGet**](docs/PostApi.md#adminpostviewget) | **Get** /admin/post/view | 获取指定信息
*PostApi* | [**AdminRolePostEditPost**](docs/PostApi.md#adminroleposteditpost) | **Post** /admin/role/post_edit | 岗位权限菜单编辑
*PostfeeApi* | [**AdminPostfeeAddPost**](docs/PostfeeApi.md#adminpostfeeaddpost) | **Post** /admin/postfee/add | 运费模板新增/编辑
*PostfeeApi* | [**AdminPostfeeListPost**](docs/PostfeeApi.md#adminpostfeelistpost) | **Post** /admin/postfee/list | 运费列表
*PostfeeApi* | [**AdminPostfeeModeAddPost**](docs/PostfeeApi.md#adminpostfeemodeaddpost) | **Post** /admin/postfee/mode-add | 运费方式新增/编辑
*PostfeeApi* | [**AdminPostfeeModeListPost**](docs/PostfeeApi.md#adminpostfeemodelistpost) | **Post** /admin/postfee/mode-list | 运费方式列表
*ResourceApi* | [**AdminResourceGroupAddPost**](docs/ResourceApi.md#adminresourcegroupaddpost) | **Post** /admin/resource/group_add | 添加资源分组
*ResourceApi* | [**AdminResourceGroupAddResourcePost**](docs/ResourceApi.md#adminresourcegroupaddresourcepost) | **Post** /admin/resource/group_add_resource | 分组添加资源
*ResourceApi* | [**AdminResourceGroupBatchAddResourcePost**](docs/ResourceApi.md#adminresourcegroupbatchaddresourcepost) | **Post** /admin/resource/group_batch_add_resource | 分组批量添加资源
*ResourceApi* | [**AdminResourceGroupListPost**](docs/ResourceApi.md#adminresourcegrouplistpost) | **Post** /admin/resource/group_list | 获取所有API分组列表
*ResourceApi* | [**AdminResourceGroupResourceListPost**](docs/ResourceApi.md#adminresourcegroupresourcelistpost) | **Post** /admin/resource/group_resource_list | 分组资源列表
*ResourceApi* | [**AdminResourceListPost**](docs/ResourceApi.md#adminresourcelistpost) | **Post** /admin/resource/list | 获取API列表
*ReviewApi* | [**AdminReviewListPost**](docs/ReviewApi.md#adminreviewlistpost) | **Post** /admin/review/list | 订单评论列表
*RoleApi* | [**AdminRoleDynamicGet**](docs/RoleApi.md#adminroledynamicget) | **Get** /admin/role/dynamic | 获取动态路由
*RoleApi* | [**AdminRoleEditPost**](docs/RoleApi.md#adminroleeditpost) | **Post** /admin/role/edit | 编辑角色菜单权限
*RoleApi* | [**AdminRoleListGet**](docs/RoleApi.md#adminrolelistget) | **Get** /admin/role/list | 获取角色列表
*RoleApi* | [**AdminRoleMemberListGet**](docs/RoleApi.md#adminrolememberlistget) | **Get** /admin/role/member_list | 获取角色下的会员列表
*RoleApi* | [**AdminRoleShopAddPost**](docs/RoleApi.md#adminroleshopaddpost) | **Post** /admin/role/shop/add | 增加店铺授权
*RoleApi* | [**AdminRoleShopDelPost**](docs/RoleApi.md#adminroleshopdelpost) | **Post** /admin/role/shop/del | 移除店铺授权
*RoleApi* | [**AdminRoleShopListPost**](docs/RoleApi.md#adminroleshoplistpost) | **Post** /admin/role/shop/list | 列表店铺授权
*SearchCustomsItemInfoApi* | [**AdminGoodCustomsItemSearchPost**](docs/SearchCustomsItemInfoApi.md#admingoodcustomsitemsearchpost) | **Post** /admin/good/customs_item/search | 查询备案商品信息
*SearchHsCodeInfoApi* | [**AdminGoodHscodeSearchPost**](docs/SearchHsCodeInfoApi.md#admingoodhscodesearchpost) | **Post** /admin/good/hscode/search | 查询hscode信息
*SearchSkuInfoApi* | [**AdminGoodSkuSearchPost**](docs/SearchSkuInfoApi.md#admingoodskusearchpost) | **Post** /admin/good/sku/search | 查询sku信息
*ServiceApi* | [**AdminServiceAddPost**](docs/ServiceApi.md#adminserviceaddpost) | **Post** /admin/service/add | 增值服务添加/编辑
*ServiceApi* | [**AdminServiceDetailPost**](docs/ServiceApi.md#adminservicedetailpost) | **Post** /admin/service/detail | 增值服务详情
*ServiceApi* | [**AdminServiceListPost**](docs/ServiceApi.md#adminservicelistpost) | **Post** /admin/service/list | 增值服务列表
*ServiceApi* | [**AdminServicePausePost**](docs/ServiceApi.md#adminservicepausepost) | **Post** /admin/service/pause | 增值服务暂停/开启
*ShippingpolicyApi* | [**AdminShippingpolicyPriorityCreatePost**](docs/ShippingpolicyApi.md#adminshippingpolicyprioritycreatepost) | **Post** /admin/shippingpolicy/priority/create | 创建优先发货策略
*ShippingpolicyApi* | [**AdminShippingpolicyPriorityUpdatePost**](docs/ShippingpolicyApi.md#adminshippingpolicypriorityupdatepost) | **Post** /admin/shippingpolicy/priority/update | 修改优先发货策略
*ShopApi* | [**AdminRechargeGet**](docs/ShopApi.md#adminrechargeget) | **Get** /admin/recharge | 充值
*ShopApi* | [**AdminShopAddPost**](docs/ShopApi.md#adminshopaddpost) | **Post** /admin/shop/add | 新增店铺
*ShopApi* | [**AdminShopChangeShopPost**](docs/ShopApi.md#adminshopchangeshoppost) | **Post** /admin/shop/change_shop | 切换店铺
*ShopApi* | [**AdminShopFundsGet**](docs/ShopApi.md#adminshopfundsget) | **Get** /admin/shop/funds | 查询店铺资金
*ShopApi* | [**AdminShopGetBindShopPost**](docs/ShopApi.md#adminshopgetbindshoppost) | **Post** /admin/shop/get_bind_shop | 已绑定店铺列表
*ShopApi* | [**AdminShopGetinfoPost**](docs/ShopApi.md#adminshopgetinfopost) | **Post** /admin/shop/getinfo | 店铺详细信息
*ShopApi* | [**AdminShopInfoPost**](docs/ShopApi.md#adminshopinfopost) | **Post** /admin/shop/info | 店铺信息
*ShopApi* | [**AdminShopListPost**](docs/ShopApi.md#adminshoplistpost) | **Post** /admin/shop/list | 店铺列表
*ShopApi* | [**AdminShopPaywaysGet**](docs/ShopApi.md#adminshoppaywaysget) | **Get** /admin/shop/payways | 查询支付方式
*ShopApi* | [**AdminShopUpdateBindPost**](docs/ShopApi.md#adminshopupdatebindpost) | **Post** /admin/shop/update_bind | 商户子账户绑定店铺
*ShopConfigApi* | [**AdminMallConfigGetPost**](docs/ShopConfigApi.md#adminmallconfiggetpost) | **Post** /admin/mall_config/get | 获取店铺配置
*ShopConfigApi* | [**AdminMallConfigSavePost**](docs/ShopConfigApi.md#adminmallconfigsavepost) | **Post** /admin/mall_config/save | 保存店铺配置
*ShopViewApi* | [**AdminShopViewAddPost**](docs/ShopViewApi.md#adminshopviewaddpost) | **Post** /admin/shop/view/add | 修改/新增
*ShopViewApi* | [**AdminShopViewCloseAutoTransPost**](docs/ShopViewApi.md#adminshopviewcloseautotranspost) | **Post** /admin/shop/view/close_auto_trans | 关闭自动翻译
*ShopViewApi* | [**AdminShopViewInfoPost**](docs/ShopViewApi.md#adminshopviewinfopost) | **Post** /admin/shop/view/info | 详情
*ShopViewApi* | [**AdminShopViewListPost**](docs/ShopViewApi.md#adminshopviewlistpost) | **Post** /admin/shop/view/list | 列表
*ShopViewApi* | [**AdminShopViewOpenAutoTransPost**](docs/ShopViewApi.md#adminshopviewopenautotranspost) | **Post** /admin/shop/view/open_auto_trans | 开启自动翻译
*SiteMapApi* | [**MallSitemapXmlGet**](docs/SiteMapApi.md#mallsitemapxmlget) | **Get** /mall/sitemap.xml | 生成siteMap
*StandardCatApi* | [**AdminGoodInfoAddPost**](docs/StandardCatApi.md#admingoodinfoaddpost) | **Post** /admin/good/info/add | 标准类目参数新增/编辑
*StandardCatApi* | [**AdminGoodInfoDeleteOnePost**](docs/StandardCatApi.md#admingoodinfodeleteonepost) | **Post** /admin/good/info/delete/one | 标准类目参数删除
*StandardCatApi* | [**AdminScatSavePost**](docs/StandardCatApi.md#adminscatsavepost) | **Post** /admin/scat/save | 标准类目列表
*StandardCatApi* | [**AdminStandardCatPost**](docs/StandardCatApi.md#adminstandardcatpost) | **Post** /admin/standard/cat | 标准类目列表
*SubmemberApi* | [**AdminSubmemberDeletePost**](docs/SubmemberApi.md#adminsubmemberdeletepost) | **Post** /admin/submember/delete | 删除子账户
*SubmemberApi* | [**AdminSubmemberEditPost**](docs/SubmemberApi.md#adminsubmembereditpost) | **Post** /admin/submember/edit | 修改/新增子账户
*SubmemberApi* | [**AdminSubmemberEmailUniqueGet**](docs/SubmemberApi.md#adminsubmemberemailuniqueget) | **Get** /admin/submember/email_unique | 邮箱是否唯一
*SubmemberApi* | [**AdminSubmemberListGet**](docs/SubmemberApi.md#adminsubmemberlistget) | **Get** /admin/submember/list | 获取会员列表
*SubmemberApi* | [**AdminSubmemberMobileUniqueGet**](docs/SubmemberApi.md#adminsubmembermobileuniqueget) | **Get** /admin/submember/mobile_unique | 手机号是否唯一
*SubmemberApi* | [**AdminSubmemberNameUniqueGet**](docs/SubmemberApi.md#adminsubmembernameuniqueget) | **Get** /admin/submember/name_unique | 会员名称是否唯一
*SubmemberApi* | [**AdminSubmemberResetPwdPost**](docs/SubmemberApi.md#adminsubmemberresetpwdpost) | **Post** /admin/submember/reset_pwd | 重置密码
*SubmemberApi* | [**AdminSubmemberViewGet**](docs/SubmemberApi.md#adminsubmemberviewget) | **Get** /admin/submember/view | 获取指定信息
*SyncGoodsApi* | [**AdminInitDataPost**](docs/SyncGoodsApi.md#admininitdatapost) | **Post** /admin/init/data | 初始化数据
*SyncGoodsApi* | [**AdminSynconfigDelPost**](docs/SyncGoodsApi.md#adminsynconfigdelpost) | **Post** /admin/synconfig/del | 删除配置
*SyncGoodsApi* | [**AdminSynconfigListPost**](docs/SyncGoodsApi.md#adminsynconfiglistpost) | **Post** /admin/synconfig/list | 配置列表
*SyncGoodsApi* | [**AdminSynconfigSavePost**](docs/SyncGoodsApi.md#adminsynconfigsavepost) | **Post** /admin/synconfig/save | 保存配置
*SyncGoodsApi* | [**AdminSyngoodsTriggerPost**](docs/SyncGoodsApi.md#adminsyngoodstriggerpost) | **Post** /admin/syngoods/trigger | 同步商品
*TongjiApi* | [**AdminTongjiGoodsNumPost**](docs/TongjiApi.md#admintongjigoodsnumpost) | **Post** /admin/tongji/goods-num | 商品总数
*TongjiApi* | [**AdminTongjiGoodsSalesPost**](docs/TongjiApi.md#admintongjigoodssalespost) | **Post** /admin/tongji/goods-sales | 商品销售额
*TongjiApi* | [**AdminTongjiOrderDataPost**](docs/TongjiApi.md#admintongjiorderdatapost) | **Post** /admin/tongji/order-data | 订单数据
*TongjiApi* | [**AdminTongjiOrderMoneyPost**](docs/TongjiApi.md#admintongjiordermoneypost) | **Post** /admin/tongji/order-money | 订单金额
*TongjiApi* | [**AdminTongjiRealTimeDataPost**](docs/TongjiApi.md#admintongjirealtimedatapost) | **Post** /admin/tongji/real-time-data | 实时数据
*TongjiApi* | [**AdminTongjiUserBuyPost**](docs/TongjiApi.md#admintongjiuserbuypost) | **Post** /admin/tongji/user-buy | 用户购买统计
*TransApi* | [**AdminTransAddPost**](docs/TransApi.md#admintransaddpost) | **Post** /admin/trans/add | 后端翻译添加
*TransTaskApi* | [**AdminTransTaskAddGoodJobPost**](docs/TransTaskApi.md#admintranstaskaddgoodjobpost) | **Post** /admin/trans_task/add_good_job | 投递商品级翻译任务
*TransTaskApi* | [**AdminTransTaskAddShopJobPost**](docs/TransTaskApi.md#admintranstaskaddshopjobpost) | **Post** /admin/trans_task/add_shop_job | 投递商店级翻译任务
*UploadApi* | [**AdminUploadFileDeletePost**](docs/UploadApi.md#adminuploadfiledeletepost) | **Post** /admin/upload/file/delete | 删除文件
*UploadApi* | [**AdminUploadFileListPost**](docs/UploadApi.md#adminuploadfilelistpost) | **Post** /admin/upload/file/list | 文件列表
*UploadApi* | [**AdminUploadFilePost**](docs/UploadApi.md#adminuploadfilepost) | **Post** /admin/upload/file | 上传文件
*WarehouseApi* | [**AdminWarehouseCreatePost**](docs/WarehouseApi.md#adminwarehousecreatepost) | **Post** /admin/warehouse/create | 创建仓库
*WarehouseApi* | [**AdminWarehouseDeletePost**](docs/WarehouseApi.md#adminwarehousedeletepost) | **Post** /admin/warehouse/delete | 删除仓库
*WarehouseApi* | [**AdminWarehouseReadGet**](docs/WarehouseApi.md#adminwarehousereadget) | **Get** /admin/warehouse/read | 读取仓库列表
*WarehouseApi* | [**AdminWarehouseStockCreatePost**](docs/WarehouseApi.md#adminwarehousestockcreatepost) | **Post** /admin/warehouse/stock/create | 创建仓库商品库存
*WarehouseApi* | [**AdminWarehouseStockPagingPost**](docs/WarehouseApi.md#adminwarehousestockpagingpost) | **Post** /admin/warehouse/stock/paging | 读取仓库商品库存
*WarehouseApi* | [**AdminWarehouseStockReadPost**](docs/WarehouseApi.md#adminwarehousestockreadpost) | **Post** /admin/warehouse/stock/read | 读取仓库商品库存
*WarehouseApi* | [**AdminWarehouseStockUpdatePost**](docs/WarehouseApi.md#adminwarehousestockupdatepost) | **Post** /admin/warehouse/stock/update | 更新仓库商品库存
*WarehouseApi* | [**AdminWarehouseUpdatePost**](docs/WarehouseApi.md#adminwarehouseupdatepost) | **Post** /admin/warehouse/update | 更新仓库
*_Api* | [**MallEmailSendPost**](docs/_Api.md#mallemailsendpost) | **Post** /mall/email/send | 发送邮件

## Documentation For Models

 - [GithubComGogfGfV2NetGhttpUploadFile](docs/GithubComGogfGfV2NetGhttpUploadFile.md)
 - [InlineResponse200](docs/InlineResponse200.md)
 - [InlineResponse2001](docs/InlineResponse2001.md)
 - [InlineResponse20010](docs/InlineResponse20010.md)
 - [InlineResponse200100](docs/InlineResponse200100.md)
 - [InlineResponse200101](docs/InlineResponse200101.md)
 - [InlineResponse200102](docs/InlineResponse200102.md)
 - [InlineResponse200102Data](docs/InlineResponse200102Data.md)
 - [InlineResponse200103](docs/InlineResponse200103.md)
 - [InlineResponse200103Data](docs/InlineResponse200103Data.md)
 - [InlineResponse200104](docs/InlineResponse200104.md)
 - [InlineResponse200104Data](docs/InlineResponse200104Data.md)
 - [InlineResponse200105](docs/InlineResponse200105.md)
 - [InlineResponse200106](docs/InlineResponse200106.md)
 - [InlineResponse200107](docs/InlineResponse200107.md)
 - [InlineResponse200108](docs/InlineResponse200108.md)
 - [InlineResponse200109](docs/InlineResponse200109.md)
 - [InlineResponse200109Data](docs/InlineResponse200109Data.md)
 - [InlineResponse20010Data](docs/InlineResponse20010Data.md)
 - [InlineResponse20011](docs/InlineResponse20011.md)
 - [InlineResponse200110](docs/InlineResponse200110.md)
 - [InlineResponse200111](docs/InlineResponse200111.md)
 - [InlineResponse200112](docs/InlineResponse200112.md)
 - [InlineResponse200112Data](docs/InlineResponse200112Data.md)
 - [InlineResponse200113](docs/InlineResponse200113.md)
 - [InlineResponse200114](docs/InlineResponse200114.md)
 - [InlineResponse200114Data](docs/InlineResponse200114Data.md)
 - [InlineResponse200115](docs/InlineResponse200115.md)
 - [InlineResponse200115Data](docs/InlineResponse200115Data.md)
 - [InlineResponse200116](docs/InlineResponse200116.md)
 - [InlineResponse200117](docs/InlineResponse200117.md)
 - [InlineResponse200117Data](docs/InlineResponse200117Data.md)
 - [InlineResponse200118](docs/InlineResponse200118.md)
 - [InlineResponse200118Data](docs/InlineResponse200118Data.md)
 - [InlineResponse200119](docs/InlineResponse200119.md)
 - [InlineResponse200119Data](docs/InlineResponse200119Data.md)
 - [InlineResponse20011Data](docs/InlineResponse20011Data.md)
 - [InlineResponse20012](docs/InlineResponse20012.md)
 - [InlineResponse200120](docs/InlineResponse200120.md)
 - [InlineResponse200120Data](docs/InlineResponse200120Data.md)
 - [InlineResponse200121](docs/InlineResponse200121.md)
 - [InlineResponse200121Data](docs/InlineResponse200121Data.md)
 - [InlineResponse200122](docs/InlineResponse200122.md)
 - [InlineResponse200122Data](docs/InlineResponse200122Data.md)
 - [InlineResponse200123](docs/InlineResponse200123.md)
 - [InlineResponse200123Data](docs/InlineResponse200123Data.md)
 - [InlineResponse200124](docs/InlineResponse200124.md)
 - [InlineResponse200124Data](docs/InlineResponse200124Data.md)
 - [InlineResponse200125](docs/InlineResponse200125.md)
 - [InlineResponse200125Data](docs/InlineResponse200125Data.md)
 - [InlineResponse200126](docs/InlineResponse200126.md)
 - [InlineResponse200126Data](docs/InlineResponse200126Data.md)
 - [InlineResponse200127](docs/InlineResponse200127.md)
 - [InlineResponse200127Data](docs/InlineResponse200127Data.md)
 - [InlineResponse200128](docs/InlineResponse200128.md)
 - [InlineResponse200128Data](docs/InlineResponse200128Data.md)
 - [InlineResponse200129](docs/InlineResponse200129.md)
 - [InlineResponse200129Data](docs/InlineResponse200129Data.md)
 - [InlineResponse20012Data](docs/InlineResponse20012Data.md)
 - [InlineResponse20013](docs/InlineResponse20013.md)
 - [InlineResponse200130](docs/InlineResponse200130.md)
 - [InlineResponse200130Data](docs/InlineResponse200130Data.md)
 - [InlineResponse200131](docs/InlineResponse200131.md)
 - [InlineResponse200131Data](docs/InlineResponse200131Data.md)
 - [InlineResponse200132](docs/InlineResponse200132.md)
 - [InlineResponse200132Data](docs/InlineResponse200132Data.md)
 - [InlineResponse200133](docs/InlineResponse200133.md)
 - [InlineResponse200133Data](docs/InlineResponse200133Data.md)
 - [InlineResponse200134](docs/InlineResponse200134.md)
 - [InlineResponse200134Data](docs/InlineResponse200134Data.md)
 - [InlineResponse200135](docs/InlineResponse200135.md)
 - [InlineResponse200135Data](docs/InlineResponse200135Data.md)
 - [InlineResponse200136](docs/InlineResponse200136.md)
 - [InlineResponse200136Data](docs/InlineResponse200136Data.md)
 - [InlineResponse200137](docs/InlineResponse200137.md)
 - [InlineResponse200137Data](docs/InlineResponse200137Data.md)
 - [InlineResponse200138](docs/InlineResponse200138.md)
 - [InlineResponse200139](docs/InlineResponse200139.md)
 - [InlineResponse200139Data](docs/InlineResponse200139Data.md)
 - [InlineResponse20013Data](docs/InlineResponse20013Data.md)
 - [InlineResponse20014](docs/InlineResponse20014.md)
 - [InlineResponse200140](docs/InlineResponse200140.md)
 - [InlineResponse200141](docs/InlineResponse200141.md)
 - [InlineResponse200141Data](docs/InlineResponse200141Data.md)
 - [InlineResponse200142](docs/InlineResponse200142.md)
 - [InlineResponse200142Data](docs/InlineResponse200142Data.md)
 - [InlineResponse200143](docs/InlineResponse200143.md)
 - [InlineResponse200144](docs/InlineResponse200144.md)
 - [InlineResponse200144Data](docs/InlineResponse200144Data.md)
 - [InlineResponse200145](docs/InlineResponse200145.md)
 - [InlineResponse200145Data](docs/InlineResponse200145Data.md)
 - [InlineResponse200146](docs/InlineResponse200146.md)
 - [InlineResponse200146Data](docs/InlineResponse200146Data.md)
 - [InlineResponse200147](docs/InlineResponse200147.md)
 - [InlineResponse200147Data](docs/InlineResponse200147Data.md)
 - [InlineResponse200148](docs/InlineResponse200148.md)
 - [InlineResponse200148Data](docs/InlineResponse200148Data.md)
 - [InlineResponse200149](docs/InlineResponse200149.md)
 - [InlineResponse200149Data](docs/InlineResponse200149Data.md)
 - [InlineResponse20014Data](docs/InlineResponse20014Data.md)
 - [InlineResponse20015](docs/InlineResponse20015.md)
 - [InlineResponse200150](docs/InlineResponse200150.md)
 - [InlineResponse200150Data](docs/InlineResponse200150Data.md)
 - [InlineResponse200151](docs/InlineResponse200151.md)
 - [InlineResponse200151Data](docs/InlineResponse200151Data.md)
 - [InlineResponse200152](docs/InlineResponse200152.md)
 - [InlineResponse200152Data](docs/InlineResponse200152Data.md)
 - [InlineResponse200153](docs/InlineResponse200153.md)
 - [InlineResponse200153Data](docs/InlineResponse200153Data.md)
 - [InlineResponse200154](docs/InlineResponse200154.md)
 - [InlineResponse200154Data](docs/InlineResponse200154Data.md)
 - [InlineResponse200155](docs/InlineResponse200155.md)
 - [InlineResponse200155Data](docs/InlineResponse200155Data.md)
 - [InlineResponse200156](docs/InlineResponse200156.md)
 - [InlineResponse200156Data](docs/InlineResponse200156Data.md)
 - [InlineResponse200157](docs/InlineResponse200157.md)
 - [InlineResponse200157Data](docs/InlineResponse200157Data.md)
 - [InlineResponse200158](docs/InlineResponse200158.md)
 - [InlineResponse200158Data](docs/InlineResponse200158Data.md)
 - [InlineResponse200159](docs/InlineResponse200159.md)
 - [InlineResponse200159Data](docs/InlineResponse200159Data.md)
 - [InlineResponse20015Data](docs/InlineResponse20015Data.md)
 - [InlineResponse20016](docs/InlineResponse20016.md)
 - [InlineResponse200160](docs/InlineResponse200160.md)
 - [InlineResponse200160Data](docs/InlineResponse200160Data.md)
 - [InlineResponse200161](docs/InlineResponse200161.md)
 - [InlineResponse200161Data](docs/InlineResponse200161Data.md)
 - [InlineResponse200162](docs/InlineResponse200162.md)
 - [InlineResponse200162Data](docs/InlineResponse200162Data.md)
 - [InlineResponse200163](docs/InlineResponse200163.md)
 - [InlineResponse200163Data](docs/InlineResponse200163Data.md)
 - [InlineResponse200164](docs/InlineResponse200164.md)
 - [InlineResponse200164Data](docs/InlineResponse200164Data.md)
 - [InlineResponse200165](docs/InlineResponse200165.md)
 - [InlineResponse200165Data](docs/InlineResponse200165Data.md)
 - [InlineResponse200166](docs/InlineResponse200166.md)
 - [InlineResponse200166Data](docs/InlineResponse200166Data.md)
 - [InlineResponse200167](docs/InlineResponse200167.md)
 - [InlineResponse200167Data](docs/InlineResponse200167Data.md)
 - [InlineResponse200168](docs/InlineResponse200168.md)
 - [InlineResponse200168Data](docs/InlineResponse200168Data.md)
 - [InlineResponse200169](docs/InlineResponse200169.md)
 - [InlineResponse200169Data](docs/InlineResponse200169Data.md)
 - [InlineResponse20016Data](docs/InlineResponse20016Data.md)
 - [InlineResponse20017](docs/InlineResponse20017.md)
 - [InlineResponse200170](docs/InlineResponse200170.md)
 - [InlineResponse200170Data](docs/InlineResponse200170Data.md)
 - [InlineResponse200171](docs/InlineResponse200171.md)
 - [InlineResponse200172](docs/InlineResponse200172.md)
 - [InlineResponse200172Data](docs/InlineResponse200172Data.md)
 - [InlineResponse200173](docs/InlineResponse200173.md)
 - [InlineResponse200173Data](docs/InlineResponse200173Data.md)
 - [InlineResponse200174](docs/InlineResponse200174.md)
 - [InlineResponse200174Data](docs/InlineResponse200174Data.md)
 - [InlineResponse200175](docs/InlineResponse200175.md)
 - [InlineResponse200175Data](docs/InlineResponse200175Data.md)
 - [InlineResponse200176](docs/InlineResponse200176.md)
 - [InlineResponse200176Data](docs/InlineResponse200176Data.md)
 - [InlineResponse200177](docs/InlineResponse200177.md)
 - [InlineResponse200177Data](docs/InlineResponse200177Data.md)
 - [InlineResponse200178](docs/InlineResponse200178.md)
 - [InlineResponse200178Data](docs/InlineResponse200178Data.md)
 - [InlineResponse200179](docs/InlineResponse200179.md)
 - [InlineResponse200179Data](docs/InlineResponse200179Data.md)
 - [InlineResponse20017Data](docs/InlineResponse20017Data.md)
 - [InlineResponse20018](docs/InlineResponse20018.md)
 - [InlineResponse200180](docs/InlineResponse200180.md)
 - [InlineResponse200180Data](docs/InlineResponse200180Data.md)
 - [InlineResponse200181](docs/InlineResponse200181.md)
 - [InlineResponse200181Data](docs/InlineResponse200181Data.md)
 - [InlineResponse200182](docs/InlineResponse200182.md)
 - [InlineResponse200182Data](docs/InlineResponse200182Data.md)
 - [InlineResponse200183](docs/InlineResponse200183.md)
 - [InlineResponse200183Data](docs/InlineResponse200183Data.md)
 - [InlineResponse200184](docs/InlineResponse200184.md)
 - [InlineResponse200184Data](docs/InlineResponse200184Data.md)
 - [InlineResponse200185](docs/InlineResponse200185.md)
 - [InlineResponse200185Data](docs/InlineResponse200185Data.md)
 - [InlineResponse200186](docs/InlineResponse200186.md)
 - [InlineResponse200186Data](docs/InlineResponse200186Data.md)
 - [InlineResponse200187](docs/InlineResponse200187.md)
 - [InlineResponse200187Data](docs/InlineResponse200187Data.md)
 - [InlineResponse200188](docs/InlineResponse200188.md)
 - [InlineResponse200188Data](docs/InlineResponse200188Data.md)
 - [InlineResponse200189](docs/InlineResponse200189.md)
 - [InlineResponse200189Data](docs/InlineResponse200189Data.md)
 - [InlineResponse20018Data](docs/InlineResponse20018Data.md)
 - [InlineResponse20019](docs/InlineResponse20019.md)
 - [InlineResponse200190](docs/InlineResponse200190.md)
 - [InlineResponse200190Data](docs/InlineResponse200190Data.md)
 - [InlineResponse200191](docs/InlineResponse200191.md)
 - [InlineResponse200191Data](docs/InlineResponse200191Data.md)
 - [InlineResponse200192](docs/InlineResponse200192.md)
 - [InlineResponse200192Data](docs/InlineResponse200192Data.md)
 - [InlineResponse200193](docs/InlineResponse200193.md)
 - [InlineResponse200193Data](docs/InlineResponse200193Data.md)
 - [InlineResponse200194](docs/InlineResponse200194.md)
 - [InlineResponse200194Data](docs/InlineResponse200194Data.md)
 - [InlineResponse200195](docs/InlineResponse200195.md)
 - [InlineResponse200195Data](docs/InlineResponse200195Data.md)
 - [InlineResponse200196](docs/InlineResponse200196.md)
 - [InlineResponse200196Data](docs/InlineResponse200196Data.md)
 - [InlineResponse200197](docs/InlineResponse200197.md)
 - [InlineResponse200197Data](docs/InlineResponse200197Data.md)
 - [InlineResponse200198](docs/InlineResponse200198.md)
 - [InlineResponse200198Data](docs/InlineResponse200198Data.md)
 - [InlineResponse200199](docs/InlineResponse200199.md)
 - [InlineResponse200199Data](docs/InlineResponse200199Data.md)
 - [InlineResponse20019Data](docs/InlineResponse20019Data.md)
 - [InlineResponse2001Data](docs/InlineResponse2001Data.md)
 - [InlineResponse2002](docs/InlineResponse2002.md)
 - [InlineResponse20020](docs/InlineResponse20020.md)
 - [InlineResponse200200](docs/InlineResponse200200.md)
 - [InlineResponse200200Data](docs/InlineResponse200200Data.md)
 - [InlineResponse200201](docs/InlineResponse200201.md)
 - [InlineResponse200201Data](docs/InlineResponse200201Data.md)
 - [InlineResponse200202](docs/InlineResponse200202.md)
 - [InlineResponse200202Data](docs/InlineResponse200202Data.md)
 - [InlineResponse200203](docs/InlineResponse200203.md)
 - [InlineResponse200203Data](docs/InlineResponse200203Data.md)
 - [InlineResponse200204](docs/InlineResponse200204.md)
 - [InlineResponse200204Data](docs/InlineResponse200204Data.md)
 - [InlineResponse200205](docs/InlineResponse200205.md)
 - [InlineResponse200205Data](docs/InlineResponse200205Data.md)
 - [InlineResponse200206](docs/InlineResponse200206.md)
 - [InlineResponse200206Data](docs/InlineResponse200206Data.md)
 - [InlineResponse200207](docs/InlineResponse200207.md)
 - [InlineResponse200207Data](docs/InlineResponse200207Data.md)
 - [InlineResponse200208](docs/InlineResponse200208.md)
 - [InlineResponse200208Data](docs/InlineResponse200208Data.md)
 - [InlineResponse200209](docs/InlineResponse200209.md)
 - [InlineResponse200209Data](docs/InlineResponse200209Data.md)
 - [InlineResponse20020Data](docs/InlineResponse20020Data.md)
 - [InlineResponse20021](docs/InlineResponse20021.md)
 - [InlineResponse200210](docs/InlineResponse200210.md)
 - [InlineResponse200210Data](docs/InlineResponse200210Data.md)
 - [InlineResponse200211](docs/InlineResponse200211.md)
 - [InlineResponse200211Data](docs/InlineResponse200211Data.md)
 - [InlineResponse200212](docs/InlineResponse200212.md)
 - [InlineResponse200212Data](docs/InlineResponse200212Data.md)
 - [InlineResponse200213](docs/InlineResponse200213.md)
 - [InlineResponse200213Data](docs/InlineResponse200213Data.md)
 - [InlineResponse200214](docs/InlineResponse200214.md)
 - [InlineResponse200214Data](docs/InlineResponse200214Data.md)
 - [InlineResponse200215](docs/InlineResponse200215.md)
 - [InlineResponse200215Data](docs/InlineResponse200215Data.md)
 - [InlineResponse200216](docs/InlineResponse200216.md)
 - [InlineResponse200216Data](docs/InlineResponse200216Data.md)
 - [InlineResponse200217](docs/InlineResponse200217.md)
 - [InlineResponse200217Data](docs/InlineResponse200217Data.md)
 - [InlineResponse200218](docs/InlineResponse200218.md)
 - [InlineResponse200218Data](docs/InlineResponse200218Data.md)
 - [InlineResponse20021Data](docs/InlineResponse20021Data.md)
 - [InlineResponse20022](docs/InlineResponse20022.md)
 - [InlineResponse20022Data](docs/InlineResponse20022Data.md)
 - [InlineResponse20023](docs/InlineResponse20023.md)
 - [InlineResponse20023Data](docs/InlineResponse20023Data.md)
 - [InlineResponse20024](docs/InlineResponse20024.md)
 - [InlineResponse20024Data](docs/InlineResponse20024Data.md)
 - [InlineResponse20025](docs/InlineResponse20025.md)
 - [InlineResponse20026](docs/InlineResponse20026.md)
 - [InlineResponse20027](docs/InlineResponse20027.md)
 - [InlineResponse20027Data](docs/InlineResponse20027Data.md)
 - [InlineResponse20028](docs/InlineResponse20028.md)
 - [InlineResponse20029](docs/InlineResponse20029.md)
 - [InlineResponse20029Data](docs/InlineResponse20029Data.md)
 - [InlineResponse2002Data](docs/InlineResponse2002Data.md)
 - [InlineResponse2003](docs/InlineResponse2003.md)
 - [InlineResponse20030](docs/InlineResponse20030.md)
 - [InlineResponse20030Data](docs/InlineResponse20030Data.md)
 - [InlineResponse20031](docs/InlineResponse20031.md)
 - [InlineResponse20031Data](docs/InlineResponse20031Data.md)
 - [InlineResponse20032](docs/InlineResponse20032.md)
 - [InlineResponse20032Data](docs/InlineResponse20032Data.md)
 - [InlineResponse20033](docs/InlineResponse20033.md)
 - [InlineResponse20033Data](docs/InlineResponse20033Data.md)
 - [InlineResponse20034](docs/InlineResponse20034.md)
 - [InlineResponse20034Data](docs/InlineResponse20034Data.md)
 - [InlineResponse20035](docs/InlineResponse20035.md)
 - [InlineResponse20035Data](docs/InlineResponse20035Data.md)
 - [InlineResponse20036](docs/InlineResponse20036.md)
 - [InlineResponse20036Data](docs/InlineResponse20036Data.md)
 - [InlineResponse20037](docs/InlineResponse20037.md)
 - [InlineResponse20037Data](docs/InlineResponse20037Data.md)
 - [InlineResponse20038](docs/InlineResponse20038.md)
 - [InlineResponse20038Data](docs/InlineResponse20038Data.md)
 - [InlineResponse20039](docs/InlineResponse20039.md)
 - [InlineResponse20039Data](docs/InlineResponse20039Data.md)
 - [InlineResponse2003Data](docs/InlineResponse2003Data.md)
 - [InlineResponse2004](docs/InlineResponse2004.md)
 - [InlineResponse20040](docs/InlineResponse20040.md)
 - [InlineResponse20040Data](docs/InlineResponse20040Data.md)
 - [InlineResponse20041](docs/InlineResponse20041.md)
 - [InlineResponse20041Data](docs/InlineResponse20041Data.md)
 - [InlineResponse20042](docs/InlineResponse20042.md)
 - [InlineResponse20042Data](docs/InlineResponse20042Data.md)
 - [InlineResponse20043](docs/InlineResponse20043.md)
 - [InlineResponse20043Data](docs/InlineResponse20043Data.md)
 - [InlineResponse20044](docs/InlineResponse20044.md)
 - [InlineResponse20044Data](docs/InlineResponse20044Data.md)
 - [InlineResponse20045](docs/InlineResponse20045.md)
 - [InlineResponse20045Data](docs/InlineResponse20045Data.md)
 - [InlineResponse20046](docs/InlineResponse20046.md)
 - [InlineResponse20046Data](docs/InlineResponse20046Data.md)
 - [InlineResponse20047](docs/InlineResponse20047.md)
 - [InlineResponse20047Data](docs/InlineResponse20047Data.md)
 - [InlineResponse20048](docs/InlineResponse20048.md)
 - [InlineResponse20048Data](docs/InlineResponse20048Data.md)
 - [InlineResponse20049](docs/InlineResponse20049.md)
 - [InlineResponse20049Data](docs/InlineResponse20049Data.md)
 - [InlineResponse2004Data](docs/InlineResponse2004Data.md)
 - [InlineResponse2005](docs/InlineResponse2005.md)
 - [InlineResponse20050](docs/InlineResponse20050.md)
 - [InlineResponse20050Data](docs/InlineResponse20050Data.md)
 - [InlineResponse20051](docs/InlineResponse20051.md)
 - [InlineResponse20051Data](docs/InlineResponse20051Data.md)
 - [InlineResponse20052](docs/InlineResponse20052.md)
 - [InlineResponse20052Data](docs/InlineResponse20052Data.md)
 - [InlineResponse20053](docs/InlineResponse20053.md)
 - [InlineResponse20053Data](docs/InlineResponse20053Data.md)
 - [InlineResponse20054](docs/InlineResponse20054.md)
 - [InlineResponse20054Data](docs/InlineResponse20054Data.md)
 - [InlineResponse20055](docs/InlineResponse20055.md)
 - [InlineResponse20055Data](docs/InlineResponse20055Data.md)
 - [InlineResponse20056](docs/InlineResponse20056.md)
 - [InlineResponse20056Data](docs/InlineResponse20056Data.md)
 - [InlineResponse20057](docs/InlineResponse20057.md)
 - [InlineResponse20057Data](docs/InlineResponse20057Data.md)
 - [InlineResponse20058](docs/InlineResponse20058.md)
 - [InlineResponse20058Data](docs/InlineResponse20058Data.md)
 - [InlineResponse20059](docs/InlineResponse20059.md)
 - [InlineResponse20059Data](docs/InlineResponse20059Data.md)
 - [InlineResponse2005Data](docs/InlineResponse2005Data.md)
 - [InlineResponse2006](docs/InlineResponse2006.md)
 - [InlineResponse20060](docs/InlineResponse20060.md)
 - [InlineResponse20060Data](docs/InlineResponse20060Data.md)
 - [InlineResponse20061](docs/InlineResponse20061.md)
 - [InlineResponse20061Data](docs/InlineResponse20061Data.md)
 - [InlineResponse20062](docs/InlineResponse20062.md)
 - [InlineResponse20062Data](docs/InlineResponse20062Data.md)
 - [InlineResponse20063](docs/InlineResponse20063.md)
 - [InlineResponse20064](docs/InlineResponse20064.md)
 - [InlineResponse20065](docs/InlineResponse20065.md)
 - [InlineResponse20066](docs/InlineResponse20066.md)
 - [InlineResponse20066Data](docs/InlineResponse20066Data.md)
 - [InlineResponse20067](docs/InlineResponse20067.md)
 - [InlineResponse20067Data](docs/InlineResponse20067Data.md)
 - [InlineResponse20068](docs/InlineResponse20068.md)
 - [InlineResponse20068Data](docs/InlineResponse20068Data.md)
 - [InlineResponse20069](docs/InlineResponse20069.md)
 - [InlineResponse20069Data](docs/InlineResponse20069Data.md)
 - [InlineResponse2006Data](docs/InlineResponse2006Data.md)
 - [InlineResponse2007](docs/InlineResponse2007.md)
 - [InlineResponse20070](docs/InlineResponse20070.md)
 - [InlineResponse20070Data](docs/InlineResponse20070Data.md)
 - [InlineResponse20071](docs/InlineResponse20071.md)
 - [InlineResponse20071Data](docs/InlineResponse20071Data.md)
 - [InlineResponse20072](docs/InlineResponse20072.md)
 - [InlineResponse20073](docs/InlineResponse20073.md)
 - [InlineResponse20074](docs/InlineResponse20074.md)
 - [InlineResponse20074Data](docs/InlineResponse20074Data.md)
 - [InlineResponse20075](docs/InlineResponse20075.md)
 - [InlineResponse20075Data](docs/InlineResponse20075Data.md)
 - [InlineResponse20076](docs/InlineResponse20076.md)
 - [InlineResponse20076Data](docs/InlineResponse20076Data.md)
 - [InlineResponse20077](docs/InlineResponse20077.md)
 - [InlineResponse20077Data](docs/InlineResponse20077Data.md)
 - [InlineResponse20078](docs/InlineResponse20078.md)
 - [InlineResponse20078Data](docs/InlineResponse20078Data.md)
 - [InlineResponse20079](docs/InlineResponse20079.md)
 - [InlineResponse20079Data](docs/InlineResponse20079Data.md)
 - [InlineResponse2008](docs/InlineResponse2008.md)
 - [InlineResponse20080](docs/InlineResponse20080.md)
 - [InlineResponse20080Data](docs/InlineResponse20080Data.md)
 - [InlineResponse20081](docs/InlineResponse20081.md)
 - [InlineResponse20082](docs/InlineResponse20082.md)
 - [InlineResponse20083](docs/InlineResponse20083.md)
 - [InlineResponse20084](docs/InlineResponse20084.md)
 - [InlineResponse20085](docs/InlineResponse20085.md)
 - [InlineResponse20086](docs/InlineResponse20086.md)
 - [InlineResponse20087](docs/InlineResponse20087.md)
 - [InlineResponse20088](docs/InlineResponse20088.md)
 - [InlineResponse20088Data](docs/InlineResponse20088Data.md)
 - [InlineResponse20089](docs/InlineResponse20089.md)
 - [InlineResponse20089Data](docs/InlineResponse20089Data.md)
 - [InlineResponse2008Data](docs/InlineResponse2008Data.md)
 - [InlineResponse2009](docs/InlineResponse2009.md)
 - [InlineResponse20090](docs/InlineResponse20090.md)
 - [InlineResponse20090Data](docs/InlineResponse20090Data.md)
 - [InlineResponse20091](docs/InlineResponse20091.md)
 - [InlineResponse20091Data](docs/InlineResponse20091Data.md)
 - [InlineResponse20092](docs/InlineResponse20092.md)
 - [InlineResponse20093](docs/InlineResponse20093.md)
 - [InlineResponse20094](docs/InlineResponse20094.md)
 - [InlineResponse20095](docs/InlineResponse20095.md)
 - [InlineResponse20095Data](docs/InlineResponse20095Data.md)
 - [InlineResponse20096](docs/InlineResponse20096.md)
 - [InlineResponse20096Data](docs/InlineResponse20096Data.md)
 - [InlineResponse20097](docs/InlineResponse20097.md)
 - [InlineResponse20097Data](docs/InlineResponse20097Data.md)
 - [InlineResponse20098](docs/InlineResponse20098.md)
 - [InlineResponse20099](docs/InlineResponse20099.md)
 - [InlineResponse2009Data](docs/InlineResponse2009Data.md)
 - [InlineResponse200Data](docs/InlineResponse200Data.md)
 - [ModelInterface](docs/ModelInterface.md)
 - [XbuyAppDaoAfterSales](docs/XbuyAppDaoAfterSales.md)
 - [XbuyAppDaoAllReview](docs/XbuyAppDaoAllReview.md)
 - [XbuyAppDaoAllShopMemberList](docs/XbuyAppDaoAllShopMemberList.md)
 - [XbuyAppDaoArticleCatDetail](docs/XbuyAppDaoArticleCatDetail.md)
 - [XbuyAppDaoBenefitCode](docs/XbuyAppDaoBenefitCode.md)
 - [XbuyAppDaoBenefitWith](docs/XbuyAppDaoBenefitWith.md)
 - [XbuyAppDaoCatDetail](docs/XbuyAppDaoCatDetail.md)
 - [XbuyAppDaoCatTree](docs/XbuyAppDaoCatTree.md)
 - [XbuyAppDaoChild](docs/XbuyAppDaoChild.md)
 - [XbuyAppDaoCurrencyConversion](docs/XbuyAppDaoCurrencyConversion.md)
 - [XbuyAppDaoCurrencySetter](docs/XbuyAppDaoCurrencySetter.md)
 - [XbuyAppDaoDiscount](docs/XbuyAppDaoDiscount.md)
 - [XbuyAppDaoEffectSkuList](docs/XbuyAppDaoEffectSkuList.md)
 - [XbuyAppDaoEffectWith](docs/XbuyAppDaoEffectWith.md)
 - [XbuyAppDaoEmailSetter](docs/XbuyAppDaoEmailSetter.md)
 - [XbuyAppDaoExchangeRate](docs/XbuyAppDaoExchangeRate.md)
 - [XbuyAppDaoExpress](docs/XbuyAppDaoExpress.md)
 - [XbuyAppDaoFav](docs/XbuyAppDaoFav.md)
 - [XbuyAppDaoFilingResponse](docs/XbuyAppDaoFilingResponse.md)
 - [XbuyAppDaoGetSitemapResponse](docs/XbuyAppDaoGetSitemapResponse.md)
 - [XbuyAppDaoGiftCategory](docs/XbuyAppDaoGiftCategory.md)
 - [XbuyAppDaoGoodInfosDetail](docs/XbuyAppDaoGoodInfosDetail.md)
 - [XbuyAppDaoGoodListContent](docs/XbuyAppDaoGoodListContent.md)
 - [XbuyAppDaoGoodsDetails](docs/XbuyAppDaoGoodsDetails.md)
 - [XbuyAppDaoItemsLevelSummary](docs/XbuyAppDaoItemsLevelSummary.md)
 - [XbuyAppDaoMallGoodAndSkusInfo](docs/XbuyAppDaoMallGoodAndSkusInfo.md)
 - [XbuyAppDaoModeListContent](docs/XbuyAppDaoModeListContent.md)
 - [XbuyAppDaoOrder](docs/XbuyAppDaoOrder.md)
 - [XbuyAppDaoOutputSkuInfo](docs/XbuyAppDaoOutputSkuInfo.md)
 - [XbuyAppDaoOutputValue](docs/XbuyAppDaoOutputValue.md)
 - [XbuyAppDaoPaymentSetter](docs/XbuyAppDaoPaymentSetter.md)
 - [XbuyAppDaoPreorderResult](docs/XbuyAppDaoPreorderResult.md)
 - [XbuyAppDaoPromo](docs/XbuyAppDaoPromo.md)
 - [XbuyAppDaoPromoDetail](docs/XbuyAppDaoPromoDetail.md)
 - [XbuyAppDaoRows](docs/XbuyAppDaoRows.md)
 - [XbuyAppDaoSearchAnalyticsQueryResponse](docs/XbuyAppDaoSearchAnalyticsQueryResponse.md)
 - [XbuyAppDaoSearchCustomsItemRes](docs/XbuyAppDaoSearchCustomsItemRes.md)
 - [XbuyAppDaoShopAndOrderCont](docs/XbuyAppDaoShopAndOrderCont.md)
 - [XbuyAppDaoSiteEntry](docs/XbuyAppDaoSiteEntry.md)
 - [XbuyAppDaoSitemap](docs/XbuyAppDaoSitemap.md)
 - [XbuyAppDaoSitemapContent](docs/XbuyAppDaoSitemapContent.md)
 - [XbuyAppDaoSitemapGoogleList](docs/XbuyAppDaoSitemapGoogleList.md)
 - [XbuyAppDaoSkuAttr](docs/XbuyAppDaoSkuAttr.md)
 - [XbuyAppDaoSkuDetail](docs/XbuyAppDaoSkuDetail.md)
 - [XbuyAppDaoSkuDetails](docs/XbuyAppDaoSkuDetails.md)
 - [XbuyAppDaoSkusContent](docs/XbuyAppDaoSkusContent.md)
 - [XbuyAppDaoStatistics](docs/XbuyAppDaoStatistics.md)
 - [XbuyAppDaoStock](docs/XbuyAppDaoStock.md)
 - [XbuyAppDaoSubCats](docs/XbuyAppDaoSubCats.md)
 - [XbuyAppDaoTargetWith](docs/XbuyAppDaoTargetWith.md)
 - [XbuyAppDaoWithCond](docs/XbuyAppDaoWithCond.md)
 - [XbuyAppDomainOrderdomainPaymentH5Result](docs/XbuyAppDomainOrderdomainPaymentH5Result.md)
 - [XbuyAppDomainOrderdomainPaymentJsapiResult](docs/XbuyAppDomainOrderdomainPaymentJsapiResult.md)
 - [XbuyAppDomainOrderdomainPaymentNativeResult](docs/XbuyAppDomainOrderdomainPaymentNativeResult.md)
 - [XbuyAppFormAdminFormAddGoodsTransJobReq](docs/XbuyAppFormAdminFormAddGoodsTransJobReq.md)
 - [XbuyAppFormAdminFormAddGoodsTransJobRes](docs/XbuyAppFormAdminFormAddGoodsTransJobRes.md)
 - [XbuyAppFormAdminFormAddMemberGradeReq](docs/XbuyAppFormAdminFormAddMemberGradeReq.md)
 - [XbuyAppFormAdminFormAddMemberGradeRes](docs/XbuyAppFormAdminFormAddMemberGradeRes.md)
 - [XbuyAppFormAdminFormAddMemberGroupRelationReq](docs/XbuyAppFormAdminFormAddMemberGroupRelationReq.md)
 - [XbuyAppFormAdminFormAddMemberGroupRelationRes](docs/XbuyAppFormAdminFormAddMemberGroupRelationRes.md)
 - [XbuyAppFormAdminFormAddMemberGroupReq](docs/XbuyAppFormAdminFormAddMemberGroupReq.md)
 - [XbuyAppFormAdminFormAddMemberGroupRes](docs/XbuyAppFormAdminFormAddMemberGroupRes.md)
 - [XbuyAppFormAdminFormAddMemberRemarkReq](docs/XbuyAppFormAdminFormAddMemberRemarkReq.md)
 - [XbuyAppFormAdminFormAddMemberRemarkRes](docs/XbuyAppFormAdminFormAddMemberRemarkRes.md)
 - [XbuyAppFormAdminFormAddMemberTagReq](docs/XbuyAppFormAdminFormAddMemberTagReq.md)
 - [XbuyAppFormAdminFormAddMemberTagRes](docs/XbuyAppFormAdminFormAddMemberTagRes.md)
 - [XbuyAppFormAdminFormAddShopReq](docs/XbuyAppFormAdminFormAddShopReq.md)
 - [XbuyAppFormAdminFormAddShopRes](docs/XbuyAppFormAdminFormAddShopRes.md)
 - [XbuyAppFormAdminFormAddShopTransJobReq](docs/XbuyAppFormAdminFormAddShopTransJobReq.md)
 - [XbuyAppFormAdminFormAddShopTransJobRes](docs/XbuyAppFormAdminFormAddShopTransJobRes.md)
 - [XbuyAppFormAdminFormAddSkuTaxReq](docs/XbuyAppFormAdminFormAddSkuTaxReq.md)
 - [XbuyAppFormAdminFormAddSkuTaxRes](docs/XbuyAppFormAdminFormAddSkuTaxRes.md)
 - [XbuyAppFormAdminFormAfterSalesApplyReq](docs/XbuyAppFormAdminFormAfterSalesApplyReq.md)
 - [XbuyAppFormAdminFormAfterSalesApplyRes](docs/XbuyAppFormAdminFormAfterSalesApplyRes.md)
 - [XbuyAppFormAdminFormAfterSalesListReq](docs/XbuyAppFormAdminFormAfterSalesListReq.md)
 - [XbuyAppFormAdminFormAfterSalesListRes](docs/XbuyAppFormAdminFormAfterSalesListRes.md)
 - [XbuyAppFormAdminFormAfterSalesSwitchReq](docs/XbuyAppFormAdminFormAfterSalesSwitchReq.md)
 - [XbuyAppFormAdminFormAfterSalesSwitchRes](docs/XbuyAppFormAdminFormAfterSalesSwitchRes.md)
 - [XbuyAppFormAdminFormAiTransReq](docs/XbuyAppFormAdminFormAiTransReq.md)
 - [XbuyAppFormAdminFormAiTransRes](docs/XbuyAppFormAdminFormAiTransRes.md)
 - [XbuyAppFormAdminFormArticleAddReq](docs/XbuyAppFormAdminFormArticleAddReq.md)
 - [XbuyAppFormAdminFormArticleAddRes](docs/XbuyAppFormAdminFormArticleAddRes.md)
 - [XbuyAppFormAdminFormArticleCatAddReq](docs/XbuyAppFormAdminFormArticleCatAddReq.md)
 - [XbuyAppFormAdminFormArticleCatAddRes](docs/XbuyAppFormAdminFormArticleCatAddRes.md)
 - [XbuyAppFormAdminFormArticleCatInfoReq](docs/XbuyAppFormAdminFormArticleCatInfoReq.md)
 - [XbuyAppFormAdminFormArticleCatInfoRes](docs/XbuyAppFormAdminFormArticleCatInfoRes.md)
 - [XbuyAppFormAdminFormArticleCatListReq](docs/XbuyAppFormAdminFormArticleCatListReq.md)
 - [XbuyAppFormAdminFormArticleCatListRes](docs/XbuyAppFormAdminFormArticleCatListRes.md)
 - [XbuyAppFormAdminFormArticleDeleteReq](docs/XbuyAppFormAdminFormArticleDeleteReq.md)
 - [XbuyAppFormAdminFormArticleDeleteRes](docs/XbuyAppFormAdminFormArticleDeleteRes.md)
 - [XbuyAppFormAdminFormArticleGetReq](docs/XbuyAppFormAdminFormArticleGetReq.md)
 - [XbuyAppFormAdminFormArticleGetRes](docs/XbuyAppFormAdminFormArticleGetRes.md)
 - [XbuyAppFormAdminFormArticleListReq](docs/XbuyAppFormAdminFormArticleListReq.md)
 - [XbuyAppFormAdminFormArticleListRes](docs/XbuyAppFormAdminFormArticleListRes.md)
 - [XbuyAppFormAdminFormArticleSubAddReq](docs/XbuyAppFormAdminFormArticleSubAddReq.md)
 - [XbuyAppFormAdminFormArticleSubAddRes](docs/XbuyAppFormAdminFormArticleSubAddRes.md)
 - [XbuyAppFormAdminFormArticlesCatSubAddReq](docs/XbuyAppFormAdminFormArticlesCatSubAddReq.md)
 - [XbuyAppFormAdminFormArticlesCatSubAddRes](docs/XbuyAppFormAdminFormArticlesCatSubAddRes.md)
 - [XbuyAppFormAdminFormBenefitCodeAddReq](docs/XbuyAppFormAdminFormBenefitCodeAddReq.md)
 - [XbuyAppFormAdminFormBenefitCodeAddRes](docs/XbuyAppFormAdminFormBenefitCodeAddRes.md)
 - [XbuyAppFormAdminFormBenefitCodeListReq](docs/XbuyAppFormAdminFormBenefitCodeListReq.md)
 - [XbuyAppFormAdminFormBenefitCodeListRes](docs/XbuyAppFormAdminFormBenefitCodeListRes.md)
 - [XbuyAppFormAdminFormBenefitCodeStatisticsReq](docs/XbuyAppFormAdminFormBenefitCodeStatisticsReq.md)
 - [XbuyAppFormAdminFormBenefitCodeStatisticsRes](docs/XbuyAppFormAdminFormBenefitCodeStatisticsRes.md)
 - [XbuyAppFormAdminFormBindMemberGradeReq](docs/XbuyAppFormAdminFormBindMemberGradeReq.md)
 - [XbuyAppFormAdminFormBindMemberGradeRes](docs/XbuyAppFormAdminFormBindMemberGradeRes.md)
 - [XbuyAppFormAdminFormBindMemberTagReq](docs/XbuyAppFormAdminFormBindMemberTagReq.md)
 - [XbuyAppFormAdminFormBindMemberTagRes](docs/XbuyAppFormAdminFormBindMemberTagRes.md)
 - [XbuyAppFormAdminFormBindMemberToShopReq](docs/XbuyAppFormAdminFormBindMemberToShopReq.md)
 - [XbuyAppFormAdminFormBindMemberToShopRes](docs/XbuyAppFormAdminFormBindMemberToShopRes.md)
 - [XbuyAppFormAdminFormBlackListAddReq](docs/XbuyAppFormAdminFormBlackListAddReq.md)
 - [XbuyAppFormAdminFormBlackListAddRes](docs/XbuyAppFormAdminFormBlackListAddRes.md)
 - [XbuyAppFormAdminFormBlackListRemoveReq](docs/XbuyAppFormAdminFormBlackListRemoveReq.md)
 - [XbuyAppFormAdminFormBlackListRemoveRes](docs/XbuyAppFormAdminFormBlackListRemoveRes.md)
 - [XbuyAppFormAdminFormBlackListReq](docs/XbuyAppFormAdminFormBlackListReq.md)
 - [XbuyAppFormAdminFormBlackListRes](docs/XbuyAppFormAdminFormBlackListRes.md)
 - [XbuyAppFormAdminFormCancelAfterSalesReq](docs/XbuyAppFormAdminFormCancelAfterSalesReq.md)
 - [XbuyAppFormAdminFormCancelAfterSalesRes](docs/XbuyAppFormAdminFormCancelAfterSalesRes.md)
 - [XbuyAppFormAdminFormCatAddReq](docs/XbuyAppFormAdminFormCatAddReq.md)
 - [XbuyAppFormAdminFormCatAddRes](docs/XbuyAppFormAdminFormCatAddRes.md)
 - [XbuyAppFormAdminFormCatDeleteReq](docs/XbuyAppFormAdminFormCatDeleteReq.md)
 - [XbuyAppFormAdminFormCatDeleteRes](docs/XbuyAppFormAdminFormCatDeleteRes.md)
 - [XbuyAppFormAdminFormCatInfoReq](docs/XbuyAppFormAdminFormCatInfoReq.md)
 - [XbuyAppFormAdminFormCatInfoRes](docs/XbuyAppFormAdminFormCatInfoRes.md)
 - [XbuyAppFormAdminFormCatListReq](docs/XbuyAppFormAdminFormCatListReq.md)
 - [XbuyAppFormAdminFormCatListRes](docs/XbuyAppFormAdminFormCatListRes.md)
 - [XbuyAppFormAdminFormCatParentListReq](docs/XbuyAppFormAdminFormCatParentListReq.md)
 - [XbuyAppFormAdminFormCatParentListRes](docs/XbuyAppFormAdminFormCatParentListRes.md)
 - [XbuyAppFormAdminFormCatSubAddReq](docs/XbuyAppFormAdminFormCatSubAddReq.md)
 - [XbuyAppFormAdminFormCatSubAddRes](docs/XbuyAppFormAdminFormCatSubAddRes.md)
 - [XbuyAppFormAdminFormCatsList](docs/XbuyAppFormAdminFormCatsList.md)
 - [XbuyAppFormAdminFormChangeShopReq](docs/XbuyAppFormAdminFormChangeShopReq.md)
 - [XbuyAppFormAdminFormChangeShopRes](docs/XbuyAppFormAdminFormChangeShopRes.md)
 - [XbuyAppFormAdminFormCheckBindMemberGradeExistReq](docs/XbuyAppFormAdminFormCheckBindMemberGradeExistReq.md)
 - [XbuyAppFormAdminFormCheckBindMemberGradeExistRes](docs/XbuyAppFormAdminFormCheckBindMemberGradeExistRes.md)
 - [XbuyAppFormAdminFormCloseShopViewAutoTransReq](docs/XbuyAppFormAdminFormCloseShopViewAutoTransReq.md)
 - [XbuyAppFormAdminFormCloseShopViewAutoTransRes](docs/XbuyAppFormAdminFormCloseShopViewAutoTransRes.md)
 - [XbuyAppFormAdminFormConfigDeleteReq](docs/XbuyAppFormAdminFormConfigDeleteReq.md)
 - [XbuyAppFormAdminFormConfigDeleteRes](docs/XbuyAppFormAdminFormConfigDeleteRes.md)
 - [XbuyAppFormAdminFormConfigEditReq](docs/XbuyAppFormAdminFormConfigEditReq.md)
 - [XbuyAppFormAdminFormConfigEditRes](docs/XbuyAppFormAdminFormConfigEditRes.md)
 - [XbuyAppFormAdminFormConfigGetValueReq](docs/XbuyAppFormAdminFormConfigGetValueReq.md)
 - [XbuyAppFormAdminFormConfigGetValueRes](docs/XbuyAppFormAdminFormConfigGetValueRes.md)
 - [XbuyAppFormAdminFormConfigListReq](docs/XbuyAppFormAdminFormConfigListReq.md)
 - [XbuyAppFormAdminFormConfigListRes](docs/XbuyAppFormAdminFormConfigListRes.md)
 - [XbuyAppFormAdminFormConfigMaxSortReq](docs/XbuyAppFormAdminFormConfigMaxSortReq.md)
 - [XbuyAppFormAdminFormConfigMaxSortRes](docs/XbuyAppFormAdminFormConfigMaxSortRes.md)
 - [XbuyAppFormAdminFormConfigNameUniqueReq](docs/XbuyAppFormAdminFormConfigNameUniqueReq.md)
 - [XbuyAppFormAdminFormConfigNameUniqueRes](docs/XbuyAppFormAdminFormConfigNameUniqueRes.md)
 - [XbuyAppFormAdminFormConfigSysGetReq](docs/XbuyAppFormAdminFormConfigSysGetReq.md)
 - [XbuyAppFormAdminFormConfigSysGetRes](docs/XbuyAppFormAdminFormConfigSysGetRes.md)
 - [XbuyAppFormAdminFormConfigSysSetReq](docs/XbuyAppFormAdminFormConfigSysSetReq.md)
 - [XbuyAppFormAdminFormConfigSysSetRes](docs/XbuyAppFormAdminFormConfigSysSetRes.md)
 - [XbuyAppFormAdminFormConfigViewReq](docs/XbuyAppFormAdminFormConfigViewReq.md)
 - [XbuyAppFormAdminFormConfigViewRes](docs/XbuyAppFormAdminFormConfigViewRes.md)
 - [XbuyAppFormAdminFormCopyrightReq](docs/XbuyAppFormAdminFormCopyrightReq.md)
 - [XbuyAppFormAdminFormCopyrightRes](docs/XbuyAppFormAdminFormCopyrightRes.md)
 - [XbuyAppFormAdminFormCountryListReq](docs/XbuyAppFormAdminFormCountryListReq.md)
 - [XbuyAppFormAdminFormCountryListRes](docs/XbuyAppFormAdminFormCountryListRes.md)
 - [XbuyAppFormAdminFormDeleteMemberGradeReq](docs/XbuyAppFormAdminFormDeleteMemberGradeReq.md)
 - [XbuyAppFormAdminFormDeleteMemberGradeRes](docs/XbuyAppFormAdminFormDeleteMemberGradeRes.md)
 - [XbuyAppFormAdminFormDeleteMemberGroupRelationReq](docs/XbuyAppFormAdminFormDeleteMemberGroupRelationReq.md)
 - [XbuyAppFormAdminFormDeleteMemberGroupRelationRes](docs/XbuyAppFormAdminFormDeleteMemberGroupRelationRes.md)
 - [XbuyAppFormAdminFormDeleteMemberGroupReq](docs/XbuyAppFormAdminFormDeleteMemberGroupReq.md)
 - [XbuyAppFormAdminFormDeleteMemberGroupRes](docs/XbuyAppFormAdminFormDeleteMemberGroupRes.md)
 - [XbuyAppFormAdminFormDeleteMemberTagReq](docs/XbuyAppFormAdminFormDeleteMemberTagReq.md)
 - [XbuyAppFormAdminFormDeleteMemberTagRes](docs/XbuyAppFormAdminFormDeleteMemberTagRes.md)
 - [XbuyAppFormAdminFormDeleteSkuTaxReq](docs/XbuyAppFormAdminFormDeleteSkuTaxReq.md)
 - [XbuyAppFormAdminFormDeleteSkuTaxRes](docs/XbuyAppFormAdminFormDeleteSkuTaxRes.md)
 - [XbuyAppFormAdminFormDeptDeleteReq](docs/XbuyAppFormAdminFormDeptDeleteReq.md)
 - [XbuyAppFormAdminFormDeptDeleteRes](docs/XbuyAppFormAdminFormDeptDeleteRes.md)
 - [XbuyAppFormAdminFormDeptEditReq](docs/XbuyAppFormAdminFormDeptEditReq.md)
 - [XbuyAppFormAdminFormDeptEditRes](docs/XbuyAppFormAdminFormDeptEditRes.md)
 - [XbuyAppFormAdminFormDeptListReq](docs/XbuyAppFormAdminFormDeptListReq.md)
 - [XbuyAppFormAdminFormDeptListTreeReq](docs/XbuyAppFormAdminFormDeptListTreeReq.md)
 - [XbuyAppFormAdminFormDeptMaxSortReq](docs/XbuyAppFormAdminFormDeptMaxSortReq.md)
 - [XbuyAppFormAdminFormDeptMaxSortRes](docs/XbuyAppFormAdminFormDeptMaxSortRes.md)
 - [XbuyAppFormAdminFormDeptNameUniqueReq](docs/XbuyAppFormAdminFormDeptNameUniqueReq.md)
 - [XbuyAppFormAdminFormDeptNameUniqueRes](docs/XbuyAppFormAdminFormDeptNameUniqueRes.md)
 - [XbuyAppFormAdminFormDeptViewReq](docs/XbuyAppFormAdminFormDeptViewReq.md)
 - [XbuyAppFormAdminFormDeptViewRes](docs/XbuyAppFormAdminFormDeptViewRes.md)
 - [XbuyAppFormAdminFormDictAttributeReq](docs/XbuyAppFormAdminFormDictAttributeReq.md)
 - [XbuyAppFormAdminFormDictDataDeleteReq](docs/XbuyAppFormAdminFormDictDataDeleteReq.md)
 - [XbuyAppFormAdminFormDictDataDeleteRes](docs/XbuyAppFormAdminFormDictDataDeleteRes.md)
 - [XbuyAppFormAdminFormDictDataEditReq](docs/XbuyAppFormAdminFormDictDataEditReq.md)
 - [XbuyAppFormAdminFormDictDataEditRes](docs/XbuyAppFormAdminFormDictDataEditRes.md)
 - [XbuyAppFormAdminFormDictDataListReq](docs/XbuyAppFormAdminFormDictDataListReq.md)
 - [XbuyAppFormAdminFormDictDataListRes](docs/XbuyAppFormAdminFormDictDataListRes.md)
 - [XbuyAppFormAdminFormDictDataMaxSortReq](docs/XbuyAppFormAdminFormDictDataMaxSortReq.md)
 - [XbuyAppFormAdminFormDictDataMaxSortRes](docs/XbuyAppFormAdminFormDictDataMaxSortRes.md)
 - [XbuyAppFormAdminFormDictDataUniqueReq](docs/XbuyAppFormAdminFormDictDataUniqueReq.md)
 - [XbuyAppFormAdminFormDictDataUniqueRes](docs/XbuyAppFormAdminFormDictDataUniqueRes.md)
 - [XbuyAppFormAdminFormDictDataViewReq](docs/XbuyAppFormAdminFormDictDataViewReq.md)
 - [XbuyAppFormAdminFormDictDataViewRes](docs/XbuyAppFormAdminFormDictDataViewRes.md)
 - [XbuyAppFormAdminFormDictTypeDeleteReq](docs/XbuyAppFormAdminFormDictTypeDeleteReq.md)
 - [XbuyAppFormAdminFormDictTypeDeleteRes](docs/XbuyAppFormAdminFormDictTypeDeleteRes.md)
 - [XbuyAppFormAdminFormDictTypeEditReq](docs/XbuyAppFormAdminFormDictTypeEditReq.md)
 - [XbuyAppFormAdminFormDictTypeEditRes](docs/XbuyAppFormAdminFormDictTypeEditRes.md)
 - [XbuyAppFormAdminFormDictTypeExportReq](docs/XbuyAppFormAdminFormDictTypeExportReq.md)
 - [XbuyAppFormAdminFormDictTypeExportRes](docs/XbuyAppFormAdminFormDictTypeExportRes.md)
 - [XbuyAppFormAdminFormDictTypeListReq](docs/XbuyAppFormAdminFormDictTypeListReq.md)
 - [XbuyAppFormAdminFormDictTypeListRes](docs/XbuyAppFormAdminFormDictTypeListRes.md)
 - [XbuyAppFormAdminFormDictTypeRefreshCacheReq](docs/XbuyAppFormAdminFormDictTypeRefreshCacheReq.md)
 - [XbuyAppFormAdminFormDictTypeRefreshCacheRes](docs/XbuyAppFormAdminFormDictTypeRefreshCacheRes.md)
 - [XbuyAppFormAdminFormDictTypeUniqueReq](docs/XbuyAppFormAdminFormDictTypeUniqueReq.md)
 - [XbuyAppFormAdminFormDictTypeUniqueRes](docs/XbuyAppFormAdminFormDictTypeUniqueRes.md)
 - [XbuyAppFormAdminFormDictTypeViewReq](docs/XbuyAppFormAdminFormDictTypeViewReq.md)
 - [XbuyAppFormAdminFormDictTypeViewRes](docs/XbuyAppFormAdminFormDictTypeViewRes.md)
 - [XbuyAppFormAdminFormDomainAddReq](docs/XbuyAppFormAdminFormDomainAddReq.md)
 - [XbuyAppFormAdminFormDomainAddRes](docs/XbuyAppFormAdminFormDomainAddRes.md)
 - [XbuyAppFormAdminFormDomainDeleteReq](docs/XbuyAppFormAdminFormDomainDeleteReq.md)
 - [XbuyAppFormAdminFormDomainDeleteRes](docs/XbuyAppFormAdminFormDomainDeleteRes.md)
 - [XbuyAppFormAdminFormDomainGetMainReq](docs/XbuyAppFormAdminFormDomainGetMainReq.md)
 - [XbuyAppFormAdminFormDomainGetReq](docs/XbuyAppFormAdminFormDomainGetReq.md)
 - [XbuyAppFormAdminFormDomainGetRes](docs/XbuyAppFormAdminFormDomainGetRes.md)
 - [XbuyAppFormAdminFormDomainListReq](docs/XbuyAppFormAdminFormDomainListReq.md)
 - [XbuyAppFormAdminFormDomains](docs/XbuyAppFormAdminFormDomains.md)
 - [XbuyAppFormAdminFormDomianListRes](docs/XbuyAppFormAdminFormDomianListRes.md)
 - [XbuyAppFormAdminFormEmailBindReq](docs/XbuyAppFormAdminFormEmailBindReq.md)
 - [XbuyAppFormAdminFormEmailBindRes](docs/XbuyAppFormAdminFormEmailBindRes.md)
 - [XbuyAppFormAdminFormEmailListReq](docs/XbuyAppFormAdminFormEmailListReq.md)
 - [XbuyAppFormAdminFormEmailListRes](docs/XbuyAppFormAdminFormEmailListRes.md)
 - [XbuyAppFormAdminFormEmailPullReq](docs/XbuyAppFormAdminFormEmailPullReq.md)
 - [XbuyAppFormAdminFormEmailPullRes](docs/XbuyAppFormAdminFormEmailPullRes.md)
 - [XbuyAppFormAdminFormEmailSendReq](docs/XbuyAppFormAdminFormEmailSendReq.md)
 - [XbuyAppFormAdminFormEmailSendRes](docs/XbuyAppFormAdminFormEmailSendRes.md)
 - [XbuyAppFormAdminFormEmailSettingReq](docs/XbuyAppFormAdminFormEmailSettingReq.md)
 - [XbuyAppFormAdminFormEmailSettingRes](docs/XbuyAppFormAdminFormEmailSettingRes.md)
 - [XbuyAppFormAdminFormExportOrderReq](docs/XbuyAppFormAdminFormExportOrderReq.md)
 - [XbuyAppFormAdminFormExportOrderRes](docs/XbuyAppFormAdminFormExportOrderRes.md)
 - [XbuyAppFormAdminFormExpressAddReq](docs/XbuyAppFormAdminFormExpressAddReq.md)
 - [XbuyAppFormAdminFormExpressAddRes](docs/XbuyAppFormAdminFormExpressAddRes.md)
 - [XbuyAppFormAdminFormExpressDeleteReq](docs/XbuyAppFormAdminFormExpressDeleteReq.md)
 - [XbuyAppFormAdminFormExpressDeleteRes](docs/XbuyAppFormAdminFormExpressDeleteRes.md)
 - [XbuyAppFormAdminFormExpressGetReq](docs/XbuyAppFormAdminFormExpressGetReq.md)
 - [XbuyAppFormAdminFormExpressGetRes](docs/XbuyAppFormAdminFormExpressGetRes.md)
 - [XbuyAppFormAdminFormExpressListReq](docs/XbuyAppFormAdminFormExpressListReq.md)
 - [XbuyAppFormAdminFormExpressListRes](docs/XbuyAppFormAdminFormExpressListRes.md)
 - [XbuyAppFormAdminFormFileDeleteReq](docs/XbuyAppFormAdminFormFileDeleteReq.md)
 - [XbuyAppFormAdminFormFileDeleteRes](docs/XbuyAppFormAdminFormFileDeleteRes.md)
 - [XbuyAppFormAdminFormFileListReq](docs/XbuyAppFormAdminFormFileListReq.md)
 - [XbuyAppFormAdminFormFileListRes](docs/XbuyAppFormAdminFormFileListRes.md)
 - [XbuyAppFormAdminFormGetBindMemberGradeListByGradeReq](docs/XbuyAppFormAdminFormGetBindMemberGradeListByGradeReq.md)
 - [XbuyAppFormAdminFormGetBindMemberGradeListByGradeRes](docs/XbuyAppFormAdminFormGetBindMemberGradeListByGradeRes.md)
 - [XbuyAppFormAdminFormGetBindMemberGradeListByMemberReq](docs/XbuyAppFormAdminFormGetBindMemberGradeListByMemberReq.md)
 - [XbuyAppFormAdminFormGetBindMemberGradeListByMemberRes](docs/XbuyAppFormAdminFormGetBindMemberGradeListByMemberRes.md)
 - [XbuyAppFormAdminFormGetBindShopReq](docs/XbuyAppFormAdminFormGetBindShopReq.md)
 - [XbuyAppFormAdminFormGetBindShopRes](docs/XbuyAppFormAdminFormGetBindShopRes.md)
 - [XbuyAppFormAdminFormGetCurrencyReq](docs/XbuyAppFormAdminFormGetCurrencyReq.md)
 - [XbuyAppFormAdminFormGetCurrencyRes](docs/XbuyAppFormAdminFormGetCurrencyRes.md)
 - [XbuyAppFormAdminFormGetFundsReq](docs/XbuyAppFormAdminFormGetFundsReq.md)
 - [XbuyAppFormAdminFormGetFundsRes](docs/XbuyAppFormAdminFormGetFundsRes.md)
 - [XbuyAppFormAdminFormGetMallShopConfigsReq](docs/XbuyAppFormAdminFormGetMallShopConfigsReq.md)
 - [XbuyAppFormAdminFormGetMallShopConfigsRes](docs/XbuyAppFormAdminFormGetMallShopConfigsRes.md)
 - [XbuyAppFormAdminFormGetMemberBoundTagListReq](docs/XbuyAppFormAdminFormGetMemberBoundTagListReq.md)
 - [XbuyAppFormAdminFormGetMemberBoundTagListRes](docs/XbuyAppFormAdminFormGetMemberBoundTagListRes.md)
 - [XbuyAppFormAdminFormGetMemberByGroupIdAndShopIdReq](docs/XbuyAppFormAdminFormGetMemberByGroupIdAndShopIdReq.md)
 - [XbuyAppFormAdminFormGetMemberByGroupIdAndShopIdRes](docs/XbuyAppFormAdminFormGetMemberByGroupIdAndShopIdRes.md)
 - [XbuyAppFormAdminFormGetMemberByLevelIdAndShopIdReq](docs/XbuyAppFormAdminFormGetMemberByLevelIdAndShopIdReq.md)
 - [XbuyAppFormAdminFormGetMemberByLevelIdAndShopIdRes](docs/XbuyAppFormAdminFormGetMemberByLevelIdAndShopIdRes.md)
 - [XbuyAppFormAdminFormGetMemberByTagIdAndShopIdReq](docs/XbuyAppFormAdminFormGetMemberByTagIdAndShopIdReq.md)
 - [XbuyAppFormAdminFormGetMemberByTagIdAndShopIdRes](docs/XbuyAppFormAdminFormGetMemberByTagIdAndShopIdRes.md)
 - [XbuyAppFormAdminFormGetMemberGradeListByShopIdReq](docs/XbuyAppFormAdminFormGetMemberGradeListByShopIdReq.md)
 - [XbuyAppFormAdminFormGetMemberGradeListByShopIdRes](docs/XbuyAppFormAdminFormGetMemberGradeListByShopIdRes.md)
 - [XbuyAppFormAdminFormGetMemberGradeListReq](docs/XbuyAppFormAdminFormGetMemberGradeListReq.md)
 - [XbuyAppFormAdminFormGetMemberGradeListRes](docs/XbuyAppFormAdminFormGetMemberGradeListRes.md)
 - [XbuyAppFormAdminFormGetMemberGroupListByIdsReq](docs/XbuyAppFormAdminFormGetMemberGroupListByIdsReq.md)
 - [XbuyAppFormAdminFormGetMemberGroupListByIdsRes](docs/XbuyAppFormAdminFormGetMemberGroupListByIdsRes.md)
 - [XbuyAppFormAdminFormGetMemberGroupListByShopIdReq](docs/XbuyAppFormAdminFormGetMemberGroupListByShopIdReq.md)
 - [XbuyAppFormAdminFormGetMemberGroupListByShopIdRes](docs/XbuyAppFormAdminFormGetMemberGroupListByShopIdRes.md)
 - [XbuyAppFormAdminFormGetMemberGroupRelationByIdReq](docs/XbuyAppFormAdminFormGetMemberGroupRelationByIdReq.md)
 - [XbuyAppFormAdminFormGetMemberGroupRelationByIdRes](docs/XbuyAppFormAdminFormGetMemberGroupRelationByIdRes.md)
 - [XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdNotBindReq](docs/XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdNotBindReq.md)
 - [XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdNotBindRes](docs/XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdNotBindRes.md)
 - [XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdReq](docs/XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdReq.md)
 - [XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdRes](docs/XbuyAppFormAdminFormGetMemberGroupRelationListByGroupIdRes.md)
 - [XbuyAppFormAdminFormGetMemberGroupRelationListReq](docs/XbuyAppFormAdminFormGetMemberGroupRelationListReq.md)
 - [XbuyAppFormAdminFormGetMemberGroupRelationListRes](docs/XbuyAppFormAdminFormGetMemberGroupRelationListRes.md)
 - [XbuyAppFormAdminFormGetMemberListByGradeIdReq](docs/XbuyAppFormAdminFormGetMemberListByGradeIdReq.md)
 - [XbuyAppFormAdminFormGetMemberListByGradeIdRes](docs/XbuyAppFormAdminFormGetMemberListByGradeIdRes.md)
 - [XbuyAppFormAdminFormGetMemberListByGradeIdWithBindReq](docs/XbuyAppFormAdminFormGetMemberListByGradeIdWithBindReq.md)
 - [XbuyAppFormAdminFormGetMemberListByGradeIdWithBindRes](docs/XbuyAppFormAdminFormGetMemberListByGradeIdWithBindRes.md)
 - [XbuyAppFormAdminFormGetMemberListReq](docs/XbuyAppFormAdminFormGetMemberListReq.md)
 - [XbuyAppFormAdminFormGetMemberListRes](docs/XbuyAppFormAdminFormGetMemberListRes.md)
 - [XbuyAppFormAdminFormGetMemberRemarkListByMemberIdReq](docs/XbuyAppFormAdminFormGetMemberRemarkListByMemberIdReq.md)
 - [XbuyAppFormAdminFormGetMemberRemarkListByMemberIdRes](docs/XbuyAppFormAdminFormGetMemberRemarkListByMemberIdRes.md)
 - [XbuyAppFormAdminFormGetMemberRemarkListReq](docs/XbuyAppFormAdminFormGetMemberRemarkListReq.md)
 - [XbuyAppFormAdminFormGetMemberRemarkListRes](docs/XbuyAppFormAdminFormGetMemberRemarkListRes.md)
 - [XbuyAppFormAdminFormGetMemberTagBoundListByTagIdReq](docs/XbuyAppFormAdminFormGetMemberTagBoundListByTagIdReq.md)
 - [XbuyAppFormAdminFormGetMemberTagBoundListByTagIdRes](docs/XbuyAppFormAdminFormGetMemberTagBoundListByTagIdRes.md)
 - [XbuyAppFormAdminFormGetMemberTagBoundListByTagIdWithMemberInfoReq](docs/XbuyAppFormAdminFormGetMemberTagBoundListByTagIdWithMemberInfoReq.md)
 - [XbuyAppFormAdminFormGetMemberTagBoundListByTagIdWithMemberInfoRes](docs/XbuyAppFormAdminFormGetMemberTagBoundListByTagIdWithMemberInfoRes.md)
 - [XbuyAppFormAdminFormGetMemberTagBoundListReq](docs/XbuyAppFormAdminFormGetMemberTagBoundListReq.md)
 - [XbuyAppFormAdminFormGetMemberTagBoundListRes](docs/XbuyAppFormAdminFormGetMemberTagBoundListRes.md)
 - [XbuyAppFormAdminFormGetMemberTagListByShopIdReq](docs/XbuyAppFormAdminFormGetMemberTagListByShopIdReq.md)
 - [XbuyAppFormAdminFormGetMemberTagListByShopIdRes](docs/XbuyAppFormAdminFormGetMemberTagListByShopIdRes.md)
 - [XbuyAppFormAdminFormGetMemberTagListReq](docs/XbuyAppFormAdminFormGetMemberTagListReq.md)
 - [XbuyAppFormAdminFormGetMemberTagListRes](docs/XbuyAppFormAdminFormGetMemberTagListRes.md)
 - [XbuyAppFormAdminFormGetMemberTagUnBoundListByTagIdReq](docs/XbuyAppFormAdminFormGetMemberTagUnBoundListByTagIdReq.md)
 - [XbuyAppFormAdminFormGetMemberTagUnBoundListByTagIdRes](docs/XbuyAppFormAdminFormGetMemberTagUnBoundListByTagIdRes.md)
 - [XbuyAppFormAdminFormGetPayWaysReq](docs/XbuyAppFormAdminFormGetPayWaysReq.md)
 - [XbuyAppFormAdminFormGetPayWaysRes](docs/XbuyAppFormAdminFormGetPayWaysRes.md)
 - [XbuyAppFormAdminFormGetShopInfoReq](docs/XbuyAppFormAdminFormGetShopInfoReq.md)
 - [XbuyAppFormAdminFormGetShopInfoRes](docs/XbuyAppFormAdminFormGetShopInfoRes.md)
 - [XbuyAppFormAdminFormGetShopReq](docs/XbuyAppFormAdminFormGetShopReq.md)
 - [XbuyAppFormAdminFormGetShopRes](docs/XbuyAppFormAdminFormGetShopRes.md)
 - [XbuyAppFormAdminFormGetSkuStockReq](docs/XbuyAppFormAdminFormGetSkuStockReq.md)
 - [XbuyAppFormAdminFormGetSkuStockRes](docs/XbuyAppFormAdminFormGetSkuStockRes.md)
 - [XbuyAppFormAdminFormGetTaxRateByIdReq](docs/XbuyAppFormAdminFormGetTaxRateByIdReq.md)
 - [XbuyAppFormAdminFormGetTaxRateByIdRes](docs/XbuyAppFormAdminFormGetTaxRateByIdRes.md)
 - [XbuyAppFormAdminFormGetUnBindMemberListByGradeIdReq](docs/XbuyAppFormAdminFormGetUnBindMemberListByGradeIdReq.md)
 - [XbuyAppFormAdminFormGetUnBindMemberListByGradeIdRes](docs/XbuyAppFormAdminFormGetUnBindMemberListByGradeIdRes.md)
 - [XbuyAppFormAdminFormGoodAddReq](docs/XbuyAppFormAdminFormGoodAddReq.md)
 - [XbuyAppFormAdminFormGoodAddRes](docs/XbuyAppFormAdminFormGoodAddRes.md)
 - [XbuyAppFormAdminFormGoodDeleteReq](docs/XbuyAppFormAdminFormGoodDeleteReq.md)
 - [XbuyAppFormAdminFormGoodDeleteRes](docs/XbuyAppFormAdminFormGoodDeleteRes.md)
 - [XbuyAppFormAdminFormGoodFilingsReq](docs/XbuyAppFormAdminFormGoodFilingsReq.md)
 - [XbuyAppFormAdminFormGoodFilingsRes](docs/XbuyAppFormAdminFormGoodFilingsRes.md)
 - [XbuyAppFormAdminFormGoodInfoAddReq](docs/XbuyAppFormAdminFormGoodInfoAddReq.md)
 - [XbuyAppFormAdminFormGoodInfoAddRes](docs/XbuyAppFormAdminFormGoodInfoAddRes.md)
 - [XbuyAppFormAdminFormGoodInfoDeleteOneReq](docs/XbuyAppFormAdminFormGoodInfoDeleteOneReq.md)
 - [XbuyAppFormAdminFormGoodInfoDeleteOneRes](docs/XbuyAppFormAdminFormGoodInfoDeleteOneRes.md)
 - [XbuyAppFormAdminFormGoodInfoListReq](docs/XbuyAppFormAdminFormGoodInfoListReq.md)
 - [XbuyAppFormAdminFormGoodInfoListRes](docs/XbuyAppFormAdminFormGoodInfoListRes.md)
 - [XbuyAppFormAdminFormGoodInfoReq](docs/XbuyAppFormAdminFormGoodInfoReq.md)
 - [XbuyAppFormAdminFormGoodInfoRes](docs/XbuyAppFormAdminFormGoodInfoRes.md)
 - [XbuyAppFormAdminFormGoodInfoValueDeleteOneReq](docs/XbuyAppFormAdminFormGoodInfoValueDeleteOneReq.md)
 - [XbuyAppFormAdminFormGoodInfoValueDeleteOneRes](docs/XbuyAppFormAdminFormGoodInfoValueDeleteOneRes.md)
 - [XbuyAppFormAdminFormGoodInfoValuesAdd](docs/XbuyAppFormAdminFormGoodInfoValuesAdd.md)
 - [XbuyAppFormAdminFormGoodListReq](docs/XbuyAppFormAdminFormGoodListReq.md)
 - [XbuyAppFormAdminFormGoodListRes](docs/XbuyAppFormAdminFormGoodListRes.md)
 - [XbuyAppFormAdminFormGoodNumReq](docs/XbuyAppFormAdminFormGoodNumReq.md)
 - [XbuyAppFormAdminFormGoodNumRes](docs/XbuyAppFormAdminFormGoodNumRes.md)
 - [XbuyAppFormAdminFormGoodSubAddReq](docs/XbuyAppFormAdminFormGoodSubAddReq.md)
 - [XbuyAppFormAdminFormGoodSubAddRes](docs/XbuyAppFormAdminFormGoodSubAddRes.md)
 - [XbuyAppFormAdminFormGoodsList](docs/XbuyAppFormAdminFormGoodsList.md)
 - [XbuyAppFormAdminFormGoodsSalesReq](docs/XbuyAppFormAdminFormGoodsSalesReq.md)
 - [XbuyAppFormAdminFormGoodsSalesRes](docs/XbuyAppFormAdminFormGoodsSalesRes.md)
 - [XbuyAppFormAdminFormGoogleAddReq](docs/XbuyAppFormAdminFormGoogleAddReq.md)
 - [XbuyAppFormAdminFormGoogleAddRes](docs/XbuyAppFormAdminFormGoogleAddRes.md)
 - [XbuyAppFormAdminFormGoogleListReq](docs/XbuyAppFormAdminFormGoogleListReq.md)
 - [XbuyAppFormAdminFormGoogleListRes](docs/XbuyAppFormAdminFormGoogleListRes.md)
 - [XbuyAppFormAdminFormGoogleOauthReq](docs/XbuyAppFormAdminFormGoogleOauthReq.md)
 - [XbuyAppFormAdminFormGoogleOauthRes](docs/XbuyAppFormAdminFormGoogleOauthRes.md)
 - [XbuyAppFormAdminFormGoogleSearchAnalyticsReq](docs/XbuyAppFormAdminFormGoogleSearchAnalyticsReq.md)
 - [XbuyAppFormAdminFormGoogleSearchAnalyticsRes](docs/XbuyAppFormAdminFormGoogleSearchAnalyticsRes.md)
 - [XbuyAppFormAdminFormGoogleSubmitReq](docs/XbuyAppFormAdminFormGoogleSubmitReq.md)
 - [XbuyAppFormAdminFormGoogleSubmitRes](docs/XbuyAppFormAdminFormGoogleSubmitRes.md)
 - [XbuyAppFormAdminFormGoogleUrlIndexReq](docs/XbuyAppFormAdminFormGoogleUrlIndexReq.md)
 - [XbuyAppFormAdminFormGoogleUrlIndexRes](docs/XbuyAppFormAdminFormGoogleUrlIndexRes.md)
 - [XbuyAppFormAdminFormGroupAddResourceReq](docs/XbuyAppFormAdminFormGroupAddResourceReq.md)
 - [XbuyAppFormAdminFormGroupAddResourceRes](docs/XbuyAppFormAdminFormGroupAddResourceRes.md)
 - [XbuyAppFormAdminFormGroupBatchAddResourceReq](docs/XbuyAppFormAdminFormGroupBatchAddResourceReq.md)
 - [XbuyAppFormAdminFormGroupBatchAddResourceRes](docs/XbuyAppFormAdminFormGroupBatchAddResourceRes.md)
 - [XbuyAppFormAdminFormGroupResourceListReq](docs/XbuyAppFormAdminFormGroupResourceListReq.md)
 - [XbuyAppFormAdminFormImportOrderReq](docs/XbuyAppFormAdminFormImportOrderReq.md)
 - [XbuyAppFormAdminFormImportOrderRes](docs/XbuyAppFormAdminFormImportOrderRes.md)
 - [XbuyAppFormAdminFormIsShopBlackListReq](docs/XbuyAppFormAdminFormIsShopBlackListReq.md)
 - [XbuyAppFormAdminFormIsShopBlackListRes](docs/XbuyAppFormAdminFormIsShopBlackListRes.md)
 - [XbuyAppFormAdminFormLangReq](docs/XbuyAppFormAdminFormLangReq.md)
 - [XbuyAppFormAdminFormLangRes](docs/XbuyAppFormAdminFormLangRes.md)
 - [XbuyAppFormAdminFormListReq](docs/XbuyAppFormAdminFormListReq.md)
 - [XbuyAppFormAdminFormListRes](docs/XbuyAppFormAdminFormListRes.md)
 - [XbuyAppFormAdminFormLogClearReq](docs/XbuyAppFormAdminFormLogClearReq.md)
 - [XbuyAppFormAdminFormLogClearRes](docs/XbuyAppFormAdminFormLogClearRes.md)
 - [XbuyAppFormAdminFormLogExportReq](docs/XbuyAppFormAdminFormLogExportReq.md)
 - [XbuyAppFormAdminFormLogExportRes](docs/XbuyAppFormAdminFormLogExportRes.md)
 - [XbuyAppFormAdminFormLogListReq](docs/XbuyAppFormAdminFormLogListReq.md)
 - [XbuyAppFormAdminFormLogListRes](docs/XbuyAppFormAdminFormLogListRes.md)
 - [XbuyAppFormAdminFormLoginCaptchaReq](docs/XbuyAppFormAdminFormLoginCaptchaReq.md)
 - [XbuyAppFormAdminFormLoginCaptchaRes](docs/XbuyAppFormAdminFormLoginCaptchaRes.md)
 - [XbuyAppFormAdminFormLoginLogoutReq](docs/XbuyAppFormAdminFormLoginLogoutReq.md)
 - [XbuyAppFormAdminFormLoginLogoutRes](docs/XbuyAppFormAdminFormLoginLogoutRes.md)
 - [XbuyAppFormAdminFormLoginReq](docs/XbuyAppFormAdminFormLoginReq.md)
 - [XbuyAppFormAdminFormLoginRes](docs/XbuyAppFormAdminFormLoginRes.md)
 - [XbuyAppFormAdminFormLoginTokenReq](docs/XbuyAppFormAdminFormLoginTokenReq.md)
 - [XbuyAppFormAdminFormMakeTestDataReq](docs/XbuyAppFormAdminFormMakeTestDataReq.md)
 - [XbuyAppFormAdminFormMakeTestDataRes](docs/XbuyAppFormAdminFormMakeTestDataRes.md)
 - [XbuyAppFormAdminFormMemberDeleteReq](docs/XbuyAppFormAdminFormMemberDeleteReq.md)
 - [XbuyAppFormAdminFormMemberDeleteRes](docs/XbuyAppFormAdminFormMemberDeleteRes.md)
 - [XbuyAppFormAdminFormMemberEditReq](docs/XbuyAppFormAdminFormMemberEditReq.md)
 - [XbuyAppFormAdminFormMemberEditRes](docs/XbuyAppFormAdminFormMemberEditRes.md)
 - [XbuyAppFormAdminFormMemberEmailUniqueReq](docs/XbuyAppFormAdminFormMemberEmailUniqueReq.md)
 - [XbuyAppFormAdminFormMemberEmailUniqueRes](docs/XbuyAppFormAdminFormMemberEmailUniqueRes.md)
 - [XbuyAppFormAdminFormMemberGroupListReq](docs/XbuyAppFormAdminFormMemberGroupListReq.md)
 - [XbuyAppFormAdminFormMemberGroupListRes](docs/XbuyAppFormAdminFormMemberGroupListRes.md)
 - [XbuyAppFormAdminFormMemberInfoReq](docs/XbuyAppFormAdminFormMemberInfoReq.md)
 - [XbuyAppFormAdminFormMemberInfoRes](docs/XbuyAppFormAdminFormMemberInfoRes.md)
 - [XbuyAppFormAdminFormMemberListReq](docs/XbuyAppFormAdminFormMemberListReq.md)
 - [XbuyAppFormAdminFormMemberListRes](docs/XbuyAppFormAdminFormMemberListRes.md)
 - [XbuyAppFormAdminFormMemberMaxSortReq](docs/XbuyAppFormAdminFormMemberMaxSortReq.md)
 - [XbuyAppFormAdminFormMemberMaxSortRes](docs/XbuyAppFormAdminFormMemberMaxSortRes.md)
 - [XbuyAppFormAdminFormMemberMobileUniqueReq](docs/XbuyAppFormAdminFormMemberMobileUniqueReq.md)
 - [XbuyAppFormAdminFormMemberMobileUniqueRes](docs/XbuyAppFormAdminFormMemberMobileUniqueRes.md)
 - [XbuyAppFormAdminFormMemberNameUniqueReq](docs/XbuyAppFormAdminFormMemberNameUniqueReq.md)
 - [XbuyAppFormAdminFormMemberNameUniqueRes](docs/XbuyAppFormAdminFormMemberNameUniqueRes.md)
 - [XbuyAppFormAdminFormMemberNewRegisterReq](docs/XbuyAppFormAdminFormMemberNewRegisterReq.md)
 - [XbuyAppFormAdminFormMemberNewRegisterRes](docs/XbuyAppFormAdminFormMemberNewRegisterRes.md)
 - [XbuyAppFormAdminFormMemberProfileReq](docs/XbuyAppFormAdminFormMemberProfileReq.md)
 - [XbuyAppFormAdminFormMemberProfileRes](docs/XbuyAppFormAdminFormMemberProfileRes.md)
 - [XbuyAppFormAdminFormMemberResetPwdReq](docs/XbuyAppFormAdminFormMemberResetPwdReq.md)
 - [XbuyAppFormAdminFormMemberResetPwdRes](docs/XbuyAppFormAdminFormMemberResetPwdRes.md)
 - [XbuyAppFormAdminFormMemberUpdateAvatarReq](docs/XbuyAppFormAdminFormMemberUpdateAvatarReq.md)
 - [XbuyAppFormAdminFormMemberUpdateAvatarRes](docs/XbuyAppFormAdminFormMemberUpdateAvatarRes.md)
 - [XbuyAppFormAdminFormMemberUpdateProfileReq](docs/XbuyAppFormAdminFormMemberUpdateProfileReq.md)
 - [XbuyAppFormAdminFormMemberUpdateProfileRes](docs/XbuyAppFormAdminFormMemberUpdateProfileRes.md)
 - [XbuyAppFormAdminFormMemberUpdatePwdReq](docs/XbuyAppFormAdminFormMemberUpdatePwdReq.md)
 - [XbuyAppFormAdminFormMemberUpdatePwdRes](docs/XbuyAppFormAdminFormMemberUpdatePwdRes.md)
 - [XbuyAppFormAdminFormMemberViewReq](docs/XbuyAppFormAdminFormMemberViewReq.md)
 - [XbuyAppFormAdminFormMemberViewRes](docs/XbuyAppFormAdminFormMemberViewRes.md)
 - [XbuyAppFormAdminFormMenuCodeUniqueReq](docs/XbuyAppFormAdminFormMenuCodeUniqueReq.md)
 - [XbuyAppFormAdminFormMenuCodeUniqueRes](docs/XbuyAppFormAdminFormMenuCodeUniqueRes.md)
 - [XbuyAppFormAdminFormMenuDeleteReq](docs/XbuyAppFormAdminFormMenuDeleteReq.md)
 - [XbuyAppFormAdminFormMenuDeleteRes](docs/XbuyAppFormAdminFormMenuDeleteRes.md)
 - [XbuyAppFormAdminFormMenuEditReq](docs/XbuyAppFormAdminFormMenuEditReq.md)
 - [XbuyAppFormAdminFormMenuEditRes](docs/XbuyAppFormAdminFormMenuEditRes.md)
 - [XbuyAppFormAdminFormMenuListReq](docs/XbuyAppFormAdminFormMenuListReq.md)
 - [XbuyAppFormAdminFormMenuListRes](docs/XbuyAppFormAdminFormMenuListRes.md)
 - [XbuyAppFormAdminFormMenuMaxSortReq](docs/XbuyAppFormAdminFormMenuMaxSortReq.md)
 - [XbuyAppFormAdminFormMenuMaxSortRes](docs/XbuyAppFormAdminFormMenuMaxSortRes.md)
 - [XbuyAppFormAdminFormMenuNameUniqueReq](docs/XbuyAppFormAdminFormMenuNameUniqueReq.md)
 - [XbuyAppFormAdminFormMenuNameUniqueRes](docs/XbuyAppFormAdminFormMenuNameUniqueRes.md)
 - [XbuyAppFormAdminFormMenuRoleListReq](docs/XbuyAppFormAdminFormMenuRoleListReq.md)
 - [XbuyAppFormAdminFormMenuRoleListRes](docs/XbuyAppFormAdminFormMenuRoleListRes.md)
 - [XbuyAppFormAdminFormMenuSearchListReq](docs/XbuyAppFormAdminFormMenuSearchListReq.md)
 - [XbuyAppFormAdminFormMenuSubRoleListReq](docs/XbuyAppFormAdminFormMenuSubRoleListReq.md)
 - [XbuyAppFormAdminFormMenuViewReq](docs/XbuyAppFormAdminFormMenuViewReq.md)
 - [XbuyAppFormAdminFormMenuViewRes](docs/XbuyAppFormAdminFormMenuViewRes.md)
 - [XbuyAppFormAdminFormModeListReq](docs/XbuyAppFormAdminFormModeListReq.md)
 - [XbuyAppFormAdminFormModeListRes](docs/XbuyAppFormAdminFormModeListRes.md)
 - [XbuyAppFormAdminFormNoticeDeleteReq](docs/XbuyAppFormAdminFormNoticeDeleteReq.md)
 - [XbuyAppFormAdminFormNoticeDeleteRes](docs/XbuyAppFormAdminFormNoticeDeleteRes.md)
 - [XbuyAppFormAdminFormNoticeEditReq](docs/XbuyAppFormAdminFormNoticeEditReq.md)
 - [XbuyAppFormAdminFormNoticeEditRes](docs/XbuyAppFormAdminFormNoticeEditRes.md)
 - [XbuyAppFormAdminFormNoticeListReq](docs/XbuyAppFormAdminFormNoticeListReq.md)
 - [XbuyAppFormAdminFormNoticeListRes](docs/XbuyAppFormAdminFormNoticeListRes.md)
 - [XbuyAppFormAdminFormNoticeMaxSortReq](docs/XbuyAppFormAdminFormNoticeMaxSortReq.md)
 - [XbuyAppFormAdminFormNoticeMaxSortRes](docs/XbuyAppFormAdminFormNoticeMaxSortRes.md)
 - [XbuyAppFormAdminFormNoticeNameUniqueReq](docs/XbuyAppFormAdminFormNoticeNameUniqueReq.md)
 - [XbuyAppFormAdminFormNoticeNameUniqueRes](docs/XbuyAppFormAdminFormNoticeNameUniqueRes.md)
 - [XbuyAppFormAdminFormNoticeViewReq](docs/XbuyAppFormAdminFormNoticeViewReq.md)
 - [XbuyAppFormAdminFormNoticeViewRes](docs/XbuyAppFormAdminFormNoticeViewRes.md)
 - [XbuyAppFormAdminFormOnShelvesReq](docs/XbuyAppFormAdminFormOnShelvesReq.md)
 - [XbuyAppFormAdminFormOnShelvesRes](docs/XbuyAppFormAdminFormOnShelvesRes.md)
 - [XbuyAppFormAdminFormOpenShopViewAutoTransReq](docs/XbuyAppFormAdminFormOpenShopViewAutoTransReq.md)
 - [XbuyAppFormAdminFormOpenShopViewAutoTransRes](docs/XbuyAppFormAdminFormOpenShopViewAutoTransRes.md)
 - [XbuyAppFormAdminFormOrderDataReq](docs/XbuyAppFormAdminFormOrderDataReq.md)
 - [XbuyAppFormAdminFormOrderDataRes](docs/XbuyAppFormAdminFormOrderDataRes.md)
 - [XbuyAppFormAdminFormOrderInsuredReq](docs/XbuyAppFormAdminFormOrderInsuredReq.md)
 - [XbuyAppFormAdminFormOrderInsuredRes](docs/XbuyAppFormAdminFormOrderInsuredRes.md)
 - [XbuyAppFormAdminFormOrderItemRes](docs/XbuyAppFormAdminFormOrderItemRes.md)
 - [XbuyAppFormAdminFormOrderItemsReq](docs/XbuyAppFormAdminFormOrderItemsReq.md)
 - [XbuyAppFormAdminFormOrderListReq](docs/XbuyAppFormAdminFormOrderListReq.md)
 - [XbuyAppFormAdminFormOrderListRes](docs/XbuyAppFormAdminFormOrderListRes.md)
 - [XbuyAppFormAdminFormOrderMoneyReq](docs/XbuyAppFormAdminFormOrderMoneyReq.md)
 - [XbuyAppFormAdminFormOrderMoneyRes](docs/XbuyAppFormAdminFormOrderMoneyRes.md)
 - [XbuyAppFormAdminFormOrderPushLogReq](docs/XbuyAppFormAdminFormOrderPushLogReq.md)
 - [XbuyAppFormAdminFormOrderPushLogRes](docs/XbuyAppFormAdminFormOrderPushLogRes.md)
 - [XbuyAppFormAdminFormOrderPushReq](docs/XbuyAppFormAdminFormOrderPushReq.md)
 - [XbuyAppFormAdminFormOrderPushRes](docs/XbuyAppFormAdminFormOrderPushRes.md)
 - [XbuyAppFormAdminFormOrderShipReq](docs/XbuyAppFormAdminFormOrderShipReq.md)
 - [XbuyAppFormAdminFormOrderShipRes](docs/XbuyAppFormAdminFormOrderShipRes.md)
 - [XbuyAppFormAdminFormOrderWaybillExpressReq](docs/XbuyAppFormAdminFormOrderWaybillExpressReq.md)
 - [XbuyAppFormAdminFormOrderWaybillExpressRes](docs/XbuyAppFormAdminFormOrderWaybillExpressRes.md)
 - [XbuyAppFormAdminFormPayWays](docs/XbuyAppFormAdminFormPayWays.md)
 - [XbuyAppFormAdminFormPortalConfig](docs/XbuyAppFormAdminFormPortalConfig.md)
 - [XbuyAppFormAdminFormPostCodeUniqueReq](docs/XbuyAppFormAdminFormPostCodeUniqueReq.md)
 - [XbuyAppFormAdminFormPostCodeUniqueRes](docs/XbuyAppFormAdminFormPostCodeUniqueRes.md)
 - [XbuyAppFormAdminFormPostDeleteReq](docs/XbuyAppFormAdminFormPostDeleteReq.md)
 - [XbuyAppFormAdminFormPostDeleteRes](docs/XbuyAppFormAdminFormPostDeleteRes.md)
 - [XbuyAppFormAdminFormPostEditReq](docs/XbuyAppFormAdminFormPostEditReq.md)
 - [XbuyAppFormAdminFormPostEditRes](docs/XbuyAppFormAdminFormPostEditRes.md)
 - [XbuyAppFormAdminFormPostFeeAddReq](docs/XbuyAppFormAdminFormPostFeeAddReq.md)
 - [XbuyAppFormAdminFormPostFeeAddRes](docs/XbuyAppFormAdminFormPostFeeAddRes.md)
 - [XbuyAppFormAdminFormPostFeeListReq](docs/XbuyAppFormAdminFormPostFeeListReq.md)
 - [XbuyAppFormAdminFormPostFeeListRes](docs/XbuyAppFormAdminFormPostFeeListRes.md)
 - [XbuyAppFormAdminFormPostFeeModeAddReq](docs/XbuyAppFormAdminFormPostFeeModeAddReq.md)
 - [XbuyAppFormAdminFormPostFeeModeAddRes](docs/XbuyAppFormAdminFormPostFeeModeAddRes.md)
 - [XbuyAppFormAdminFormPostListReq](docs/XbuyAppFormAdminFormPostListReq.md)
 - [XbuyAppFormAdminFormPostListRes](docs/XbuyAppFormAdminFormPostListRes.md)
 - [XbuyAppFormAdminFormPostMaxSortReq](docs/XbuyAppFormAdminFormPostMaxSortReq.md)
 - [XbuyAppFormAdminFormPostMaxSortRes](docs/XbuyAppFormAdminFormPostMaxSortRes.md)
 - [XbuyAppFormAdminFormPostMenuEditReq](docs/XbuyAppFormAdminFormPostMenuEditReq.md)
 - [XbuyAppFormAdminFormPostMenuEditRes](docs/XbuyAppFormAdminFormPostMenuEditRes.md)
 - [XbuyAppFormAdminFormPostNameUniqueReq](docs/XbuyAppFormAdminFormPostNameUniqueReq.md)
 - [XbuyAppFormAdminFormPostNameUniqueRes](docs/XbuyAppFormAdminFormPostNameUniqueRes.md)
 - [XbuyAppFormAdminFormPostViewReq](docs/XbuyAppFormAdminFormPostViewReq.md)
 - [XbuyAppFormAdminFormPostViewRes](docs/XbuyAppFormAdminFormPostViewRes.md)
 - [XbuyAppFormAdminFormPriorityPolicyCreateReq](docs/XbuyAppFormAdminFormPriorityPolicyCreateReq.md)
 - [XbuyAppFormAdminFormPriorityPolicyCreateRes](docs/XbuyAppFormAdminFormPriorityPolicyCreateRes.md)
 - [XbuyAppFormAdminFormPriorityPolicyUpdateReq](docs/XbuyAppFormAdminFormPriorityPolicyUpdateReq.md)
 - [XbuyAppFormAdminFormPriorityPolicyUpdateRes](docs/XbuyAppFormAdminFormPriorityPolicyUpdateRes.md)
 - [XbuyAppFormAdminFormPropId](docs/XbuyAppFormAdminFormPropId.md)
 - [XbuyAppFormAdminFormReChargeReq](docs/XbuyAppFormAdminFormReChargeReq.md)
 - [XbuyAppFormAdminFormReChargeRes](docs/XbuyAppFormAdminFormReChargeRes.md)
 - [XbuyAppFormAdminFormReChargeResData](docs/XbuyAppFormAdminFormReChargeResData.md)
 - [XbuyAppFormAdminFormReadSkuTaxReq](docs/XbuyAppFormAdminFormReadSkuTaxReq.md)
 - [XbuyAppFormAdminFormReadSkuTaxRes](docs/XbuyAppFormAdminFormReadSkuTaxRes.md)
 - [XbuyAppFormAdminFormRealTimeDataReq](docs/XbuyAppFormAdminFormRealTimeDataReq.md)
 - [XbuyAppFormAdminFormRealTimeDataRes](docs/XbuyAppFormAdminFormRealTimeDataRes.md)
 - [XbuyAppFormAdminFormRefundReq](docs/XbuyAppFormAdminFormRefundReq.md)
 - [XbuyAppFormAdminFormRefundRes](docs/XbuyAppFormAdminFormRefundRes.md)
 - [XbuyAppFormAdminFormResourceGroupAddReq](docs/XbuyAppFormAdminFormResourceGroupAddReq.md)
 - [XbuyAppFormAdminFormResourceGroupListReq](docs/XbuyAppFormAdminFormResourceGroupListReq.md)
 - [XbuyAppFormAdminFormResourceGroupListRes](docs/XbuyAppFormAdminFormResourceGroupListRes.md)
 - [XbuyAppFormAdminFormResourceListReq](docs/XbuyAppFormAdminFormResourceListReq.md)
 - [XbuyAppFormAdminFormResourceListRes](docs/XbuyAppFormAdminFormResourceListRes.md)
 - [XbuyAppFormAdminFormReviewListReq](docs/XbuyAppFormAdminFormReviewListReq.md)
 - [XbuyAppFormAdminFormReviewListRes](docs/XbuyAppFormAdminFormReviewListRes.md)
 - [XbuyAppFormAdminFormRoleDynamicBase](docs/XbuyAppFormAdminFormRoleDynamicBase.md)
 - [XbuyAppFormAdminFormRoleDynamicMenu](docs/XbuyAppFormAdminFormRoleDynamicMenu.md)
 - [XbuyAppFormAdminFormRoleDynamicMeta](docs/XbuyAppFormAdminFormRoleDynamicMeta.md)
 - [XbuyAppFormAdminFormRoleDynamicReq](docs/XbuyAppFormAdminFormRoleDynamicReq.md)
 - [XbuyAppFormAdminFormRoleListReq](docs/XbuyAppFormAdminFormRoleListReq.md)
 - [XbuyAppFormAdminFormRoleListRes](docs/XbuyAppFormAdminFormRoleListRes.md)
 - [XbuyAppFormAdminFormRoleMemberListReq](docs/XbuyAppFormAdminFormRoleMemberListReq.md)
 - [XbuyAppFormAdminFormRoleMemberListRes](docs/XbuyAppFormAdminFormRoleMemberListRes.md)
 - [XbuyAppFormAdminFormRoleMenuEditReq](docs/XbuyAppFormAdminFormRoleMenuEditReq.md)
 - [XbuyAppFormAdminFormRoleMenuEditRes](docs/XbuyAppFormAdminFormRoleMenuEditRes.md)
 - [XbuyAppFormAdminFormRoleShopAddReq](docs/XbuyAppFormAdminFormRoleShopAddReq.md)
 - [XbuyAppFormAdminFormRoleShopAddRes](docs/XbuyAppFormAdminFormRoleShopAddRes.md)
 - [XbuyAppFormAdminFormRoleShopDelReq](docs/XbuyAppFormAdminFormRoleShopDelReq.md)
 - [XbuyAppFormAdminFormRoleShopDelRes](docs/XbuyAppFormAdminFormRoleShopDelRes.md)
 - [XbuyAppFormAdminFormRoleShopListReq](docs/XbuyAppFormAdminFormRoleShopListReq.md)
 - [XbuyAppFormAdminFormRoleShopListRes](docs/XbuyAppFormAdminFormRoleShopListRes.md)
 - [XbuyAppFormAdminFormSaveCatReq](docs/XbuyAppFormAdminFormSaveCatReq.md)
 - [XbuyAppFormAdminFormSaveCatRes](docs/XbuyAppFormAdminFormSaveCatRes.md)
 - [XbuyAppFormAdminFormSaveMallShopConfigsReq](docs/XbuyAppFormAdminFormSaveMallShopConfigsReq.md)
 - [XbuyAppFormAdminFormSaveMallShopConfigsRes](docs/XbuyAppFormAdminFormSaveMallShopConfigsRes.md)
 - [XbuyAppFormAdminFormSaveReq](docs/XbuyAppFormAdminFormSaveReq.md)
 - [XbuyAppFormAdminFormSaveRes](docs/XbuyAppFormAdminFormSaveRes.md)
 - [XbuyAppFormAdminFormScatsList](docs/XbuyAppFormAdminFormScatsList.md)
 - [XbuyAppFormAdminFormSearchCustomsItemInfoReq](docs/XbuyAppFormAdminFormSearchCustomsItemInfoReq.md)
 - [XbuyAppFormAdminFormSearchCustomsItemInfoRes](docs/XbuyAppFormAdminFormSearchCustomsItemInfoRes.md)
 - [XbuyAppFormAdminFormSearchHsCodeInfoReq](docs/XbuyAppFormAdminFormSearchHsCodeInfoReq.md)
 - [XbuyAppFormAdminFormSearchHsCodeInfoRes](docs/XbuyAppFormAdminFormSearchHsCodeInfoRes.md)
 - [XbuyAppFormAdminFormSearchSkuInfoReq](docs/XbuyAppFormAdminFormSearchSkuInfoReq.md)
 - [XbuyAppFormAdminFormSearchSkuInfoRes](docs/XbuyAppFormAdminFormSearchSkuInfoRes.md)
 - [XbuyAppFormAdminFormServiceAddReq](docs/XbuyAppFormAdminFormServiceAddReq.md)
 - [XbuyAppFormAdminFormServiceAddRes](docs/XbuyAppFormAdminFormServiceAddRes.md)
 - [XbuyAppFormAdminFormServiceDetailReq](docs/XbuyAppFormAdminFormServiceDetailReq.md)
 - [XbuyAppFormAdminFormServiceDetailRes](docs/XbuyAppFormAdminFormServiceDetailRes.md)
 - [XbuyAppFormAdminFormServiceDetailsRes](docs/XbuyAppFormAdminFormServiceDetailsRes.md)
 - [XbuyAppFormAdminFormServiceListReq](docs/XbuyAppFormAdminFormServiceListReq.md)
 - [XbuyAppFormAdminFormServiceListRes](docs/XbuyAppFormAdminFormServiceListRes.md)
 - [XbuyAppFormAdminFormServicePauseReq](docs/XbuyAppFormAdminFormServicePauseReq.md)
 - [XbuyAppFormAdminFormServicePauseRes](docs/XbuyAppFormAdminFormServicePauseRes.md)
 - [XbuyAppFormAdminFormSetCurrencyRatesReq](docs/XbuyAppFormAdminFormSetCurrencyRatesReq.md)
 - [XbuyAppFormAdminFormSetCurrencyRatesRes](docs/XbuyAppFormAdminFormSetCurrencyRatesRes.md)
 - [XbuyAppFormAdminFormShopBlackListAddReq](docs/XbuyAppFormAdminFormShopBlackListAddReq.md)
 - [XbuyAppFormAdminFormShopBlackListAddRes](docs/XbuyAppFormAdminFormShopBlackListAddRes.md)
 - [XbuyAppFormAdminFormShopBlackListRemoveReq](docs/XbuyAppFormAdminFormShopBlackListRemoveReq.md)
 - [XbuyAppFormAdminFormShopBlackListRemoveRes](docs/XbuyAppFormAdminFormShopBlackListRemoveRes.md)
 - [XbuyAppFormAdminFormShopBlackListReq](docs/XbuyAppFormAdminFormShopBlackListReq.md)
 - [XbuyAppFormAdminFormShopBlackListRes](docs/XbuyAppFormAdminFormShopBlackListRes.md)
 - [XbuyAppFormAdminFormShopListReq](docs/XbuyAppFormAdminFormShopListReq.md)
 - [XbuyAppFormAdminFormShopListRes](docs/XbuyAppFormAdminFormShopListRes.md)
 - [XbuyAppFormAdminFormShopMemberListReq](docs/XbuyAppFormAdminFormShopMemberListReq.md)
 - [XbuyAppFormAdminFormShopMemberListRes](docs/XbuyAppFormAdminFormShopMemberListRes.md)
 - [XbuyAppFormAdminFormShopViewAddReq](docs/XbuyAppFormAdminFormShopViewAddReq.md)
 - [XbuyAppFormAdminFormShopViewAddRes](docs/XbuyAppFormAdminFormShopViewAddRes.md)
 - [XbuyAppFormAdminFormShopViewInfoReq](docs/XbuyAppFormAdminFormShopViewInfoReq.md)
 - [XbuyAppFormAdminFormShopViewInfoRes](docs/XbuyAppFormAdminFormShopViewInfoRes.md)
 - [XbuyAppFormAdminFormShopViewListReq](docs/XbuyAppFormAdminFormShopViewListReq.md)
 - [XbuyAppFormAdminFormShopViewListRes](docs/XbuyAppFormAdminFormShopViewListRes.md)
 - [XbuyAppFormAdminFormSkuAddParam](docs/XbuyAppFormAdminFormSkuAddParam.md)
 - [XbuyAppFormAdminFormSkuAddReq](docs/XbuyAppFormAdminFormSkuAddReq.md)
 - [XbuyAppFormAdminFormSkuAddRes](docs/XbuyAppFormAdminFormSkuAddRes.md)
 - [XbuyAppFormAdminFormSkuDeleteReq](docs/XbuyAppFormAdminFormSkuDeleteReq.md)
 - [XbuyAppFormAdminFormSkuDeleteRes](docs/XbuyAppFormAdminFormSkuDeleteRes.md)
 - [XbuyAppFormAdminFormSkuList2Res](docs/XbuyAppFormAdminFormSkuList2Res.md)
 - [XbuyAppFormAdminFormSkuListReq](docs/XbuyAppFormAdminFormSkuListReq.md)
 - [XbuyAppFormAdminFormSkuPropsAddReq](docs/XbuyAppFormAdminFormSkuPropsAddReq.md)
 - [XbuyAppFormAdminFormSkuPropsAddRes](docs/XbuyAppFormAdminFormSkuPropsAddRes.md)
 - [XbuyAppFormAdminFormSkuPropsDeleteReq](docs/XbuyAppFormAdminFormSkuPropsDeleteReq.md)
 - [XbuyAppFormAdminFormSkuPropsDeleteRes](docs/XbuyAppFormAdminFormSkuPropsDeleteRes.md)
 - [XbuyAppFormAdminFormSkuPropsListReq](docs/XbuyAppFormAdminFormSkuPropsListReq.md)
 - [XbuyAppFormAdminFormSkuPropsListRes](docs/XbuyAppFormAdminFormSkuPropsListRes.md)
 - [XbuyAppFormAdminFormSkuPropsSubAddReq](docs/XbuyAppFormAdminFormSkuPropsSubAddReq.md)
 - [XbuyAppFormAdminFormSkuPropsSubAddRes](docs/XbuyAppFormAdminFormSkuPropsSubAddRes.md)
 - [XbuyAppFormAdminFormSkuSubAddParam](docs/XbuyAppFormAdminFormSkuSubAddParam.md)
 - [XbuyAppFormAdminFormSkuSubAddParamPropIds](docs/XbuyAppFormAdminFormSkuSubAddParamPropIds.md)
 - [XbuyAppFormAdminFormSkuSubAddReq](docs/XbuyAppFormAdminFormSkuSubAddReq.md)
 - [XbuyAppFormAdminFormSkuSubAddRes](docs/XbuyAppFormAdminFormSkuSubAddRes.md)
 - [XbuyAppFormAdminFormSkuTaxRateListReq](docs/XbuyAppFormAdminFormSkuTaxRateListReq.md)
 - [XbuyAppFormAdminFormSkuTaxRateListRes](docs/XbuyAppFormAdminFormSkuTaxRateListRes.md)
 - [XbuyAppFormAdminFormSkuValueAddReq](docs/XbuyAppFormAdminFormSkuValueAddReq.md)
 - [XbuyAppFormAdminFormSkuValueAddRes](docs/XbuyAppFormAdminFormSkuValueAddRes.md)
 - [XbuyAppFormAdminFormSkuValueDeleteReq](docs/XbuyAppFormAdminFormSkuValueDeleteReq.md)
 - [XbuyAppFormAdminFormSkuValueDeleteRes](docs/XbuyAppFormAdminFormSkuValueDeleteRes.md)
 - [XbuyAppFormAdminFormSkuValueSubAddReq](docs/XbuyAppFormAdminFormSkuValueSubAddReq.md)
 - [XbuyAppFormAdminFormSkuValueSubAddRes](docs/XbuyAppFormAdminFormSkuValueSubAddRes.md)
 - [XbuyAppFormAdminFormStandardCatReq](docs/XbuyAppFormAdminFormStandardCatReq.md)
 - [XbuyAppFormAdminFormStandardCatRes](docs/XbuyAppFormAdminFormStandardCatRes.md)
 - [XbuyAppFormAdminFormStockDetails](docs/XbuyAppFormAdminFormStockDetails.md)
 - [XbuyAppFormAdminFormSubLoginReq](docs/XbuyAppFormAdminFormSubLoginReq.md)
 - [XbuyAppFormAdminFormSubLoginRes](docs/XbuyAppFormAdminFormSubLoginRes.md)
 - [XbuyAppFormAdminFormSubmemberDeleteReq](docs/XbuyAppFormAdminFormSubmemberDeleteReq.md)
 - [XbuyAppFormAdminFormSubmemberDeleteRes](docs/XbuyAppFormAdminFormSubmemberDeleteRes.md)
 - [XbuyAppFormAdminFormSubmemberEditReq](docs/XbuyAppFormAdminFormSubmemberEditReq.md)
 - [XbuyAppFormAdminFormSubmemberEditRes](docs/XbuyAppFormAdminFormSubmemberEditRes.md)
 - [XbuyAppFormAdminFormSubmemberEmailUniqueReq](docs/XbuyAppFormAdminFormSubmemberEmailUniqueReq.md)
 - [XbuyAppFormAdminFormSubmemberEmailUniqueRes](docs/XbuyAppFormAdminFormSubmemberEmailUniqueRes.md)
 - [XbuyAppFormAdminFormSubmemberListReq](docs/XbuyAppFormAdminFormSubmemberListReq.md)
 - [XbuyAppFormAdminFormSubmemberListRes](docs/XbuyAppFormAdminFormSubmemberListRes.md)
 - [XbuyAppFormAdminFormSubmemberMobileUniqueReq](docs/XbuyAppFormAdminFormSubmemberMobileUniqueReq.md)
 - [XbuyAppFormAdminFormSubmemberMobileUniqueRes](docs/XbuyAppFormAdminFormSubmemberMobileUniqueRes.md)
 - [XbuyAppFormAdminFormSubmemberNameUniqueReq](docs/XbuyAppFormAdminFormSubmemberNameUniqueReq.md)
 - [XbuyAppFormAdminFormSubmemberNameUniqueRes](docs/XbuyAppFormAdminFormSubmemberNameUniqueRes.md)
 - [XbuyAppFormAdminFormSubmemberResetPwdReq](docs/XbuyAppFormAdminFormSubmemberResetPwdReq.md)
 - [XbuyAppFormAdminFormSubmemberResetPwdRes](docs/XbuyAppFormAdminFormSubmemberResetPwdRes.md)
 - [XbuyAppFormAdminFormSubmemberViewReq](docs/XbuyAppFormAdminFormSubmemberViewReq.md)
 - [XbuyAppFormAdminFormSubmemberViewRes](docs/XbuyAppFormAdminFormSubmemberViewRes.md)
 - [XbuyAppFormAdminFormSynAddReq](docs/XbuyAppFormAdminFormSynAddReq.md)
 - [XbuyAppFormAdminFormSynAddRes](docs/XbuyAppFormAdminFormSynAddRes.md)
 - [XbuyAppFormAdminFormSynDeleteReq](docs/XbuyAppFormAdminFormSynDeleteReq.md)
 - [XbuyAppFormAdminFormSynDeleteRes](docs/XbuyAppFormAdminFormSynDeleteRes.md)
 - [XbuyAppFormAdminFormSynGoodsReq](docs/XbuyAppFormAdminFormSynGoodsReq.md)
 - [XbuyAppFormAdminFormSynGoodsRes](docs/XbuyAppFormAdminFormSynGoodsRes.md)
 - [XbuyAppFormAdminFormSynListReq](docs/XbuyAppFormAdminFormSynListReq.md)
 - [XbuyAppFormAdminFormSynListRes](docs/XbuyAppFormAdminFormSynListRes.md)
 - [XbuyAppFormAdminFormTransAddReq](docs/XbuyAppFormAdminFormTransAddReq.md)
 - [XbuyAppFormAdminFormTransAddRes](docs/XbuyAppFormAdminFormTransAddRes.md)
 - [XbuyAppFormAdminFormUnBindMemberGradeReq](docs/XbuyAppFormAdminFormUnBindMemberGradeReq.md)
 - [XbuyAppFormAdminFormUnBindMemberGradeRes](docs/XbuyAppFormAdminFormUnBindMemberGradeRes.md)
 - [XbuyAppFormAdminFormUnBindMemberToShopReq](docs/XbuyAppFormAdminFormUnBindMemberToShopReq.md)
 - [XbuyAppFormAdminFormUnBindMemberToShopRes](docs/XbuyAppFormAdminFormUnBindMemberToShopRes.md)
 - [XbuyAppFormAdminFormUnBoundMemberTagReq](docs/XbuyAppFormAdminFormUnBoundMemberTagReq.md)
 - [XbuyAppFormAdminFormUnBoundMemberTagRes](docs/XbuyAppFormAdminFormUnBoundMemberTagRes.md)
 - [XbuyAppFormAdminFormUpdateBindMemberGradeReq](docs/XbuyAppFormAdminFormUpdateBindMemberGradeReq.md)
 - [XbuyAppFormAdminFormUpdateBindMemberGradeRes](docs/XbuyAppFormAdminFormUpdateBindMemberGradeRes.md)
 - [XbuyAppFormAdminFormUpdateBindShopReq](docs/XbuyAppFormAdminFormUpdateBindShopReq.md)
 - [XbuyAppFormAdminFormUpdateBindShopRes](docs/XbuyAppFormAdminFormUpdateBindShopRes.md)
 - [XbuyAppFormAdminFormUpdateMemberGradeReq](docs/XbuyAppFormAdminFormUpdateMemberGradeReq.md)
 - [XbuyAppFormAdminFormUpdateMemberGradeRes](docs/XbuyAppFormAdminFormUpdateMemberGradeRes.md)
 - [XbuyAppFormAdminFormUpdateMemberGroupRelationReq](docs/XbuyAppFormAdminFormUpdateMemberGroupRelationReq.md)
 - [XbuyAppFormAdminFormUpdateMemberGroupRelationRes](docs/XbuyAppFormAdminFormUpdateMemberGroupRelationRes.md)
 - [XbuyAppFormAdminFormUpdateMemberGroupReq](docs/XbuyAppFormAdminFormUpdateMemberGroupReq.md)
 - [XbuyAppFormAdminFormUpdateMemberGroupRes](docs/XbuyAppFormAdminFormUpdateMemberGroupRes.md)
 - [XbuyAppFormAdminFormUpdateMemberTagNameReq](docs/XbuyAppFormAdminFormUpdateMemberTagNameReq.md)
 - [XbuyAppFormAdminFormUpdateMemberTagNameRes](docs/XbuyAppFormAdminFormUpdateMemberTagNameRes.md)
 - [XbuyAppFormAdminFormUpdateReceiverInfoReq](docs/XbuyAppFormAdminFormUpdateReceiverInfoReq.md)
 - [XbuyAppFormAdminFormUpdateReceiverInfoRes](docs/XbuyAppFormAdminFormUpdateReceiverInfoRes.md)
 - [XbuyAppFormAdminFormUpdateSkuStockReq](docs/XbuyAppFormAdminFormUpdateSkuStockReq.md)
 - [XbuyAppFormAdminFormUpdateSkuStockRes](docs/XbuyAppFormAdminFormUpdateSkuStockRes.md)
 - [XbuyAppFormAdminFormUpdateSkuTaxReq](docs/XbuyAppFormAdminFormUpdateSkuTaxReq.md)
 - [XbuyAppFormAdminFormUpdateSkuTaxRes](docs/XbuyAppFormAdminFormUpdateSkuTaxRes.md)
 - [XbuyAppFormAdminFormUploadFileReq](docs/XbuyAppFormAdminFormUploadFileReq.md)
 - [XbuyAppFormAdminFormUploadFileRes](docs/XbuyAppFormAdminFormUploadFileRes.md)
 - [XbuyAppFormAdminFormUserBuyReq](docs/XbuyAppFormAdminFormUserBuyReq.md)
 - [XbuyAppFormAdminFormUserBuyRes](docs/XbuyAppFormAdminFormUserBuyRes.md)
 - [XbuyAppFormAdminFormVerifyEmailApplyReq](docs/XbuyAppFormAdminFormVerifyEmailApplyReq.md)
 - [XbuyAppFormAdminFormVerifyEmailApplyRes](docs/XbuyAppFormAdminFormVerifyEmailApplyRes.md)
 - [XbuyAppFormAdminFormWareHouseCreateReq](docs/XbuyAppFormAdminFormWareHouseCreateReq.md)
 - [XbuyAppFormAdminFormWareHouseCreateRes](docs/XbuyAppFormAdminFormWareHouseCreateRes.md)
 - [XbuyAppFormAdminFormWareHouseDeleteReq](docs/XbuyAppFormAdminFormWareHouseDeleteReq.md)
 - [XbuyAppFormAdminFormWareHouseDeleteRes](docs/XbuyAppFormAdminFormWareHouseDeleteRes.md)
 - [XbuyAppFormAdminFormWareHouseReadReq](docs/XbuyAppFormAdminFormWareHouseReadReq.md)
 - [XbuyAppFormAdminFormWareHouseReadRes](docs/XbuyAppFormAdminFormWareHouseReadRes.md)
 - [XbuyAppFormAdminFormWareHouseStockCreateReq](docs/XbuyAppFormAdminFormWareHouseStockCreateReq.md)
 - [XbuyAppFormAdminFormWareHouseStockCreateRes](docs/XbuyAppFormAdminFormWareHouseStockCreateRes.md)
 - [XbuyAppFormAdminFormWareHouseStockPagingReq](docs/XbuyAppFormAdminFormWareHouseStockPagingReq.md)
 - [XbuyAppFormAdminFormWareHouseStockPagingRes](docs/XbuyAppFormAdminFormWareHouseStockPagingRes.md)
 - [XbuyAppFormAdminFormWareHouseStockReadReq](docs/XbuyAppFormAdminFormWareHouseStockReadReq.md)
 - [XbuyAppFormAdminFormWareHouseStockReadRes](docs/XbuyAppFormAdminFormWareHouseStockReadRes.md)
 - [XbuyAppFormAdminFormWareHouseStockUpdateReq](docs/XbuyAppFormAdminFormWareHouseStockUpdateReq.md)
 - [XbuyAppFormAdminFormWareHouseStockUpdateRes](docs/XbuyAppFormAdminFormWareHouseStockUpdateRes.md)
 - [XbuyAppFormAdminFormWareHouseUpdateReq](docs/XbuyAppFormAdminFormWareHouseUpdateReq.md)
 - [XbuyAppFormAdminFormWareHouseUpdateRes](docs/XbuyAppFormAdminFormWareHouseUpdateRes.md)
 - [XbuyAppFormAdminFormWeightSetting](docs/XbuyAppFormAdminFormWeightSetting.md)
 - [XbuyAppFormAdminFormZendCountryListReq](docs/XbuyAppFormAdminFormZendCountryListReq.md)
 - [XbuyAppFormAdminFormZendCountryListRes](docs/XbuyAppFormAdminFormZendCountryListRes.md)
 - [XbuyAppFormApiFormBaseLangReq](docs/XbuyAppFormApiFormBaseLangReq.md)
 - [XbuyAppFormApiFormBaseLangRes](docs/XbuyAppFormApiFormBaseLangRes.md)
 - [XbuyAppFormApiFormExportReq](docs/XbuyAppFormApiFormExportReq.md)
 - [XbuyAppFormApiFormExportRes](docs/XbuyAppFormApiFormExportRes.md)
 - [XbuyAppFormApiFormIpLocationReq](docs/XbuyAppFormApiFormIpLocationReq.md)
 - [XbuyAppFormApiFormIpLocationRes](docs/XbuyAppFormApiFormIpLocationRes.md)
 - [XbuyAppFormApiFormLogClearReq](docs/XbuyAppFormApiFormLogClearReq.md)
 - [XbuyAppFormApiFormLogClearRes](docs/XbuyAppFormApiFormLogClearRes.md)
 - [XbuyAppFormApiFormLogExportReq](docs/XbuyAppFormApiFormLogExportReq.md)
 - [XbuyAppFormApiFormLogExportRes](docs/XbuyAppFormApiFormLogExportRes.md)
 - [XbuyAppFormApiFormLogListReq](docs/XbuyAppFormApiFormLogListReq.md)
 - [XbuyAppFormApiFormLogListRes](docs/XbuyAppFormApiFormLogListRes.md)
 - [XbuyAppFormApiFormLoginCheckReq](docs/XbuyAppFormApiFormLoginCheckReq.md)
 - [XbuyAppFormApiFormLoginCheckRes](docs/XbuyAppFormApiFormLoginCheckRes.md)
 - [XbuyAppFormApiFormLoginLogoutReq](docs/XbuyAppFormApiFormLoginLogoutReq.md)
 - [XbuyAppFormApiFormLoginLogoutRes](docs/XbuyAppFormApiFormLoginLogoutRes.md)
 - [XbuyAppFormApiFormLoginReq](docs/XbuyAppFormApiFormLoginReq.md)
 - [XbuyAppFormApiFormLoginRes](docs/XbuyAppFormApiFormLoginRes.md)
 - [XbuyAppFormApiFormMemberProfileReq](docs/XbuyAppFormApiFormMemberProfileReq.md)
 - [XbuyAppFormApiFormMemberProfileRes](docs/XbuyAppFormApiFormMemberProfileRes.md)
 - [XbuyAppFormApiFormReceiveBarCodeRateReq](docs/XbuyAppFormApiFormReceiveBarCodeRateReq.md)
 - [XbuyAppFormApiFormReceiveBarCodeRateRes](docs/XbuyAppFormApiFormReceiveBarCodeRateRes.md)
 - [XbuyAppFormApiFormReceiveExamineReq](docs/XbuyAppFormApiFormReceiveExamineReq.md)
 - [XbuyAppFormApiFormReceiveExamineRes](docs/XbuyAppFormApiFormReceiveExamineRes.md)
 - [XbuyAppFormApiFormReceiveOrderShipReq](docs/XbuyAppFormApiFormReceiveOrderShipReq.md)
 - [XbuyAppFormApiFormReceiveOrderShipRes](docs/XbuyAppFormApiFormReceiveOrderShipRes.md)
 - [XbuyAppFormDocFormGetAdminApiDocReq](docs/XbuyAppFormDocFormGetAdminApiDocReq.md)
 - [XbuyAppFormDocFormGetAdminApiDocRes](docs/XbuyAppFormDocFormGetAdminApiDocRes.md)
 - [XbuyAppFormDocFormGetMallApiDocReq](docs/XbuyAppFormDocFormGetMallApiDocReq.md)
 - [XbuyAppFormDocFormGetMallApiDocRes](docs/XbuyAppFormDocFormGetMallApiDocRes.md)
 - [XbuyAppFormDocFormGetOpenApiDocReq](docs/XbuyAppFormDocFormGetOpenApiDocReq.md)
 - [XbuyAppFormDocFormGetOpenApiDocRes](docs/XbuyAppFormDocFormGetOpenApiDocRes.md)
 - [XbuyAppFormInputAdminDeptListModel](docs/XbuyAppFormInputAdminDeptListModel.md)
 - [XbuyAppFormInputAdminDeptListTreeDept](docs/XbuyAppFormInputAdminDeptListTreeDept.md)
 - [XbuyAppFormInputAdminDeptListTreeModel](docs/XbuyAppFormInputAdminDeptListTreeModel.md)
 - [XbuyAppFormInputAdminDeptTreeDept](docs/XbuyAppFormInputAdminDeptTreeDept.md)
 - [XbuyAppFormInputAdminDeptViewModel](docs/XbuyAppFormInputAdminDeptViewModel.md)
 - [XbuyAppFormInputAdminMemberListModel](docs/XbuyAppFormInputAdminMemberListModel.md)
 - [XbuyAppFormInputAdminMemberViewModel](docs/XbuyAppFormInputAdminMemberViewModel.md)
 - [XbuyAppFormInputAdminNoticeListModel](docs/XbuyAppFormInputAdminNoticeListModel.md)
 - [XbuyAppFormInputAdminPostListModel](docs/XbuyAppFormInputAdminPostListModel.md)
 - [XbuyAppFormInputAdminRoleListModel](docs/XbuyAppFormInputAdminRoleListModel.md)
 - [XbuyAppFormInputAdminSubmemberListModel](docs/XbuyAppFormInputAdminSubmemberListModel.md)
 - [XbuyAppFormInputAdminSubmemberViewModel](docs/XbuyAppFormInputAdminSubmemberViewModel.md)
 - [XbuyAppFormInputLogListModel](docs/XbuyAppFormInputLogListModel.md)
 - [XbuyAppFormInputSysConfigListModel](docs/XbuyAppFormInputSysConfigListModel.md)
 - [XbuyAppFormMallFormAddressAddReq](docs/XbuyAppFormMallFormAddressAddReq.md)
 - [XbuyAppFormMallFormAddressAddRes](docs/XbuyAppFormMallFormAddressAddRes.md)
 - [XbuyAppFormMallFormAddressDeleteReq](docs/XbuyAppFormMallFormAddressDeleteReq.md)
 - [XbuyAppFormMallFormAddressDeleteRes](docs/XbuyAppFormMallFormAddressDeleteRes.md)
 - [XbuyAppFormMallFormAddressListReq](docs/XbuyAppFormMallFormAddressListReq.md)
 - [XbuyAppFormMallFormAddressListRes](docs/XbuyAppFormMallFormAddressListRes.md)
 - [XbuyAppFormMallFormAfterSaleSkuInfo](docs/XbuyAppFormMallFormAfterSaleSkuInfo.md)
 - [XbuyAppFormMallFormAfterSalesApplyReq](docs/XbuyAppFormMallFormAfterSalesApplyReq.md)
 - [XbuyAppFormMallFormAfterSalesApplyRes](docs/XbuyAppFormMallFormAfterSalesApplyRes.md)
 - [XbuyAppFormMallFormArticleCatListReq](docs/XbuyAppFormMallFormArticleCatListReq.md)
 - [XbuyAppFormMallFormArticleCatListRes](docs/XbuyAppFormMallFormArticleCatListRes.md)
 - [XbuyAppFormMallFormArticleGetReq](docs/XbuyAppFormMallFormArticleGetReq.md)
 - [XbuyAppFormMallFormArticleGetRes](docs/XbuyAppFormMallFormArticleGetRes.md)
 - [XbuyAppFormMallFormArticleListReq](docs/XbuyAppFormMallFormArticleListReq.md)
 - [XbuyAppFormMallFormArticleListRes](docs/XbuyAppFormMallFormArticleListRes.md)
 - [XbuyAppFormMallFormArticleNewListReq](docs/XbuyAppFormMallFormArticleNewListReq.md)
 - [XbuyAppFormMallFormArticleNewListRes](docs/XbuyAppFormMallFormArticleNewListRes.md)
 - [XbuyAppFormMallFormBenefitCodeOperateReq](docs/XbuyAppFormMallFormBenefitCodeOperateReq.md)
 - [XbuyAppFormMallFormBenefitCodeOperateRes](docs/XbuyAppFormMallFormBenefitCodeOperateRes.md)
 - [XbuyAppFormMallFormCartAddReq](docs/XbuyAppFormMallFormCartAddReq.md)
 - [XbuyAppFormMallFormCartAddRes](docs/XbuyAppFormMallFormCartAddRes.md)
 - [XbuyAppFormMallFormCartClearReq](docs/XbuyAppFormMallFormCartClearReq.md)
 - [XbuyAppFormMallFormCartClearRes](docs/XbuyAppFormMallFormCartClearRes.md)
 - [XbuyAppFormMallFormCartDetail](docs/XbuyAppFormMallFormCartDetail.md)
 - [XbuyAppFormMallFormCartGetReq](docs/XbuyAppFormMallFormCartGetReq.md)
 - [XbuyAppFormMallFormCartGetRes](docs/XbuyAppFormMallFormCartGetRes.md)
 - [XbuyAppFormMallFormCartNumReq](docs/XbuyAppFormMallFormCartNumReq.md)
 - [XbuyAppFormMallFormCartNumRes](docs/XbuyAppFormMallFormCartNumRes.md)
 - [XbuyAppFormMallFormCartRemoveReq](docs/XbuyAppFormMallFormCartRemoveReq.md)
 - [XbuyAppFormMallFormCartRemoveRes](docs/XbuyAppFormMallFormCartRemoveRes.md)
 - [XbuyAppFormMallFormCartUpdateReq](docs/XbuyAppFormMallFormCartUpdateReq.md)
 - [XbuyAppFormMallFormCartUpdateRes](docs/XbuyAppFormMallFormCartUpdateRes.md)
 - [XbuyAppFormMallFormCatInfoReq](docs/XbuyAppFormMallFormCatInfoReq.md)
 - [XbuyAppFormMallFormCatInfoRes](docs/XbuyAppFormMallFormCatInfoRes.md)
 - [XbuyAppFormMallFormCatListReq](docs/XbuyAppFormMallFormCatListReq.md)
 - [XbuyAppFormMallFormCatListRes](docs/XbuyAppFormMallFormCatListRes.md)
 - [XbuyAppFormMallFormCountryListReq](docs/XbuyAppFormMallFormCountryListReq.md)
 - [XbuyAppFormMallFormCountryListRes](docs/XbuyAppFormMallFormCountryListRes.md)
 - [XbuyAppFormMallFormExpress](docs/XbuyAppFormMallFormExpress.md)
 - [XbuyAppFormMallFormFavAddReq](docs/XbuyAppFormMallFormFavAddReq.md)
 - [XbuyAppFormMallFormFavAddRes](docs/XbuyAppFormMallFormFavAddRes.md)
 - [XbuyAppFormMallFormFavDeleteReq](docs/XbuyAppFormMallFormFavDeleteReq.md)
 - [XbuyAppFormMallFormFavDeleteRes](docs/XbuyAppFormMallFormFavDeleteRes.md)
 - [XbuyAppFormMallFormFavListReq](docs/XbuyAppFormMallFormFavListReq.md)
 - [XbuyAppFormMallFormFavListRes](docs/XbuyAppFormMallFormFavListRes.md)
 - [XbuyAppFormMallFormFavNumReq](docs/XbuyAppFormMallFormFavNumReq.md)
 - [XbuyAppFormMallFormFavNumRes](docs/XbuyAppFormMallFormFavNumRes.md)
 - [XbuyAppFormMallFormGeneratePaymentCredentialReq](docs/XbuyAppFormMallFormGeneratePaymentCredentialReq.md)
 - [XbuyAppFormMallFormGeneratePaymentCredentialRes](docs/XbuyAppFormMallFormGeneratePaymentCredentialRes.md)
 - [XbuyAppFormMallFormGetMallAllShopConfigsReq](docs/XbuyAppFormMallFormGetMallAllShopConfigsReq.md)
 - [XbuyAppFormMallFormGetMallAllShopConfigsRes](docs/XbuyAppFormMallFormGetMallAllShopConfigsRes.md)
 - [XbuyAppFormMallFormGetMallShopConfigsReq](docs/XbuyAppFormMallFormGetMallShopConfigsReq.md)
 - [XbuyAppFormMallFormGetMallShopConfigsRes](docs/XbuyAppFormMallFormGetMallShopConfigsRes.md)
 - [XbuyAppFormMallFormGoodInfoReq](docs/XbuyAppFormMallFormGoodInfoReq.md)
 - [XbuyAppFormMallFormGoodInfoRes](docs/XbuyAppFormMallFormGoodInfoRes.md)
 - [XbuyAppFormMallFormGoodListReq](docs/XbuyAppFormMallFormGoodListReq.md)
 - [XbuyAppFormMallFormGoodListRes](docs/XbuyAppFormMallFormGoodListRes.md)
 - [XbuyAppFormMallFormGoodSameCatReq](docs/XbuyAppFormMallFormGoodSameCatReq.md)
 - [XbuyAppFormMallFormGoodSameCatRes](docs/XbuyAppFormMallFormGoodSameCatRes.md)
 - [XbuyAppFormMallFormInsuredReq](docs/XbuyAppFormMallFormInsuredReq.md)
 - [XbuyAppFormMallFormInsuredRes](docs/XbuyAppFormMallFormInsuredRes.md)
 - [XbuyAppFormMallFormLoginCaptchaReq](docs/XbuyAppFormMallFormLoginCaptchaReq.md)
 - [XbuyAppFormMallFormLoginCaptchaRes](docs/XbuyAppFormMallFormLoginCaptchaRes.md)
 - [XbuyAppFormMallFormMallConfigAllReq](docs/XbuyAppFormMallFormMallConfigAllReq.md)
 - [XbuyAppFormMallFormMallConfigAllRes](docs/XbuyAppFormMallFormMallConfigAllRes.md)
 - [XbuyAppFormMallFormMallConfigListReq](docs/XbuyAppFormMallFormMallConfigListReq.md)
 - [XbuyAppFormMallFormMallConfigListRes](docs/XbuyAppFormMallFormMallConfigListRes.md)
 - [XbuyAppFormMallFormMallMemberChangePassWordReq](docs/XbuyAppFormMallFormMallMemberChangePassWordReq.md)
 - [XbuyAppFormMallFormMallMemberChangePassWordRes](docs/XbuyAppFormMallFormMallMemberChangePassWordRes.md)
 - [XbuyAppFormMallFormMallMemberLoginReq](docs/XbuyAppFormMallFormMallMemberLoginReq.md)
 - [XbuyAppFormMallFormMallMemberLoginRes](docs/XbuyAppFormMallFormMallMemberLoginRes.md)
 - [XbuyAppFormMallFormMallMemberLogoutReq](docs/XbuyAppFormMallFormMallMemberLogoutReq.md)
 - [XbuyAppFormMallFormMallMemberLogoutRes](docs/XbuyAppFormMallFormMallMemberLogoutRes.md)
 - [XbuyAppFormMallFormMallMemberProfileReq](docs/XbuyAppFormMallFormMallMemberProfileReq.md)
 - [XbuyAppFormMallFormMallMemberProfileRes](docs/XbuyAppFormMallFormMallMemberProfileRes.md)
 - [XbuyAppFormMallFormMallMemberRegisterReq](docs/XbuyAppFormMallFormMallMemberRegisterReq.md)
 - [XbuyAppFormMallFormMallMemberRegisterRes](docs/XbuyAppFormMallFormMallMemberRegisterRes.md)
 - [XbuyAppFormMallFormMallMemberUpdateReq](docs/XbuyAppFormMallFormMallMemberUpdateReq.md)
 - [XbuyAppFormMallFormMallMemberUpdateRes](docs/XbuyAppFormMallFormMallMemberUpdateRes.md)
 - [XbuyAppFormMallFormMallReFreshTokenReq](docs/XbuyAppFormMallFormMallReFreshTokenReq.md)
 - [XbuyAppFormMallFormMallReFreshTokenRes](docs/XbuyAppFormMallFormMallReFreshTokenRes.md)
 - [XbuyAppFormMallFormOrderCreateReq](docs/XbuyAppFormMallFormOrderCreateReq.md)
 - [XbuyAppFormMallFormOrderCreateRes](docs/XbuyAppFormMallFormOrderCreateRes.md)
 - [XbuyAppFormMallFormOrderListReq](docs/XbuyAppFormMallFormOrderListReq.md)
 - [XbuyAppFormMallFormOrderListRes](docs/XbuyAppFormMallFormOrderListRes.md)
 - [XbuyAppFormMallFormOrderPayReq](docs/XbuyAppFormMallFormOrderPayReq.md)
 - [XbuyAppFormMallFormOrderPayRes](docs/XbuyAppFormMallFormOrderPayRes.md)
 - [XbuyAppFormMallFormOrderWaybillExpressReq](docs/XbuyAppFormMallFormOrderWaybillExpressReq.md)
 - [XbuyAppFormMallFormOrderWaybillExpressRes](docs/XbuyAppFormMallFormOrderWaybillExpressRes.md)
 - [XbuyAppFormMallFormPayWays](docs/XbuyAppFormMallFormPayWays.md)
 - [XbuyAppFormMallFormPaymentMethodGetReq](docs/XbuyAppFormMallFormPaymentMethodGetReq.md)
 - [XbuyAppFormMallFormPaymentMethodGetRes](docs/XbuyAppFormMallFormPaymentMethodGetRes.md)
 - [XbuyAppFormMallFormPreorderReq](docs/XbuyAppFormMallFormPreorderReq.md)
 - [XbuyAppFormMallFormPreorderRes](docs/XbuyAppFormMallFormPreorderRes.md)
 - [XbuyAppFormMallFormPrepayOrderListReq](docs/XbuyAppFormMallFormPrepayOrderListReq.md)
 - [XbuyAppFormMallFormPrepayOrderListRes](docs/XbuyAppFormMallFormPrepayOrderListRes.md)
 - [XbuyAppFormMallFormQueryOrderReq](docs/XbuyAppFormMallFormQueryOrderReq.md)
 - [XbuyAppFormMallFormQueryOrderRes](docs/XbuyAppFormMallFormQueryOrderRes.md)
 - [XbuyAppFormMallFormResetPasswordApplyReq](docs/XbuyAppFormMallFormResetPasswordApplyReq.md)
 - [XbuyAppFormMallFormResetPasswordApplyRes](docs/XbuyAppFormMallFormResetPasswordApplyRes.md)
 - [XbuyAppFormMallFormResetPasswordConfirmReq](docs/XbuyAppFormMallFormResetPasswordConfirmReq.md)
 - [XbuyAppFormMallFormResetPasswordConfirmRes](docs/XbuyAppFormMallFormResetPasswordConfirmRes.md)
 - [XbuyAppFormMallFormReviewAddReq](docs/XbuyAppFormMallFormReviewAddReq.md)
 - [XbuyAppFormMallFormReviewAddRes](docs/XbuyAppFormMallFormReviewAddRes.md)
 - [XbuyAppFormMallFormReviewListReq](docs/XbuyAppFormMallFormReviewListReq.md)
 - [XbuyAppFormMallFormReviewListRes](docs/XbuyAppFormMallFormReviewListRes.md)
 - [XbuyAppFormMallFormSalesOrderCreateReq](docs/XbuyAppFormMallFormSalesOrderCreateReq.md)
 - [XbuyAppFormMallFormSalesOrderCreateRes](docs/XbuyAppFormMallFormSalesOrderCreateRes.md)
 - [XbuyAppFormMallFormSalesOrderListReq](docs/XbuyAppFormMallFormSalesOrderListReq.md)
 - [XbuyAppFormMallFormSalesOrderListRes](docs/XbuyAppFormMallFormSalesOrderListRes.md)
 - [XbuyAppFormMallFormSceneInfo](docs/XbuyAppFormMallFormSceneInfo.md)
 - [XbuyAppFormMallFormSendEmailReq](docs/XbuyAppFormMallFormSendEmailReq.md)
 - [XbuyAppFormMallFormSendEmailRes](docs/XbuyAppFormMallFormSendEmailRes.md)
 - [XbuyAppFormMallFormShopListReq](docs/XbuyAppFormMallFormShopListReq.md)
 - [XbuyAppFormMallFormShopListRes](docs/XbuyAppFormMallFormShopListRes.md)
 - [XbuyAppFormMallFormShopViewListReq](docs/XbuyAppFormMallFormShopViewListReq.md)
 - [XbuyAppFormMallFormShopViewListRes](docs/XbuyAppFormMallFormShopViewListRes.md)
 - [XbuyAppFormMallFormShoppingList](docs/XbuyAppFormMallFormShoppingList.md)
 - [XbuyAppFormMallFormSiteMapReq](docs/XbuyAppFormMallFormSiteMapReq.md)
 - [XbuyAppFormMallFormSiteMapRes](docs/XbuyAppFormMallFormSiteMapRes.md)
 - [XbuyAppFormMallFormStandardCatReq](docs/XbuyAppFormMallFormStandardCatReq.md)
 - [XbuyAppFormMallFormStandardCatRes](docs/XbuyAppFormMallFormStandardCatRes.md)
 - [XbuyAppFormMallFormSubscribesAddReq](docs/XbuyAppFormMallFormSubscribesAddReq.md)
 - [XbuyAppFormMallFormSubscribesAddRes](docs/XbuyAppFormMallFormSubscribesAddRes.md)
 - [XbuyAppFormMallFormSubscribesDeleteReq](docs/XbuyAppFormMallFormSubscribesDeleteReq.md)
 - [XbuyAppFormMallFormSubscribesDeleteRes](docs/XbuyAppFormMallFormSubscribesDeleteRes.md)
 - [XbuyAppFormMallFormSubscribesListReq](docs/XbuyAppFormMallFormSubscribesListReq.md)
 - [XbuyAppFormMallFormSubscribesListRes](docs/XbuyAppFormMallFormSubscribesListRes.md)
 - [XbuyAppFormMallFormTransAddReq](docs/XbuyAppFormMallFormTransAddReq.md)
 - [XbuyAppFormMallFormTransAddRes](docs/XbuyAppFormMallFormTransAddRes.md)
 - [XbuyAppFormMallFormUpdatePayerInfoReq](docs/XbuyAppFormMallFormUpdatePayerInfoReq.md)
 - [XbuyAppFormMallFormUpdatePayerInfoRes](docs/XbuyAppFormMallFormUpdatePayerInfoRes.md)
 - [XbuyAppFormMallFormVerifyEmailApplyReq](docs/XbuyAppFormMallFormVerifyEmailApplyReq.md)
 - [XbuyAppFormMallFormVerifyEmailApplyRes](docs/XbuyAppFormMallFormVerifyEmailApplyRes.md)
 - [XbuyAppFormMallFormVerifyEmailConfirmReq](docs/XbuyAppFormMallFormVerifyEmailConfirmReq.md)
 - [XbuyAppFormMallFormVerifyEmailConfirmRes](docs/XbuyAppFormMallFormVerifyEmailConfirmRes.md)
 - [XbuyAppFormMallFormVerifyEmailReq](docs/XbuyAppFormMallFormVerifyEmailReq.md)
 - [XbuyAppFormMallFormVerifyEmailRes](docs/XbuyAppFormMallFormVerifyEmailRes.md)
 - [XbuyAppFormMallFormVerifyPhoneApplyReq](docs/XbuyAppFormMallFormVerifyPhoneApplyReq.md)
 - [XbuyAppFormMallFormVerifyPhoneApplyRes](docs/XbuyAppFormMallFormVerifyPhoneApplyRes.md)
 - [XbuyAppFormNotifyFormPaymentNotifyReq](docs/XbuyAppFormNotifyFormPaymentNotifyReq.md)
 - [XbuyAppFormNotifyFormPaymentNotifyRes](docs/XbuyAppFormNotifyFormPaymentNotifyRes.md)
 - [XbuyAppFormNotifyFormRefundNotifyReq](docs/XbuyAppFormNotifyFormRefundNotifyReq.md)
 - [XbuyAppFormNotifyFormRefundNotifyRes](docs/XbuyAppFormNotifyFormRefundNotifyRes.md)
 - [XbuyAppFormThirdFormFacebookOauthCallbackReq](docs/XbuyAppFormThirdFormFacebookOauthCallbackReq.md)
 - [XbuyAppFormThirdFormFacebookOauthCallbackRes](docs/XbuyAppFormThirdFormFacebookOauthCallbackRes.md)
 - [XbuyAppFormThirdFormFacebookOauthReq](docs/XbuyAppFormThirdFormFacebookOauthReq.md)
 - [XbuyAppFormThirdFormFacebookOauthRes](docs/XbuyAppFormThirdFormFacebookOauthRes.md)
 - [XbuyAppFormThirdFormGoogleOauthCallbackReq](docs/XbuyAppFormThirdFormGoogleOauthCallbackReq.md)
 - [XbuyAppFormThirdFormGoogleOauthCallbackRes](docs/XbuyAppFormThirdFormGoogleOauthCallbackRes.md)
 - [XbuyAppFormThirdFormGoogleOauthReq](docs/XbuyAppFormThirdFormGoogleOauthReq.md)
 - [XbuyAppFormThirdFormGoogleOauthRes](docs/XbuyAppFormThirdFormGoogleOauthRes.md)
 - [XbuyAppFormThirdFormMicrosoftOauthCallbackReq](docs/XbuyAppFormThirdFormMicrosoftOauthCallbackReq.md)
 - [XbuyAppFormThirdFormMicrosoftOauthCallbackRes](docs/XbuyAppFormThirdFormMicrosoftOauthCallbackRes.md)
 - [XbuyAppFormThirdFormMicrosoftOauthReq](docs/XbuyAppFormThirdFormMicrosoftOauthReq.md)
 - [XbuyAppFormThirdFormMicrosoftOauthRes](docs/XbuyAppFormThirdFormMicrosoftOauthRes.md)
 - [XbuyAppFormThirdFormTwitterOauthCallbackReq](docs/XbuyAppFormThirdFormTwitterOauthCallbackReq.md)
 - [XbuyAppFormThirdFormTwitterOauthCallbackRes](docs/XbuyAppFormThirdFormTwitterOauthCallbackRes.md)
 - [XbuyAppFormThirdFormTwitterOauthReq](docs/XbuyAppFormThirdFormTwitterOauthReq.md)
 - [XbuyAppFormThirdFormTwitterOauthRes](docs/XbuyAppFormThirdFormTwitterOauthRes.md)
 - [XbuyAppFormThirdFormWechatOauthCallbackReq](docs/XbuyAppFormThirdFormWechatOauthCallbackReq.md)
 - [XbuyAppFormThirdFormWechatOauthCallbackRes](docs/XbuyAppFormThirdFormWechatOauthCallbackRes.md)
 - [XbuyAppFormThirdFormWechatOauthCheckReq](docs/XbuyAppFormThirdFormWechatOauthCheckReq.md)
 - [XbuyAppFormThirdFormWechatOauthCheckRes](docs/XbuyAppFormThirdFormWechatOauthCheckRes.md)
 - [XbuyAppFormThirdFormWechatOauthReq](docs/XbuyAppFormThirdFormWechatOauthReq.md)
 - [XbuyAppFormThirdFormWechatOauthRes](docs/XbuyAppFormThirdFormWechatOauthRes.md)
 - [XbuyAppModelEntityAdminExpressCompany](docs/XbuyAppModelEntityAdminExpressCompany.md)
 - [XbuyAppModelEntityAdminMenu](docs/XbuyAppModelEntityAdminMenu.md)
 - [XbuyAppModelEntityAdminNotice](docs/XbuyAppModelEntityAdminNotice.md)
 - [XbuyAppModelEntityAdminRoleShop](docs/XbuyAppModelEntityAdminRoleShop.md)
 - [XbuyAppModelEntityAdminSubmemberShop](docs/XbuyAppModelEntityAdminSubmemberShop.md)
 - [XbuyAppModelEntityAdminSynConfig](docs/XbuyAppModelEntityAdminSynConfig.md)
 - [XbuyAppModelEntityMallAddress](docs/XbuyAppModelEntityMallAddress.md)
 - [XbuyAppModelEntityMallAftersaleItems](docs/XbuyAppModelEntityMallAftersaleItems.md)
 - [XbuyAppModelEntityMallArticlesSub](docs/XbuyAppModelEntityMallArticlesSub.md)
 - [XbuyAppModelEntityMallBenefitCondition](docs/XbuyAppModelEntityMallBenefitCondition.md)
 - [XbuyAppModelEntityMallBenefitEnvCustomInfo](docs/XbuyAppModelEntityMallBenefitEnvCustomInfo.md)
 - [XbuyAppModelEntityMallBillingAddress](docs/XbuyAppModelEntityMallBillingAddress.md)
 - [XbuyAppModelEntityMallCarrySubMode](docs/XbuyAppModelEntityMallCarrySubMode.md)
 - [XbuyAppModelEntityMallCats](docs/XbuyAppModelEntityMallCats.md)
 - [XbuyAppModelEntityMallConfigs](docs/XbuyAppModelEntityMallConfigs.md)
 - [XbuyAppModelEntityMallCountry](docs/XbuyAppModelEntityMallCountry.md)
 - [XbuyAppModelEntityMallEmailLog](docs/XbuyAppModelEntityMallEmailLog.md)
 - [XbuyAppModelEntityMallFareTemplate](docs/XbuyAppModelEntityMallFareTemplate.md)
 - [XbuyAppModelEntityMallGoodInfos](docs/XbuyAppModelEntityMallGoodInfos.md)
 - [XbuyAppModelEntityMallGoods](docs/XbuyAppModelEntityMallGoods.md)
 - [XbuyAppModelEntityMallMemberGrades](docs/XbuyAppModelEntityMallMemberGrades.md)
 - [XbuyAppModelEntityMallMemberGroups](docs/XbuyAppModelEntityMallMemberGroups.md)
 - [XbuyAppModelEntityMallMemberTags](docs/XbuyAppModelEntityMallMemberTags.md)
 - [XbuyAppModelEntityMallMembers](docs/XbuyAppModelEntityMallMembers.md)
 - [XbuyAppModelEntityMallOrderInsured](docs/XbuyAppModelEntityMallOrderInsured.md)
 - [XbuyAppModelEntityMallOrderItems](docs/XbuyAppModelEntityMallOrderItems.md)
 - [XbuyAppModelEntityMallOrderWaybillExpress](docs/XbuyAppModelEntityMallOrderWaybillExpress.md)
 - [XbuyAppModelEntityMallPushOrderLogs](docs/XbuyAppModelEntityMallPushOrderLogs.md)
 - [XbuyAppModelEntityMallScats](docs/XbuyAppModelEntityMallScats.md)
 - [XbuyAppModelEntityMallSeoGoogle](docs/XbuyAppModelEntityMallSeoGoogle.md)
 - [XbuyAppModelEntityMallService](docs/XbuyAppModelEntityMallService.md)
 - [XbuyAppModelEntityMallShopView](docs/XbuyAppModelEntityMallShopView.md)
 - [XbuyAppModelEntityMallShops](docs/XbuyAppModelEntityMallShops.md)
 - [XbuyAppModelEntityMallSkus](docs/XbuyAppModelEntityMallSkus.md)
 - [XbuyAppModelEntityMallSubscribes](docs/XbuyAppModelEntityMallSubscribes.md)
 - [XbuyAppModelEntityMallUploads](docs/XbuyAppModelEntityMallUploads.md)
 - [XbuyAppModelEntityMallWarehouse](docs/XbuyAppModelEntityMallWarehouse.md)
 - [XbuyAppModelEntityMallWarehouseStock](docs/XbuyAppModelEntityMallWarehouseStock.md)
 - [XbuyAppModelEntitySysDictData](docs/XbuyAppModelEntitySysDictData.md)
 - [XbuyAppModelEntitySysDictType](docs/XbuyAppModelEntitySysDictType.md)
 - [XbuyAppModelEntitySysResourceGroups](docs/XbuyAppModelEntitySysResourceGroups.md)
 - [XbuyAppModelEntitySysResources](docs/XbuyAppModelEntitySysResources.md)
 - [XbuyAppModelIdentity](docs/XbuyAppModelIdentity.md)
 - [XbuyAppModelLabelTreeMenu](docs/XbuyAppModelLabelTreeMenu.md)
 - [XbuyAppModelTreeMenu](docs/XbuyAppModelTreeMenu.md)
 - [XbuyAppServiceMsgserviceMailserviceMessage](docs/XbuyAppServiceMsgserviceMailserviceMessage.md)
 - [XbuyEtcAllowCustomMade](docs/XbuyEtcAllowCustomMade.md)
 - [XbuyEtcAllowMerchant](docs/XbuyEtcAllowMerchant.md)
 - [XbuyEtcAllowStores](docs/XbuyEtcAllowStores.md)
 - [XbuyEtcCopyrightLink](docs/XbuyEtcCopyrightLink.md)
 - [XbuyEtcDepartmentAuthority](docs/XbuyEtcDepartmentAuthority.md)
 - [XbuyEtcDomainBinding](docs/XbuyEtcDomainBinding.md)
 - [XbuyEtcOtherInfo](docs/XbuyEtcOtherInfo.md)
 - [XbuyEtcSaasModel](docs/XbuyEtcSaasModel.md)
 - [XbuyEtcTechnicalSupport](docs/XbuyEtcTechnicalSupport.md)
 - [XbuyXerrorError](docs/XbuyXerrorError.md)

## Documentation For Authorization
 Endpoints do not require authorization.


## Author


