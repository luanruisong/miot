package auth

import (
	"fmt"
	"github.com/luanruisong/miot/consts"
	"testing"
)

func TestLogin(t *testing.T) {
	fmt.Println(Login(consts.SID_XIAOMIIO, "user", "pass"), "====")
}
