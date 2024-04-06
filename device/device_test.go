package device

import (
	"fmt"
	"testing"
)

func TestDeviceList(t *testing.T) {
	ret, _ := List(false, 0)
	for _, v := range ret.List {
		fmt.Println(v.Model)
	}
}

func TestAction(t *testing.T) {
	fmt.Println(Action(ActionDetail{
		Did:  "{did}",
		Siid: 5,
		Aiid: 4,
		In:   []any{"今天天气", 1}, // 0 : silent-execution
	}))

}
