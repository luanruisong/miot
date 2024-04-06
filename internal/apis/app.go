package apis

import (
	"errors"
	"fmt"
	"github.com/luanruisong/miot/internal/token"
	"github.com/luanruisong/miot/internal/utils"
)

func SignAppPost[T any](sid, uri string, data any) (T, error) {
	var res T
	if err := token.CheckLogin(sid); err != nil {
		return res, err
	}
	singer := token.GetToken().GetSubToken(sid).Singer()
	resp, err := AppReq(sid).SetFormData(singer.SignData(uri, data)).Post(AppURI(uri))
	if err != nil {
		return res, err
	}
	if !resp.IsSuccess() {
		return res, fmt.Errorf("resp err:%d", resp.StatusCode())
	}
	return Decode[T](resp.Body())
}

type AppRet[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  T      `json:"result"`
}

func Decode[T any](data []byte) (T, error) {
	ret, err := utils.Decode[AppRet[T]](data)
	if ret.Code != 0 {
		var zero T
		return zero, errors.New(ret.Message)
	}
	return ret.Result, err
}
