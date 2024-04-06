package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"math/rand"
	"os"
	"strings"
)

func SRand(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = rune(rand.Intn(26) + 65)
		fmt.Println(b[i], string(b))
	}
	return string(b)
}

func GetMD5Hash(password string) string {
	hash := md5.Sum([]byte(password))
	hexString := hex.EncodeToString(hash[:])
	return strings.ToUpper(hexString)
}

func Decode[T any](b []byte) (T, error) {
	p := new(T)
	b, _ = bytes.CutPrefix(b, []byte("&&&START&&&")) //&&&START&&&
	return *p, jsoniter.Unmarshal(b, &p)
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

func CheckEnv() error {
	envs := []string{EnvUser, EnvPass}
	for _, v := range envs {
		if len(os.Getenv(v)) == 0 {
			return fmt.Errorf("env %s not found", v)
		}
	}
	return nil
}

func GetUser() string {
	return os.Getenv(EnvUser)
}

func GetPass() string {
	return os.Getenv(EnvPass)
}
