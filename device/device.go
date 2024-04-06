package device

import (
	"github.com/luanruisong/miot/internal/apis"
	utils2 "github.com/luanruisong/miot/internal/utils"
)

func List(getVirtualModel bool, getHuamiDevices int) (DeviceListResult, error) {
	return apis.SignAppPost[DeviceListResult](utils2.SID_XIAOMIIO, "/home/device_list", map[string]any{
		"getVirtualModel": getVirtualModel,
		"getHuamiDevices": getHuamiDevices,
	})
}

func Action(action ActionDetail) (ActionResult, error) {
	return apis.SignAppPost[ActionResult](utils2.SID_XIAOMIIO, "/miotspec/action", map[string]any{
		"params": action,
	})
}
