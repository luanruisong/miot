package device

import (
	"fmt"
	"github.com/luanruisong/miot/apis"
	"github.com/luanruisong/miot/auth"
	"github.com/luanruisong/miot/token"
	"github.com/luanruisong/miot/utils"
)

func List(getVirtualModel bool, getHuamiDevices int) ([]DeviceInfo, error) {
	tk := token.GetToken()
	if !tk.IsLogin() || !tk.IsSubTokenLogin(utils.SID_XIAOMIIO) {
		if err := auth.Login(utils.SID_XIAOMIIO, utils.GetUser(), utils.GetPass()); err != nil {
			panic(err)
		}
	}
	subToken := tk.GetSubToken(utils.SID_XIAOMIIO)
	uri := "/home/device_list"
	data := map[string]interface{}{
		"getVirtualModel": getVirtualModel,
		"getHuamiDevices": getHuamiDevices,
	}
	singer := subToken.Singer()
	resp, err := apis.ApiReq(utils.SID_XIAOMIIO).SetFormData(singer.SignData(uri, data)).Post(apis.AppURI(uri))
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
	if !tk.IsLogin() || !tk.IsSubTokenLogin(utils.SID_XIAOMIIO) {
		if err := auth.Login(utils.SID_XIAOMIIO, utils.GetUser(), utils.GetPass()); err != nil {
			panic(err)
		}
	}
	subToken := tk.GetSubToken(utils.SID_XIAOMIIO)
	uri := "/miotspec/action"
	data := map[string]any{
		"params": action,
	}
	singer := subToken.Singer()
	resp, err := apis.ApiReq(utils.SID_XIAOMIIO).SetFormData(singer.SignData(uri, data)).Post(apis.AppURI(uri))
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
