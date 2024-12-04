# golang获取最新行政区划信息

## 一、背景

在[乐琪药品流向数据查询管理系统（三）功能篇](https://www.qipanet.com/project/279.html)中有提到过系统支持自动更新行政区划信息，那怎么将省、市、区县信息按需随便更新呢？当时在网上找了很久，找到了一些公开的资料或者开源的小项目，有的很久没维护，有的是区域信息错漏百出，总之没找到适合我们使用的，靠别人的都不太靠谱。既然找不到，那就自己造个轮子吧！反正自己项目也是需要用到的。

## 二、原理

在==民政部==官网上（xzqh.mca.gov.cn），可以查询到各省行政区划信息，但是却不提供文件下载，需要一个个手动查询。3202的今天，竟然不给人下载数据，有点理解不了。抽查了几个近年有变更的区域，发现数据还挺准的，毕竟是国家级网站，还是有一定的权威性，比那些阿猫阿狗的付费接口精准多了。

抓包分析了下发现查询接口还是比较简陋的，于是做了个简单的爬虫模块，丢到quartz里做定时任务，就实现了定时更新行政区划信息，整个过程还是比较简单的。


核对最新数据: [2022年中华人民共和国县以上行政区划代码](https://www.mca.gov.cn/mzsj/xzqh/2022/202201xzqh.html)


## 三、开源

这么个小需求，网上竟找不到合适的解决方法，于是做完这个模块后就想着把它开源出来，希望能帮助到有需要的同学。但是这个模块是在spring boot项目里用的，不方便拿出来，因此用golang写了个简单的爬虫程序，自动下载全国各省份（暂不包含台湾省）的行政区划数据。

`main.exe`为主程序，直接运行即可输出的文件有json、csv两种格式的数据，路径为主程序根目录。==在使用时请做好延时处理，避免采集过于频繁导致网站崩溃！==

> 爬虫模块在实际项目中仅每个月使用一次，如有侵权请联系我删除！

## 四、演示

* 采集过程显示进度

  ![image-20231004164211980](README/image-20231004164211980.png)

* json格式

![image-20231004155359822](README/image-20231004155359822.png)

## 五、其他

- 源码：https://github.com/root6819/usgApiWithGo
- 官网： [http://www.qipanet.com](http://www.qipanet.com/)
- 微信 root6819 | QQ 302777528


```
{
"P": "海南省",
"N": "省直辖县级行政单位",
"D": "省直辖县级行政单位",
"PY": "shengzhixiaxianjixingzhengdanwei",
"C": "s"
},
[{"children":[],"diji":"","quHuaDaiMa":"469001","quhao":"0898","shengji":"","xianji":"五指山市"},{"children":[],"diji":"","quHuaDaiMa":"469002","quhao":"0898","shengji":"","xianji":"琼海市"},{"children":[],"diji":"","quHuaDaiMa":"469005","quhao":"0898","shengji":"","xianji":"文昌市"},{"children":[],"diji":"","quHuaDaiMa":"469006","quhao":"0898","shengji":"","xianji":"万宁市"},{"children":[],"diji":"","quHuaDaiMa":"469007","quhao":"0898","shengji":"","xianji":"东方市"},{"children":[],"diji":"","quHuaDaiMa":"469021","quhao":"0898","shengji":"","xianji":"定安县"},{"children":[],"diji":"","quHuaDaiMa":"469022","quhao":"0898","shengji":"","xianji":"屯昌县"},{"children":[],"diji":"","quHuaDaiMa":"469023","quhao":"0898","shengji":"","xianji":"澄迈县"},{"children":[],"diji":"","quHuaDaiMa":"469024","quhao":"0898","shengji":"","xianji":"临高县"},{"children":[],"diji":"","quHuaDaiMa":"469025","quhao":"0898","shengji":"","xianji":"白沙黎族自治县"},{"children":[],"diji":"","quHuaDaiMa":"469026","quhao":"0898","shengji":"","xianji":"昌江黎族自治县"},{"children":[],"diji":"","quHuaDaiMa":"469027","quhao":"0898","shengji":"","xianji":"乐东黎族自治县"},{"children":[],"diji":"","quHuaDaiMa":"469028","quhao":"0898","shengji":"","xianji":"陵水黎族自治县"},{"children":[],"diji":"","quHuaDaiMa":"469029","quhao":"0898","shengji":"","xianji":"保亭黎族苗族自治县"},{"children":[],"diji":"","quHuaDaiMa":"469030","quhao":"0898","shengji":"","xianji":"琼中黎族苗族自治县"}]
```

```
{
"P": "河南省",
"N": "省直辖县级行政单位",
"D": "省直辖县级行政单位",
"PY": "shengzhixiaxianjixingzhengdanwei",
"C": "s"
},
```


```
{
"P": "湖北省",
"N": "省直辖县级行政单位",
"D": "省直辖县级行政单位",
"PY": "shengzhixiaxianjixingzhengdanwei",
"C": "s"
}
[{"children":[],"diji":"","quHuaDaiMa":"429004","quhao":"0728","shengji":"","xianji":"仙桃市"},{"children":[],"diji":"","quHuaDaiMa":"429005","quhao":"0728","shengji":"","xianji":"潜江市"},{"children":[],"diji":"","quHuaDaiMa":"429006","quhao":"0728","shengji":"","xianji":"天门市"},{"children":[],"diji":"","quHuaDaiMa":"429021","quhao":"0719","shengji":"","xianji":"神农架林区"}]
```


```
{
"P": "新疆维吾尔自治区",
"N": "自治区直辖县级行政单位",
"D": "自治区直辖县级行政单位",
"PY": "zizhiquzhixiaxianjixingzhengdanwei",
"C": "z"
}
[{"children":[],"diji":"","quHuaDaiMa":"659001","quhao":"0993","shengji":"","xianji":"石河子市"},{"children":[],"diji":"","quHuaDaiMa":"659002","quhao":"0997","shengji":"","xianji":"阿拉尔市"},{"children":[],"diji":"","quHuaDaiMa":"659003","quhao":"0998","shengji":"","xianji":"图木舒克市"},{"children":[],"diji":"","quHuaDaiMa":"659004","quhao":"0994","shengji":"","xianji":"五家渠市"},{"children":[],"diji":"","quHuaDaiMa":"659005","quhao":"0906","shengji":"","xianji":"北屯市"},{"children":[],"diji":"","quHuaDaiMa":"659006","quhao":"0906","shengji":"","xianji":"铁门关市"},{"children":[],"diji":"","quHuaDaiMa":"659007","quhao":"0909","shengji":"","xianji":"双河市"},{"children":[],"diji":"","quHuaDaiMa":"659008","quhao":"0999","shengji":"","xianji":"可克达拉市"},{"children":[],"diji":"","quHuaDaiMa":"659009","quhao":"0903","shengji":"","xianji":"昆玉市"},{"children":[],"diji":"","quHuaDaiMa":"659010","quhao":"0992","shengji":"","xianji":"胡杨河市"},{"children":[],"diji":"","quHuaDaiMa":"659011","quhao":"0902","shengji":"","xianji":"新星市"},{"children":[],"diji":"","quHuaDaiMa":"659012","quhao":"0901","shengji":"","xianji":"白杨市"}]
```
