package wechat

import (
	"encoding/base64"
	"errors"
)

// WXAPI_ENT 企业号接口
const (
	WXAPI_ENT        = "https://qyapi.weixin.qq.com/cgi-bin/"
	WXAPI_TOKEN_ENT  = WXAPI_ENT + "gettoken?corpid=%s&corpsecret=%s"
	WXAPI_MSG_ENT    = WXAPI_ENT + "message/send?access_token="
	WXAPI_IP_ENT     = WXAPI_ENT + "getcallbackip?access_token="
	WXAPI_UPLOAD_ENT = WXAPI_ENT + "media/upload?access_token=%s&type=%s"
)

// SetEnt 初始化企业号，设置token,corpid,secrat,aesKey
func SetEnt(tk, id, sec, key string) (err error) {
	if len(key) != 43 {
		return errors.New("非法的AesKey")
	}
	safeMode = true
	entMode = true
	token, appId, secret = tk, id, sec
	msgUrl = WXAPI_MSG_ENT
	uploadUrl = WXAPI_UPLOAD_ENT
	aesKey, err = base64.StdEncoding.DecodeString(key + "=")
	if err != nil {
		return err
	}
	FetchAccessToken(WXAPI_TOKEN_ENT)
	return nil
}