package auth

import (
	"fmt"
	"github.com/luanruisong/miot/internal/utils"
	"testing"
)

func TestLogin(t *testing.T) {
	fmt.Println(Login(utils.SID_XIAOMIIO, "user", "pass"), "====")
}
