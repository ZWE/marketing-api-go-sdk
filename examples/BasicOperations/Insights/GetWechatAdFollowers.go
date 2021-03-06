/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.3
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/antihax/optional"
	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
	"github.com/tencentad/marketing-api-go-sdk/pkg/model"
)

type WechatAdFollowersGetExample struct {
	TAds                     *ads.SDKClient
	AccessToken              string
	TimeRange                model.TimeRange
	WechatAdFollowersGetOpts *api.WechatAdFollowersGetOpts
}

func (e *WechatAdFollowersGetExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.TimeRange = model.TimeRange{
		StartTime: int64(0),
		EndTime:   int64(0),
	}
	e.WechatAdFollowersGetOpts = &api.WechatAdFollowersGetOpts{

		Fields: optional.NewInterface([]string{"openid", "wechat_adgroup_id", "wechat_campaign_id", "wechat_account_id", "wechat_agency_id", "subscribe_time", "position_type"}),
	}
}

func (e *WechatAdFollowersGetExample) RunExample() (model.WechatAdFollowersGetResponseData, http.Header, error) {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	return tads.WechatAdFollowers().Get(ctx, e.TimeRange, e.WechatAdFollowersGetOpts)
}

func main() {
	e := &WechatAdFollowersGetExample{}
	e.Init()
	response, headers, err := e.RunExample()
	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Response error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Headers:", headers)
}
