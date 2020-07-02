/*
 * Marketing API
 *
 * Marketing API
 *
 * API version: 1.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"encoding/json"
	"fmt"

	"github.com/tencentad/marketing-api-go-sdk/pkg/ads"
	"github.com/tencentad/marketing-api-go-sdk/pkg/api"
	"github.com/tencentad/marketing-api-go-sdk/pkg/config"
	"github.com/tencentad/marketing-api-go-sdk/pkg/errors"
)

type OauthAuthorizeExample struct {
	TAds               *ads.SDKClient
	AccessToken        string
	ClientId           int64
	RedirectUri        string
	OauthAuthorizeOpts *api.OauthAuthorizeOpts
}

func (e *OauthAuthorizeExample) Init() {
	e.AccessToken = "YOUR ACCESS TOKEN"
	e.TAds = ads.Init(&config.SDKConfig{
		AccessToken: e.AccessToken,
		IsDebug:     true,
	})
	e.ClientId = 789
	e.RedirectUri = "redirectUri_example"
	e.OauthAuthorizeOpts = &api.OauthAuthorizeOpts{}
}

func (e *OauthAuthorizeExample) RunExample() {
	tads := e.TAds
	// change ctx as needed
	ctx := *tads.Ctx
	response, httpResponse, err := tads.Oauth().Authorize(ctx, e.ClientId, e.RedirectUri, e.OauthAuthorizeOpts)

	if err != nil {
		if resErr, ok := err.(errors.ResponseError); ok {
			errStr, _ := json.Marshal(resErr)
			fmt.Println("Resopnse error:", string(errStr))
		} else {
			fmt.Println("Error:", err)
		}
	}
	fmt.Println("Response data:", response)
	fmt.Println("Http response:", httpResponse)
}

func main() {
	e := &OauthAuthorizeExample{}
	e.Init()
	e.RunExample()
}