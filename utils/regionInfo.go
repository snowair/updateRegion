package utils

type ProvinceInfo struct {
	Name     string // 官方名称
	Display  string // 外显名称
	Code     string `json:"-"`
	PY       string // 拼音
	C        string // 首字母
	CityInfo []*CityInfo
}
type CityInfo struct {
	Province string //ProvinceName
	Name     string
	Display  string // 外显名称
	PY       string // 拼音
	C        string // 首字母
	Code     string `json:"-"`
	//AreaInfo []AreaInfo
}
type AreaInfo struct {
	Name    string
	Display string // 外显名称
	Code    string `json:"-"`
}

type CityList struct {
	Name     string // 官方名称
	Display  string // 外显名称
	Code     string `json:"-"`
	CityInfo []CityInfo
}

type CitySelect struct {
	Hot  []CityInfo    `json:"hot,omitempty"`
	List []CitySection `json:"list"`
}

type CitySection struct {
	C    string      `json:"capital"`
	List []*CityInfo `json:"list"`
}
