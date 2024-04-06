package device

type DeviceInfo struct {
	Did         string `json:"did"`
	Token       string `json:"token"`
	Longitude   string `json:"longitude"`
	Latitude    string `json:"latitude"`
	Name        string `json:"name"`
	Pid         string `json:"pid"`
	Localip     string `json:"localip"`
	Mac         string `json:"mac"`
	Ssid        string `json:"ssid"`
	Bssid       string `json:"bssid"`
	ParentId    string `json:"parent_id"`
	ParentModel string `json:"parent_model"`
	ShowMode    int    `json:"show_mode"`
	Model       string `json:"model"`
	AdminFlag   int    `json:"adminFlag"`
	ShareFlag   int    `json:"shareFlag"`
	PermitLevel int    `json:"permitLevel"`
	IsOnline    bool   `json:"isOnline"`
	Desc        string `json:"desc"`
	Uid         int    `json:"uid"`
	PdId        int    `json:"pd_id"`
	Password    string `json:"password"`
	P2PId       string `json:"p2p_id"`
	Rssi        int    `json:"rssi"`
	FamilyId    int    `json:"family_id"`
	ResetFlag   int    `json:"reset_flag"`
	DescNew     string `json:"desc_new,omitempty"`
	DescTime    []int  `json:"desc_time,omitempty"`
}

type DeviceListResult struct {
	List []DeviceInfo `json:"list"`
}

type ActionDetail struct {
	Did  string `json:"did"`
	Siid int    `json:"siid"`
	Aiid int    `json:"aiid"`
	In   []any  `json:"in"`
}

type ActionResult struct {
	Did         string `json:"did"`
	Miid        int    `json:"miid"`
	Siid        int    `json:"siid"`
	Aiid        int    `json:"aiid"`
	Code        int    `json:"code"`
	ExeTime     int    `json:"exe_time"`
	NetCost     int    `json:"net_cost"`
	Otlocalts   int64  `json:"otlocalts"`
	OaCost      int    `json:"oa_cost"`
	OaRpcCost   int    `json:"_oa_rpc_cost"`
	WithLatency int    `json:"withLatency"`
}
