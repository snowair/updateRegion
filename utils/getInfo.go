package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/mozillazg/go-pinyin"
)

var py pinyin.Args

func init() {
	py = pinyin.NewArgs()
}

func DoGetInfo() {

	var cInfos = []*CityInfo{}
	provinces := []string{
		"北京市（京）:110000",
		"天津市（津）:120000",
		"河北省（冀）:130000",
		"山西省（晋）:140000",
		"内蒙古自治区（内蒙古）:150000",
		"辽宁省（辽）:210000",
		"吉林省（吉）:220000",
		"黑龙江省（黑）:230000",
		"上海市（沪）:310000",
		"江苏省（苏）:320000",
		"浙江省（浙）:330000",
		"安徽省（皖）:340000",
		"福建省（闽）:350000",
		"江西省（赣）:360000",
		"山东省（鲁）:370000",
		"河南省（豫）:410000",
		"湖北省（鄂）:420000",
		"湖南省（湘）:430000",
		"广东省（粤）:440000",
		"广西壮族自治区（桂）:450000",
		"海南省（琼）:460000",
		"重庆市（渝）:500000",
		"四川省（川、蜀）:510000",
		"贵州省（黔、贵）:520000",
		"云南省（滇、云）:530000",
		"西藏自治区（藏）:540000",
		"陕西省（陕、秦）:610000",
		"甘肃省（甘、陇）:620000",
		"青海省（青）:630000",
		"宁夏回族自治区（宁）:640000",
		"新疆维吾尔自治区（新）:650000",
		"香港特别行政区（港）:810000",
		"澳门特别行政区（澳）:820000",
		"台湾省（台）:710000",
	}
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

		pInfo.N = p
		pInfo.D = simple
		pInfo.Code = provinceInfo[1]

		spinyin := pinyin.Pinyin(pInfo.D, py)
		for _, v := range spinyin {
			pInfo.PY += strings.Join(v, "")
		}
		if len(pInfo.PY) > 0 {
			pInfo.C = string([]rune(pInfo.PY)[0])
		} else {
			pInfo.C = "z"
		}

		fmt.Printf("进度: %d/%d 正在查询省份%s...\r\n", (index + 1), len(provinces), province)
		citys, err := getRegionInfo(provinceInfo[0], "")
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		cityInfos := getCitys(citys, &pInfo)
		cInfos = append(cInfos, cityInfos...)

		time.Sleep(time.Second * 1)
	}

	//3. 输出json,csv
	writeJsonFile(makeCityList(cInfos))
	//writeJsonFile(cInfos)

	log.Println("查询完成，已输出json、csv文件到：", GetExeDir())
}

func makeCityList(citys []*CityInfo) CitySelect {
	capitals := map[string][]*CityInfo{} // 记录相同首字母的城市
	capitalsList := [][]*CityInfo{}

	// 遍历所有城市
	for _, v := range citys {
		if _, ok := capitals[v.C]; !ok {
			capitals[v.C] = []*CityInfo{} // 新建
		}

		capitals[v.C] = append(capitals[v.C], v)
	}

	for _, v := range capitals {
		sort.Slice(v, func(i, j int) bool {
			if v[i].PY < v[j].PY {
				return true
			}
			return false
		})
		capitalsList = append(capitalsList, v)
	}

	sort.Slice(capitalsList, func(i, j int) bool {
		if capitalsList[i][0].C < capitalsList[j][0].C {
			return true
		}
		return false
	})

	tmp := CitySelect{
		Hot:  []CityInfo{},
		List: []CitySection{},
	}

	for _, c := range capitalsList {
		section := CitySection{}
		for _, v := range c {
			section.C = v.C
			section.List = c
			if bigCityMap[v.N] {
				tmp.Hot = append(tmp.Hot, *v)
			}
		}
		tmp.List = append(tmp.List, section)
	}

	return tmp
}

var bigCityMap = map[string]bool{
	"北京市": true,
	"天津市": true,
	"上海市": true,
	"重庆市": true,
	"广州市": true,
	"深圳市": true,
	"杭州市": true,
	"西安市": true,
	"武汉市": true,
	"成都市": true,
}

