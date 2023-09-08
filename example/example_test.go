package example

import (
	sdk "github.com/gcfguo/bee-hole-api-sdk-golang"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/glog"
	"golang.org/x/net/context"
	"os"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	GetToken()
	PublicCats()
}

// GetToken 获取token
func GetToken() {
	authCode := "2TxqktZJHmmDk4YbeMwJEtoHqcE" //授权码 联系管理员获取
	client := sdk.NewAPIClient(&sdk.Configuration{
		BasePath: "http://localhost:8299", //接口地址
	})
	timeStamp := time.Now().String()
	authResp, _, err := client.LoginApi.AdminLoginSubAuthPost(context.Background(), sdk.XbuyAppFormAdminFormLoginTokenReq{
		Username: "api",                                        //子商户 用户名
		Password: "abc123",                                     //子商户 密码
		Sign:     gmd5.MustEncryptString(timeStamp + authCode), //签名
		Time:     timeStamp,                                    //时间戳
		ShopId:   1,                                            //店铺id
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

	client := sdk.NewAPIClient(&sdk.Configuration{
		BasePath: "",
		DefaultHeader: map[string]string{
			"Authorization": "Bearer " + string(token),
		},
	})
	resp, _, err := client.StandardCatApi.AdminStandardCatPost(context.Background(), sdk.XbuyAppFormAdminFormStandardCatReq{})
	if err != nil {
		glog.Error(context.Background(), "err", err)
		return
	}
	glog.Info(context.Background(), "resp", resp)
}
