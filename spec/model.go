package spec

type (
	SpecDetail struct {
		Urn         string                 `json:"urn"`
		Type        string                 `json:"type"`
		Name        string                 `json:"name"`
		Prop        string                 `json:"prop"`
		Description string                 `json:"description"`
		Services    map[string]SpecService `json:"services"`
	}
	Spec struct {
		Status  string `json:"status"`
		Model   string `json:"model"`
		Version int    `json:"version"`
		Type    string `json:"type"`
		Ts      int    `json:"ts"`
	}
	Info struct {
		Specs []Spec `json:"specs"`
	}

	SpecProperty struct {
		Iid         int      `json:"iid"`
		Urn         string   `json:"urn"`
		Type        string   `json:"type"`
		Name        string   `json:"name"`
		Prop        string   `json:"prop"`
		Description string   `json:"description"`
		Format      string   `json:"format"`
		Access      []string `json:"access"`
		Unit        string   `json:"unit"`
		ValueRange  []int    `json:"value-range"`
		Source      int      `json:"source"`
		GattAccess  []string `json:"gatt-access"`
	}
	SpecAction struct {
		Iid         int           `json:"iid"`
		Urn         string        `json:"urn"`
		Type        string        `json:"type"`
		Name        string        `json:"name"`
		Prop        string        `json:"prop"`
		Description string        `json:"description"`
		In          []interface{} `json:"in"`
		Out         []interface{} `json:"out"`
	}
	SpecService struct {
		Iid         int                     `json:"iid"`
		Urn         string                  `json:"urn"`
		Type        string                  `json:"type"`
		Name        string                  `json:"name"`
		Prop        string                  `json:"prop"`
		Description string                  `json:"description"`
		Properties  map[string]SpecProperty `json:"properties"`
		Actions     map[string]SpecAction   `json:"actions"`
	}
)
