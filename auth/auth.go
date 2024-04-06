package auth

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/luanruisong/miot/consts"
	"github.com/luanruisong/miot/internal/apis"
	"github.com/luanruisong/miot/internal/token"
	"github.com/luanruisong/miot/internal/utils"
	"net/url"
)

func Login(sid, user, pass string) error {
	resp, err := apis.AuthReq().SetQueryParams(map[string]string{
		"sid":   sid,
		"_json": "true",
	}).Get(apis.AuthURI("/pass/serviceLogin"))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	ret, err := utils.Decode[ServiceLoginRet](resp.Body())
	if err != nil {
		return err
	}
	if ret.Code != 0 {
		ret, err = serverLoginAuth(ret, user, pass)
		if err != nil {
			return err
		}
	}
	tk := token.GetToken()
	tk.UserId = ret.UserId
	tk.PassToken = ret.PassToken
	serviceToken, err := generateServiceToken(ret.Location, ret.Nonce, ret.Ssecurity)
	if err != nil {
		return err
	}
	return tk.SetSubToken(sid, ret.Ssecurity, serviceToken).Sync()
}

func serverLoginAuth(req ServiceLoginRet, user, pass string) (ServiceLoginRet, error) {
	data := map[string]string{
		"_json":    "true",
		"qs":       req.Qs,
		"sid":      req.Sid,
		"_sign":    req.Sign,
		"callback": req.Callback,
		"user":     user,
		"hash":     utils.GetMD5Hash(pass),
	}
	resp, err := apis.AuthReq().SetFormData(data).Post(apis.AuthURI("/pass/serviceLoginAuth2"))
	if err != nil {
		return ServiceLoginRet{}, err
	}
	if !resp.IsSuccess() {
		return ServiceLoginRet{}, fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	return utils.Decode[ServiceLoginRet](resp.Body())
}

func generateServiceToken(location string, nonce int64, ssecurity string) (string, error) {
	nsec := fmt.Sprintf("nonce=%d&%s", nonce, ssecurity)
	hash := sha1.Sum([]byte(nsec))
	encoded := base64.StdEncoding.EncodeToString(hash[:])
	u, _ := url.Parse(location)
	query := u.Query()
	query.Set("clientSign", encoded)
	u.RawQuery = query.Encode()
	resp, err := apis.AuthReq().Get(u.String())
	if err != nil {
		return "", err
	}
	if !resp.IsSuccess() {
		return "", fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	for _, v := range resp.Cookies() {
		if v.Name == "serviceToken" {
			return v.Value, nil
		}
	}
	return "", errors.New("can not find service Token")
}

func AutoLogin(sid string) error {
	if err := token.CheckLogin(sid); err != nil {
		if err = consts.CheckEnv(); err != nil {
			return err
		}
		return Login(sid, consts.GetUser(), consts.GetPass())
	}
	return nil
}
