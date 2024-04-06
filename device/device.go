package device

import (
	"github.com/luanruisong/miot/consts"
	"github.com/luanruisong/miot/internal/apis"
)

func List(getVirtualModel bool, getHuamiDevices int) (DeviceListResult, error) {
	return apis.SignAppPost[DeviceListResult](consts.SID_XIAOMIIO, "/home/device_list", map[string]any{
		"getVirtualModel": getVirtualModel,
		"getHuamiDevices": getHuamiDevices,
	})
}

func Action(action ActionDetail) (ActionResult, error) {
	return apis.SignAppPost[ActionResult](consts.SID_XIAOMIIO, "/miotspec/action", map[string]any{
		"params": action,
	})
}
