package apis

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/luanruisong/miot/token"
	"github.com/luanruisong/miot/utils"
	"net/http"
	"net/url"
	"path"
)

var (
	req *resty.Client
)

func init() {
	req = resty.New()
}

func AuthReq() *resty.Request {
	app := token.GetToken()
	cookies := append([]*http.Cookie{}, &http.Cookie{
		Name:  "sdkVersion",
		Value: "3.9",
	}, &http.Cookie{
		Name:  "deviceId",
		Value: app.DeviceId,
	})
	if app.IsLogin() {
		cookies = append(cookies, &http.Cookie{
			Name:  "userId",
			Value: fmt.Sprintf("%d", app.UserId),
		}, &http.Cookie{
			Name:  "passToken",
			Value: app.PassToken,
		})
	}
	header := map[string]string{
		"User-Agent": "APP/com.xiaomi.mihome APPV/6.0.103 iosPassportSDK/3.9.0 iOS/14.4 miHSTS",
	}
	return req.R().SetCookies(cookies).SetHeaders(header)
}

func ApiReq(sid string) *resty.Request {
	app := token.GetToken()
	if !app.IsLogin() {
		panic(errors.New("need login"))
	}
	cookies := append([]*http.Cookie{}, &http.Cookie{
		Name:  "PassportDeviceId",
		Value: app.DeviceId,
	}, &http.Cookie{
		Name:  "userId",
		Value: fmt.Sprintf("%d", app.UserId),
	}, &http.Cookie{
		Name:  "serviceToken",
		Value: app.GetSubToken(sid).ServiceToken,
	})
	header := map[string]string{
		"User-Agent":                 "APP/com.xiaomi.mihome APPV/6.0.103 iosPassportSDK/3.9.0 iOS/14.4 miHSTS",
		"x-xiaomi-protocal-flag-cli": "PROTOCAL-HTTP2",
	}
	return req.R().SetCookies(cookies).SetHeaders(header)
}

func AuthURI(uri string) string {
	return _uri(utils.AuthHost, uri)
}

func AppURI(uri string) string {
	return _uri(utils.AppHost, uri)
}

func _uri(host, uri string) string {
	u, _ := url.Parse(host)
	u.Path = path.Join(u.Path, uri)
	return u.String()
}
