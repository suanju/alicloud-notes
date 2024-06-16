package backend

import (
	_const "alicloud-notes/backend/const"
	"alicloud-notes/backend/types"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"

	"github.com/tickstep/aliyunpan-api/aliyunpan_web"
)

//lint:ignore U1000 全局实例
var panInstance *aliyunpan_web.WebPanClient

// GenerateQrcode 阿里云盘获取登录二维码
func (a *App) GenerateQrcode() (resp *types.GenerateQrcodeResp, err error) {
	clientResp, err := resty.New().R().
		SetQueryParams(map[string]string{
			"appName":     "aliyun_drive",
			"appEntrance": "web",
			"isMobile":    "false",
		}).
		Get(_const.AliyunDriveGenerate)
	if err != err {
		return nil, err
	}
	httpResp := new(types.GenerateResponse)
	err = json.Unmarshal(clientResp.Body(), httpResp)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err %s", clientResp.Body())
	}
	//生成二维码
	qrImg, err := qrcode.Encode(httpResp.Content.Data.CodeContent, qrcode.Medium, 256)
	if err != nil {
		return nil, fmt.Errorf("failed to generate qrcode %s", err)
	}
	base64Str := "data:image/png;base64," + base64.StdEncoding.EncodeToString(qrImg)
	return &types.GenerateQrcodeResp{
		Qrcode: base64Str,
		T:      httpResp.Content.Data.T,
		Ck:     httpResp.Content.Data.Ck,
	}, nil
}

// QrcodeState 获取二维码状态
func (a *App) QrcodeState(req *types.QrcodeStateReq) (resp *types.QrcodeStateResp, err error) {
	params := url.Values{}
	params.Set("t", strconv.FormatInt(req.T, 10))
	params.Set("ck", req.Ck)
	params.Set("appName", "aliyun_drive")
	params.Set("appEntrance", "web")
	body := params.Encode()
	clientResp, err := resty.New().R().
		SetHeader(_const.HeaderContentType, _const.FormURLEncodedContentType).
		SetBody(body).
		Post(_const.AliyunQrcodeQuery)
	if err != err {
		return nil, err
	}
	httpResp := new(types.QrcodeStateResponse)
	err = json.Unmarshal(clientResp.Body(), httpResp)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal err %s", clientResp.Body())
	}
	if httpResp.Content.Data.QrCodeStatus == _const.QrCodeStatusConfirmed {
		//成功扫码授权
		decodedBytes, _ := base64.StdEncoding.DecodeString(httpResp.Content.Data.BizExt)
		bizExt := new(types.QrcodeStateBizExtData)
		_ = json.Unmarshal(decodedBytes, &bizExt)
		return &types.QrcodeStateResp{
			UserName:     bizExt.PdsLoginResult.UserName,
			UserID:       bizExt.PdsLoginResult.UserID,
			TokenType:    bizExt.PdsLoginResult.TokenType,
			RefreshToken: bizExt.PdsLoginResult.RefreshToken,
		}, nil
	}
	return nil, fmt.Errorf(_const.QrCodeStatus[httpResp.Content.Data.QrCodeStatus])
}

// CreatInstance 创建实例
func (a *App) CreatInstance(req *types.CreatInstanceReq) (resp *types.CreatInstanceResp, err error) {
	webToken, apierror := aliyunpan_web.GetAccessTokenFromRefreshToken(req.RefreshToken)
	if apierror != nil {
		return nil, fmt.Errorf("获取AcccesToken失败")
	}
	appConfig := aliyunpan_web.AppConfig{
		DeviceId: uuid.New().String(),
	}
	panInstance := aliyunpan_web.NewWebPanClient(*webToken, aliyunpan_web.AppLoginToken{}, appConfig, aliyunpan_web.SessionConfig{
		DeviceName: "云笔记",
		ModelName:  "云笔记",
	})
	// create session
	panInstance.CreateSession(&aliyunpan_web.CreateSessionParam{
		DeviceName: "云笔记",
		ModelName:  "云笔记",
	})
	userInfo, apierror := panInstance.GetUserInfo()
	if apierror != nil {
		return nil, fmt.Errorf("获取用户信息失败")
	}
	return userInfo, nil
}
