package main

import (
	"client/swagger"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/glog"
	"golang.org/x/net/context"
	"os"
	"time"
)

func main() {
	GetToken()
	PublicCats()
}

// GetToken 获取token
func GetToken() {
	authCode := "" //授权码 联系管理员获取
	client := swagger.NewAPIClient(&swagger.Configuration{
		BasePath: "", //接口地址
	})
	timeStamp := time.Now().String()
	authResp, _, err := client.LoginApi.AdminLoginSubAuthPost(context.Background(), swagger.XbuyAppFormAdminFormLoginTokenReq{
		Username: "",                                           //子商户 用户名
		Password: "",                                           //子商户 密码
		Sign:     gmd5.MustEncryptString(timeStamp + authCode), //签名
		Time:     timeStamp,                                    //时间戳
		ShopId:   0,                                            //店铺id
	})
	if err != nil {
		glog.Error(context.Background(), "err", err)
		return
	}
	glog.Info(context.Background(), "authResp", authResp)
	err = os.WriteFile("token.txt", []byte(authResp.Data.Token), 0666)
	if err != nil {
		glog.Error(context.Background(), "err", err)
		return
	}
}

// PublicCats 公共分类查询
func PublicCats() {
	token, err := os.ReadFile("token.txt")
	if err != nil {
		glog.Error(context.Background(), "err", err)
		return
	}

	client := swagger.NewAPIClient(&swagger.Configuration{
		BasePath: "",
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + string(token),
		},
	})
	resp, _, err := client.StandardCatApi.AdminStandardCatPost(context.Background(), swagger.XbuyAppFormAdminFormStandardCatReq{})
	if err != nil {
		glog.Error(context.Background(), "err", err)
		return
	}
	glog.Info(context.Background(), "resp", resp)
}
