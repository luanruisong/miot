package apis

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/luanruisong/miot/consts"
	"os"
	"path"
	"strings"
)

type SubToken struct {
	ServiceToken string
	Ssecurity    string
}

type Token struct {
	DeviceId  string
	UserId    int
	PassToken string
	Tks       map[string]*SubToken
}

func (tk *Token) IsLogin() bool {
	return tk.UserId > 0 && len(tk.PassToken) > 0
}

func (tk *Token) GetSubToken(sid string) *SubToken {
	return tk.Tks[sid]
}

func (tk *Token) SetSubToken(sid, ssecurity, serviceToken string) {
	tk.Tks[sid] = &SubToken{
		ServiceToken: serviceToken,
		Ssecurity:    ssecurity,
	}
}

func (tk *Token) Sync() error {
	b, _ := jsoniter.Marshal(tk)
	return os.WriteFile(filePath(), b, os.ModePerm)
}

func getToken() *Token {
	if _tks == nil {
		_tks = &Token{
			DeviceId:  strings.ToUpper(RandStr(16)),
			UserId:    0,
			PassToken: "",
			Tks:       make(map[string]*SubToken),
		}
	}
	return _tks
}

func filePath() string {
	home := os.Getenv(consts.EnvHome)
	if len(home) == 0 {
		home = path.Join(os.Getenv("HOME"), "/.miot/")
	}
	if ok, _ := PathExists(home); !ok {
		if err := os.MkdirAll(home, os.ModePerm); err != nil {
			panic(err)
		}
	}
	return path.Join(home, "tks.json")
}

var (
	_tks *Token
)

func init() {
	fp := filePath()
	fmt.Println("load file", fp)
	ok, _ := PathExists(fp)
	if ok {
		b, err := os.ReadFile(fp)
		if err != nil {
			panic(err)
		}
		if len(b) > 0 {
			_tks = &Token{}
			err = jsoniter.Unmarshal(b, &_tks)
			if err != nil {
				panic(err)
			}
		}
	}

}
