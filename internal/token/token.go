package token

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"github.com/luanruisong/miot/internal/utils"
	"os"
	"path"
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

func (st *SubToken) Singer() Singer {
	return NewSinger(st.Ssecurity)
}

func (tk *Token) IsLogin() bool {
	return tk.UserId > 0 && len(tk.PassToken) > 0
}

func (tk *Token) IsSubTokenLogin(sid string) bool {
	return tk.GetSubToken(sid) != nil
}

func (tk *Token) GetSubToken(sid string) *SubToken {
	return tk.Tks[sid]
}

func (tk *Token) SetSubToken(sid, ssecurity, serviceToken string) *Token {
	tk.Tks[sid] = &SubToken{
		ServiceToken: serviceToken,
		Ssecurity:    ssecurity,
	}
	return tk
}

func (tk *Token) Sync() error {
	b, _ := jsoniter.Marshal(tk)
	return os.WriteFile(filePath(), b, os.ModePerm)
}

func GetToken() *Token {
	if _tks == nil {
		_tks = &Token{
			DeviceId:  utils.SRand(16),
			UserId:    0,
			PassToken: "",
			Tks:       make(map[string]*SubToken),
		}
	}
	return _tks
}

func filePath() string {
	home := os.Getenv(utils.EnvHome)
	if len(home) == 0 {
		home = path.Join(os.Getenv("HOME"), "/.miot/")
	}
	if ok, _ := utils.PathExists(home); !ok {
		if err := os.MkdirAll(home, os.ModePerm); err != nil {
			panic(err)
		}
	}
	return path.Join(home, "tks.json")
}

func CheckLogin(sid string) error {
	if tk := GetToken(); !tk.IsLogin() || tk.GetSubToken(sid) == nil {
		return NoLoginErr
	}
	return nil
}

var (
	_tks       *Token
	NoLoginErr = errors.New("not login")
)

func init() {
	fp := filePath()
	ok, _ := utils.PathExists(fp)
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