// 获取某个省份下所有城市
func getCitys(citys []map[string]interface{}, pInfo *ProvinceInfo) []*CityInfo {
	for _, city := range citys {
		cName := city["diji"].(string)

		cCode := city["quHuaDaiMa"].(string)

		var cInfo = CityInfo{}
		cInfo.N = cName
		cInfo.P = pInfo.N
		cInfo.D = strings.TrimSuffix(cName, "市")
		cInfo.D = strings.TrimSuffix(cName, "盟")
		cInfo.D = strings.TrimSuffix(cName, "盟")
		cInfo.D = strings.TrimSuffix(cName, "藏族羌族自治州")
		cInfo.D = strings.TrimSuffix(cName, "蒙古自治州")
		cInfo.D = strings.TrimSuffix(cName, "回族自治州")
		cInfo.D = strings.TrimSuffix(cName, "彝族自治州")
		cInfo.D = strings.TrimSuffix(cName, "白族自治州")
		cInfo.D = strings.TrimSuffix(cName, "傣族景颇族自治州")
		cInfo.D = strings.TrimSuffix(cName, "藏族自治州")
		cInfo.D = strings.TrimSuffix(cName, "土家族苗族自治州")
		cInfo.D = strings.TrimSuffix(cName, "蒙古族藏族自治州☆")
		cInfo.D = strings.TrimSuffix(cName, "哈尼族彝族自治州")
		cInfo.D = strings.TrimSuffix(cName, "柯尔克孜自治州")
		cInfo.D = strings.TrimSuffix(cName, "回族自治州")
		cInfo.D = strings.TrimSuffix(cName, "苗族侗族自治州")
		cInfo.D = strings.TrimSuffix(cName, "布依族苗族自治州")
		cInfo.D = strings.TrimSuffix(cName, "壮族苗族自治州")
		cInfo.D = strings.TrimSuffix(cName, "傣族自治州")
		cInfo.D = strings.TrimSuffix(cName, "朝鲜族自治州")
		cInfo.D = strings.TrimSuffix(cName, "哈萨克自治州☆")
		cInfo.D = strings.TrimSuffix(cName, "傈僳族自治州")
		cInfo.D = strings.TrimSuffix(cName, "地区")
		cInfo.Code = cCode

		spinyin := pinyin.Pinyin(cInfo.D, py)
		for _, v := range spinyin {
			cInfo.PY += strings.Join(v, "")
		}
		if len(cInfo.PY) > 0 {
			cInfo.C = string([]rune(cInfo.PY)[0])
		} else {
			cInfo.C = "z"
		}

		if cName == "" {
			// TODO: 新疆直辖县

		}

		// 查询区县行政区
		//areas, err := getRegionInfo(province, cName)
		//if err != nil {
		//fmt.Println(err.Error())
		//continue
		//} else {
		//fmt.Println(spinyin)
		//}

		//getAreas(areas, &cInfo)
		pInfo.CityInfo = append(pInfo.CityInfo, &cInfo)
	}
	return pInfo.CityInfo
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
		//cInfo.AreaInfo = append(cInfo.AreaInfo, aInfo)

	}
}
func getRegionInfo(province string, city string) (jsonArr []map[string]interface{}, err error) {
	target := "http://xzqh.mca.gov.cn/selectJson"
	pData := url.Values{}
	if province != "" {
		pData.Set("shengji", province)
	}
	if city != "" {
		pData.Set("diji", city)
	}
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded; charset=UTF-8"
	headers["Origin"] = "http://xzqh.mca.gov.cn"
	headers["Referer"] = "http://xzqh.mca.gov.cn/map"
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36 Edg/114.0.1788.0"
	resp, err := Execute(target, "POST", pData, headers)
	if err != nil {
		log.Println("Execute err>>", err.Error())
		return nil, err
	}

	err = json.Unmarshal([]byte(resp), &jsonArr)
	if err != nil {
		fmt.Println(resp)
		return jsonArr, err
	}

	//log.Println(jsonArr)
	return jsonArr, nil
}
