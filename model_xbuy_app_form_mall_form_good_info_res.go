/*
 * xbuy
 *
 * xbuy大黄蜂买手项目
 *
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// 数据集
type XbuyAppFormMallFormGoodInfoRes struct {
	// 编号
	GoodId int64 `json:"good_id,omitempty"`
	// 来源平台
	FromPlatform string `json:"from_platform,omitempty"`
	// 来源平台商品id
	FromGoodId int64 `json:"from_good_id,omitempty"`
	// 所属门店
	ShopId int64 `json:"shop_id,omitempty"`
	// 所属类目
	CatId int64 `json:"cat_id,omitempty"`
	// 标题
	Title string `json:"title,omitempty"`
	// 描述
	Desc string `json:"desc,omitempty"`
	// 备注
	Note string `json:"note,omitempty"`
	// 商品视频
	VideoUrl string `json:"video_url,omitempty"`
	// 商品主图
	ImageUrl string `json:"image_url,omitempty"`
	// 商品主图
	ImageUrl1 string `json:"image_url_1,omitempty"`
	// 商品主图
	ImageUrl2 string `json:"image_url_2,omitempty"`
	// 商品主图
	ImageUrl3 string `json:"image_url_3,omitempty"`
	// 商品主图
	ImageUrl4 string `json:"image_url_4,omitempty"`
	// 是否上架,默认1(上架)2下架
	OnShelves int32 `json:"on_shelves,omitempty"`
	// 创建时间
	CreateTime string `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime string `json:"update_time,omitempty"`
	ViewId int64 `json:"view_id,omitempty"`
	Slug string `json:"slug,omitempty"`
	Cid int64 `json:"cid,omitempty"`
	Rate float64 `json:"rate,omitempty"`
	SubTitle string `json:"sub_title,omitempty"`
	SalesNum int64 `json:"sales_num,omitempty"`
	Weight float64 `json:"weight,omitempty"`
	IsGift int32 `json:"is_gift,omitempty"`
	Ai int32 `json:"ai,omitempty"`
	PriceMin float64 `json:"price_min,omitempty"`
	PriceMax float64 `json:"price_max,omitempty"`
	Skus []XbuyAppDaoSkusContent `json:"skus,omitempty"`
	SkuProps []XbuyAppDaoOutputSkuInfo `json:"sku_props,omitempty"`
	GoodInfos []XbuyAppDaoGoodInfosDetail `json:"good_infos,omitempty"`
}
