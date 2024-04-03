package apis

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"net/url"
	"os"
)

func Login(sid string) error {
	tk := getToken()
	resp, err := AuthReq().SetQueryParams(map[string]string{
		"sid":   sid,
		"_json": "true",
	}).Get(AuthURI("/pass/serviceLogin"))
	if err != nil {
		return err
	}
	if !resp.IsSuccess() {
		return fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	ret := ServiceLoginRet{}
	if err = decode(resp.Body(), &ret); err != nil {
		return err
	}
	if ret.Code != 0 {
		data := map[string]string{
			"_json":    "true",
			"qs":       ret.Qs,
			"sid":      ret.Sid,
			"_sign":    ret.Sign,
			"callback": ret.Callback,
			"user":     os.Getenv("MI_USER"),
			"hash":     GetMD5Hash(os.Getenv("MI_PASS")),
		}
		resp, err = AuthReq().SetFormData(data).Post(AuthURI("/pass/serviceLoginAuth2"))
		if err != nil {
			return err
		}
		if !resp.IsSuccess() {
			return fmt.Errorf("resp err:%d", resp.StatusCode())
		}
		if err = decode(resp.Body(), &ret); err != nil {
			return err
		}
	}
	tk.UserId = ret.UserId
	tk.PassToken = ret.PassToken
	serviceToken, err := securityTokenService(ret.Location, ret.Nonce, ret.Ssecurity)
	if err != nil {
		return err
	}
	tk.SetSubToken(sid, ret.Ssecurity, serviceToken)
	return tk.Sync()
}

func securityTokenService(location string, nonce int64, ssecurity string) (string, error) {
	nsec := fmt.Sprintf("nonce=%d&%s", nonce, ssecurity)
	hash := sha1.Sum([]byte(nsec))
	encoded := base64.StdEncoding.EncodeToString(hash[:])
	u, _ := url.Parse(location)
	query := u.Query()
	query.Set("clientSign", encoded)
	u.RawQuery = query.Encode()
	resp, err := AuthReq().Get(u.String())
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

type ServiceLoginRet struct {
	Code      int    `json:"code"`
	Desc      string `json:"desc"`
	Qs        string `json:"qs"`
	Sid       string `json:"sid"`
	Sign      string `json:"_sign"`
	Callback  string `json:"callback"`
	PassToken string `json:"passToken"`
	UserId    int    `json:"userId"`
	Location  string `json:"location"`
	Nonce     int64  `json:"nonce"`
	Ssecurity string `json:"ssecurity"`
}
