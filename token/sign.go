package token

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	jsoniter "github.com/json-iterator/go"
	"strings"
	"time"
)

func NewSinger(ssecurity string) *Singer {
	return &Singer{ssecurity}
}

type Singer struct {
	ssecurity string
}

func (s *Singer) SignData(uri string, data any) map[string]string {
	var jsonData string
	switch data.(type) {
	case string:
		jsonData = data.(string)
	default:
		jsonData, _ = jsoniter.MarshalToString(data)
	}
	nonce := s.GenerateNonce()
	snonce := s.SignNonce(s.ssecurity, nonce)
	msg := strings.Join([]string{uri, snonce, nonce, "data=" + jsonData}, "&")
	// 解码snonce
	key, err := base64.StdEncoding.DecodeString(snonce)
	if err != nil {
		panic(err)
	}
	// 创建HMAC实例
	h := hmac.New(sha256.New, key)
	// 写入消息
	h.Write([]byte(msg))
	// 计算HMAC
	macSum := h.Sum(nil)
	sign := base64.StdEncoding.EncodeToString(macSum)
	return map[string]string{
		"_nonce":    nonce,
		"data":      jsonData,
		"signature": sign,
	}
}

func (s *Singer) SignNonce(ssecurity, nonce string) string {
	// 解码ssecurity和nonce
	ssecurityBytes, _ := base64.StdEncoding.DecodeString(ssecurity)
	nonceBytes, _ := base64.StdEncoding.DecodeString(nonce)

	// 创建一个新的SHA256哈希
	h := sha256.New()

	// 更新哈希值
	h.Write(ssecurityBytes)
	h.Write(nonceBytes)

	// 计算哈希值
	hash := h.Sum(nil)

	// 将哈希值编码为Base64字符串
	encoded := base64.StdEncoding.EncodeToString(hash)

	return encoded
}

func (s *Singer) GenerateNonce() string {
	// 生成8字节的随机数
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// 获取当前时间，精确到分钟，并转换为4字节的大端序字节
	now := time.Now()
	minute := int64(now.Unix() / 60)
	minuteBytes := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		minuteBytes[i] = byte(minute & 0xff)
		minute >>= 8
	}

	// 合并随机字节和时间字节
	nonceBytes := append(randomBytes, minuteBytes[:4]...)

	// 将结果编码为Base64字符串
	encoded := base64.StdEncoding.EncodeToString(nonceBytes)

	return encoded
}
