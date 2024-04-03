package apis

import (
	"crypto/md5"
	"encoding/hex"
	jsoniter "github.com/json-iterator/go"
	"math/rand"
	"net/url"
	"os"
	"path"
	"strings"
)

func RandStr(length int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetMD5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	hexString := hex.EncodeToString(hash[:])
	return strings.ToUpper(hexString)
}

func decode(b []byte, p any) error {
	b = b[11:]
	return jsoniter.Unmarshal(b, p)
}

func URI(host, uri string) string {
	u, _ := url.Parse(host)
	u.Path = path.Join(u.Path, uri)
	return u.String()
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
