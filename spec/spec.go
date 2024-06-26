package spec

import "github.com/go-resty/resty/v2"

const (
	HOST = "https://home.miot-spec.com"
)

var req *resty.Client

func init() {
	req = resty.New().SetBaseURL(HOST)
}

// Search for spec by keyWord
func Search(keyWord string) ([]Spec, error) {
	resp, err := req.R().Get("/s/" + keyWord)
	if err != nil {
		return nil, err
	}
	return decodeSpecs(resp.Body())
}

// Detail get SpecDetail by model
func Detail(model string) (*SpecDetail, error) {
	resp, err := req.R().Get("/spec/" + model)
	if err != nil {
		return nil, err
	}
	return decodeDetail(resp.Body())
}
