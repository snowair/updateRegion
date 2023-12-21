package utils

type ProvinceInfo struct {
	Name     string  // 官方名称
	Display  string  // 外显名称
	Code     string `json:"-"`
	CityInfo []CityInfo
}
type CityInfo struct {
	Name     string
	Display  string  // 外显名称
	Code     string `json:"-"`
	AreaInfo []AreaInfo
}
type AreaInfo struct {
	Name string
	Display  string  // 外显名称
	Code     string `json:"-"`
}
