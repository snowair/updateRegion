package utils

type ProvinceInfo struct {
	Name     string
	Code     string `json:"-"`
	CityInfo []CityInfo
}
type CityInfo struct {
	Name     string
	Code     string `json:"-"`
	AreaInfo []AreaInfo
}
type AreaInfo struct {
	Name string
	Code     string `json:"-"`
}
