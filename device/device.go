package device

import (
	"fmt"
	"github.com/luanruisong/miot/auth"
	"github.com/luanruisong/miot/internal/apis"
	"github.com/luanruisong/miot/internal/token"
	utils2 "github.com/luanruisong/miot/internal/utils"
)

func List(getVirtualModel bool, getHuamiDevices int) ([]DeviceInfo, error) {
	tk := token.GetToken()
	if !tk.IsLogin() || !tk.IsSubTokenLogin(utils2.SID_XIAOMIIO) {
		if err := auth.Login(utils2.SID_XIAOMIIO, utils2.GetUser(), utils2.GetPass()); err != nil {
			panic(err)
		}
	}
	subToken := tk.GetSubToken(utils2.SID_XIAOMIIO)
	uri := "/home/device_list"
	data := map[string]interface{}{
		"getVirtualModel": getVirtualModel,
		"getHuamiDevices": getHuamiDevices,
	}
	singer := subToken.Singer()
	resp, err := apis.ApiReq(utils2.SID_XIAOMIIO).SetFormData(singer.SignData(uri, data)).Post(apis.AppURI(uri))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	ret, err := Decode[DeviceListResult](resp.Body())
	if err != nil {
		return nil, err
	}
	return ret.List, nil
}

func Action(action *ActionDetail) (*ActionResult, error) {
	tk := token.GetToken()
	if !tk.IsLogin() || !tk.IsSubTokenLogin(utils2.SID_XIAOMIIO) {
		if err := auth.Login(utils2.SID_XIAOMIIO, utils2.GetUser(), utils2.GetPass()); err != nil {
			panic(err)
		}
	}
	subToken := tk.GetSubToken(utils2.SID_XIAOMIIO)
	uri := "/miotspec/action"
	data := map[string]any{
		"params": action,
	}
	singer := subToken.Singer()
	resp, err := apis.ApiReq(utils2.SID_XIAOMIIO).SetFormData(singer.SignData(uri, data)).Post(apis.AppURI(uri))
	if err != nil {
		return nil, err
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	ret, err := Decode[ActionResult](resp.Body())
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
