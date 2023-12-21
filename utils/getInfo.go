package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

func DoGetInfo() {
	var pInfos = []ProvinceInfo{}
	//暂不包含台湾省的数据
	provinces := strings.Split("上海市（沪）:310000,江苏省（苏）:320000,浙江省（浙）:330000,安徽省（皖）:340000,福建省（闽）:350000,江西省（赣）:360000,山东省（鲁）:370000,河南省（豫）:410000,湖北省（鄂）:420000,湖南省（湘）:430000,广东省（粤）:440000,广西壮族自治区（桂）:450000,海南省（琼）:460000,重庆市（渝）:500000,四川省（川、蜀）:510000,贵州省（黔、贵）:520000,云南省（滇、云）:530000,西藏自治区（藏）:540000,陕西省（陕、秦）:610000,甘肃省（甘、陇）:620000,青海省（青）:630000,宁夏回族自治区（宁）:640000,新疆维吾尔自治区（新）:650000,香港特别行政区（港）:810000,澳门特别行政区（澳）:820000,台湾省（台）:710000", ",")

	//1.遍历省份,获取每个省份下的地级市
	for index, province := range provinces {
		// if index >= 3 {
		// 	break
		// }
		provinceInfo := strings.Split(province, ":")
		var pInfo = ProvinceInfo{}

		p := provinceInfo[0][:strings.IndexRune(provinceInfo[0], rune('（'))]

		simple := strings.TrimSuffix(p, "省")
		simple = strings.TrimSuffix(simple, "市")
		simple = strings.TrimSuffix(simple, "自治区")
		simple = strings.TrimSuffix(simple, "特别行政区")
		simple = strings.TrimSuffix(simple, "维吾尔")
		simple = strings.TrimSuffix(simple, "回族")
		simple = strings.TrimSuffix(simple, "壮族")

		pInfo.Name = p
		pInfo.Display = simple
		pInfo.Code = provinceInfo[1]

		fmt.Printf("进度: %d/%d 正在查询省份%s...\r\n", (index + 1), len(provinces), province)
		citys, err := getRegionInfo(provinceInfo[0], "")
		if err != nil {
			continue
		}
		//2. 遍历地级市，获取每个地级市下的区县
		//地级市编码
		getCitys(citys, provinceInfo[0], &pInfo)
		pInfos = append(pInfos, pInfo)
		time.Sleep(time.Second * 1)
	}

	//3. 输出json,csv
	writeJsonFile(pInfos)
	writeCsvFile(pInfos)
	log.Println("查询完成，已输出json、csv文件到：", GetExeDir())
}

var bigCityMap = map[string]bool{
	"北京市": true,
	"天津市": true,
	"上海市": true,
	"重庆市": true,
	"广州市": true,
	"深圳市": true,
}

// 获取某个省份下所有城市
func getCitys(citys []map[string]interface{}, province string, pInfo *ProvinceInfo) {
	for _, city := range citys {
		cName := city["diji"].(string)

		cCode := city["quHuaDaiMa"].(string)

		var cInfo = CityInfo{}
		cInfo.Name = cName
		cInfo.Display = strings.TrimSuffix(cName, "市")
		cInfo.Code = cCode

		if !bigCityMap[cName] {
			pInfo.CityInfo = append(pInfo.CityInfo, cInfo)
			continue
		} else {
			areas, err := getRegionInfo(province, cName)
			if err != nil {
				continue
			}

			getAreas(areas, &cInfo)
			pInfo.CityInfo = append(pInfo.CityInfo, cInfo)
		}
	}
}

// 获取某个城市下所有的区县
func getAreas(areas []map[string]interface{}, cInfo *CityInfo) {
	for _, area := range areas {
		aName := area["xianji"].(string)
		aCode := area["quHuaDaiMa"].(string)

		var aInfo = AreaInfo{}
		aInfo.Name = aName
		aInfo.Display = aName
		aInfo.Code = aCode
		cInfo.AreaInfo = append(cInfo.AreaInfo, aInfo)

	}
}
func getRegionInfo(province string, city string) (jsonArr []map[string]interface{}, err error) {
	url := "http://xzqh.mca.gov.cn/selectJson"
	pData := ""
	//var myMap map[string] string
	// var myMap = make(map[string]string)
	if province != "" {
		pData = "shengji=" + province
	}
	if city != "" {
		pData += "&diji=" + city
	}
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded; charset=utf-8"
	resp, err := Execute(url, "POST", ([]byte)(pData), headers)
	if err != nil {
		log.Println("Execute err>>", err.Error())
		return nil, err
	}

	json.Unmarshal(resp.([]byte), &jsonArr)

	//log.Println(jsonArr)
	return jsonArr, nil
}
