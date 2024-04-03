package miot_spec

import (
	"bytes"
	"github.com/PuerkitoBio/goquery"
	jsoniter "github.com/json-iterator/go"
)

func decodeApp(b []byte) (string, error) {
	io := bytes.NewReader(b)
	doc, err := goquery.NewDocumentFromReader(io)
	if err != nil {
		return "", err
	}
	val, _ := doc.Find("#app").Attr("data-page")
	return val, nil
}

func decodeSpecs(b []byte) ([]Spec, error) {
	val, err := decodeApp(b)
	if err != nil {
		return nil, err
	}
	ret := make([]Spec, 0)
	list := make([]Info, 0)
	props := jsoniter.Get([]byte(val), "props")
	props.Get("list").ToVal(&list)
	if len(list) == 0 {
		props.Get("specs").ToVal(&list)
	}
	for _, v := range list {
		ret = append(ret, v.Specs...)
	}
	return ret, nil
}

func decodeDetail(b []byte) (*SpecDetail, error) {
	val, err := decodeApp(b)
	if err != nil {
		return nil, err
	}
	detail := &SpecDetail{}

	jsoniter.Get([]byte(val), "props").Get("spec").ToVal(detail)
	return detail, nil
}
